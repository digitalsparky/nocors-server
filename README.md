# nocors-server

NoCORS Server is a basic golang web file server that automatically sends the required
headers for CORS exemptions for local testing.

Compiled binaries are already available in the release section for:

 - Linux (32 and 64 bit)
 - Mac (32 and 64 bit)
 - Windows (32 and 64 bit)

However you can build it yourself if you like.

Simply run:

```go build -ldflags "-s -w" -o nocors-server```

Binaries are also compressed using UPX.

By default the server listens on 127.0.0.1:4000 and serves the current working directory.

This can be changed using flags:

```./nocors-server -a <your-listen_address> -p <your-port> -s <path-to-serve>```

The following would serve /var/www/vhoss/test.com/html to all interfaces on port 6000.

```./nocors-server -a 0.0.0.0 -p 6000 -s /var/www/vhosts/test.com/html```


# Like my stuff?

Would you like to buy me a coffee or send me a tip?
While it's not expected, I would really appreciate it.

[![Paypal](https://www.paypalobjects.com/webstatic/mktg/Logo/pp-logo-100px.png)](https://paypal.me/MattSpurrier) <a href="https://www.buymeacoffee.com/digitalsparky" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/white_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>
