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

func Usuarios(w http.ResponseWriter, r *http.Request) { //Pronto, nao ha necessidade de validação
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

func Buscaid(w http.ResponseWriter, r *http.Request) { //pronto, com validação
	id := chi.URLParam(r, "id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "id inválido")
		return
	}

	if idint <= 0 {

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

func Buscanome(w http.ResponseWriter, r *http.Request) { //pronto, com validação
	nome := chi.URLParam(r, "nome")

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
	var usuario Pessoa
	for rows.Next() {
		err := rows.Scan(&usuario.Id, &usuario.Nome, &usuario.Senha)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		if usuario.Nome == nome {
			json.NewEncoder(w).Encode(usuario)
			w.WriteHeader(http.StatusOK)
			return
		} else {

			json.NewEncoder(w).Encode("Verifique o nome digitado e tente novamente")
			w.WriteHeader(http.StatusBadRequest)
		}

	}

}

func CriarUsuario(w http.ResponseWriter, r *http.Request) { //Pronto, com validação

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
		fmt.Fprintf(w, "Id deve ser maior que ZERO")
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
	fmt.Fprintf(w, "Usuário Criado com sucesso")
	w.WriteHeader(http.StatusOK)

}
func AtualizarUsuario(w http.ResponseWriter, r *http.Request) { //Pronto , com validação
	var usuario Pessoa
	json.NewDecoder(r.Body).Decode(&usuario)

	db, err := sql.Open("mysql", "root:wedeju180587@tcp(localhost:3306)/safisa")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if usuario.Senha == "" || usuario.Nome == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Todos os Campos sao obrigatorios")
		return
	}
	if usuario.Id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Id deve ser maior que ZERO")
		return
	}

	rows, err := db.Query("select * from usuarios") //selec usado somente para validar a duplicidade de Id
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var usuariox Pessoa
	for rows.Next() {

		err := rows.Scan(&usuariox.Id, &usuariox.Nome, &usuariox.Senha)
		if err != nil {
			panic(err)
		}
	}

	if usuario.Id != usuariox.Id {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Usuario nao cadastrado!\nVerifique o Id digitado e tente novamente")
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
	fmt.Fprintf(w, "Usuário atualizado com sucesso")
	w.WriteHeader(http.StatusOK)

}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {//Pronto, com validaçõa

	id := chi.URLParam(r, "id")
	idint, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if idint <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Id deve ser maior que ZERO")
		return
	}

	db, err := sql.Open("mysql", "root:wedeju180587@tcp(localhost:3306)/safisa")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("select * from usuarios") //selec usado somente para validar a duplicidade de Id
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var usuariox Pessoa

	for rows.Next() {

		err := rows.Scan(&usuariox.Id, &usuariox.Nome, &usuariox.Senha)
		if err != nil {
			panic(err)
		}
	}

	if idint != usuariox.Id {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Usuario nao cadastrado!\nVerifique o Id digitado e tente novamente")
		return
	}

	stmt, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(idint)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Usuario Excluido com sucesso")

	w.WriteHeader(http.StatusOK)
}
