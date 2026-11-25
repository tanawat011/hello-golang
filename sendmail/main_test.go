package main

import (
	"bytes"
	"strings"
	"testing"
)

func env(vars map[string]string) func(string) string {
	return func(k string) string { return vars[k] }
}

func TestLoadConfigNamesWhatIsMissing(t *testing.T) {
	_, err := loadConfig(env(map[string]string{"MAIL_FROM": "a@example.com"}), true)
	if err == nil {
		t.Fatal("expected an error")
	}
	for _, name := range []string{"SMTP_HOST", "SMTP_USERNAME", "SMTP_PASSWORD", "MAIL_TO"} {
		if !strings.Contains(err.Error(), name) {
			t.Errorf("error %q does not name %s", err, name)
		}
	}
	if strings.Contains(err.Error(), "MAIL_FROM") {
		t.Errorf("error %q names MAIL_FROM, which was set", err)
	}
}

func TestLoadConfigDryRunNeedsNoSMTP(t *testing.T) {
	c, err := loadConfig(env(map[string]string{"MAIL_FROM": "a@example.com", "MAIL_TO": "b@example.com"}), false)
	if err != nil {
		t.Fatal(err)
	}
	if c.Port != 587 {
		t.Errorf("default port = %d, want 587", c.Port)
	}
}

func TestLoadConfigRejectsBadPort(t *testing.T) {
	for _, p := range []string{"abc", "0", "70000"} {
		_, err := loadConfig(env(map[string]string{
			"MAIL_FROM": "a@example.com", "MAIL_TO": "b@example.com", "SMTP_PORT": p,
		}), false)
		if err == nil {
			t.Errorf("SMTP_PORT=%s: expected an error", p)
		}
	}
}

func TestBuildMessage(t *testing.T) {
	c := config{From: "a@example.com", To: "b@example.com", CcAddr: "c@example.com", CcName: "TA"}
	m, err := buildMessage(c, "Hello!", "Hello <b>there</b>")
	if err != nil {
		t.Fatal(err)
	}
	var buf bytes.Buffer
	if _, err := m.WriteTo(&buf); err != nil {
		t.Fatal(err)
	}
	out := buf.String()
	for _, want := range []string{
		"From: <a@example.com>",
		"To: <b@example.com>",
		`Cc: "TA" <c@example.com>`,
		"Subject: Hello!",
		"Content-Type: text/html",
		"Hello <b>there</b>",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("message lacks %q:\n%s", want, out)
		}
	}
}

func TestBuildMessageRejectsBadAddress(t *testing.T) {
	if _, err := buildMessage(config{From: "not an address", To: "b@example.com"}, "s", "b"); err == nil {
		t.Error("expected an error for a bad From address")
	}
}
