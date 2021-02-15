package main

import (
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nurali-techie/kurls/repo"
	"github.com/nurali-techie/kurls/utils"
)

func main() {
	if len(os.Args) < 2 {
		helpAndExit(0)
	}

	// setup db
	db := utils.GetDatabase()
	if db != nil {
		defer db.Close()
	}
	err := utils.InitDB(db)
	if err != nil {
		panic(err)
	}

	// setup repo
	kurlRepo := repo.NewKurlRepo(db)

	// get cmd
	cmd := utils.GetCommand(os.Args[1:])
	if cmd == nil {
		helpAndExit(0)
	}

	// run cmd
	if cmd != nil {
		cmd.Run(kurlRepo)
	}
}

func helpAndExit(errCode int) {
	fmt.Println("help")
	os.Exit(errCode)
}
