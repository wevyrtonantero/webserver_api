package pessoas

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/wevyrton/exercicio/internal/alertas"

	"github.com/go-chi/chi"
)

type Pessoa struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Senha string `json:"senha"`
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
	db, err := sql.Open("mysql", "root:wedeju180587@tcp(localhost:3306)/safisa")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	rows, err := db.Query("select * from usuarios")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var usuarios []Pessoa

	for rows.Next() {

		var usuario Pessoa

		err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Senha)
		if err != nil {
			panic(err)

		}
		usuarios = append(usuarios, usuario)
	}
	json.NewEncoder(w).Encode(usuarios)

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

	db, err := sql.Open("mysql", "root:wedeju180587@tcp(localhost:3306)/safisa")
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("insert into usuarios(id, nome, senha) values(?,?,?)")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	defer stmt.Close()

	_, err = stmt.Exec(
		usuario.Id,
		usuario.Nome,
		usuario.Senha,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	var usuario Pessoa
	json.NewDecoder(r.Body).Decode(&usuario)

	db, err := sql.Open("mysql", "root:wedeju180587@tcp(localhost:3306)/safisa")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("UPDATE usuarios SET nome = ?, senha = ? WHERE id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(usuario.Nome, usuario.Senha, usuario.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	db, err := sql.Open("mysql", "root:wedeju180587@tcp(localhost:3306)/safisa")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
