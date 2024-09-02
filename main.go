package main

import (
	"fmt"

	"log"
	"net/http"
	"real-time-rofu/back/database"
	"real-time-rofu/back/server"
	"real-time-rofu/back/websockets"

	_ "github.com/mattn/go-sqlite3"
)

type settingStruct struct {
	port string
}

var settings = settingStruct{port: ":8080"}

func main() {

	loading()
	database.DataBaseCreateIfNotExist()

	myhttp := http.NewServeMux()
	websockets.SetupWebSockets(myhttp)
	server.LaunchWebPAge(myhttp)
	go func() {
		if err := http.ListenAndServe(settings.port, myhttp); err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("L=> Server is listening on", settings.port)

	select {}
}

func loading() {
	fmt.Println("====================")
	database.LoadingDatabase()
	fmt.Println("====================")
}
