> __Un voyage de mille lieues a commencé par un pas.__
> *-Proverbe chinois*

# Firebase Conference

__Guide pour la conférence de Firebase, pour toutes les personnes qui assisteront à cette conférence dans l'optique que ceux-ci puissent revenir sur les propos de la conférence et aussi pratiquer par des exemples__


# Plan de cette session sur Firebase
  - Ce qu'il faut savoir sur Firebase
    - Historique de Firebase
    - Qu'est ce que Firebase ?
    - Philosophie de Firebase
  - Un peu de pratique avec Firebase(FireStore)
    - Créer une base de données Firebase
    - Intégration de firebase à son projet
  - Exemple
  - D'autre sources
  - Remerciement

# Ce qu'il faut savoir sur Firebase

> __Ouvre un livre, c'est lui qui t'ouvrira.__
> *-Proverbe chinois*


## Historique de Firebase

Lancé en 2011 sous le nom d'Envolve, par Andrew Lee et par James Templin, le service est racheté par Google en octobre 2014. Il appartient aujourd'hui à la maison mère de Google : [Alphabet](https://fr.wikipedia.org/wiki/Alphabet_(entreprise)).


## Qu'est ce que Firebase ?

Firebase est un ensemble de services d'hébergement pour n'importe quel type d'application (Android, iOS, Javascript, Node.js, Java, Unity, PHP, C++ ...). Il propose d'héberger en [NoSQL](https://fr.wikipedia.org/wiki/NoSQL) et en temps réel des bases de données, du contenu, de l'authentification sociale (Google, Facebook, Twitter et Github), et des notifications, ou encore des services, tel que par exemple un serveur de communication temps réel.


## Philosophie de Firebase

L’objectif de la création de [Firebase.google.com](https://firebase.google.com) en 2011 par James Tamplin et Andrew Lee est d’éviter aux professionnels et aux particuliers de s’engager dans un processus complexe de création et de maintenance d’une architecture serveur.

De plus, la plateforme peut être exploitée par plusieurs utilisateurs en même temps sans connaître un quelconque bug. La praticité est également au rendez-vous grâce à ses fonctionnalités intuitives. Depuis le rachat de la plateforme par Google en 2014, Firebase sdks a connu de nombreuses améliorations et n’a de cesse de satisfaire ses utilisateurs.


## Quels services utiliser avec Firebase ?

[Firebase](https://firebase.google.com) met à votre disposition différents services pouvant être répartis en deux catégories :
  - Les outils de développement et de test de son application
  - Les outils permettant d’augmenter et d’engager ses cibles.
  
 
 
# Un peu de pratique avec Firebase (Cloud Firestore)
 
 > **Le savoir n'est pas difficile, seule sa mise en pratique l'est.**
 > *-Proverbe chinois*
 
 Bien avant de passer à la phase pratique nous allons consulter les différentes offres que nous propose les services [Firebase](https://firebase.google.com/pricing/)

 
 ## Créer une base de données Firebase
  
  - Prérequis pour avoir accès à Firebase
    Il faut tout simplement disposer d'un compte *__google(Gmail)__*
    
 
  *N.B:* Pour toute la suite de cette session nous utiliserons l'offre gratuite de firebase, donc vous pourriez suivre sans aucun souci
 
  * Pour la création d'une base de données avec Firebase, il faut suivre les étapes suivantes:
    1. Se rendre sur le site [console.firebase.google.com](https://console.firebase.google.com/)
    2. Cliquer sur __+ ajouter un projet__(__+ add project__ en anglais)
    3. Donner un nom au projet puis continuer
    4. Décider d'activer les services **Google Analytics** ou pas pour ce projet
    5. Configurer **Google Analytics** puis créer le projet
    6. Une fois le projet créer, rendez-vous dans le menu latéral puis **Database**
    7. Cliquer sur le boutton **Créer une base de donnée**
    8. Sélectionner le mode de votre application
    9. Définir l'emplacement du **Cloud Firestore**
    10. Vous pouvez commencer à utiliser votre base de données (Vous pouvez choisir entre deux options de base de données **Cloud Firestore** ou **realtime database**)
  
  
 ## Intégration de firebase à son projet
 
  * Pour intégrer Firebase à son projet il faut suivre les étapes suivantes:
    1. Se rendre sur le site [console.firebase.google.com](https://console.firebase.google.com/)
    2. Sélectionner le projet
    3. Sélectionner l'application (**Web, IOS ou Android**)
    4. Suivre les instructions selon le types de votre projet et les besoin de votre projet
    
    
    
# Exemple

> **L’exemple n’est pas le meilleur moyen de convaincre, c’est le seul.**
> *-Gandhi*

Pour cet exemple nous allons nous inspirer du tutoriel sur [firestore-web](https://codelabs.developers.google.com/codelabs/firestore-web/). A la différence que cete fois nous allons essayer d'implementer avec d'autres langages ou framework autre que le javascript classique.

![Image du projet](https://github.com/firebase/friendlyeats-web/raw/master/docs/finished_image.png)

L'objectif de cette approche est de nous permettre d'avoir un aperçu sur plusieurs façon d'intégrer firebase à notre projet.
Les fichiers de base pour pouvoir suivre les différentes implémentations se situe dans `Base`
Pour démarrez commencer par cloner ce repository grâce, au gestionnaire de package [git](https://git-scm.com/downloads) si vous n'avez pas encore ça sur votre machine.

```bash
git clone https://github.com/GreyCoderK/FirebaseConference.git
```

Puis taper la commande précédente à l'interieur **d'un invite de commande** pour récuperer cela localement chez vous.

* Un aperçu des differents langages qui seront utilisés : 

  * [golang](./golang)
  * [nodejs](./nodejs)
  * [python](./python)



# D'autre sources

> **Avec le temps et la patience, la feuille du mûrier devient de la soie.**
> *-Proverbe chinois*

* [La documentation](https://firebase.google.com/docs/firestore/)
* [une playlist sur l'implementation web](https://www.youtube.com/watch?v=2Vf1D-rUMwE&list=PLl-K7zZEsYLmnJ_FpMOZgyg6XcIGBu2OX)
* [Un ensemble de ressource sur Firebase](https://github.com/jthegedus/awesome-firebase)



# Remerciements

> **Dans la vie, entourez vous des personnes qui illuminent votre chemin**
> *-plaisir des yeux sept hanou*

* *__@ The Day__* Pour tous ces enseignements et pour ce qu'il fais pour la communauté*
* *__@ EZFRICA__* Pour m'avoir donné cette occasion de vous parlez de Firebase*
* *__@ Donutson__* Pour les relectures, corrections et suggestions*
* *__@ Mr Mathématiques__* Pour le plan de ce document* 
