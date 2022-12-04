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

		// Route /cep

		r.Get("/cep", cep.ListarCep)
		r.Get("/cep/{id}", cep.BuscarCep)
		r.Post("/cep", cep.CriarCep)
		r.Put("/cep", cep.AtualizarCep)
		r.Delete("/cep/{bairro}", cep.DeletarCep)
	})

	r.Get("/", pessoas.Inicial)

	println("servidor rodando na porta 8080")

	http.ListenAndServe(":8080", r)

}
