# pihole-api-go# 

**WIP** : pihole-api-go is a go Client for interacting with Pihole by https://pi-hole.net/.

## Installation ##

pihole-api-go is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NicoFgrx/pihole-api-go
```

will resolve and add the package to the current development module, along with its dependencies.


Alternatively the same can be achieved if you use import in a package:

```go
import "github.com/NicoFgrx/pihole-api-go/api"
```

and run `go get` without parameters.

## Usage ##

```go
import pihole "github.com/NicoFgrx/pihole-api-go/api"
```

Construct a new pihole client, then use the various services on the client to
access different parts of the Pihole API. For example:

```go

	url := "http://localhost:8080/admin/api.php" // must be http[s]://<IP>:<port>/admin/api.php
	key := "xxx" // find the token on the web UI in Settings>API

	client := pihole.NewClient(url, key)

    // Get all custom dns defined on the pihole
    customdns_lst, err := client.GetAllCustomDNS()

```

You can find more examples in examples folder.


## License ##
This library is distributed under the MIT licence found in the [LICENSE](./LICENSE)
file.

