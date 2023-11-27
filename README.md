# PUSH SWAP
Bienvenue dans le projet <font color = "orange"> PUSH SWAP </font> ! <br>
Si c'est votre première fois ici, nous sommes heureux de vous accueillir !<br> Nous espérons que vous trouverez ce projet utile. Nous vous encourageons à contribuer, poser des questions et collaborer avec nous.
## Table des matières

- [Description](#description)
- [Auteurs](#auteurs)
- [Usage](#usage)
- [Détails de l'implémentation : Algorithme](#algorithm)
- [Contributions](#contributions)
- [Licence](#licence)

## Description

Push-Swap est un projet qui utilise un Algorithme de Tri Non Comparatif. Nous avons à notre disposition une liste de int valeurs, deux piles (a et b) et un ensemble d'instructions.<br>
le projet est divisé en deux programmes :<br>
`push-swap :` qui calcule et affiche sur la sortie standard le plus petit programme en utilisant le langage d'instruction push-swap qui trie les arguments entiers reçus.<br>
`Checker :` qui prend des arguments entiers et lit des instructions sur l'entrée standard. Une fois lus, le vérificateur les exécute et les affiche OK si les entiers sont triés. Sinon, il s'affichera KO.
Comme dit précédemment, nous aurons à notre disposition deux piles . Notre objectif est de trier la pile a, par ordre croissant, en utilisant les deux piles et un ensemble d'instructions.

Voici les instructions que vous pouvez utiliser pour trier la pile :

- pa poussez le premier élément supérieur de la pile b se mettre en pile a
- pb poussez le premier élément supérieur de la pile a se mettre en pile b
- sa permuter les 2 premiers éléments de la pile a
- sb permuter les 2 premiers éléments de la pile b
- ss exécuter sa et sb
- ra tourner la pile a (décalage de tous les éléments de la pile a par 1, le premier élément devient le dernier)
- rb tourner la pile b
- rr exécuter ra et rb
- rra rotation inverse a ( déplacer tous les éléments de la pile a par 1, le dernier élément devient le premier )
- rrb rotation inverse b
- rrr exécuter rra et rrb


## Auteurs

- [edieng](https://learn.zone01dakar.sn/git/edieng)
- [jgoudiab](https://learn.zone01dakar.sn/git/jgoudiab)
- [babacandiaye](https://learn.zone01dakar.sn/git/babacandiaye)

## Usage 

Avant d'utiliser ce programme, assurez-vous d'avoir Go installé sur votre système. Pour l'installer, suivez les instructions sur le site officiel de Go : [https://golang.org/doc/install](https://golang.org/doc/install).

Une fois Go installé, suivez ces étapes pour installer ce programme :

1. Clonez ce dépôt sur votre machine :

   ```shell
   git clone https://learn.zone01dakar.sn/git/edieng/push-swap.git
   ```
2. Accédez au répertoire du projet :

   ```shell
   cd push-swap
   ```
3. Exécutez le push-swap avec la commande suivante :

   ```shell
   ./push-swap "<nombres>"
   ```
4. Exécutez le checker avec la commande suivante :

   ```shell
   ./checker "<nombres>"
   ```
## Détails de l'implémentation : Algorithme

**Structure du projet :**
   
```
├── push-swap/
│   ├── push-swap/
│   │   ├──
|   |   ├──
│   │   ├──
|   |   ├── 
│   ├── checker/
│   │   ├──
|   |   ├──
|   |   ├──
|   |   ├──
|   ├──main.go
|   ├──go.mod
|   ├──README.md
├── - - - -
```
**Implémentation technique :**

**fonctionnement du Push-Swap :**

l'algorithme verifie d'abord la valadité des arguments ensuite pour optimiser le traitement on a eu à le subdiviser en des lots (3, 6 et 50) mais le fonctonnement reste la même l'algo cherche d'abord si un des plus petit nombre est à la première position si c'est le cas il verifie si c'est nécessaire d'effectuer un `pb` dans le cas inverse il cherche si un des plus petit nombre est proche du haut du pile `ra` ou il est plus proche du bas `rra`. <br>
Si pile A est ordonnée il rammene tous les nombres étant sur pile b

**fonctionnement du Checker :**
Le checker verifie des instructions sur l'entrée standard, chaque instruction sera suivie par \n. Une fois que toutes les instructions ont été lues, checker les exécutera sur la pile reçue en argument.<br>
Si après avoir exécuté ces instructions, empilez a est en fait triée et b est vide, alors le vérificateur doit afficher "OK" suivi d'un \n sur la sortie standard. Dans tous les autres cas, le vérificateur doit afficher "KO" suivi d'un \n sur la sortie standard.<br>
En cas d'erreur, vous devez afficher Erreur suivie d'une \n sur l'erreur standard. Les erreurs incluent par exemple: certains arguments ne sont pas des entiers, il y a des doublons, une instruction n'existe pas et/ou est mal formatée.<br>
Dans le cas où il n'y a pas d'arguments, le programme n'affiche rien.
## Contributions

Les contributions à ce projet sont les bienvenues ! Si vous souhaitez contribuer, veuillez ouvrir une demande d'extraction (pull request) et nous examinerons vos modifications.

## Licence

Ce projet est sous licence [zone01](LICENSE).