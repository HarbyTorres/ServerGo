package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(method string, path string) (http.HandlerFunc, bool, bool) {
	//asignacion de el valor en el mapa de reglas a las variables handler y exist
	_, exist := r.rules[path]
	handler, existMethod := r.rules[path][method]
	return handler, existMethod, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	/*el FindHandler compara el request con el mapa de reglas para saber si existe o no.
	los valores son asignados a las variables 'handler' y 'exist'*/
	handler, existMethod, exist := r.FindHandler(request.URL.Path, request.Method)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !existMethod {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}
