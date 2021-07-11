//godoc.com para encontrar bibliotecas em go

package main

import (
	"net/http"

	"github.com/cmllmd/AplicacaoWeb-Loja/routes"
)

func main() {

	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
