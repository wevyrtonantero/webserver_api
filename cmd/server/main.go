package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/wevyrton/exercicio/internal/pessoas"
)

func main() {

	r := chi.NewRouter()
	r.Get("/", pessoas.Inicial)
	r.Get("/usuarios", pessoas.Usuarios)
	r.Get("/usuarios-id/{id}", pessoas.Buscaid)
	r.Get("/usuarios-nome/{nome}", pessoas.Buscanome)

	println("servidor rodando na porta 8080")

	http.ListenAndServe(":8080", r)
}
