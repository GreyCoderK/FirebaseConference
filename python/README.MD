![logo python](https://miro.medium.com/max/1400/0*9qrxZ3xKv0dm814Q.png)

# Presentation

Python est un langage de programmation interprété, multi-paradigme et multiplateformes.
Il favorise la programmation impérative structurée, fonctionnelle et orientée objet.
Il est doté d'un typage dynamique fort, d'une gestion automatique de la mémoire par ramasse-miettes et d'un système de gestion d'exceptions ; il est ainsi similaire à Perl, Ruby, Scheme, Smalltalk et Tcl. 

# Objectif

En tant que langage de programmation de haut niveau, Python permet aux programmeurs de se focaliser sur ce qu’ils font plutôt que sur la façon dont ils le font. Ainsi, écrire des progammes prend moins de temps que dans un autre langage. Il s’agit d’un langage idéal pour les débutants.

# Installation

Pour installer Python sur sa machine, il faut se rendre sur le site suivant [https://www.python.org/downloads/](https://www.python.org/downloads/) puis télécharger un éxécutable que vous installerez chez vous

# Quelques note avant de débuter 

Pour cette implementation avec python nous allons utiliser [Flask](https://flask.palletsprojects.com/en/1.1.x/).

Flask est un framework open-source de développement web en Python. Son but principal est d'être léger, afin de garder la souplesse de la programmation Python, associé à un système de templates. Il est distribué sous *licence BSD*. 

Flask a été créé initialement par Armin Ronacher comme étant un poisson d'avril. Le souhait de Ronacher était de réaliser un framework web contenu dans un seul fichier Python mais pouvant maintenir des applications très demandées.

En 2018, Flask était élu "Framework web le plus populaire" par le Python Developers Survey. En janvier 2020, il cumulait plus de 49000 étoiles sur Github, plus que n'importe quel autre framework de développement web Python.

# Création d'un projet

Pour créer notre projet python nous allons utliser un environnement virtuel. On utilise un environnement virtuel généralement pour ne pas que les dépendances de notre projet python soit installés de façon global, mais dans notre projet. Lors du deploiement on pourra alors récupérer facilement ces dépendances pour refaire la même opération sur notre serveur.

Quoi de mieux que la [documentation](https://docs.python.org/fr/3/library/venv.html) pour nous montrer comment installer un environnement virtuel.

Installons maintenant Flask apres avoir activé et créer notre environnement virtuel:

```shell
pip3 install flask
```

ou 

```shell
python3 -m pip install flask
```

Après avoir créer notre environnement virutel nous pouvons passé à la phase suivante.

# Intégrer Firebase dans son projet

Une fois que vous avez installé Python sur votre machine, installer Firebase dans votre projet.

Pour cela déplacer vous dans le repertoire contenant votre projet, puis ouvrez un invite de commande dans le projet.

Tapez ensuite la commande suivante :

```shell
pip install --upgrade firebase-admin
```

# Création d'un serveur local

```python
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
    # flask render all file in templates folder
    # flask affiche tous les fichier dans le repertoire templates
    return render_template('index.html')

# port setup
# configuration du port
port = int(os.environ.get('PORT', 8080))

# Launch the server on port 3000
# Lancer le serveur sur le port 3000
if __name__ == '__main__':
    app.run(threaded=True, host='localhost', port=port)
```

 
 - *N.B* :

	Lorque que vous lancerez le serveur, allez à cette adresse `localhost:3000/` pour avoir une vue de l'affichage.

	Mais je suis au regret de vous dis que vous ne verrez presque rien pour l'instant car l'affichage est intimement lié au JavaScript donc pour avoir un bref aperçu sur le serveur. Donc pour le fait, enlever l'attribut `hidden` de quelques `div` puis relancer le serveur.

	Nous venons de connecter la base de donnée *Cloud FireStore* à notre projet mais pour être sûr que cela fonctionne nous devons effectuer une requête chose que nous ne ferons pas d'aussi tôt après avoir réorganiser le fichier `index.html` afin que nous puissons voir de façon visuel les changements.

	Pour télécharger le fichier `ServiceAccount.json` correspondant à votre projet je vous conseil de vous rendre dans les *paramètres* de votre projet puis *compte et service* et la vous pouvez télécharger le fichier que vous placerez dans le même dossier que le fichier principal de votre application.

Puis rendez-vous dans votre terminal, ensuite tapez :

```
python3 app.py
```

Pour comprendre en détail ce que nous avons effectué avec flask je vous recommende la [documentation](https://flask.palletsprojects.com/en/1.1.x/)


# Réorganisation des fichiers template
