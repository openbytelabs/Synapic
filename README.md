# Synapic

Synapic, kullanıcı gizliliğini ön planda tutan, birden fazla kaynaktan veri toplayan açık kaynak bir meta arama motorudur. Tüm kullanıcı verileri cihazınızda yerel olarak saklanır ve hiçbir kişisel bilgi üçüncü taraflarla paylaşılmaz.

## Özellikler

- **Gizlilik Odaklı**: Tüm veriler localStorage ile cihazınızda saklanır
- **Çoklu Arama Desteği**: Web, görsel, video, haber ve harita araması
- **Yapay Zeka Entegrasyonu**: OpenRouter üzerinden çoklu ücretsiz AI modelleri (Llama, Qwen, DeepSeek, Gemini)
- **Çok Dilli Destek**: 30'dan fazla dil ve bölge desteği
- **Hava Durumu Entegrasyonu**: Konum tabanlı hava durumu bilgisi
- **Modern Arayüz**: Vue.js 3 ile geliştirilmiş reaktif kullanıcı arayüzü

## Teknoloji Yığını

- **Frontend**: Vue.js 3 (Composition API)
- **Backend**: Node.js + Express
- **Arama Motoru**: DuckDuckGo
- **AI Servisi**: OpenRouter API
- **Hava Durumu**: WeatherAPI
- **Konum Tespiti**: IPInfo

## Kurulum

### Gereksinimler

- Node.js (v14 veya üzeri)
- Go (uptime backend için)
- npm veya yarn

### Proje Yapısı

Synapic üç ana bileşenden oluşur:
- **consus**: Ana backend servisi
- **frontend**: Kullanıcı arayüzü
- **uptime**: Sistem izleme servisi (backend + frontend)

### Kurulum Adımları

1. Projeyi klonlayın:
```bash
git clone https://github.com/openbytelabs/Synapic
cd Synapic
```

2. **Consus Backend** kurulumu:
```bash
cd consus
npm install
npm run start
```

3. **Frontend** kurulumu:
```bash
cd ../frontend
npm install
npm run serve
```

4. **Uptime Backend** kurulumu:
```bash
cd ../uptime/backend
go mod tidy
go run main.go
```

5. **Uptime Frontend** kurulumu:
```bash
cd ../frontend
npm install
npm run serve
```

## Kullanım

Synapic'i açtıktan sonra arama çubuğuna sorgunuzu yazın ve aramak istediğiniz türü seçin (web, görsel, video, haber veya AI). Dil ve bölge tercihleriniz otomatik olarak kaydedilir ve bir sonraki aramalarınızda kullanılır.

## Gizlilik

Synapic hiçbir kişisel veriyi sunucularda saklamaz. Tüm arama geçmişi, tercihler ve ayarlar yalnızca tarayıcınızın localStorage'ında tutulur. İstediğiniz zaman tarayıcı önbelleğinizi temizleyerek tüm verileri silebilirsiniz.

## Lisans

Bu proje BSD 3-Clause Lisansı altında lisanslanmıştır.

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

## İletişim

Sorularınız için: yigitkabak@tuta.io

---

<div align="center">

### Teşekkürler

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
