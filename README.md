# regru-api-go module for Reg.ru API v2

https://www.reg.ru provide API access to control users, billing, domains etc.

Currently only several zone (domain) control functions implemented in this module.
```bash
"zone/get_resource_records"
"zone/add_txt"
"zone/add_alias"
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

	regru "github.com/drengskapr/regru-api-go"
)

func main() {
	c := regru.New(
		os.Getenv("API_USERNAME"),
		os.Getenv("API_PASSWORD"),
	)

	// Get DNS resource records
	rec, err := c.GetZones("mydomain.com")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("result: %s\n", rec.Result)

	// Add a TXT record
	// c.AddTxtRr("mydomain.com", "_acme_foo_bar", "txt-record-content")

	// Add an A record
	// c.AddARr("mydomain.com", "www", "1.2.3.4")

	// Remove a record by subdomain and type (removes all matching records)
	// c.RmRr("mydomain.com", "_acme_example", "TXT", "")

	// Remove a specific record by subdomain, type, and content
	// c.RmRr("mydomain.com", "_acme_example", "TXT", "txt-record-content")
}
```

## Testing

Unit and integration tests are included. Integration tests require live credentials and a whitelisted IP — they skip automatically when env vars are unset.

```bash
# Unit tests only (no credentials needed)
go test ./...

# Integration tests against the live API
export API_USERNAME="your-regru-username"
export API_PASSWORD="your-regru-password"
export API_DOMAIN="yourdomain.com"

go test ./internal/zonecontrol/ -v
```

Integration tests create and remove records under fixed subdomains (`_regru_api_test`, `_regru_api_test_content`, `test-a-record`). Existing records are not touched.