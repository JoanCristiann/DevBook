package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// CarregarTelaDeLogin exibe a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "login.html", nil)
}

// CarregarTelaDeLogin exibe a tela de cadastro
func CarregarPaginaDeCadastroDeUsuarios(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}
