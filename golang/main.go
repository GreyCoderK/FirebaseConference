//namespace of the package, Default package name main
//(package du fichier, par défaut le nom du package est main)
package main

// import file
//(importation des fichier)
import (
	"context"
	"html/template"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

//variable of Template type for get frontend template file
//(Variable de type Template pour avoir acces au fichier front)
var tpl *template.Template

//Handler variable for file serve
//(variable de type http.Handler pour utiliser les fichiers static)
var fileServer http.Handler

//init function it execute only one time when the file is run
//(La fonction est initialisé une seule fois au démarrage du serveur)
func init() {
	//(for get all file in folder)
	//pour atteindre un dossier
	//tpl = template.Must(template.ParseGlob("./template/*"))

	// pour atteindre un fichier
	// (for get a file)
	tpl = template.Must(template.ParseFiles("./template/index.html"))
}

//HandleError function for handle an error
//(La fonction pour gérer les erreurs)
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

//hello function for show index page
//(La fonction pour afficher la page d'index)
func hello(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, tpl.Name(), fileServer)
	HandleError(w, err)
}

//endpoint
//(point d'entrée de l'application)
func main() {

	//set up for firestore configuration
	//configuration pour l'integration de firebase
	//ServiceAccount.json is the file set up of the firebase database, put this file in the same folder that main.go
	//ServiceAccount.json est le fichier de configuration de notre base de donnée il faut le mettre dans le même dossier que le main.go
	projectID := "restaurant-83ea4"
	opt := option.WithCredentialsFile("./ServiceAccount.json")
	config := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(context.Background(), config, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Close client when done with
	// se termine quand le client en a fini
	defer client.Close()

	//Create a custom server
	//(Création d'un server)
	mux := http.NewServeMux()

	//Give access to assets files
	//(Donner un accès au fichier image, style, ...)
	mux.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./static/images"))))
	mux.Handle("/styles/", http.StripPrefix("/styles", http.FileServer(http.Dir("./static/styles"))))
	//mux.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./static/js"))))

	//Execute a function on a select url
	//(Executer une fonction sur quand une url est défini)
	mux.HandleFunc("/", hello)

	//Launch the server on port 3000
	//Lancer le serveur sur le port 3000
	http.ListenAndServe(":3000", mux)
}
