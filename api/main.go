package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe("10.16.57.176:9696", r)

	//listen->registerhandlers->handlers
	//handler->validation(1.request 2..user)->business logic->response
}
