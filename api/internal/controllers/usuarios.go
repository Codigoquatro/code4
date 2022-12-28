package controllers

import (
	"api/internal/banco"
	"api/internal/modelo"
	"api/internal/repositorio"
	"api/internal/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	request, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}
	var usuario modelo.Usuario
	if err = json.Unmarshal(request, &usuario); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	if err = usuario.Preparar(); err != nil {
		respostas.Erro(w, http.StatusBadRequest, err)
		return
	}
	db, err := banco.Con()
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	repo := repositorio.NovoRepositoUsuarios(db)
	usuario.ID, err = repo.Criar(usuario)
	if err != nil {
		respostas.Erro(w, http.StatusInternalServerError, err)
		return
	}

	respostas.JSON(w, http.StatusCreated, usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscar todos os usuarios!!"))
}

func BuscarUsuarioByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BuscarUsuarioByID!!"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuario atualizado!!"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuario deletado!!"))
}
