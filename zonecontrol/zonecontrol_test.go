package zonecontrol

import (
	"os"
	"testing"
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

func TestGetZones(t *testing.T) {
	username, password, domain := credentials(t)

	rec, err := GetZones(username, password, domain)
	if err != nil {
		t.Fatalf("GetZones returned error: %v", err)
	}
	if rec.Result != "success" {
		t.Errorf("expected result 'success', got: %s (error: %s)", rec.Result, rec.ErrorText)
	}
}

func TestTxtRrLifecycle(t *testing.T) {
	username, password, domain := credentials(t)

	t.Cleanup(func() {
		RmTxtRr(username, password, domain, "_regru_api_test", "TXT", "")
	})

	rec, err := AddTxtRr(username, password, domain, "_regru_api_test", "regru-api-go-test")
	if err != nil {
		t.Fatalf("AddTxtRr returned error: %v", err)
	}
	if rec.Result != "success" {
		t.Errorf("expected result 'success' after Add, got: %s (error: %s)", rec.Result, rec.ErrorText)
	}

	rec, err = RmTxtRr(username, password, domain, "_regru_api_test", "TXT", "")
	if err != nil {
		t.Fatalf("RmTxtRr returned error: %v", err)
	}
	if rec.Result != "success" {
		t.Errorf("expected result 'success' after Rm, got: %s (error: %s)", rec.Result, rec.ErrorText)
	}
}

func TestTxtRrLifecycleWithContent(t *testing.T) {
	username, password, domain := credentials(t)

	t.Cleanup(func() {
		RmTxtRr(username, password, domain, "_regru_api_test_content", "TXT", "regru-api-go-test")
	})

	rec, err := AddTxtRr(username, password, domain, "_regru_api_test_content", "regru-api-go-test")
	if err != nil {
		t.Fatalf("AddTxtRr returned error: %v", err)
	}
	if rec.Result != "success" {
		t.Errorf("expected result 'success' after Add, got: %s (error: %s)", rec.Result, rec.ErrorText)
	}

	rec, err = RmTxtRr(username, password, domain, "_regru_api_test_content", "TXT", "regru-api-go-test")
	if err != nil {
		t.Fatalf("RmTxtRr returned error: %v", err)
	}
	if rec.Result != "success" {
		t.Errorf("expected result 'success' after Rm with content, got: %s (error: %s)", rec.Result, rec.ErrorText)
	}
}

func TestAddARrLifecycle(t *testing.T) {
	username, password, domain := credentials(t)

	t.Cleanup(func() {
		RmTxtRr(username, password, domain, "test-a-record", "A", "")
	})

	rec, err := AddARr(username, password, domain, "test-a-record", "1.2.3.4")
	if err != nil {
		t.Fatalf("AddARr returned error: %v", err)
	}
	if rec.Result != "success" {
		t.Errorf("expected result 'success' after AddARr, got: %s (error: %s)", rec.Result, rec.ErrorText)
	}

	rec, err = RmTxtRr(username, password, domain, "test-a-record", "A", "")
	if err != nil {
		t.Fatalf("RmTxtRr returned error: %v", err)
	}
	if rec.Result != "success" {
		t.Errorf("expected result 'success' after Rm A record, got: %s (error: %s)", rec.Result, rec.ErrorText)
	}
}
