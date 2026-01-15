package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/kk-ami/GoApi/internal/middleware"
)

func Handler(r *chi.Mux) {
	//Global middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account",)
}