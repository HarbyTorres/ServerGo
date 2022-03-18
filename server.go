package main

import (
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

func (s *Server) Handle(path string, method string, handler http.HandlerFunc) {
	//Asociacion del handler con la ruta, es decir, el mapa con la llave path asignado al handler
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
	//asi el servidor es capaz de agregar la ruta especifica a un handler especifico
}

//Toma una lista de middlewares y un handler, ejecutará todos los middleware y si todos pasan ejecutará el handler
func (s *Server) AddMidleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func (s *Server) Listen() error {
	//el router va a ser el encargado de tomar las urls y procesarlas como se debe, crea los entry points
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	}
	return nil
}
