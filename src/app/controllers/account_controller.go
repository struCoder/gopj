package controllers

import (
	"io"
	"net/http"
  "html/template"
)

func HandleLogin(res http.ResponseWriter, req *http.Request) {
	reqMethod = req.Method
  if reqMethod == "GET" {
    template.ParseFiles("views")
  } else if reqMethod = "POST" {

  } else {
    io.WriteString(w, "FORBIDDEN")
  }
}

func getLogin(res http.ResponseWriter, req *http.Request) {
  
}

func postLogin(res http.ResponseWriter, req *http.Request) {
  
}

func Logout(res http.ResponseWriter, req *http.Request) {
  
}

func Home(res http.ResponseWriter, req *http.Request) {

}