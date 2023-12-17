package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/apoorva2537/go-bookstore/pkg/models"
)
var NewBook models.Books
func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks:=models.GetAllBooks()
	res,_:=json.Marshal(newBooks)
	w.Header().set('content-type','pkglication/json')
	w.WriteHeader(http.StatusOK)
	w.write(res)	
}