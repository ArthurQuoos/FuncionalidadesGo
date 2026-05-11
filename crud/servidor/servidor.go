package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type usuario struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// CriarUsuario é a função responsável por criar um novo usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := io.ReadAll(r.Body)
	if erro != nil {
		http.Error(w, "Erro ao ler corpo da requisição", http.StatusBadRequest)
		return
	}

	var usuario usuario
	erro = json.Unmarshal(corpoRequisicao, &usuario)
	if erro != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	fmt.Println(usuario)

	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Erro ao conectar ao banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) VALUES (?, ?)")
	if erro != nil {
		http.Error(w, "Erro ao preparar a consulta", http.StatusInternalServerError)
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil {
		http.Error(w, "Erro ao executar o statement", http.StatusInternalServerError)
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil {
		http.Error(w, "Erro ao obter o ID inserido", http.StatusInternalServerError)
		return
	}

	//STATUS CODES
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com ID: %d", idInserido)))

}

// BuscarUsuarios é a função responsável por buscar todos os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	db, erro := banco.Conectar()
	if erro != nil {
		http.Error(w, "Erro ao conectar ao banco de dados", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	linhas, erro := db.Query("SELECT * from usuarios")
	if erro != nil {
		http.Error(w, "Erro ao executar a consulta", http.StatusInternalServerError)
		return
	}
	defer linhas.Close()

	var usuarios []usuario
	for linhas.Next() {
		var usuario usuario
		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil {
			w.Write([]byte("Erro ao escanear o usuario de dados"))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil {
		http.Error(w, "Erro ao codificar os usuários para JSON", http.StatusInternalServerError)
		return
	}
}

// BuscarUsuario é a função responsável por buscar um usuário específico no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {

}
