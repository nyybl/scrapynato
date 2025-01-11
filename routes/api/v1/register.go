package v1

import "github.com/go-chi/chi/v5"

func RegisterV1(r *chi.Mux) {
	r.Mount("/api/v1/manga", MangaRouter())
}