package main

import (
	"bolao/src/app/handler"
	"html/template"
	"log"
	"net/http"
)

type PageVariables struct {
	Title  string
	Header string
}

func main() {
	templates := template.Must(template.ParseFiles(
		"templates/home.html",
		"templates/menu.html",
		"templates/partidas.html",
		"templates/formularioPartida.html",
	))

	// Rota para a página inicial
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		pageVariables := PageVariables{
			Title:  "Página Inicial",
			Header: "Página Inicial",
		}
		err := templates.ExecuteTemplate(w, "home.html", pageVariables)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Rota para exibir a página de partidas
	http.HandleFunc("/partidas", func(w http.ResponseWriter, r *http.Request) {
		pageVariables := PageVariables{
			Title:  "Lista de Partidas",
			Header: "Lista de Partidas",
		}
		err := templates.ExecuteTemplate(w, "partidas.html", pageVariables)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Rota para o formulário de criação/edição de partida
	http.HandleFunc("/partida/formulario", func(w http.ResponseWriter, r *http.Request) {
		pageVariables := PageVariables{
			Title:  "Criar/Editar Partida",
			Header: "Criar/Editar Partida",
		}
		err := templates.ExecuteTemplate(w, "formularioPartida.html", pageVariables)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	// Controllers para a API de partidas (GET, POST, PUT)
	http.HandleFunc("/api/partidas", handler.PartidaHandler)
	http.HandleFunc("/api/times", handler.GetTimes)
	http.HandleFunc("/api/partidas/validar", handler.ValidarPartidaHandler)

	// Servir arquivos estáticos
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))

	log.Println("Servidor ouvindo na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
