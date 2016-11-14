package main

import (
	"log"
	"net/http"

	"github.com/parkn-co/parkn-server/src/config"
	"github.com/parkn-co/parkn-server/src/router"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := router.InitRouter()

	log.Println("Listening at 0.0.0.0:", config.App.Server.Port)
	log.Println(http.ListenAndServe("0.0.0.0:"+config.App.Server.Port, r))
}
