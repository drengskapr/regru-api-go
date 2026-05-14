# regru-api-go module for Reg.ru API v2

https://www.reg.ru provide API access to control users, billing, domains etc.

Currently only several zone (domain) control functions implemented in this module.
```bash
"zone/get_resource_records"
"zone/add_txt"
"zone/remove_record"
```
API documentation https://www.reg.ru/reseller/api2doc#common

# Access from known IP-address only

Access configuration https://www.reg.ru/user/account/#/settings/api/

```json
{
   "charset" : "utf-8",
   "error_code" : "ACCESS_DENIED_FROM_IP",
   "error_params" : {
      "command_name" : "zone/get_resource_records"
   },
   "error_text" : "Access to API from this IP denied",
   "messagestore" : null,
   "result" : "error"
}
```

## Examples

```go
package main

import (
	"fmt"
	"os"
	"github.com/drengskapr/regru-api-go/zonecontrol"
)

func main() {
	username := os.Getenv("API_USERNAME")
	password := os.Getenv("API_PASSWORD")
	domainName := "mydomain.com"

	// Get DNS resource records
	rec, err := zonecontrol.GetZones(username, password, domainName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("result: %s\n", rec.Result)

	// Add a TXT record
	// zonecontrol.AddTxtRr(username, password, domainName, "_acme_foo_bar", "txt-record-content")

	// Add an A record
	// zonecontrol.AddARr(username, password, domainName, "www", "1.2.3.4")

	// Remove a record
	// zonecontrol.RmRr(username, password, domainName, "_acme_example", "TXT", "")
}
```