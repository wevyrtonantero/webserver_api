package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/wevyrton/exercicio/internal/cep"
	"github.com/wevyrton/exercicio/internal/pessoas"
)

func main() {

	r := chi.NewRouter()
	r.Route("/usuarios", func(r chi.Router) {
		r.Get("/", pessoas.Usuarios)
		r.Get("/nome/{nome}", pessoas.Buscanome)
		r.Get("/{id}", pessoas.Buscaid)
		r.Post("/", pessoas.CriarUsuario)
		r.Put("/", pessoas.AtualizarUsuario)
		r.Delete("/{id}", pessoas.DeletarUsuario)
	})

	r.Route("/cep", func(r chi.Router) {
		r.Get("/", cep.ListarCep)
		r.Get("/{id}", cep.BuscarCep)
		r.Post("/", cep.CriarCep)
		r.Put("/", cep.AtualizarCep)
		r.Delete("/{id}", cep.DeletarCep)
	})

	r.Get("/", pessoas.Inicial)

	println("servidor rodando na porta 8080")

	http.ListenAndServe(":8080", r)

}
