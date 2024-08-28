# GoGetWooOrders

Utility that fetches Woocommerce order data written in Go. Currently supports fetchich Billing emails based on order id's. Results are saved in a csv file in the same directory.

## Get the binary

To use on windows, you can download a precompiled binary from thi repo.
To do so, run the following command from CMD:
```
C:\Users\Kikinko> curl https://github.com/michal0000000/GoGetWooOrders/releases/download/v1.0.0/gowoo-v1.0.0.exe -o gowoo.exe
```

You can now see the binary in your users folder. Next, create a fille named `.env` with and paste the necessarry information to access Woocommerce API:
```
SHOP_URL=https://example.com
CUSTOMER_KEY=ck_xxxxx
CUSTOMER_SECRET=cs_xxxxx
```

Congratulations, you're all set!

## Fetch order data

Working with `GoGetWooOrders` is easy, but in case you forget you can use the `--help` flag to display usage information:
```
C:\Users\Kikinko> .\gowoo.exe --help

Usage of .\gowoo-v1.0.0.exe:
  --orders string
        List of orders separated with ',', i.e. 123,456,789
```

Following this advice, run the tool using the `--orders` flag. Results will be displayed in the CMD and written to a file as well:
```
C:\Users\Kikinko> .\gowoo.exe --orders 39123,39122,39121

ORDER_ID   EMAIL
39123      p.gildan.t@gmail.sk
39122      robo.mikla@azet.com
39121      m.behuncik@gmail.com

Done. Wrote file orders-20240828.csv
```

And that's all there is to it. Enjoy!
