package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nyybl/scrapynato/lib"
)

func MangaRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/{id}", lib.HandleHttp(GetMangaByID))
	r.Get("/search", lib.HandleHttp(GetMangaBySearch))
	return r
}

func GetMangaByID(w http.ResponseWriter, r *http.Request) lib.ResponseSchema {
	id := chi.URLParam(r, "id")
	if id == "" {
		return lib.NewErrorResponse(http.StatusBadRequest, errors.New("missing required parameter: id"))
	}
	d, err := lib.ScrapeManga(id)
	if err != nil {
		if err == lib.ErrNotFound {
			return lib.NewErrorResponse(http.StatusNotFound, fmt.Errorf("could not manga with ID: ", id))
		} else {
			return lib.NewErrorResponse(http.StatusInternalServerError, err)
		}
	}
	return lib.NewResponse(http.StatusOK, d)
}

func GetMangaBySearch(w http.ResponseWriter, r *http.Request) lib.ResponseSchema {
	query := r.URL.Query().Get("query")
	if query == "" {
		return lib.NewErrorResponse(http.StatusBadRequest, errors.New("missing required query"))
	}
	d, err := lib.SearchManga(query)

	if err != nil {
		return lib.NewErrorResponse(http.StatusInternalServerError, err)
	}
	return lib.NewResponse(http.StatusOK, d)
}