> __Un voyage de mille lieues a commencé par un pas.__
> *-Proverbe chinois*

# Firebase Conference

__Guide pour la conférence de Firebase, pour toutes les personnes qui assisteront à cette conférence dans l'optique que ceux-ci puissent revenir sur les propos de la conférence et aussi pratiquer par des exemples__


# Plan de cette session sur Firebase
  - Ce qu'il faut savoir sur Firebase
    - Historique de Firebase
    - Qu'est ce que Firebase ?
    - Philosophie de Firebase
    - Quels services utiliser avec Firebase ?
    - Les types de Base de Donnée que nous offre Firebase
      - Qu'est ce que FireStore ?
      - Qu'est ce que Realtime database?
      - Conclusion
    - Fonctionnement de FireStore
      - Qu'est ce que c'est le NoSQL ?
      - Plusieurs espèces de NoSQL
  - Un peu de pratique avec Firebase(FireStore)
    - Créer une base de données Firebase
    - Intégration de firebase à son projet
  - Exemple
    - Quelques notes avant de debuter
    - Objectif de notre approche
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
  
## Les types de Base de Donnée que nous offre Firebase
Firebase nous propose 2 types de base de donnée, à savoir :
  - Cloud Firestore 
  - Realtime Database
  
### Qu'est ce que FireStore ?
  
Cloud Firestore est une base de données de documents NoSQL sans serveur, cloud native, entièrement gérée et rapide. Ce service simplifie le stockage, la synchronisation et l'interrogation des données pour les applications Web, mobiles et IoT à l'échelle mondiale. Ses bibliothèques clientes permettent une synchronisation en direct et un fonctionnement hors connexion, tandis que ses fonctionnalités de sécurité et ses intégrations à Firebase et Google Cloud Platform (GCP) accélèrent le développement d'applications véritablement sans serveur.

### Qu'est ce que Realtime database?

Stockez et synchronisez les données avec notre base de données cloud NoSQL. Les données sont synchronisées sur tous les clients en temps réel et restent disponibles lorsque votre application se déconnecte.
La base de données en temps réel Firebase est une base de données hébergée dans le cloud. Les données sont stockées au format JSON et synchronisées en temps réel avec chaque client connecté. Lorsque vous créez des applications multiplates-formes avec nos SDK iOS, Android et JavaScript, tous vos clients partagent une instance de base de données en temps réel et reçoivent automatiquement des mises à jour avec les données les plus récentes

### conclusion

Pour la suite de cette section nous allons nous consacrer entièrement à *Cloud FireStore* et essayer de voir en détail les différentes fonctionnalités.

## Fonctionnement de FireStore

Comme nous l'avons souligné précédement Firebase est une base de donnée NoSQL. Donc nous allons chercher à savoir comment celle-ci fonctionne pour pouvoir comprendre le fonctionnement de FireStore.

### Qu'est ce que c'est le NoSQL ?
Le NoSQL, pour "not only SQL", désigne les bases de données qui ne sont pas fondées sur l'architecture classique des bases de données relationnelles. Développé à l'origine pour gérer du big data, l'utilisation de base de données NoSQL a explosée depuis quelques années. Mais qu'est-ce que réellement le NoSQL ?

### Plusieurs espèces de NoSQL

Lorsque l'on parle de NoSQL, on regroupe des systèmes de base de données qui ne sont pas relationnels, mais il faut savoir qu'il existe plusieurs types de bases de données "NoSQL".

  - Les bases *clef/valeur*, permettent de stocker des informations sous forme d'un couple clef/valeur où la valeur peut être une chaine de caractère, un entier ou un objet sérialisé. Une base de données [Redis](https://redis.io/). Ce type de base de données offre de très bonne performances par sa simplicité et peut même être utilisé pour stocker les sessions utilisateur ou le cache de votre site par exemple.

  - Les bases orientées *colonnes*, ressemble aux bases de données relationnelles, car les données sont sauvegardées sous forme de ligne avec des colonnes, mais se distingue par le fait que le nombre de colonnes peut varier d'une ligne à l'autre. Les solutions les plus connues sont [HBase](https://hbase.apache.org/) ou [Cassandra](https://cassandra.apache.org/).

  - Les bases orientées *document*, représente les informations sous forme d'objet XML ou JSON. L'avantage est de pouvoir récupérer simplement des informations structurées de manière hiérarchique. Les solutions les plus connues sont [CouchDB](https://couchdb.apache.org/), [RavenDB](https://ravendb.net/) et [MongoDB](https://www.mongodb.com/fr).

  - Les bases orientées *graphe*, présentent les données sous forme de noeud et de relation. Cette structure permet de récupérer simplement des relations complexes. Un exemple de base graphe est [Neo4J](https://neo4j.com/).

  *N.B* : FireStore est une base de donnée orientées document

  ![Firestore data organisation](https://static.javatpoint.com/tutorial/firebase/images/firebase-data-organization-in-firestore.png)

  - Explication du schéma

  Tous les documents doivent être stockés dans des collections. Les documents peuvent contenir des sous-collections et des objets imbriqués, qui peuvent inclure des champs primitifs comme des chaînes ou des objets complexes comme des listes.

  Un document est un enregistrement léger qui contient un champ, qui correspond aux valeurs. Un nom identifie chaque document et nous pouvons traiter les documents comme des enregistrements JSON légers.

  - Un aperçu de l'organisation de donnée dans Firebase
  ![Aperçu dans Firebase](https://miro.medium.com/max/2050/1*8WFTGIkbejzIYw2w8VBr2Q.png)


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

Pour cet exemple nous allons nous inspirer du tutoriel sur [firestore-web](https://codelabs.developers.google.com/codelabs/firestore-web/). A la différence que cette fois nous allons essayer d'implementer avec d'autres langages ou framework autre que le javascript classique.

![Image du projet](https://github.com/firebase/friendlyeats-web/raw/master/docs/finished_image.png)

## Quelques notes avant de debuter

Avant de commencer l'implémentation de dans d'autres langages, je vous conseille vivement de suivre ce [tutoriel](https://codelabs.developers.google.com/codelabs/firestore-web/) pour avoir un aperçu de l'utilisation de Firestore dans votre projet en utilisant le JavaScript classique. 

Nous soulignerons également que nous partons de 0 pour obtenir le resultat de ce tutoriel et cela dans les différent langages sélectionnés. Pour obtenir le resultat espèré nous nous atelons dans la mesure du possible à complèter chaque jour ce repertoire pour la conférence. Il pourrait être possible à l'instant  de la session que l'implémentation complète ne soit pas faite.

Dans l'exemple du tutoriel nous avons une Single Page Application (*SPA*), et la logique est entièrement gèrer depuis le javascript. Il pourrait être possible alors que nous fassions quelques changement majeur dans le code de base html qui lui proviens du tutoriel précédent. Donc nous pourrions peut-être passé d'un SPA à une application sur plusieurs pages.

## Objectif de notre approche

L'objectif de cette approche est de nous permettre d'avoir un aperçu sur plusieurs façon d'intégrer firebase à notre projet.
Les fichiers de base pour pouvoir suivre les différentes implémentations se situe dans `Base`
Pour démarrez commencer par cloner ce repository grâce, au gestionnaire de package [git](https://git-scm.com/downloads) si vous n'avez pas encore ça sur votre machine.

```bash
git clone https://github.com/GreyCoderK/FirebaseConference.git
```

Puis taper la commande précédente à l'interieur **d'un invite de commande** pour récuperer cela localement chez vous.

* Un aperçu des différents langages qui seront utilisés : 

  * [Golang](./golang)
  * [NodeJs](./nodejs)
  * [Python](./python)

* Explication du choix des langages :
  
  * _Golang_
  
    Pour répondre à ce besoin de simplicité, et ne pas avoir un langage qui empile les fonctionnalités pour faire plaisir au premier quidam, il se veut simple à lire, à comprendre et à prendre en main.
    Libre ensuite au développeur de penser son programme comme bon lui semble.
    Le projet Golang est à l’origine un projet interne chez Google dont le but est de résoudre les problèmes de montée en charge. Et, puisque ce problème est commun à beaucoup d’entreprises, le projet est devenu un langage à part entière avec le concept de concurrence asynchrone en point de mire.

  * _NodeJs_

    NodeJS est une plateforme qui s’appuie sur le moteur JavaScript de Chrome permettant de développer rapidement des applications réseau rapides et évolutives. NodeJS utilise un modèle événementiel, aux E/S non bloquantes qui le rend léger et efficace, idéal pour les applications gérant d’importants volumes de données en temps réel sur des dispositifs distribués.
    NodeJS utilise JavaScript, un langage très utilisé par les développeurs web dans les applications front-end, côté navigateur. L’interpréteur de syntaxe JavaScript NodeJS a été réécrit et compilé en C++ ce qui en fait un interpréteur JavaScript bien plus rapide que celui du navigateur.

  * _Python_

    Si vous débutez dans le codage, commencez par le langage Python, car il est puissant sans être trop compliqué. Python est un langage relativement nouveau. Il est donc plus simple que les anciens, ce qui le rend plus intuitif et plus rapide à maîtriser.
    À peu près tout ce que vous pouvez créer avec d’autres langages de programmation, tels que C ++ ou Ruby, peut être créé avec Python. Vous pouvez également faire de l’apprentissage machine, analyser des données, traiter le langage naturel, et bien d'autres choses encore.


# D'autre sources

> **La sagesse des siècles nous apprend qu'il suffit d'approfondir une chose pour en connaître plusieurs autres.**
> *-Dan Millman*

* [une playlist sur l'implementation web](https://www.youtube.com/watch?v=2Vf1D-rUMwE&list=PLl-K7zZEsYLmnJ_FpMOZgyg6XcIGBu2OX)
* [Un ensemble de ressource sur Firebase](https://github.com/jthegedus/awesome-firebase)



# Remerciements

> **Dans la vie, entourez vous des personnes qui illuminent votre chemin**
> *-plaisir des yeux sept hanou*

* @The_Day *Pour tous ces enseignements et pour ce qu'il fais pour la communauté*
* @EZFRICA *Pour m'avoir donné cette occasion de vous parlez de Firebase*
* @Donutson *Pour les relectures, corrections et suggestions*
* @Mr_Mathematiques *Pour le plan de ce document* 
