package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Don't forget the driver!
	"github.com/xmarlem/simplebank/api"
	db "github.com/xmarlem/simplebank/db/sqlc"
	"github.com/xmarlem/simplebank/util"
)

func main() {

	conf, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config")
	}

	fmt.Println(conf.DBSource)

	conn, err := sql.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db")
	}
	// prima alloco uno store...
	store := db.NewStore(conn)
	// poi lo passo ad una nuova instanza del server/api ...
	server := api.NewServer(store)

	// infine faccio partire il server sull'indirizzo passato via config
	err = server.Start(conf.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
