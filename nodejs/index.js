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
