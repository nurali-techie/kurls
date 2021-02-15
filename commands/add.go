package commands

import (
	"fmt"
	"log"

	"github.com/nurali-techie/kurls/domains"
	"github.com/nurali-techie/kurls/repo"
)

type Add struct {
	Key  string
	Curl string
}

func NewAdd(key string, curl string) *Add {
	return &Add{
		Key:  key,
		Curl: curl,
	}
}

func (c *Add) Run(kurlRepo repo.KurlRepo) {
	kurl := domains.Kurl{
		Key:  c.Key,
		Curl: c.Curl,
	}
	err := kurlRepo.Add(kurl)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("curl added, key=", c.Key)
	}
}

func printParts(parts []string) {
	for ind, part := range parts {
		fmt.Println(ind, part)
	}
}
