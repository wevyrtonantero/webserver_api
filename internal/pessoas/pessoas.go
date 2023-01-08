package pessoas

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	w.WriteHeader(http.StatusOK)

}

func Buscaid(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "id inválido")
		return
	}

	if idint == 0 {

		fmt.Fprintf(w, "id deve ser maior que ZERO")
		return

	}

	db, err := sql.Open("mysql", "root:wedeju180587@tcp(localhost:3306)/safisa")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM usuarios WHERE id =?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer stmt.Close()

	var usuario Pessoa
	err = stmt.QueryRow(idint).Scan(&usuario.Id, &usuario.Nome, &usuario.Senha)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	if usuario.Id == 0 {
		json.NewEncoder(w).Encode("Usuario nao Existe")
		return
	}

	json.NewEncoder(w).Encode(usuario)

}

func Buscanome(w http.ResponseWriter, r *http.Request) {
	nome := chi.URLParam(r, "nome")
	fmt.Println(nome)

	db, err := sql.Open("mysql", "root:wedeju180587@tcp(localhost:3306)/safisa")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer db.Close()
	rows, err := db.Query("SELECT * FROM usuarios WHERE nome =?", nome)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer rows.Close()
	fmt.Println(nome)
	var usuario Pessoa
	for rows.Next() {
		err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Senha)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(usuario)
		fmt.Println(nome)

	}

}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {

	var usuario Pessoa
	//pegando dados do Bory, que foi difgitado pelo usuario
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		panic(err)
	}

	//abrindo conexão com base de dados
	db, err := sql.Open("mysql", "root:wedeju180587@tcp(localhost:3306)/safisa")
	if err != nil {
		panic(err)
	}
	//validando usuario
	if usuario.Id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Usuario deve ser maior que ZERO")
		return
	} else if usuario.Nome == "" || usuario.Senha == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Usuario e Senha não pode ser vazio")
		return
	}
	rows, err := db.Query("select * from usuarios") //selec usado somente para validar a duplicidade de Id
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {

		var usuariox Pessoa

		err := rows.Scan(&usuariox.Id, &usuariox.Nome, &usuariox.Senha)
		if err != nil {
			panic(err)
		}

		if usuario.Id == usuariox.Id {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Usuário já Cadastrado\n")
			json.NewEncoder(w).Encode(usuariox)
			return
		}
	}

	//inserindo dados digitado pelo usuario na base de dados
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
	w.WriteHeader(http.StatusOK)

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
	w.WriteHeader(http.StatusOK)
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
	w.WriteHeader(http.StatusOK)
}
