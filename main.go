package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/denizekindaglar/rssfeed/internal/config"
	"github.com/denizekindaglar/rssfeed/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func main() {
	cfgFile, err := config.Read()
	if err != nil {
		log.Fatalf("Could not read config file: %v", err)
	}

	dburl := cfgFile.DBUrl
	db, err := sql.Open("postgres", dburl)
	if err != nil {
		log.Fatalf("Could not connect to db: %v", err)
	}
	dbQueries := database.New(db)

	st := &state{
		db:     dbQueries,
		config: &cfgFile,
	}
	commands := commands{
		dict: make(map[string]func(*state, command) error),
	}
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	commands.register("browse", handlerBrowse)
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	cmd := command{name: os.Args[1], arg: os.Args[2:]}
	err = commands.run(st, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
