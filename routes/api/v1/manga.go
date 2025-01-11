package v1

import (
	"net/http"

	"github.com/nyybl/scrapynato/lib"
	"github.com/go-chi/chi/v5"
)

func MangaRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", lib.HandleHttp(GetIndex))
	return r
}

func GetIndex(w http.ResponseWriter, r *http.Request) lib.ResponseSchema {
	return lib.NewResponse(http.StatusOK, "/")
}

func GetMangaByID(w http.ResponseWriter, r *http.Request) lib.ResponseSchema {
	return lib.NewResponse(1, "")
}