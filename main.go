package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/MudassirDev/blog-aggregator/internal/config"
	"github.com/MudassirDev/blog-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	dbConn, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error forming an connection: %v", err)
	}

	queries := database.New(dbConn)

	programState := &state{
		cfg: &cfg,
		db:  queries,
	}

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerGetUsers)
	cmds.register("agg", handleAgg)
	cmds.register("addfeed", handleAddFeed)
	cmds.register("feeds", handleGetFeeds)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}
}
