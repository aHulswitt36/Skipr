package api

import (
	"net/http"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/Player", PlayerHandler)
	mux.HandleFunc("/Lineup", LineupHandler)

	return mux
}
