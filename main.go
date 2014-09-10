package main

import (
	"os"
	"fmt"
	"net/mail"
	"github.com/jhillyerd/go.enmime"
)

var Debug bool

func main() {
  m, _ := mail.ReadMessage(os.Stdin)

  body, err := enmime.ParseMIMEBody(m)
  if err != nil {
    fmt.Println("Error parsing MIME body.")
  }
  fmt.Println(m.Header.Get("Subject"))
  fmt.Printf("%v", body.Text)
}
