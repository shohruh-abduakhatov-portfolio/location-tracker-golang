package cmd

import (
	"fmt"
	"log"
	"net/http"

	cfg "gitlab.com/logitab/back-end-team/location-tracker-go/config"
	internal "gitlab.com/logitab/back-end-team/location-tracker-go/internal"
	_ "gitlab.com/logitab/back-end-team/location-tracker-go/mq"
	_ "gitlab.com/logitab/back-end-team/location-tracker-go/store"
)

// Execute runs web app and websocket
func Execute() {
	hub := internal.NewHub()
	go hub.Run()
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		log.Println(">> New Connection from: ", GetIP(r))
		internal.ServeWs(hub, w, r)
	})
	add := fmt.Sprintf("%s:%s", cfg.Cfg.Server.IP, cfg.Cfg.Server.Port)
	fmt.Println("Listening at: ", add)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
