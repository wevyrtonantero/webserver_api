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

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/safisa")
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare("insert into enderecos(rua, numero) values(?,?)")
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		cep.Enderecos[0].Rua,
		cep.Enderecos[0].Numero,
	)
	if err != nil {
		panic(err)
	}

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
