package commands

import (
	"fmt"
	"log"

	"github.com/atotto/clipboard"
	"github.com/nurali-techie/kurls/repo"
)

type Get struct {
	Key string
}

func (c *Get) Run(kurlRepo repo.KurlRepo) {
	kurl, err := kurlRepo.Get(c.Key)
	if err != nil {
		log.Fatal(err)
	}
	err = clipboard.WriteAll(kurl.Curl)
	if err != nil {
		fmt.Printf("error: paste curl to clipboard, %v\n", err)
	}
}
