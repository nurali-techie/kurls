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

func (c *Help) Run(_ repo.KurlRepo, _ repo.MemoryRepo) {
	if c.Message == "" {
		fmt.Println(helpText())
	} else {
		fmt.Println(c.Message)
	}
	os.Exit(0)
}

func helpText() string {
	help := `kurls is a curl commands management tool
	
Usage: kurls CMD [args...]

  ls|list : list keys
  ad|add  : add key with curl from clipboard
  gt|get  : get curl for given key and copy curl to clipboard

Example:
  ls|list : kurls ls  / kurl list / kurl ls MY_POST / kurl ls MY / kurl ls POST
  ad|add  : kurls ad MY_POST / kurls add MY_POST
  gt|get  : kurls gt MY_POST / kurls get MY_POST
`
	return help
}
