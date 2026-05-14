package regru

import "github.com/drengskapr/regru-api-go/internal/zonecontrol"

// DNS response types — defined in internal/zonecontrol, re-exported here for callers.
type DnsRecords = zonecontrol.DnsRecords
type DnsAnswer  = zonecontrol.DnsAnswer
type DnsDomain  = zonecontrol.DnsDomain
type DnsRr      = zonecontrol.DnsRr
