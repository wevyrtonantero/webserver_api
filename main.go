package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type Pessoa struct {
	Id    int
	Nome  string
	Senha string
}

var Pessoas = []Pessoa{
	{
		Id:    1,
		Nome:  "Elton Casacio",
		Senha: "123",
	},
	{
		Id:    2,
		Nome:  "Nenem da Silva Filho",
		Senha: "564",
	},
	{
		Id:    3,
		Nome:  "Wevyrton Antero",
		Senha: "321",
	},
}

func inicial(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Ola mundo")

}

func usuarios(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Pessoas)

}
func buscaid(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(Pessoas[idint])

}

func main() {

	r := chi.NewRouter()
	r.Get("/", inicial)
	r.Get("/usuarios", usuarios)
	r.Get("/usuarios/{id}", buscaid)

	println("servidor rodando na porta 8080")

	http.ListenAndServe(":8080", r)
}
