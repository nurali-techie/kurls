package commands

import (
	"fmt"
	"os"

	"github.com/nurali-techie/kurls/repo"
)

var DefaultHelp = &Help{}

type Help struct {
	Message string
}

func (c *Help) Run(_ repo.KurlRepo) {
	if c.Message == "" {
		fmt.Println("help")
	} else {
		fmt.Println(c.Message)
	}
	os.Exit(0)
}
