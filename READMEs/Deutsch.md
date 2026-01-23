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

Die Weitergabe und Verwendung in Quell- und Binärform, mit oder ohne
Änderungen, ist unter folgenden Bedingungen gestattet:

1. Bei der Weitergabe des Quellcodes müssen der obenstehende Copyright-Hinweis, diese
   Liste der Bedingungen und der folgende Haftungsausschluss beibehalten werden.

2. Bei der Weitergabe in Binärform müssen der obenstehende Copyright-Hinweis,
   diese Liste der Bedingungen und der folgende Haftungsausschluss in der Dokumentation
   und/oder anderen mit der Verteilung bereitgestellten Materialien enthalten.

3. Weder der Name des Urheberrechtsinhabers noch die Namen seiner
   Mitwirkenden dürfen ohne vorherige ausdrückliche schriftliche Genehmigung zur Werbung oder Verkaufsförderung für Produkte verwendet werden, die aus
   dieser Software abgeleitet sind.

DIESE SOFTWARE WIRD VON DEN URHEBERRECHTSINHABERN UND MITWIRKENDEN „WIE BESEHEN“
UND JEGLICHE AUSDRÜCKLICHEN ODER STILLSCHWEIGENDEN GEWÄHRLEISTUNGEN, EINSCHLIESSLICH, ABER NICHT BESCHRÄNKT AUF DIE
STILLSCHWEIGENDEN GEWÄHRLEISTUNGEN DER MARKTGÄNGIGKEIT UND EIGNUNG FÜR EINEN BESTIMMTEN ZWECK,
WERDEN AUSGESCHLOSSEN. IN KEINEM FALL HAFTEN DER URHEBERRECHTSINHABER ODER DIE MITWIRKENDEN
FÜR DIREKTE, INDIREKTE, ZUFÄLLIGE, BESONDERE, EXEMPLARISCHE ODER FOLGESCHÄDEN
(EINSCHLIESSLICH, ABER NICHT BESCHRÄNKT AUF DIE BESCHAFFUNG VON ERSATZWAREN ODER
DIENSTLEISTUNGEN, NUTZUNGSAUSFALL, DATENVERLUST ODER ENTGANGENE GEWINNE ODER GESCHÄFTSUNTERBRECHUNGEN), UNABHÄNGIG VON
UND AUFGRUND JEGLICHER HAFTUNGSTHEORIE, OB AUS VERTRAG, VERSCHULDENSUNABHÄNGIGER HAFTUNG
ODER UNERLAUBTER HANDLUNG (EINSCHLIESSLICH FAHRLÄSSIGKEIT ODER ANDERWEITIG), DIE SICH IN IRGENDEINER WEISE AUS DER NUTZUNG
DIESER SOFTWARE ERGIBT, SELBST WENN AUF DIE MÖGLICHKEIT SOLCHER SCHÄDEN HINGEWIESEN WURDE.
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
