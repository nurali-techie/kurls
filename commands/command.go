package commands

import "github.com/nurali-techie/kurls/repo"

type Command interface {
	Run(kurlRepo repo.KurlRepo, memRepo repo.MemoryRepo)
}
