package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/wevyrton/exercicio/internal/pessoas"
)

func main() {

	r := chi.NewRouter()
	r.Route("/usuarios", func(r chi.Router) {
		r.Get("/", pessoas.Usuarios)
		r.Get("/nome/{nome}", pessoas.Buscanome)
		r.Get("/{id}", pessoas.Buscaid)
	})

	r.Get("/", pessoas.Inicial)

	println("servidor rodando na porta 8080")

	http.ListenAndServe(":8080", r)
}
