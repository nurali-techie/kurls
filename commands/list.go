package commands

import (
	"log"

	"github.com/nurali-techie/kurls/repo"
)

type List struct {
	Filter string
}

func (c *List) Run(kurlRepo repo.KurlRepo, memRepo repo.MemoryRepo) {
	keys, err := kurlRepo.List(c.Filter)
	if err != nil {
		log.Fatal(err)
	}
	memCmd := &Memory{
		Keys: keys,
	}
	memCmd.Run(kurlRepo, memRepo)
}
