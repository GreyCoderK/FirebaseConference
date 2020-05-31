![logo nodejs](https://miro.medium.com/max/1400/1*Jw9V__6jYhm2amP74D_0lw.png)

# Presentation

Node.js est une plateforme logicielle libre en JavaScript orientée vers les applications réseau événementielles hautement concurrentes qui doivent pouvoir monter en charge.
Elle utilise la machine virtuelle V8, la librairie libuv pour sa boucle d'évènements, et implémente sous licence MIT les spécifications CommonJS.
Parmi les modules natifs de Node.js, on retrouve http qui permet le développement de serveur HTTP. Il est donc possible de se passer de serveurs web tels que Nginx ou Apache lors du déploiement de sites et d'applications web développés avec Node.js.
Concrètement, Node.js est un environnement bas niveau permettant l’exécution de JavaScript côté serveur.


# Objectif

Node.js utilise le langage JavaScript du serveur et permet de faire du JavaScript en dehors du navigateur. En se basant sur la puissance du JavaScript, Node.js propose une nouvelle façon de développer des applications web dynamiques.

Avant l’existence du Node.js, le JavaScript était toujours utilisé côté client, plus précisément côté utilisateur qui navigue sur une application web. Le navigateur web de l’utilisateur (Safari, Chrome, IE etc.) exécute le JavaScript et modifie la page HTML.

# Installation

Pour installer Node.js sur sa machine, il faut se rendre sur le site suivant [https://nodejs.org/en/download/](https://nodejs.org/en/download/) puis télécharger l'executable vous correspondant

# Quelques note avant de débuter 

Pour cette implementation avec nodejs nous allons utiliser [Express](https://expressjs.com/fr/)

Express.js est un framework pour construire des applications web basées sur Node.js. C'est de fait le framework standard pour le développement de serveur en Node.js. L'auteur original, TJ Holowaychuck, le décrit comme un serveur inspiré de [Sinatra](http://sinatrarb.com/) dans le sens qu'il est relativement minimaliste tout en permettant d'étendre ses fonctionnalités via des plugins. 

# Création du projet

Pour créer un projet avec Nodejs rien de plus simple que de créer un repertoire. puis ensuite de l'initialiser avec la commande suivante :

```shell
npm init --y
```

Installons maintenant ExpressJs, grâce à la commande suivante :

```
npm i express
```

# Intégrer Firebase à son projet

Une fois que vous avez installé Node.Js sur votre machine, installer Firebase dans votre projet.

Pour cela déplacer vous dans le repertoire contenant votre projet, puis ouvrez un invite de commande dans le projet.

Tapez ensuite la commande suivante :

```shell
npm install firebase-admin --save
```

# Création d'un serveur local

```js
/**
 * File importation
 * importation de librairie
 */
const express = require('express')
const app = express()

/**
 * Firestore Setup
 * Configuration de FireStore
 */
var admin = require("firebase-admin");

var serviceAccount = require("./ServiceAccount.json");

admin.initializeApp({
  credential: admin.credential.cert(serviceAccount),
  databaseURL: "https://restaurant-83ea4.firebaseio.com"
});

/**
 * firestore instance
 * Instance firestore
 */
let db = admin.firestore()

/**
 * Home page http handler
 * page d'accueil
 */
app.get('/', function (req, res) {
    /**
     * render file at the path __dirname + '/templates/index.html'
     * affiche le fichier à l'adresse __dirname + '/templates/index.html'
     */
    res.sendFile(__dirname + '/templates/index.html')
})

/**
 * Launch the server on port 3000
 * Lancer le serveur sur le port 3000
 */
app.listen(3000, function () {
  console.log('Restaurant app listening on port 3000!')
})
```


 - *N.B* :

	Lorque que vous lancerez le serveur, allez à cette adresse `localhost:3000/` pour avoir une vue de l'affichage.

	Mais je suis au regret de vous dis que vous ne verrez presque rien pour l'instant car l'affichage est intimement lié au JavaScript donc pour avoir un bref aperçu sur le serveur. Donc pour le fait, enlever l'attribut `hidden` de quelques `div` puis relancer le serveur.

	Nous venons de connecter la base de donnée *Cloud FireStore* à notre projet mais pour être sûr que cela fonctionne nous devons effectuer une requête chose que nous ne ferons pas d'aussi tôt après avoir réorganiser le fichier `index.html` afin que nous puissons voir de façon visuel les changements.

	Pour télécharger le fichier `ServiceAccount.json` correspondant à votre projet je vous conseil de vous rendre dans les *paramètres* de votre projet puis *compte et service* et la vous pouvez télécharger le fichier que vous placerez dans le même dossier que le fichier principal de votre application.

Puis rendez-vous dans votre terminal, ensuite tapez :

```shell
node index.js
```

Pour comprendre en détail ce que nous avons effectué avec Express.Js je vous recommende la [documentation](https://expressjs.com/fr/)
