package zonecontrol

import (
	"bytes"
	"os"
	"strings"
	"testing"

	log "github.com/sirupsen/logrus"
)

func credentials(t *testing.T) (username, password, domain string) {
	t.Helper()
	username = os.Getenv("API_USERNAME")
	password = os.Getenv("API_PASSWORD")
	domain = os.Getenv("API_DOMAIN")
	if username == "" || password == "" || domain == "" {
		t.Skip("set API_USERNAME, API_PASSWORD, API_DOMAIN to run")
	}
	return
}

func captureLog(t *testing.T) (*bytes.Buffer, func()) {
	t.Helper()
	buf := &bytes.Buffer{}
	orig := log.StandardLogger().Out
	log.SetOutput(buf)
	return buf, func() { log.SetOutput(orig) }
}

func TestGetZones(t *testing.T) {
	username, password, domain := credentials(t)
	buf, restore := captureLog(t)
	defer restore()

	GetZones(username, password, domain)

	if !strings.Contains(buf.String(), "success") {
		t.Errorf("expected log to contain 'success', got: %s", buf.String())
	}
}

func TestAddTxtRr(t *testing.T) {
	username, password, domain := credentials(t)
	buf, restore := captureLog(t)
	defer restore()

	AddTxtRr(username, password, domain, "_regru_api_test", "regru-api-go-test")

	if !strings.Contains(buf.String(), "success") {
		t.Errorf("expected log to contain 'success', got: %s", buf.String())
	}
}

func TestRmTxtRr(t *testing.T) {
	username, password, domain := credentials(t)
	buf, restore := captureLog(t)
	defer restore()

	RmTxtRr(username, password, domain, "_regru_api_test", "TXT", "")

	if !strings.Contains(buf.String(), "success") {
		t.Errorf("expected log to contain 'success', got: %s", buf.String())
	}
}
