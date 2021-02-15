package commands

import (
	"fmt"
	"log"

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
	fmt.Println(kurl.Curl)
}
