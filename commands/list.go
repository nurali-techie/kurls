package commands

import (
	"fmt"
	"log"

	"github.com/nurali-techie/kurls/repo"
)

type List struct {
	Filter string
}

func (c *List) Run(kurlRepo repo.KurlRepo) {
	keys, err := kurlRepo.List(c.Filter)
	if err != nil {
		log.Fatal(err)
	}
	for _, key := range keys {
		fmt.Println(key)
	}
}
