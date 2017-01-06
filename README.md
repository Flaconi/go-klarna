# go-klarna

[![Build Status](https://travis-ci.org/goglue/go-klarna.svg?branch=master)](https://travis-ci.org/goglue/go-klarna)

### About

This library is an implementation for [Klarna payment](https://developers.klarna.com/api/) API in go

### How to use

**Configuration**

Config struct encapsulate the client needs of information

```go
type Config struct {
	BaseURL     *url.URL
	APIUsername string
	APIPassword string
	Timeout     time.Duration
}
```

**Client**

Is the abstraction of `HTTP` client, required by each service in order to operate.

```go
import (
        klarna "github.com/goglue/go-klarna"
        "net/url"
        "time"
)

func main() {
        uri, _ := url.Parse(klarna.EuroAPI)
        conf := klarna.Config{
                BaseURL:uri,
                Timeout: time.Second * 10,
        }
        client := klarna.NewClient(conf)
}
```

Now since we have an instance of the client, we can instantiate any service instance with this client ...

**Service**

Now, since we have a client, let's instantiate a service
```go
import (
        // ...
)

func main() {
        // previous section ...
        client := klarna.NewClient(conf)
        
        // payment service
        paymentSrv := klarna.NewPaymentSrv(client)
        err := paymentSrv.CancelExistingAuthorization("string-token")
        if nil != err {
                // do something with the error
        }
}
```

### Road map
- [x] Implement Checkout API service
- [x] Cover Checkout API service with tests
- [x] Implement Payment API service
- [x] Cover Payment API service with tests
- [ ] Implement Order Management service
- [ ] Cover Order Management service with tests
- [ ] Implement Checkout API Callbacks service
- [ ] Cover Checkout API Callbacks service with tests
