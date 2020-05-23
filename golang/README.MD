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

