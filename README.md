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
