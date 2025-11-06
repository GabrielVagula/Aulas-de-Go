package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	// ... dentro da func main()

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "home.html", nil)

		// Verifique se houve um erro ao executar o template
		if err != nil {
			// Envia uma resposta 500 (Erro Interno do Servidor) para o navegador
			http.Error(w, "Erro interno ao renderizar a página.", http.StatusInternalServerError)
			// Loga o erro no terminal para você ver o que deu errado!
			log.Printf("ERRO ao executar o template 'home.html': %v", err)
			return
		}
	})

	fmt.Println("Porta 5000 opened")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
