package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/cmllmd/AplicacaoWeb-Loja/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, _ := strconv.ParseFloat(preco, 64)
		quantidadeConverida, _ := strconv.Atoi(quantidade)

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConverida)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}
		precoConvertido, _ := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}
		quantidadeConverida, _ := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		models.AtualizaProduto(idConvertido, nome, descricao, precoConvertido, quantidadeConverida)
	}
	http.Redirect(w, r, "/", 301)
}
