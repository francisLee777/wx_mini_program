package main

import (
	"log"
	"net/http"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/handler"
)

func main() {
	if err := db.Init(); err != nil {
		//panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	// 他妈的原生http好麻烦
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/api/count", handler.CounterHandler)
	http.HandleFunc("/api/listFoodMenu", handler.ListFoodMenu)

	log.Fatal(http.ListenAndServe(":80", nil))
}
