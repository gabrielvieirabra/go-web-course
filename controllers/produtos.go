package controllers

import (
	"github.com/go-web-course/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosProdutos()
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

		precoConvertidoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Fatal(err)
		}

		quantidadeConvertidoInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Fatal(err)
		}

		models.CriaNovoProduto(nome, descricao, precoConvertidoFloat, quantidadeConvertidoInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeleteProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidoParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}

		precoConvertidoFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Fatal(err)
		}

		quantidadeConvertidoInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Fatal(err)
		}

		models.AtualizaProduto(idConvertidoParaInt, nome, descricao, precoConvertidoFloat, quantidadeConvertidoInt)

	}
	http.Redirect(w, r, "/", 301)
}
