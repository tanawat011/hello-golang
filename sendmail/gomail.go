package main

import (
	"fmt"
	"io"

	"gopkg.in/gomail.v2"
)

func main() {
	// Example_noSMTP()
	Example()
}

type sendEmail struct {
	smtp            string
	smtp_port       int
	from_email      string
	from_email_pass string
	to_email        string
	header_email    string
	header_name     string
	content         string
}

func Example() {
	setting := sendEmail{
		smtp:            "smtp.example.com",
		smtp_port:       587,
		from_email:      "XXXXXXXXXX@gmail.com",
		from_email_pass: "XXXXXXXXXX",
		to_email:        "XXXXXXXXXX@gmail.com",
		header_email:    "example@example.com",
		header_name:     "TA",
		content:         "Hello <b>mr.Tanawat Pinthongpan</b>",
	}
	m := gomail.NewMessage()
	m.SetHeader("From", setting.from_email)
	m.SetHeader("To", setting.to_email)
	m.SetAddressHeader("Cc", setting.header_email, setting.header_name)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", setting.content)
	// m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer(setting.smtp, setting.smtp_port, setting.from_email, setting.from_email_pass)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

// Send an email using an API or postfix.
func Example_noSMTP() {
	m := gomail.NewMessage()
	m.SetHeader("From", "talovenit1987@gmail.com")
	m.SetHeader("To", "takissnit1987@gmail.com")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/plain", "Hello!")

	s := gomail.SendFunc(func(from string, to []string, msg io.WriterTo) error {
		// Implements you email-sending function, for example by calling
		// an API, or running postfix, etc.
		fmt.Println("From:", from)
		fmt.Println("To:", to)
		return nil
	})

	if err := gomail.Send(s, m); err != nil {
		panic(err)
	}
	// Output:
	// From: from@example.com
	// To: [to@example.com]
}

// var m *gomail.Message

// func ExampleSetCopyFunc() {
// 	m.Attach("foo.txt", gomail.SetCopyFunc(func(w io.Writer) error {
// 		_, err := w.Write([]byte("Content of foo.txt"))
// 		return err
// 	}))
// }

// func ExampleSetHeader() {
// 	h := map[string][]string{"Content-ID": {"<foo@bar.mail>"}}
// 	m.Attach("foo.jpg", gomail.SetHeader(h))
// }

// func ExampleRename() {
// 	m.Attach("/tmp/0000146.jpg", gomail.Rename("picture.jpg"))
// }

// func ExampleMessage_AddAlternative() {
// 	m.SetBody("text/plain", "Hello!")
// 	m.AddAlternative("text/html", "<p>Hello!</p>")
// }

// func ExampleMessage_AddAlternativeWriter() {
// 	t := template.Must(template.New("example").Parse("Hello {{.}}!"))
// 	m.AddAlternativeWriter("text/plain", func(w io.Writer) error {
// 		return t.Execute(w, "Bob")
// 	})
// }

// func ExampleMessage_Attach() {
// 	m.Attach("/tmp/image.jpg")
// }

// func ExampleMessage_Embed() {
// 	m.Embed("/tmp/image.jpg")
// 	m.SetBody("text/html", `<img src="cid:image.jpg" alt="My image" />`)
// }

// func ExampleMessage_FormatAddress() {
// 	m.SetHeader("To", m.FormatAddress("bob@example.com", "Bob"), m.FormatAddress("cora@example.com", "Cora"))
// }

// func ExampleMessage_FormatDate() {
// 	m.SetHeaders(map[string][]string{
// 		"X-Date": {m.FormatDate(time.Now())},
// 	})
// }

// func ExampleMessage_SetAddressHeader() {
// 	m.SetAddressHeader("To", "bob@example.com", "Bob")
// }

// func ExampleMessage_SetBody() {
// 	m.SetBody("text/plain", "Hello!")
// }

// func ExampleMessage_SetDateHeader() {
// 	m.SetDateHeader("X-Date", time.Now())
// }

// func ExampleMessage_SetHeader() {
// 	m.SetHeader("Subject", "Hello!")
// }

// func ExampleMessage_SetHeaders() {
// 	m.SetHeaders(map[string][]string{
// 		"From":    {m.FormatAddress("alex@example.com", "Alex")},
// 		"To":      {"bob@example.com", "cora@example.com"},
// 		"Subject": {"Hello"},
// 	})
// }

// func ExampleSetCharset() {
// 	m = gomail.NewMessage(gomail.SetCharset("ISO-8859-1"))
// }

// func ExampleSetEncoding() {
// 	m = gomail.NewMessage(gomail.SetEncoding(gomail.Base64))
// }

// func ExampleSetPartEncoding() {
// 	m.SetBody("text/plain", "Hello!", gomail.SetPartEncoding(gomail.Unencoded))
// }
