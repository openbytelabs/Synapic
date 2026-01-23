# Synapic

Synapic ist eine Open-Source-Metasuchmaschine, die den Datenschutz der Nutzer in den Vordergrund stellt und Daten aus mehreren Quellen sammelt. Alle Nutzerdaten werden lokal auf Ihrem Gerät gespeichert und keine persönlichen Daten werden an Dritte weitergegeben.

## Funktionen

- **Datenschutzorientiert**: Alle Daten werden mit localStorage auf Ihrem Gerät gespeichert
- **Unterstützung für mehrere Suchmaschinen**: Suche im Internet, nach Bildern, Videos, Nachrichten und Karten
- **Integration künstlicher Intelligenz**: Mehrere kostenlose KI-Modelle (Llama, Qwen, DeepSeek, Gemini) über OpenRouter
- **Mehrsprachige Unterstützung**: Unterstützung für mehr als 30 Sprachen und Regionen
- **Wetterintegration**: Standortbasierte Wetterinformationen
- **Moderne Benutzeroberfläche**: Reaktive Benutzeroberfläche, entwickelt mit Vue.js 3

## Technologie-Stack

- **Frontend**: Vue.js 3 (Composition API)
- **Backend**: Node.js + Express
- **Suchmaschine**: DuckDuckGo
- **KI-Dienst**: OpenRouter API
- **Wetter**: WeatherAPI
- **Standortbestimmung**: IPInfo

## Installation

### Anforderungen

- Node.js (v14 oder höher)
- Go (für Uptime-Backend)
- npm oder yarn

### Projektstruktur

Synapic besteht aus drei Hauptkomponenten:
- **consus**: Haupt-Backend-Dienst
- **frontend**: Benutzeroberfläche
- **uptime**: Systemüberwachungsdienst (Backend + Frontend)

### Installationsschritte

1. Klonen Sie das Projekt:
```bash
git clone https://github.com/openbytelabs/Synapic
cd Synapic
```

2. **Consus Backend** installieren:
```bash
cd consus
npm install
npm run start
```

3. **Frontend** installieren:
```bash
cd ../frontend
npm install
npm run serve
```

4. **Uptime Backend** installieren:
```bash
cd ../uptime/backend
go mod tidy
go run main.go
```

5. **Uptime Frontend** installieren:
```bash
cd ../frontend
npm install
npm run serve
```

## Verwendung

Geben Sie nach dem Öffnen von Synapic Ihre Suchanfrage in die Suchleiste ein und wählen Sie den gewünschten Suchtyp (Web, Bild, Video, Nachrichten oder KI). Ihre Sprach- und Regionseinstellungen werden automatisch gespeichert und bei Ihren nächsten Suchanfragen verwendet.

## Datenschutz

Synapic speichert keine personenbezogenen Daten auf seinen Servern. Der gesamte Suchverlauf, alle Einstellungen und Präferenzen werden ausschließlich im localStorage Ihres Browsers gespeichert. Sie können alle Daten jederzeit löschen, indem Sie den Cache Ihres Browsers leeren.

## Lizenz

Dieses Projekt ist unter der BSD 3-Clause-Lizenz lizenziert.

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

## Kontakt

Bei Fragen wenden Sie sich bitte an: yigitkabak@tuta.io

---

<div align="center">

### Vielen Dank

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
