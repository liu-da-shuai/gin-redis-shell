package router

import (
	"fmt"
	"gin-redis-shell/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	r := mux.NewRouter()
	r.HandleFunc("/quote", handlers.GetDailyQuote).Methods("GET")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Printf("HTTP server start failed,err :%v", err)
		return
	}

}
