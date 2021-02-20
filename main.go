package main

import (
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nurali-techie/kurls/commands"
	"github.com/nurali-techie/kurls/repo"
	"github.com/nurali-techie/kurls/utils"
)

func main() {
	var err error
	if len(os.Args) < 2 {
		commands.DefaultHelp.Run(nil, nil)
	}

	// setup db
	db := utils.GetDatabase()
	if db != nil {
		defer db.Close()
	}
	err = utils.InitDB(db)
	if err != nil {
		panic(err)
	}

	// setup repo
	kurlRepo := repo.NewKurlRepo(db)
	memRepo := repo.NewMemoryRepo(db)

	// get cmd
	cmd := utils.GetCommand(os.Args[1:])
	if cmd != nil {
		cmd.Run(kurlRepo, memRepo)
	}
}
