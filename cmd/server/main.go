package main

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wevyrton/exercicio/internal/cep"
	"github.com/wevyrton/exercicio/internal/pessoas"
)

func main() {

	db, err := sql.Open("mysql", "root:wedeju180587@tcp(localhost:3306)/safisa")
	if err != nil {
		panic(err)
	}


	p, err := pessoas.NovaPessoa(db)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Route("/usuarios", func(r chi.Router) {
		r.Get("/", p.Usuarios)
		r.Get("/{id}", p.BuscarPorID)
		r.Post("/", p.AtualizarUsuario)
		r.Put("/", p.AtualizarUsuario)
		r.Delete("/{id}", p.DeletarUsuario)
	})

	r.Route("/cep", func(r chi.Router) {
		r.Get("/", cep.ListarCep)
		r.Get("/{id}", cep.BuscarCep)
		r.Post("/", cep.CriarCep)
		r.Put("/", cep.AtualizarCep)
		r.Delete("/{id}", cep.DeletarCep)
	})

	println("servidor rodando na porta 8080")

	http.ListenAndServe(":8080", r)

}
