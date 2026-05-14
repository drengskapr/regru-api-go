package zonecontrol

import (
	"encoding/json"

	"github.com/drengskapr/regru-api-go/internal/client"
)

type DnsRr struct {
	Content string
	Prio    int
	Rectype string
	State   string
	Subname string
}

type DnsDomain struct {
	Dname       string
	ErrorCode   string
	ErrorText   string
	ErrorParams map[string]string
	Result      string
	Rrs         []DnsRr
	ServiceId   string
	Servtype    string
	Soa         map[string]string
}

type DnsAnswer struct {
	Domains []DnsDomain
}

type DnsRecords struct {
	Answer       DnsAnswer         `json:"answer,omitempty"`
	Charset      string            `json:"charset,omitempty"`
	Messagestore string            `json:"messagestore,omitempty"`
	Result       string            `json:"result,omitempty"`
	ErrorCode    string            `json:"error_code,omitempty"`
	ErrorText    string            `json:"error_text,omitempty"`
	ErrorParams  map[string]string `json:"error_params,omitempty"`
}

const apiUrl = "https://api.reg.ru/api/regru2/"
const zoneGetRrs = "zone/get_resource_records"
const zoneAddTxt = "zone/add_txt"
const zoneRemoveRrs = "zone/remove_record"
const zoneAddAlias = "zone/add_alias"

func parseResponse(rawData []byte) (DnsRecords, error) {
	b := DnsRecords{}
	err := json.Unmarshal(rawData, &b)
	return b, err
}

// GetZones returns resource records for domain.
func GetZones(username, password, domainName string) (DnsRecords, error) {
	postFields := map[string]string{
		"username":    username,
		"password":    password,
		"domain_name": domainName,
	}
	body, err := client.ApiRequest(apiUrl+zoneGetRrs, postFields)
	if err != nil {
		return DnsRecords{}, err
	}
	return parseResponse(body)
}

// AddTxtRr adds a TXT resource record for domain.
func AddTxtRr(username, password, domainName, subdomain, textBody string) (DnsRecords, error) {
	postFields := map[string]string{
		"username":    username,
		"password":    password,
		"domain_name": domainName,
		"subdomain":   subdomain,
		"text":        textBody,
	}
	body, err := client.ApiRequest(apiUrl+zoneAddTxt, postFields)
	if err != nil {
		return DnsRecords{}, err
	}
	return parseResponse(body)
}

// RmRr removes a resource record for domain.
func RmRr(username, password, domainName, subdomain, resourceRecordType, content string) (DnsRecords, error) {
	postFields := map[string]string{
		"username":    username,
		"password":    password,
		"domain_name": domainName,
		"subdomain":   subdomain,
		"record_type": resourceRecordType,
	}
	if content != "" {
		postFields["content"] = content
	}
	body, err := client.ApiRequest(apiUrl+zoneRemoveRrs, postFields)
	if err != nil {
		return DnsRecords{}, err
	}
	return parseResponse(body)
}

// AddARr adds an A resource record for domain.
func AddARr(username, password, domainName, subdomain, ipAddr string) (DnsRecords, error) {
	postFields := map[string]string{
		"username":    username,
		"password":    password,
		"domain_name": domainName,
		"subdomain":   subdomain,
		"ipaddr":      ipAddr,
	}
	body, err := client.ApiRequest(apiUrl+zoneAddAlias, postFields)
	if err != nil {
		return DnsRecords{}, err
	}
	return parseResponse(body)
}
