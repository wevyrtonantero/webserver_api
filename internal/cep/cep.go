package cep

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/wevyrton/exercicio/internal/pessoas"
)

type Logradouro struct {
	Rua     string         `json:"Rua"`
	Numero  int            `json:"Numero"`
	Bairro  string         `json:"Bairro"`
	Cidade  string         `json:"Cidade"`
	Uf      string         `json:"Uf"`
	Pessoas pessoas.Pessoa `json:"Morador"`
}

var morador0 pessoas.Pessoa = pessoas.Pessoas[0]
var morador1 pessoas.Pessoa = pessoas.Pessoas[1]
var morador2 pessoas.Pessoa = pessoas.Pessoas[2]

var Enderecos = []Logradouro{
	{
		Rua:     "Henriqueta Lisboa",
		Numero:  100,
		Bairro:  "Jardim Amanda",
		Cidade:  "Hortolândia",
		Uf:      "São Paulo",
		Pessoas: morador0,
	},

	{
		Rua:     "Helizia Machado Benassi",
		Numero:  333,
		Bairro:  "Nova Cidade Jardim",
		Cidade:  "Jundiai",
		Uf:      "São Paulo",
		Pessoas: morador1,
	},

	{
		Rua:     "Borba Gato",
		Numero:  28,
		Bairro:  "Comendador Soares",
		Cidade:  "Nova Iguaçu",
		Uf:      "Rio de Janeiro",
		Pessoas: morador2,
	},
	{
		Rua:     "Henriqueta Lisboa",
		Numero:  100,
		Bairro:  "Jardim Amanda",
		Cidade:  "Hortolândia",
		Uf:      "São Paulo",
		Pessoas: pessoas.Pessoa{Id: 04, Nome: "Gabriel", Senha: "2609"},
	},

	{
		Rua:     "Helizia Machado Benassi",
		Numero:  333,
		Bairro:  "Nova Cidade Jardim",
		Cidade:  "Jundiai",
		Uf:      "São Paulo",
		Pessoas: pessoas.Pessoa{},
	},

	{
		Rua:     "Afonso Pena",
		Numero:  2558,
		Bairro:  "Abolição",
		Cidade:  "Campinas",
		Uf:      "São Paulo",
		Pessoas: pessoas.Pessoa{Id: 06, Nome: "Isabelle", Senha: "bf!!@#$¨&*(DGHae22212"},
	},
}

func ListarCep(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Enderecos)
}
func BuscarCep(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("NAO EXISTE ESSE USUARIO")
	}
	for _, cep := range Enderecos {
		if cep.Pessoas.Id == idint {
			json.NewEncoder(w).Encode(cep)

		}
	}

	fmt.Println(" funcionando...BuscarCep GET...")
}
func CriarCep(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logradouro funcionando...CriarLogradouro")
}
func AtualizarCep(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" funcionando...CriarCep PUT...")
}
func DeletarCep(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" funcionando...DeletarCep DELETE...")
}
