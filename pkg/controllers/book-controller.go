 package controllers

 import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/akhil/go-bookstore/pkg/utils"
	"github.com/akhil/go-bookstore/pkg/models"

 )

 var NewBook = models.Book

 func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ :=json.Marshal(newBooks) 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
 }

 