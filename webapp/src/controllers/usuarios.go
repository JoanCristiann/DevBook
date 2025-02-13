package controllers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"webapp/src/respostas"
)

// CriarUsuario chama a API para cadastrar um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, err := json.Marshal(map[string]string{
		"nome":     r.FormValue("nome"),
		"email":    r.FormValue("email"),
		"username": r.FormValue("username"),
		"senha":    r.FormValue("senha"),
	})

	if err != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: err.Error()})
		return
	}

	response, err := http.Post("http://localhost:5000/usuarios", "application/json", bytes.NewBuffer(usuario))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCode(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}
