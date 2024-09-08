package main

import "net/http"

func NewHttpHandler(
	usePointAsDiscountHandler UsePointsAsDiscountHandler,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	return mux
}
