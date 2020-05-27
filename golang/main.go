//namespace of the package, Default package name main (package du fichier, par défaut le nom du package est main)
package main

// import file (importation des fichier)
import (
	"context"
	"html/template"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
)

//variable of Template type for get frontend template file (Variable de type Template pour avoir acces au fichier front)
var tpl *template.Template
var fileServer http.Handler

//init function it execute only one time when the file is run (La fonction est initialisé une seule fois au démarrage du serveur)
func init() {
	//tpl = template.Must(template.ParseGlob("./template/*")) // pour atteindre un dossier (for get all file in folder)
	tpl = template.Must(template.ParseFiles("./template/index.html")) // pour atteindre un fichier (for get a file)
}

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "restaurant-83ea4"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

//HandleError function for handle an error (La fonction pour gérer les erreurs)
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

//hello function for show index page (La fonction pour afficher la page d'index)
func hello(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, tpl.Name(), fileServer)
	HandleError(w, err)
}

//endpoint( point d'entrée de l'application)
func main() {
	mux := http.NewServeMux() //Create a custom server (Création d'un server)

	//Give access to assets files( Donner un accès au fichier image, style, ...)
	mux.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./static/images"))))
	mux.Handle("/styles/", http.StripPrefix("/styles", http.FileServer(http.Dir("./static/styles"))))
	//mux.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./static/js"))))

	mux.HandleFunc("/", hello)        //Execute a function on a select url (Executer une fonction sur quand une url est défini)
	http.ListenAndServe(":3000", mux) //Launch the server on port 8080
}
