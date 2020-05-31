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

# Création d'un projet

Pour créer un avec go il suffit tout simplement de créer un dossier portant le nom de votre projet dans lequel vous allez mettre les fichiers que vous coderez ensuite.

Une fois passé cette étape, tapez la commande suivante dans le repertoire du fichier:

```
go mod init
```

cette commande initialise notre projet en tant que module golang, cela facilite la gestion des dépendances.

Nous pouvons commencer à coder.

# Intégrer Firebase à son projet

Une fois que vous avez installé **Golang** sur votre machine, installer Firebase dans votre projet.

Pour cela déplacer vous dans le repertoire contenant votre projet, puis ouvrez un **invite de commande** dans le projet.

Tapez ensuite la commande suivante :


```bash
go get -u firebase.google.com/go
```

# Création d'un server local

Pour l'instant nous allons juste faire les configurations de base de notre serveur, qui tournera sur le port 3000.
Après avoir configurer le serveur, nous allons connecter notre projet golang à notre base de donnée firebase. 

```go
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

```

 - *N.B* :

	Lorque que vous lancerez le serveur, allez à cette adresse `localhost:3000/` pour avoir une vue de l'affichage.

	Mais je suis au regret de vous dis que vous ne verrez presque rien pour l'instant car l'affichage est intimement lié au JavaScript donc pour avoir un bref aperçu sur le serveur. Donc pour le fait, enlever l'attribut `hidden` de quelques `div` puis relancer le serveur.

	Nous venons de connecter la base de donnée *Cloud FireStore* à notre projet mais pour être sûr que cela fonctionne nous devons effectuer une requête chose que nous ne ferons pas d'aussi tôt après avoir réorganiser le fichier `index.html` afin que nous puissons voir de façon visuel les changements.

	Pour télécharger le fichier `ServiceAccount.json` correspondant à votre projet je vous conseil de vous rendre dans les *paramètres* de votre projet puis *compte et service* et la vous pouvez télécharger le fichier que vous placerez dans le même dossier que le fichier principal de votre application.

Puis rendez-vous dans votre terminal, ensuite tapez :

```shell
go run main.go
```

Pour savoir un peu plus sur le fonctionnement du web avec golang, je vous conseille vivement de bien vouloir lire la [documentation](https://golang.org/doc/) ou rendez-vous sur ce [github](https://github.com/GoesToEleven/golang-web-dev) de @GoesToEleven .


# Réorganisation des fichiers template

