![golang logo](https://miro.medium.com/max/1400/0*7vQ8eRc28yz9k__r.png)

# Presentation

Go est un langage de programmation compilé et concurrent inspiré de C et Pascal.
Ce langage a été développé par Google à partir d'un concept initial de Robert Griesemer, Rob Pike et Ken Thompson.
Go possède deux implémentations, la première utilise gc, le compilateur Go et la seconde utilise gccgo, « frontend » GCC écrit en C++.
Go est écrit en C en utilisant yacc et GNU Bison pour l'analyse syntaxique jusqu'à la version 1.4, et en Go lui-même pour les versions suivantes (1.5).

# Objectif

Un objectif de Go est donné par Rob Pike, l'un de ses trois créateurs, qui dit à propos des développeurs inexpérimentés :
> Ils ne sont pas capables de comprendre un langage brillant, mais nous voulons les amener à réaliser de bons programmes.
> -Rob Pike

Ainsi, le langage que nous leur donnons doit être facile à comprendre et facile à adopter
Go veut faciliter et accélérer la programmation à grande échelle : en raison de sa simplicité il est donc concevable de l'utiliser aussi bien pour écrire des applications, des scripts ou de grands systèmes.
Cette simplicité est nécessaire aussi pour assurer la maintenance et l'évolution des programmes sur plusieurs générations de développeurs.

S'il vise aussi la rapidité d'exécution, indispensable à la programmation système, il considère le multithreading comme le moyen le plus robuste d'assurer sur les processeurs actuels cette rapidité tout en rendant la maintenance facile par séparation de tâches simples exécutées indépendamment afin d'éviter de créer des « usines à gaz ».
Cette conception permet également le fonctionnement sans réécriture sur des architectures multi-cœurs en exploitant immédiatement l'augmentation de puissance correspondante.

# Installation

Pour installer Go sur sa machine, il faut se rendre sur le site suivant [https://golang.org/doc/install](https://golang.org/doc/install) puis suivre les instructions en fonction de son système d'exploitation

# Intégrer Firebase à son projet

Une fois que vous avez installé **Golang** sur votre machine, installer Firebase dans votre projet.

Pour cela déplacer vous dans le repertoire contenant votre projet, puis ouvrez un **invite de commande** dans le projet.

Tapez ensuite la commande suivante :


```bash
go get -u firebase.google.com/go
```

# Création d'un server local

Pour l'instant nous allons juste créer un simple serveur Go, sur le serveur 8080

```go
//namespace of the package, Default package name main (package du fichier, par défaut le nom du package est main)
package main

// import file (importation des fichier)
import (
	"context"
	"html/template"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

//variable of Template type for get frontend template file (Variable de type Template pour avoir acces au fichier front)
var tpl *template.Template
var fileServer http.Handler

//init function it execute only one time when the file is run (La fonction est initialisé une seule fois au démarrage du serveur)
func init() {
	//tpl = template.Must(template.ParseGlob("./template/*")) // pour atteindre un dossier (for get all file in folder)
	tpl = template.Must(template.ParseFiles("./template/index.html")) // pour atteindre un fichier (for get a file)
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
	defer client.Close()

	mux := http.NewServeMux() //Create a custom server (Création d'un server)

	//Give access to assets files( Donner un accès au fichier image, style, ...)
	mux.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./static/images"))))
	mux.Handle("/styles/", http.StripPrefix("/styles", http.FileServer(http.Dir("./static/styles"))))
	//mux.Handle("/js/", http.StripPrefix("/js", http.FileServer(http.Dir("./static/js"))))

	mux.HandleFunc("/", hello)        //Execute a function on a select url (Executer une fonction sur quand une url est défini)
	http.ListenAndServe(":3000", mux) //Launch the server on port 8080
}

```

Puis rendez-vous dans votre terminal, ensuite taper :

```shell
go run main.go
```

Pour savoir un peu plus sur les fonctions que nous avons utilisé je vous conseille vivement de bien vouloir lire la [documentation](https://golang.org/doc/).
