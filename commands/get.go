package commands

import (
	"fmt"
	"log"
	"strconv"

	"github.com/atotto/clipboard"
	"github.com/nurali-techie/kurls/repo"
)

type Get struct {
	Key string
}

func (c *Get) Run(kurlRepo repo.KurlRepo, memRepo repo.MemoryRepo) {
	seq, err := strconv.Atoi(c.Key)
	if err == nil {
		key, err := memRepo.Get(seq)
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		c.Key = key
	}

	kurl, err := kurlRepo.Get(c.Key)
	if err != nil {
		log.Fatal(err)
	}

	err = clipboard.WriteAll(kurl.Curl)
	if err != nil {
		fmt.Printf("error: paste curl to clipboard, %v\n", err)
		return
	}
	fmt.Printf("success: '%s' curl copied to clipboard\n", c.Key)
}
