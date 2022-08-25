package main

import (
	"net/http"

	"github.com/ali-ghn/go-web-learning/04-mongodb/01/controllers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	http.ListenAndServe(":8081", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}
