package properties

import "testing"

func TestFromFile(t *testing.T) {
	fn := "Reporter.properties.example"
	p := NewFromFile(fn)

	if p.UserID != "user@example.com" {
		t.Errorf("expected %q got %q", "user@example.com", p.UserID)
	}

	if p.Password != "P@$$w0rd" {
		t.Errorf("expected %q got %q", "P@$$w0rd", p.Password)
	}

	if p.Mode != "Robot.xml" {
		t.Errorf("expected %q got %q", "Robot.xml", p.Mode)
	}

	if p.Account != "123456" {
		t.Errorf("expected %q got %q", "123456", p.Account)
	}

	surl := "https://reportingitc-reporter.apple.com/reportservice/sales/v1"
	if p.SalesURL != surl {
		t.Errorf("expected %q got %q", surl, p.FinanceURL)
	}

	furl := "https://reportingitc-reporter.apple.com/reportservice/finance/v1"
	if p.FinanceURL != furl {
		t.Errorf("expected %q got %q", furl, p.FinanceURL)
	}
}
