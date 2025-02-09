package controllers

import "net/http"

// CarregarTelaDeLogin exibe a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página de login"))
}
