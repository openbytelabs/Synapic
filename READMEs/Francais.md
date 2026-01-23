# Synapic

Synapic est un moteur de recherche métatextuel open source qui privilégie la confidentialité des utilisateurs et collecte des données provenant de plusieurs sources. Toutes les données utilisateur sont stockées localement sur votre appareil et aucune information personnelle n'est partagée avec des tiers.

## Caractéristiques

- **Axé sur la confidentialité** : toutes les données sont stockées sur votre appareil via localStorage
- **Prise en charge de la recherche multiple** : recherche sur le Web, d'images, de vidéos, d'actualités et de cartes
- **Intégration de l'intelligence artificielle** : plusieurs modèles d'IA gratuits (Llama, Qwen, DeepSeek, Gemini) via OpenRouter
- **Prise en charge multilingue** : prise en charge de plus de 30 langues et régions
- **Intégration de la météo** : informations météorologiques basées sur la localisation
- **Interface moderne** : interface utilisateur réactive développée avec Vue.js 3

## pile technologique

- **Frontend** : Vue.js 3 (Composition API)
- **Backend** : Node.js + Express
- **Moteur de recherche** : DuckDuckGo
- **Service IA** : OpenRouter API
- **Météo** : WeatherAPI
- **Localisation** : IPInfo

## Installation

### Configuration requise

- Node.js (v14 ou supérieure)
- Go (pour le backend uptime)
- npm ou yarn

### Structure du projet

Synapic se compose de trois éléments principaux :
- **consus** : service backend principal
- **frontend** : interface utilisateur
- **uptime** : service de surveillance du système (backend + frontend)

### Étapes d'installation

1. Clonez le projet :
```bash
git clone https://github.com/openbytelabs/Synapic
cd Synapic
```

2. Installation de **Consus Backend** :
```bash
cd consus
npm install
npm run start
```

3. Installation de **Frontend** :
```bash
cd ../frontend
npm install
npm run serve
```

4. Installation de **Uptime Backend** :
```bash
cd ../uptime/backend
go mod tidy
go run main.go
```

5. Installation de **Uptime Frontend** :
```bash
cd ../frontend
npm install
npm run serve
```

## Utilisation

Après avoir ouvert Synapic, saisissez votre requête dans la barre de recherche et sélectionnez le type de recherche souhaité (web, image, vidéo, actualités ou IA). Vos préférences en matière de langue et de région sont automatiquement enregistrées et utilisées pour vos prochaines recherches.

## Confidentialité

Synapic ne stocke aucune donnée personnelle sur ses serveurs. Tout l'historique de recherche, les préférences et les paramètres sont uniquement conservés dans le localStorage de votre navigateur. Vous pouvez supprimer toutes les données à tout moment en vidant le cache de votre navigateur.

## Licence

Ce projet est sous licence BSD 3-Clause.

```
Licence BSD 3-Clause

Copyright (c) 2023, openbytelabs

La redistribution et l'utilisation sous forme source et binaire, avec ou sans
modification, sont autorisées à condition que les conditions suivantes soient remplies :

1. Les redistributions du code source doivent conserver la mention de copyright ci-dessus, cette
   liste de conditions et la clause de non-responsabilité suivante.

2. Les redistributions sous forme binaire doivent reproduire la mention de copyright ci-dessus,
   cette liste de conditions et la clause de non-responsabilité suivante dans la documentation
   et/ou dans d'autres documents fournis avec la distribution.

3. Ni le nom du détenteur du copyright ni les noms de ses
   contributeurs ne peuvent être utilisés pour approuver ou promouvoir des produits dérivés de
   ce logiciel sans autorisation écrite préalable spécifique.

CE LOGICIEL EST FOURNI PAR LES DÉTENTEURS DU COPYRIGHT ET LES CONTRIBUTEURS « TEL QUEL »
ET TOUTE GARANTIE EXPRESSE OU IMPLICITE, Y COMPRIS, MAIS SANS S'Y LIMITER, LES
GARANTIES IMPLICITES DE QUALITÉ MARCHANDE ET D'ADÉQUATION À UN USAGE PARTICULIER, SONT
EXCLUES. EN AUCUN CAS, LE TITULAIRE DU DROIT D'AUTEUR OU LES CONTRIBUTEURS NE PEUVENT ÊTRE TENUS RESPONSABLES
DES DOMMAGES DIRECTS, INDIRECTS, ACCESSOIRES, SPÉCIAUX, EXEMPLAIRES OU CONSÉCUTIFS
(Y COMPRIS, MAIS SANS S'Y LIMITER, L'ACHAT DE BIENS OU DE SERVICES DE REMPLACEMENT,
LA PERTE D'UTILISATION, DE DONNÉES OU DE BÉNÉFICES, OU L'INTERRUPTION D'ACTIVITÉ) QUELLE QU'EN SOIT LA CAUSE
LA CAUSE ET SUR TOUTE THÉORIE DE RESPONSABILITÉ, QU'ELLE SOIT CONTRACTUELLE, STRICTE
OU DÉLICTUELLE (Y COMPRIS LA NÉGLIGENCE OU AUTRE) DÉCOULANT DE L'UTILISATION
DE CE LOGICIEL, MÊME SI LA POSSIBILITÉ DE TELS DOMMAGES A ÉTÉ SIGNALÉE.
```

## Contact

Pour toute question : yigitkabak@tuta.io

---

<div align="center">

### Remerciements

<a href="https://www.freebsd.org/" style="text-decoration: none; border: none;">
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/d/d5/BSD_wordmark.svg/330px-BSD_wordmark.svg.png?20111103140023" width="100" alt="BSD" style="border: none;"/>
</a>
&nbsp;&nbsp;&nbsp;
<a href="https://www.linux.org/" style="text-decoration: none; border: none;">
<img src="https://upload.wikimedia.org/wikipedia/commons/3/35/Tux.svg" width="80" alt="Linux" style="border: none;"/>
</a>
&nbsp;&nbsp;&nbsp;
<a href="https://vuejs.org/" style="text-decoration: none; border: none;">
<img src="https://upload.wikimedia.org/wikipedia/commons/9/95/Vue.js_Logo_2.svg" width="80" alt="Vue.js" style="border: none;"/>
</a>
&nbsp;&nbsp;&nbsp;
<a href="https://nodejs.org/" style="text-decoration: none; border: none;">
<img src="https://upload.wikimedia.org/wikipedia/commons/d/d9/Node.js_logo.svg" width="100" alt="Node.js" style="border: none;"/>
</a>

</div>
