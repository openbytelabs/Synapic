# Synapic

Synapic is an open-source meta search engine that prioritizes user privacy and collects data from multiple sources. All user data is stored locally on your device, and no personal information is shared with third parties.

## Features

- **Privacy-Focused**: All data is stored on your device using localStorage
- **Multi-Search Support**: Web, image, video, news, and map search
- **AI Integration**: Multiple free AI models (Llama, Qwen, DeepSeek, Gemini) via OpenRouter
- **Multi-Language Support**: Support for over 30 languages and regions
- **Weather Integration**: Location-based weather information
- **Modern Interface**: Reactive user interface developed with Vue.js 3

## Technology Stack

- **Frontend**: Vue.js 3 (Composition API)
- **Backend**: Node.js + Express
- **Search Engine**: DuckDuckGo
- **AI Service**: OpenRouter API
- **Weather**: WeatherAPI
- **Location Detection**: IPInfo

## Installation

### Requirements

- Node.js (v14 or higher)
- Go (for uptime backend)
- npm or yarn

### Project Structure

Synapic consists of three main components:
- **consus**: Main backend service
- **frontend**: User interface
- **uptime**: System monitoring service (backend + frontend)

### Installation Steps

1. Clone the project:
```bash
git clone https://github.com/openbytelabs/Synapic
cd Synapic
```

2. Install **Consus Backend**:
```bash
cd consus
npm install
npm run start
```

3. Install **Frontend**:
```bash
cd ../frontend
npm install
npm run serve
```

4. Install **Uptime Backend**:
```bash
cd ../uptime/backend
go mod tidy
go run main.go
```

5. Install **Uptime Frontend**:
```bash
cd ../frontend
npm install
npm run serve
```

## Usage

After opening Synapic, type your query into the search bar and select the type you want to search for (web, image, video, news, or AI). Your language and region preferences are automatically saved and used for your next searches.

## Privacy

Synapic does not store any personal data on its servers. All search history, preferences, and settings are stored only in your browser's localStorage. You can delete all data at any time by clearing your browser cache.

## License

This project is licensed under the BSD 3-Clause License.

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

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS “AS IS”
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

For questions: yigitkabak@tuta.io

---

<div align="center">

### Thanks

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
