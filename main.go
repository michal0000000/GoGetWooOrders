package main

import (
  "github.com/dinistavares/go-woocommerce-api"
  "github.com/joho/godotenv"
  "fmt"
  "flag"
  "strings"
  "os"
  "encoding/csv"
  "time"
  "unicode"
)


type csvRow struct {
  orderId string
  email string
}

func main () {
	
  o := flag.String("orders", "", "List of orders separated with ',', i.e. 123,456,789 ")
  flag.Parse()
  orders := strings.Split(*o,",")

  shopURL := goDotEnvVariable("SHOP_URL")
  key := goDotEnvVariable("CUSTOMER_KEY")
  secret := goDotEnvVariable("CUSTOMER_SECRET")

  client, err := woocommerce.New(shopURL)

  if err != nil {
    fmt.Printf("Error: Couldn't connect to shop using url %s\n",shopURL)
    return
  }

  client.Authenticate(key, secret)

  opts := woocommerce.GetOrderParams{
  }

  csvContent := []csvRow{
	  {"order_id","email"},
  }

  for _,orderId := range orders {
	if isNumeric(orderId) == false{
	  fmt.Printf("Error: %s is not a proper order id\n",orderId)
	  continue
	}

	// handle trailing commad
	if len(orderId) == 1 && orderId[0] == ',' {
		continue
	}

  	order, _, err := client.Orders.Get(orderId,&opts)
        if err != nil {
	  fmt.Printf("Error: Couldn't fetch order %s, reason: %s\n",orderId,err)
	  continue
        }   
	email := order.Billing.Email
	newRow := csvRow{orderId,email}
	csvContent = append(csvContent, newRow)
  }

  today := time.Now()
  formattedDate := today.Format("20060102")
  filename := fmt.Sprintf("orders-%s.csv", formattedDate)

  csvFile, err := os.Create(filename)
  if err != nil {
    fmt.Println("failed creating file: %s", err)
  }
  defer csvFile.Close()
  
  w := csv.NewWriter(csvFile)
  defer w.Flush()

  var data [][]string
	
  if len(csvContent) > 1 {
	fmt.Println()
	fmt.Printf("%-10s %-30s\n", "ORDER_ID", "EMAIL")
  }

  for i, record := range csvContent {
      row := []string{record.orderId, record.email}
      data = append(data, row)

      if i > 0 {
      	fmt.Printf("%-10s %-20s\n", record.orderId,record.email)
      }
  }
  w.WriteAll(data)

  fmt.Printf("\nDone. Wrote file %s\n\n",filename)
  
}

func goDotEnvVariable(key string) string {

  // load .env file
  err := godotenv.Load(".env")

  if err != nil {
    fmt.Println("Error loading .env file")
  }

  return os.Getenv(key)
}

func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
