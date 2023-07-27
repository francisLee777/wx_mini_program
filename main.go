package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/handler"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	if err := db.Init(); err != nil {
		panic(fmt.Sprintf("mysql init failed with %+v", err))
	}
	// 他妈的原生http好麻烦
	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/api/count", handler.CounterHandler)
	http.HandleFunc("/api/listFoodMenu", handler.ListFoodMenu)
	http.HandleFunc("/api/orderFood", handler.OrderFood)
	http.HandleFunc("/api/user/getUserInfo", handler.GetUserInfo)
	http.HandleFunc("/api/user/saveNickName", handler.SaveNickName)
	http.HandleFunc("/api/user/saveIconURL", handler.SaveIconURL)
	log.Fatal(http.ListenAndServe(":80", nil))
}
