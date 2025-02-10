package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/modestprophet/gator/internal/config"
	"github.com/modestprophet/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Failed to read config: %v", err)
		return
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Printf("Failed to connect to database: %v", err)
		return
	}
	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: cfg,
	}

	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)

	if len(os.Args) < 2 {
		fmt.Println("Error: command required")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := []string{}
	if len(os.Args) > 2 {
		cmdArgs = os.Args[2:]
	}

	err = cmds.run(s, command{
		name: cmdName,
		args: cmdArgs,
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// updatedCfg, err := config.Read()
	// if err != nil {
	// 	fmt.Printf("Failed to re-read config: %v", err)
	// }

	// fmt.Printf("Final configuration:\n%+v\n", updatedCfg)
}
