package pessoas

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/wevyrton/exercicio/internal/alertas"

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

func Inicial(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ola mundo")
}

func Usuarios(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Pessoas)
}

func Buscaid(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(Pessoas); i++ {
		if Pessoas[i].Id == idint {
			json.NewEncoder(w).Encode(Pessoas[i])
			return
		}
	}
	json.NewEncoder(w).Encode(alertas.AlertaDeId)
}

func Buscanome(w http.ResponseWriter, r *http.Request) {
	nome := chi.URLParam(r, "nome")
	for i := 0; i < len(Pessoas); i++ {
		if strings.EqualFold(Pessoas[i].Nome, nome) {
			json.NewEncoder(w).Encode(Pessoas[i])
			return
		}
	}
	json.NewEncoder(w).Encode(alertas.AlertaDeNome)
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario Pessoa
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		panic(err)
	}
	
	fmt.Println( usuario, "rota Criarusuario funcionando")





}
