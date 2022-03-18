package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Middelware 1")
			if flag {
				hf(w, r)
			} else {
				return
			}
		}
	}
}
