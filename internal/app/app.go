package app

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type App struct {
	AddUserHandler     http.HandlerFunc
	GetUsersHandler    http.HandlerFunc
	GetUserByIdHandler http.HandlerFunc
}

func (a *App) Handler() http.HandlerFunc {
	router := httprouter.New()
	router.HandlerFunc(http.MethodPost, "/users", a.AddUserHandler)
	router.HandlerFunc(http.MethodGet, "/users", a.GetUsersHandler)
	h := http.HandlerFunc(router.ServeHTTP)
	return h
}

func New() App {
	return App{
		AddUserHandler:  AddUserHandler(),
		GetUsersHandler: GetAllUsersHandler(),
	}
}
