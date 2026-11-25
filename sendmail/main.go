// sendmail composes an HTML e-mail and either prints it (the default) or
// sends it through an SMTP server. Every setting comes from the environment,
// so no address or password lives in the code. See .env.example.
package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/wneessen/go-mail"
)

type config struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
	To       string
	CcAddr   string // optional
	CcName   string // optional
}

// loadConfig reads the settings through getenv (os.Getenv in main), so a
// test can hand in its own values.
func loadConfig(getenv func(string) string, needSMTP bool) (config, error) {
	c := config{
		Host:     getenv("SMTP_HOST"),
		Username: getenv("SMTP_USERNAME"),
		Password: getenv("SMTP_PASSWORD"),
		From:     getenv("MAIL_FROM"),
		To:       getenv("MAIL_TO"),
		CcAddr:   getenv("MAIL_CC"),
		CcName:   getenv("MAIL_CC_NAME"),
		Port:     587,
	}

	var missing []string
	need := map[string]string{"MAIL_FROM": c.From, "MAIL_TO": c.To}
	if needSMTP {
		need["SMTP_HOST"] = c.Host
		need["SMTP_USERNAME"] = c.Username
		need["SMTP_PASSWORD"] = c.Password
	}
	for _, name := range []string{"SMTP_HOST", "SMTP_USERNAME", "SMTP_PASSWORD", "MAIL_FROM", "MAIL_TO"} {
		if v, ok := need[name]; ok && v == "" {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		return config{}, fmt.Errorf("missing environment variables: %s", strings.Join(missing, ", "))
	}

	if p := getenv("SMTP_PORT"); p != "" {
		port, err := strconv.Atoi(p)
		if err != nil || port < 1 || port > 65535 {
			return config{}, fmt.Errorf("SMTP_PORT %q is not a port number", p)
		}
		c.Port = port
	}
	return c, nil
}

func buildMessage(c config, subject, html string) (*mail.Msg, error) {
	m := mail.NewMsg()
	if err := m.From(c.From); err != nil {
		return nil, fmt.Errorf("from: %w", err)
	}
	if err := m.To(c.To); err != nil {
		return nil, fmt.Errorf("to: %w", err)
	}
	if c.CcAddr != "" {
		if err := m.AddCcFormat(c.CcName, c.CcAddr); err != nil {
			return nil, fmt.Errorf("cc: %w", err)
		}
	}
	m.Subject(subject)
	m.SetBodyString(mail.TypeTextHTML, html)
	return m, nil
}

func send(c config, m *mail.Msg) error {
	client, err := mail.NewClient(c.Host,
		mail.WithPort(c.Port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(c.Username),
		mail.WithPassword(c.Password),
		mail.WithTLSPolicy(mail.TLSMandatory), // never send the password in the clear
	)
	if err != nil {
		return err
	}
	return client.DialAndSend(m)
}

func main() {
	really := flag.Bool("send", false, "send the message (without it the message is only printed)")
	flag.Parse()

	c, err := loadConfig(os.Getenv, *really)
	if err != nil {
		log.Fatal(err)
	}
	m, err := buildMessage(c, "Hello!", "Hello <b>from hello-golang</b>")
	if err != nil {
		log.Fatal(err)
	}

	if !*really {
		if _, err := m.WriteTo(os.Stdout); err != nil {
			log.Fatal(err)
		}
		fmt.Println("\n(dry run: add -send to send it)")
		return
	}
	if err := send(c, m); err != nil {
		var se *mail.SendError
		if errors.As(err, &se) {
			log.Fatalf("send failed: %v", se)
		}
		log.Fatal(err)
	}
	log.Printf("sent to %s", c.To)
}
