# Synapic Arama Motoru

Synapic, kullanıcı gizliliğini ön planda tutan, yapay zeka destekli sonuçlar sunan açık kaynak bir bağımsız arama motorudur. Go ve Vue.js ile geliştirilmiş olup, kendi arama indeksleme sistemine sahiptir ve web, görsel ve haber arama yetenekleri ile entegre AI yanıtları sunar.

## Özellikler

- **Gizlilik Odaklı**: Kullanıcı verisi toplama veya takip yok
- **Bağımsız İndeksleme**: Bluge ile özel tam metin arama motoru
- **Çoklu Arama Desteği**: Web, görsel ve haber araması
- **Yapay Zeka Entegrasyonu**: OpenRouter üzerinden 35'ten fazla ücretsiz AI modeli (Llama, Qwen, DeepSeek, Gemini ve daha fazlası)
- **API Kimlik Doğrulama**: Güvenli API anahtarı tabanlı erişim kontrolü
- **Hız Sınırlama**: Kullanıcı başına yapılandırılabilir günlük arama limitleri

## Teknoloji Yığını

- **Backend**: Bluge arama indeksleme ile Go
- **Frontend**: Vue.js 3 (Composition API)
- **Arama Motoru**: Bluge tam metin arama
- **AI Servisi**: 35'ten fazla ücretsiz model ile OpenRouter API
- **Depolama**: Dosya tabanlı indeks depolama

## Kurulum

### Gereksinimler

- Go 1.16 veya üzeri
- Node.js v14 veya üzeri
- OpenRouter API anahtarı (AI özellikleri için)

### Kurulum Adımları

1. Projeyi klonlayın:
```bash
git clone https://github.com/openbytelabs/Synapic
cd Synapic
```

5. Frontend'i kurun ve çalıştırın:
```bash
npm install
npm run serve
```

## API Kullanımı

AI destekli arama endpoint'i şu parametreleri kabul eder:

- `q`: Arama sorgusu (zorunlu)
- `apikey`: Kimlik doğrulama anahtarı (zorunlu)
- `type`: Arama türü - `web`, `image` veya `news` (varsayılan: `web`)

## Mimari

Synapic, şu bileşenlerle bağımsız bir arama motoru olarak çalışır:
- **Özel İndeks**: Bluge destekli tam metin arama indeksleri
- **Çoklu Tür Desteği**: Web, görseller ve haberler için ayrı indeksler
- **AI Katmanı**: Arama sonuçlarından akıllı yanıt üretimi
- **API Gateway**: Hız sınırlama ile kimlik doğrulamalı erişim

## Ünlem
Backend kodları(consus), kodları olası güvenlik sorunlarıyla karşılaşmamak için paylaşılmamıştır.

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

THIS SOFTWARE IS PROVIDED "AS IS".
```

## İletişim

Sorularınız için: yigitkabak@tuta.io

---

<div align="center">

### Bu teknolojiler kullanılarak sevgiyle geliştirildi ❤️

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

<a href="https://golang.org" style="text-decoration: none; border: none;">
 <img src="https://duckduckgo.com/i/77cac52781feeb4a.png" width="100" alt="Node.js" style="border: none;"/>
</a>

</div>
