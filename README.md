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
// Get domain information
package main

import (
	"os"
	"github.com/drengskapr/regru-api-go/zonecontrol"
)

var username, password, domainName string

func main() {

	username = os.Getenv("API_USERNAME")
	password = os.Getenv("API_PASSWORD")
	domainName = "mydomain.com"
	
   zonecontrol.GetZones(username, password, domainName)
   
   // Create TXT resource record
   //zonecontrol.AddTxtRr(username, password, domainName, "_acme_foo_bar", "txt-record-content")
   // Remove TXT resource record
	//zonecontrol.RmTxtRr(username, password, domainName, "_acme_example", "TXT", "")
}

```

# Known Issues
Test API returns some fields as strings for test access and as int for real data, and vice versa.

Response for real account:
```json
{
   "answer" : {
      "domains" : [
         {
            "dname" : "example.xyz",
            "result" : "success",
            "rrs" : [
               {
                  "content" : "111.222.111.222",
                  "prio" : 0,
                  "rectype" : "A",
                  "state" : "A",
                  "subname" : "@"
               },
            ],
            "service_id" : "12345678",
            "servtype" : "domain",
            "soa" : {
               "minimum_ttl" : "10m",
               "ttl" : "10m"
            }
         }
      ]
   },
   "charset" : "utf-8",
   "messagestore" : null,
   "result" : "success"
}

```
Response for test account:
```json
{
   "answer" : {
      "domains" : [
         {
            "dname" : "example.com",
            "result" : "success",
            "rrs" : [
               {
                  "content" : "111.222.111.222",
                  "prio" : "0",
                  "rectype" : "A",
                  "state" : "A",
                  "subname" : "www"
               }
            ],
            "service_id" : 12345,
            "servtype" : "domain",
            "soa" : {
               "minimum_ttl" : "12h",
               "ttl" : "1d"
            }
         }
      ]
   },
   "charset" : "utf-8",
   "messagestore" : null,
   "result" : "success"
}

```

```json
WARN[0000] Could not unmarshal json: json: cannot unmarshal string into Go struct field rrsData.answer.Domains.Rrs.Prio of type int 
INFO[0000] The answer is: {Answer:{Domains:[{Dname:mydomain.pro ErrorCode: ErrorText: ErrorParams:map[] Result:success Rrs:[{Content:111.222.111.222 Prio:0 Rectype:A State:A Subname:www}] ServiceId: Servtype:domain Soa:map[minimum_ttl:12h ttl:1d]}]} Charset:utf-8 Messagestore: Result:success ErrorCode: ErrorText: ErrorParams:map[]}

```