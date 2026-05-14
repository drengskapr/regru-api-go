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

// captureLog redirects logrus global output to a buffer. These tests must not
// be run in parallel (no t.Parallel) because SetOutput mutates shared state.
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

	if !strings.Contains(buf.String(), "Result:success") {
		t.Errorf("expected log to contain 'Result:success', got: %s", buf.String())
	}
}

func TestTxtRrLifecycle(t *testing.T) {
	username, password, domain := credentials(t)
	buf, restore := captureLog(t)
	defer restore()

	// Register cleanup to guarantee removal even if test fails
	t.Cleanup(func() {
		RmTxtRr(username, password, domain, "_regru_api_test", "TXT", "")
	})

	// Add TXT record
	AddTxtRr(username, password, domain, "_regru_api_test", "regru-api-go-test")

	if !strings.Contains(buf.String(), "Result:success") {
		t.Errorf("expected log to contain 'Result:success' after Add, got: %s", buf.String())
	}

	// Reset buffer for next operation
	buf.Reset()

	// Remove TXT record with empty content
	RmTxtRr(username, password, domain, "_regru_api_test", "TXT", "")

	if !strings.Contains(buf.String(), "Result:success") {
		t.Errorf("expected log to contain 'Result:success' after Rm, got: %s", buf.String())
	}
}

func TestTxtRrLifecycleWithContent(t *testing.T) {
	username, password, domain := credentials(t)
	buf, restore := captureLog(t)
	defer restore()

	// Register cleanup to guarantee removal even if test fails
	t.Cleanup(func() {
		RmTxtRr(username, password, domain, "_regru_api_test_content", "TXT", "regru-api-go-test")
	})

	// Add TXT record
	AddTxtRr(username, password, domain, "_regru_api_test_content", "regru-api-go-test")

	if !strings.Contains(buf.String(), "Result:success") {
		t.Errorf("expected log to contain 'Result:success' after Add, got: %s", buf.String())
	}

	// Reset buffer for next operation
	buf.Reset()

	// Remove TXT record with non-empty content
	RmTxtRr(username, password, domain, "_regru_api_test_content", "TXT", "regru-api-go-test")

	if !strings.Contains(buf.String(), "Result:success") {
		t.Errorf("expected log to contain 'Result:success' after Rm with content, got: %s", buf.String())
	}
}
