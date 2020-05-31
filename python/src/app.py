# Librairies importation
# Importation des librairies qui seront utilisée
import os
from flask import Flask, render_template
import firebase_admin
from firebase_admin import credentials, firestore

# Creation of Flask app
# Creation d'une application Flask
app = Flask(__name__)


# Initialize Firestore DB
# Initialisation de Firestore DB
cred = credentials.Certificate("./ServiceAccount.json")
default_app = firebase_admin.initialize_app(cred)

db = firestore.client()

# Http route handler for home page
# route pour la page d'accueil
@app.route('/')
def index():
    return render_template('index.html')

# port setup
# configuration du port
port = int(os.environ.get('PORT', 8080))

# Launch the server on port 3000
# Lancer le serveur sur le port 3000
if __name__ == '__main__':
    app.run(threaded=True, host='localhost', port=port)