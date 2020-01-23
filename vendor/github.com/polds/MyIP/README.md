# MyIP

MyIP is a dead simple Go package that only lives to return **your** External IP Address. It pulls from a list of websites known to serve IP Address info, and will iterate through every website if needed until it returns your IP Address. At the moment there are only two websites it pulls from, Pull Requests are gladly welcome to add additional websites. 

### License

MIT: do whatever you want with it.

### Usage

```
go get github.com/polds/MyIP
```

    package main

    import "github.com/polds/MyIP"

    func main() {
        ip, err := myip.GetMyIP()
        // error handling omitted for brevity

        print(ip)
    }

### Issues

I have no way to check if this works for IPv6. Only confirmed to work with IPv4.

### Todo

 * Add concurrency with buffer to handle multiple requests at the same time and return with the first response
 * Add timeouts to the transport. 

### Questions, Comments, Communication?

Tweet me: [@Peter_Olds](https://twitter.com/Peter_Olds)