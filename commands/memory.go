package commands

import (
	"fmt"

	"github.com/nurali-techie/kurls/domains"
	"github.com/nurali-techie/kurls/repo"
)

type Memory struct {
	Keys []string
}

func (c *Memory) Run(kurlRepo repo.KurlRepo, memRepo repo.MemoryRepo) {
	if len(c.Keys) == 0 {
		return
	}
	memory := domains.Memory{
		Keys: c.Keys,
	}
	err := memRepo.Add(memory)
	if err != nil {
		for _, key := range c.Keys {
			fmt.Println(key)
		}
	} else {
		for i, key := range c.Keys {
			fmt.Printf("%d) %s\n", i+1, key)
		}
	}
}
