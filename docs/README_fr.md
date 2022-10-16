edgar
=

génère une page d'accueil pour n'importe quoi avec un readme. Un remplaçant pour mes absurdités.
Cet outil est destiné à créer des pages pour des projets basés sur le fichier README.md
J'espère que c'est particulièrement utile pour les pages github.

Fondamentalement, un générateur de site statique vraiment simple qui prennent un seul fichier markdown et émet
page HTML d'apparence raisonnable pour elle.

STATUS: Ce projet est maintenu. Je répondrai aux questions, aux demandes de retrait et aux demandes de fonctionnalités dans quelques jours. Il le fait
ce que c'est censé faire.

Utilisation
---

```md
Utilisation de edgar:
-author string
L'auteur du fichier HTML (par défaut eyedeekay)
-cs string
Le fichier CSS à utiliser, un par défaut sera généré si on n'existe pas (default style.css)
-donate string
ajouter la section de don aux portefeuilles de crypto-monnaie. Utilisez les schémas d'URL d'adresse, séparés par des virgules (aucun espace). Changez-les avant de courir à moins que vous ne vouliez l'argent pour aller à moi. (default monero:4A2BwLabGUiU65C5JRfwXqFTwWPYNSmuZRjbTDjsu9wT6wV6kMFyXn83ydnVjVcR7BCsWh8B5b4Z9b6cmqjfZiFd9sBUpWT,bitcoin:1D1sDmyZ5
-filename string
Le fichier markdown pour convertir en HTML, ou une liste séparée des fichiers (par défaut README.md,USAGE.md,index.html,docs/README.md)
-i2plink
ajouter un lien i2p à la page footer. Logo courtoisie de @Shoalsteed et @mark22k (par défaut vrai)
- non
désactiver la section de don (changer les adresses de portefeuille -donate avant de le mettre à true) (par défaut true)
- le fichier d'entrée. html
Le nom du fichier de sortie (Seulement utilisé pour le premier fichier, d'autres seront nommés inputfile.html) (default index.html)
-script string
Le fichier script à utiliser.
- snowflake
ajouter un flocons de neige au pied de page (par défaut vrai)
-support chaîne
changement de message/CTA pour la section des dons. (par défaut "Appuyer le développement indépendant de l'edgar")
-title string
Le titre du fichier HTML, si blanc, il sera généré à partir du premier h1 dans le fichier markdown.
```
