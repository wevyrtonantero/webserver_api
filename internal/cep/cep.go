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
	Rua    string         `json:"rua"`
	Numero int            `json:"numero"`
	Bairro string         `json:"bairro"`
	Cidade string         `json:"cidade"`
	Uf     string         `json:"uf"`
	Pessoa pessoas.Pessoa `json:"pessoa"`
}

var morador0 pessoas.Pessoa = pessoas.Pessoas[0]
var morador1 pessoas.Pessoa = pessoas.Pessoas[1]
var morador2 pessoas.Pessoa = pessoas.Pessoas[2]

var Enderecos = []Logradouro{
	{
		Rua:    "Henriqueta Lisboa",
		Numero: 100,
		Bairro: "Jardim Amanda",
		Cidade: "Hortolândia",
		Uf:     "São Paulo",
		Pessoa: morador0,
	},

	{
		Rua:    "Helizia Machado Benassi",
		Numero: 333,
		Bairro: "Nova Cidade Jardim",
		Cidade: "Jundiai",
		Uf:     "São Paulo",
		Pessoa: morador1,
	},

	{
		Rua:    "Borba Gato",
		Numero: 28,
		Bairro: "Comendador Soares",
		Cidade: "Nova Iguaçu",
		Uf:     "Rio de Janeiro",
		Pessoa: morador2,
	},
	{
		Rua:    "Henriqueta Lisboa",
		Numero: 100,
		Bairro: "Jardim Amanda",
		Cidade: "Hortolândia",
		Uf:     "São Paulo",
		Pessoa: pessoas.Pessoa{Id: 04, Nome: "Gabriel", Senha: "2609"},
	},

	{
		Rua:    "Helizia Machado Benassi",
		Numero: 333,
		Bairro: "Nova Cidade Jardim",
		Cidade: "Jundiai",
		Uf:     "São Paulo",
		Pessoa: pessoas.Pessoa{},
	},

	{
		Rua:    "Afonso Pena",
		Numero: 2558,
		Bairro: "Abolição",
		Cidade: "Campinas",
		Uf:     "São Paulo",
		Pessoa: pessoas.Pessoa{Id: 06, Nome: "Isabelle", Senha: "bf!!@#$¨&*(DGHae22212"},
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
		if cep.Pessoa.Id == idint {
			json.NewEncoder(w).Encode(cep)

		}
	}

}

func CriarCep(w http.ResponseWriter, r *http.Request) {
	var morador Logradouro
	err := json.NewDecoder(r.Body).Decode(&morador)
	if err != nil {
		panic(err)
	}
	Enderecos = append(Enderecos, morador)
}

func AtualizarCep(w http.ResponseWriter, r *http.Request) {
	var usuario Logradouro
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		panic(err)
	}

	for indice, endereco := range Enderecos {
		if usuario.Pessoa.Id == Enderecos[indice].Pessoa.Id {
			Enderecos[indice] = usuario
			fmt.Printf("%v ALterado\n", endereco.Pessoa.Nome)
		}
	}

}

func DeletarCep(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("NAO EXISTE ESSE USUARIO")
	}
	for indice, endereco := range Enderecos {
		if endereco.Pessoa.Id == idint {
			Enderecos = append(Enderecos[:indice], Enderecos[indice+1:]...)

		}
	}
}
