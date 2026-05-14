package regru

import "github.com/drengskapr/regru-api-go/internal/zonecontrol"

// Client holds credentials for the Reg.ru API.
type Client struct {
	username string
	password string
}

// New creates a new Client with the given credentials.
func New(username, password string) *Client {
	return &Client{username: username, password: password}
}

// GetZones returns DNS resource records for the domain.
func (c *Client) GetZones(domainName string) (DnsRecords, error) {
	return zonecontrol.GetZones(c.username, c.password, domainName)
}

// AddTxtRr adds a TXT resource record for the domain.
func (c *Client) AddTxtRr(domainName, subdomain, text string) (DnsRecords, error) {
	return zonecontrol.AddTxtRr(c.username, c.password, domainName, subdomain, text)
}

// AddARr adds an A resource record for the domain.
func (c *Client) AddARr(domainName, subdomain, ipAddr string) (DnsRecords, error) {
	return zonecontrol.AddARr(c.username, c.password, domainName, subdomain, ipAddr)
}

// RmRr removes a resource record for the domain.
func (c *Client) RmRr(domainName, subdomain, recordType, content string) (DnsRecords, error) {
	return zonecontrol.RmRr(c.username, c.password, domainName, subdomain, recordType, content)
}
