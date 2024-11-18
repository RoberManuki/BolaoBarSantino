package main

import (
	"bolao/src/app/handler"
	"html/template"
	"log"
	"net/http"
	//"os"
)

type PageVariables struct {
	Title  string
	Header string
}

var templates *template.Template

func init() {
	var err error
	templates, err = template.ParseFiles(
		"templates/home.html",
		"templates/partidas.html",
		"templates/formularioPartida.html",
		"templates/menu.html",
	)
	if err != nil {
		log.Fatal("Erro ao carregar templates: ", err)
	}
}

func main() {
	setupPartidasRoutes()
	setupTimesRoutes()
	setupPagesRoutes()
	setupStaticRoutes()

	log.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupPartidasRoutes() {
	http.HandleFunc("/api/partidas", handler.PartidaHandler)
	http.HandleFunc("/api/partidas/", handler.PartidaByID)
	http.HandleFunc("/api/partidas/validar", handler.ValidarPartidaHandler)
}

func setupTimesRoutes() {
	http.HandleFunc("/api/times", handler.GetTimes)
}

func setupPagesRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/partidas", partidasPage)
	http.HandleFunc("/partida/formulario", formularioPartidaPage)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title:  "Página Inicial",
		Header: "Página Inicial",
	}
	err := templates.ExecuteTemplate(w, "home.html", pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func partidasPage(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title:  "Lista de Partidas",
		Header: "Lista de Partidas",
	}
	err := templates.ExecuteTemplate(w, "partidas.html", pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func formularioPartidaPage(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		Title:  "Criar/Editar Partida",
		Header: "Criar/Editar Partida",
	}
	err := templates.ExecuteTemplate(w, "formularioPartida.html", pageVariables)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func setupStaticRoutes() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
}
