package main

import (
	"fmt"
	"github.com/jhillyerd/go.enmime"
	"net/mail"
	"os"
)

var Debug bool

func main() {
	m, _ := mail.ReadMessage(os.Stdin)

	body, err := enmime.ParseMIMEBody(m)
	if err != nil {
		fmt.Println("Error parsing MIME body.")
	} else {
		fmt.Println(m.Header.Get("Subject"))
		fmt.Printf("%s", body.Text)
	}
}
