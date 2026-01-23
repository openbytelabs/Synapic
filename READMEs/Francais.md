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
BSD 3-Clause License

Copyright (c) 2023, openbytelabs

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution.

3. Neither the name of the copyright holder nor the names of its
   contributors may be used to endorse or promote products derived from
   this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
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
