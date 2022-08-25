package main

import (
	"fmt"
	"net/http"

	"github.com/ali-ghn/go-web-learning/04-mongodb/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.UserController.NewUserController(getSession())
	r.GET("/", index)
	r.POST("/user", uc.CreateUser)
	r.GET("/user", uc.GetUser)
	http.ListenAndServe(":8081", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintln(w, "welcome\n")
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		panic(err)
	}
	return s
}
