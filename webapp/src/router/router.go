package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar cria um router com as rotas configuradas
func Gerar() *mux.Router {
	router := mux.NewRouter()
	return rotas.Configurar(router)
}
