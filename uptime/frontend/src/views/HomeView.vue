<template>
  <div id="app">
    <main class="main">
      <div class="container">
        <div class="language-selector">
          <select 
            class="lang-select" 
            :value="currentLang"
            @change="changeLang($event.target.value)">
            <option value="tr">TR</option>
            <option value="en">EN</option>
          </select>
        </div>

        <div class="status-header">
          <div class="status-circle" :class="allOperational"></div>
          <h1 class="status-title">{{ translations[currentLang].statusTitle[allOperational] }}</h1>
          <p class="status-time">{{ currentTime }}</p>
        </div>

        <div class="services">
          <div class="service" v-for="site in sites" :key="site.url">
            <div class="service-header" @click="toggleChart(site)">
              <div class="service-info">
                <div class="service-name">{{ site.name }}</div>
                <div class="service-url">{{ site.url }}</div>
              </div>
              <div class="service-meta">
                <div class="service-stats">
                  <span class="stat">{{ site.ping }}ms</span>
                  <span class="stat">{{ site.uptime.toFixed(1) }}%</span>
                </div>
                <div class="pie-chart-container" @mouseenter="showPieTooltip(site, $event)" @mouseleave="hidePieTooltip(site)">
                  <svg width="32" height="32" viewBox="0 0 36 36" class="pie-chart">
                    <circle cx="18" cy="18" r="15.915" fill="transparent" stroke="#f87171" stroke-width="3.5"></circle>
                    <circle 
                      cx="18" 
                      cy="18" 
                      r="15.915" 
                      fill="transparent"
                      stroke="#5cdd8b"
                      stroke-width="3.5"
                      :stroke-dasharray="`${site.uptime} ${100 - site.uptime}`"
                      stroke-dashoffset="25"
                      transform="rotate(-90 18 18)"
                    ></circle>
                  </svg>
                  <div v-if="site.pieTooltip.show" class="pie-tooltip">
                    <div class="pie-tooltip-item">
                      <span class="pie-tooltip-color" style="background-color: #5cdd8b;"></span>
                      <span>{{ translations[currentLang].uptime }}: {{ site.uptime.toFixed(1) }}%</span>
                    </div>
                    <div class="pie-tooltip-item">
                      <span class="pie-tooltip-color" style="background-color: #f87171;"></span>
                      <span>{{ translations[currentLang].downtime }}: {{ (100 - site.uptime).toFixed(1) }}%</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <div class="uptime-history">
              <div class="history-label">{{ translations[currentLang].days90 }}</div>
              <div class="history-bars">
                <div 
                  v-for="(bar, index) in site.uptimeBars" 
                  :key="index" 
                  class="history-bar" 
                  :class="bar.status"
                  :title="bar.date + ' - ' + bar.statusText">
                </div>
              </div>
              <div class="history-label">{{ translations[currentLang].today }}</div>
            </div>

            <transition name="expand">
              <div v-if="site.chartExpanded" class="chart-container">
                <div class="chart-header">
                  <div class="chart-title">{{ translations[currentLang].responseTime }}</div>
                  <div class="time-buttons">
                    <button 
                      class="time-btn" 
                      :class="{ active: site.timeRange === 'minute' }"
                      @click="changeTimeRange(site, 'minute')">
                      5{{ translations[currentLang].min }}
                    </button>
                    <button 
                      class="time-btn" 
                      :class="{ active: site.timeRange === 'hour' }"
                      @click="changeTimeRange(site, 'hour')">
                      1{{ translations[currentLang].hour }}
                    </button>
                    <button 
                      class="time-btn" 
                      :class="{ active: site.timeRange === 'day' }"
                      @click="changeTimeRange(site, 'day')">
                      24{{ translations[currentLang].hour }}
                    </button>
                  </div>
                </div>
                <div class="chart-wrapper">
                  <canvas 
                    :id="'chart-' + site.url.replace(/[^a-z0-9]/gi, '')"
                    @mousemove="handleMouseMove($event, site)"
                    @mouseleave="handleMouseLeave(site)">
                  </canvas>
                  <div 
                    v-if="site.tooltip.show" 
                    class="chart-tooltip" 
                    :style="{ left: site.tooltip.x + 'px', top: site.tooltip.y + 'px' }">
                    <div class="tooltip-time">{{ site.tooltip.time }}</div>
                    <div class="tooltip-value">{{ site.tooltip.value }}ms</div>
                  </div>
                </div>
              </div>
            </transition>
          </div>
        </div>
      </div>
    </main>

    <footer class="footer">
      <div class="container">
        <span class="footer-copyright">openbyte all rights reserved</span>
      </div>
    </footer>
  </div>
</template>

<script>
import { ref, onMounted, onUnmounted, computed } from 'vue';

export default {
  name: 'App',
  setup() {
    const currentLang = ref('tr');
    const sites = ref([
      {
        url: 'https://synapic.com.tr',
        name: 'Synapic',
        status: 'unknown',
        ping: 0,
        uptime: 100,
        uptimeBars: [],
        chartData: [],
        timeRange: 'minute',
        chartExpanded: false,
        tooltip: { show: false, x: 0, y: 0, time: '', value: '' },
        pieTooltip: { show: false, x: 0, y: 0 }
      },
      {
        url: 'https://api.synapic.com.tr/api/search?q=test',
        name: 'Synapic API',
        status: 'unknown',
        ping: 0,
        uptime: 100,
        uptimeBars: [],
        chartData: [],
        timeRange: 'minute',
        chartExpanded: false,
        tooltip: { show: false, x: 0, y: 0, time: '', value: '' },
        pieTooltip: { show: false, x: 0, y: 0 }
      }
    ]);

    const currentTime = ref('');
    let updateInterval = null;
    const API_URL = 'https://uptimeapi.synapic.com.tr/api';
    const API_KEY = 'synapicsearch';

    const translations = {
      en: {
        statusTitle: {
          operational: 'All Systems Operational',
          degraded: 'Degraded Performance',
          down: 'System Down'
        },
        status: {
          operational: 'Operational',
          down: 'Down'
        },
        days90: '90 Days',
        today: 'Today',
        responseTime: 'Response Time',
        min: 'm',
        hour: 'h',
        uptime: 'Uptime',
        downtime: 'Downtime'
      },
      tr: {
        statusTitle: {
          operational: 'Tüm Sistemler Çalışıyor',
          degraded: 'Düşük Performans',
          down: 'Sistem Çalışmıyor'
        },
        status: {
          operational: 'Çalışıyor',
          down: 'Çalışmıyor'
        },
        days90: '90 Gün',
        today: 'Bugün',
        responseTime: 'Yanıt Süresi',
        min: 'dk',
        hour: 's',
        uptime: 'Çalışma',
        downtime: 'Kesinti'
      }
    };

    const allOperational = computed(() => {
      const hasData = sites.value.some(site => site.status !== 'unknown');
      if (!hasData) {
        return 'operational';
      }
      
      const hasDown = sites.value.some(site => site.status !== 'online' && site.status !== 'unknown');
      const hasDegraded = sites.value.some(site => site.uptime < 99 && site.uptime >= 95);
      
      if (hasDown) return 'down';
      if (hasDegraded) return 'degraded';
      return 'operational';
    });

    const showPieTooltip = (site, event) => {
      site.pieTooltip = {
        show: true,
        x: 0,
        y: 0
      };
    };

    const hidePieTooltip = (site) => {
      site.pieTooltip.show = false;
    };

    const updateCurrentTime = () => {
      const now = new Date();
      const days = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
      const months = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
      
      const dayName = days[now.getDay()];
      const monthName = months[now.getMonth()];
      const day = now.getDate();
      const year = now.getFullYear();
      
      let hours = now.getHours();
      const minutes = now.getMinutes().toString().padStart(2, '0');
      const ampm = hours >= 12 ? 'PM' : 'AM';
      hours = hours % 12;
      hours = hours ? hours : 12;
      const hoursStr = hours.toString().padStart(2, '0');
      
      currentTime.value = `${dayName}, ${monthName} ${day}, ${year}, ${hoursStr}:${minutes} ${ampm}`;
    };

    const changeLang = (lang) => {
      currentLang.value = lang;
      localStorage.setItem('statusLang', lang);
    };

    const generateUptimeBars = () => {
      const bars = [];
      const now = new Date();
      for (let i = 89; i >= 0; i--) {
        const date = new Date(now - i * 24 * 60 * 60 * 1000);
        bars.push({
          date: date.toLocaleDateString(),
          status: 'bar-operational',
          statusText: 'Operational'
        });
      }
      return bars;
    };

    const toggleChart = (site) => {
      site.chartExpanded = !site.chartExpanded;
      if (site.chartExpanded && site.chartData.length === 0) {
        setTimeout(() => {
          fetchMetrics(site);
        }, 100);
      } else if (site.chartExpanded) {
        setTimeout(() => {
          createChart(site);
        }, 100);
      }
    };

    const getCanvasId = (url) => {
      return 'chart-' + url.replace(/[^a-z0-9]/gi, '');
    };

    const drawChart = (canvas, data) => {
      if (!canvas || !data || data.length === 0) return;

      const ctx = canvas.getContext('2d');
      const rect = canvas.getBoundingClientRect();
      canvas.width = rect.width * 2;
      canvas.height = rect.height * 2;
      ctx.scale(2, 2);

      const width = rect.width;
      const height = rect.height;
      const padding = { top: 20, right: 20, bottom: 40, left: 50 };
      const chartWidth = width - padding.left - padding.right;
      const chartHeight = height - padding.top - padding.bottom;

      ctx.clearRect(0, 0, width, height);

      const maxPing = Math.max(...data.map(d => d.y), 100);
      const roundedMax = Math.ceil(maxPing / 100) * 100;

      ctx.strokeStyle = 'rgba(255, 255, 255, 0.05)';
      ctx.lineWidth = 1;
      for (let i = 0; i <= 4; i++) {
        const y = padding.top + (chartHeight / 4) * i;
        ctx.beginPath();
        ctx.moveTo(padding.left, y);
        ctx.lineTo(width - padding.right, y);
        ctx.stroke();
      }

      ctx.fillStyle = 'rgba(255, 255, 255, 0.5)';
      ctx.font = '11px sans-serif';
      ctx.textAlign = 'right';
      for (let i = 0; i <= 4; i++) {
        const value = roundedMax - (roundedMax / 4) * i;
        const y = padding.top + (chartHeight / 4) * i;
        ctx.fillText(Math.round(value), padding.left - 8, y + 4);
      }

      const barWidth = chartWidth / data.length;
      const barGap = Math.max(barWidth * 0.2, 1);
      const actualBarWidth = Math.max(barWidth - barGap, 1);

      data.forEach((point, index) => {
        const barHeight = (point.y / roundedMax) * chartHeight;
        const x = padding.left + index * barWidth + barGap / 2;
        const y = padding.top + chartHeight - barHeight;

        ctx.fillStyle = '#5cdd8b';
        ctx.fillRect(x, y, actualBarWidth, barHeight);
      });

      ctx.fillStyle = 'rgba(255, 255, 255, 0.4)';
      ctx.font = '10px sans-serif';
      ctx.textAlign = 'center';
      const timeStep = Math.max(Math.ceil(data.length / 6), 1);
      data.forEach((point, index) => {
        if (index % timeStep === 0 || index === data.length - 1) {
          const x = padding.left + index * barWidth + barWidth / 2;
          const date = new Date(point.x);
          let time;
          
          if (data.length <= 24) {
            const day = date.getDate();
            const month = date.getMonth() + 1;
            const hours = date.getHours();
            time = `${day}/${month} ${hours.toString().padStart(2, '0')}:00`;
          } else {
            time = date.toLocaleTimeString('en-US', { 
              hour: '2-digit', 
              minute: '2-digit'
            });
          }
          ctx.fillText(time, x, height - 20);
        }
      });
    };

    const createChart = (site) => {
      const canvasId = getCanvasId(site.url);
      const canvas = document.getElementById(canvasId);
      
      if (!canvas) return;
      if (!site.chartData || site.chartData.length === 0) return;

      drawChart(canvas, site.chartData);
    };

    const changeTimeRange = (site, range) => {
      site.timeRange = range;
      fetchMetrics(site);
    };

    const handleMouseMove = (event, site) => {
      const canvas = event.target;
      const rect = canvas.getBoundingClientRect();
      const x = event.clientX - rect.left;
      const y = event.clientY - rect.top;

      const padding = { top: 20, right: 20, bottom: 40, left: 50 };
      const chartWidth = rect.width - padding.left - padding.right;
      const chartHeight = rect.height - padding.top - padding.bottom;

      if (x < padding.left || x > rect.width - padding.right || 
          y < padding.top || y > rect.height - padding.bottom) {
        site.tooltip.show = false;
        return;
      }

      if (!site.chartData || site.chartData.length === 0) return;

      const barWidth = chartWidth / site.chartData.length;
      const index = Math.floor((x - padding.left) / barWidth);

      if (index >= 0 && index < site.chartData.length) {
        const point = site.chartData[index];
        const time = new Date(point.x).toLocaleTimeString('en-US', { 
          hour: '2-digit', 
          minute: '2-digit',
          second: '2-digit'
        });
        
        site.tooltip = {
          show: true,
          x: event.clientX - rect.left,
          y: event.clientY - rect.top - 60,
          time: time,
          value: Math.round(point.y)
        };
      }
    };

    const handleMouseLeave = (site) => {
      site.tooltip.show = false;
    };

    const fetchStats = async () => {
      try {
        const response = await fetch(`${API_URL}/stats?apiKey=${API_KEY}`);
        
        if (response.ok) {
          const data = await response.json();
          
          sites.value.forEach(site => {
            const siteData = data.find(d => d.url === site.url);
            if (siteData) {
              site.ping = siteData.ping;
              site.status = siteData.status;
              site.uptime = siteData.uptime;
            }
          });
        }
      } catch (error) {
        console.error('Stats fetch error:', error);
      }
    };

    const fetchUptimeHistory = async () => {
      try {
        const response = await fetch(`${API_URL}/uptime-history?apiKey=${API_KEY}`);
        
        if (response.ok) {
          const data = await response.json();
          
          sites.value.forEach(site => {
            const siteHistory = data.find(d => d.url === site.url);
            if (siteHistory && siteHistory.history) {
              const bars = [];
              const now = new Date();
              
              for (let i = 89; i >= 0; i--) {
                const date = new Date(now - i * 24 * 60 * 60 * 1000);
                const dateStr = date.toISOString().split('T')[0];
                const dayData = siteHistory.history[dateStr];
                
                let status = 'bar-operational';
                let statusText = 'Operational';
                
                if (dayData) {
                  if (dayData.uptime < 95) {
                    status = 'bar-down';
                    statusText = 'Down';
                  } else if (dayData.uptime < 99) {
                    status = 'bar-degraded';
                    statusText = 'Degraded';
                  }
                }
                
                bars.push({
                  date: date.toLocaleDateString(),
                  status: status,
                  statusText: statusText
                });
              }
              
              site.uptimeBars = bars;
            } else {
              site.uptimeBars = generateUptimeBars();
            }
          });
        } else {
          sites.value.forEach(site => {
            if (site.uptimeBars.length === 0) {
              site.uptimeBars = generateUptimeBars();
            }
          });
        }
      } catch (error) {
        console.error('Uptime history fetch error:', error);
        sites.value.forEach(site => {
          if (site.uptimeBars.length === 0) {
            site.uptimeBars = generateUptimeBars();
          }
        });
      }
    };

    const fetchMetrics = async (site) => {
      try {
        const url = `${API_URL}/metrics?apiKey=${API_KEY}&range=${site.timeRange}&site=${encodeURIComponent(site.url)}`;
        const response = await fetch(url);
        
        if (response.ok) {
          const data = await response.json();
          
          if (data.length > 0) {
            let chartData;

            if (site.timeRange === 'minute') {
              chartData = data.map(d => ({
                x: d.timestamp,
                y: d.ping
              }));
            } else if (site.timeRange === 'hour') {
              const hourlyData = {};
              data.forEach(d => {
                const date = new Date(d.timestamp);
                const hourKey = new Date(date.getFullYear(), date.getMonth(), date.getDate(), date.getHours()).getTime();
                
                if (!hourlyData[hourKey]) {
                  hourlyData[hourKey] = { sum: 0, count: 0 };
                }
                hourlyData[hourKey].sum += d.ping;
                hourlyData[hourKey].count += 1;
              });

              chartData = Object.keys(hourlyData).map(key => ({
                x: parseInt(key),
                y: Math.round(hourlyData[key].sum / hourlyData[key].count)
              })).sort((a, b) => a.x - b.x);
            } else if (site.timeRange === 'day') {
              const hourlyData = {};
              data.forEach(d => {
                const date = new Date(d.timestamp);
                const hourKey = new Date(date.getFullYear(), date.getMonth(), date.getDate(), date.getHours()).getTime();
                
                if (!hourlyData[hourKey]) {
                  hourlyData[hourKey] = { sum: 0, count: 0 };
                }
                hourlyData[hourKey].sum += d.ping;
                hourlyData[hourKey].count += 1;
              });

              chartData = Object.keys(hourlyData).map(key => ({
                x: parseInt(key),
                y: Math.round(hourlyData[key].sum / hourlyData[key].count)
              })).sort((a, b) => a.x - b.x);
            }
            
            site.chartData = chartData;
            createChart(site);
          }
        }
      } catch (error) {
        console.error('Metrics fetch error:', error);
      }
    };

    const updateData = async () => {
      await fetchStats();
      await fetchUptimeHistory();
      sites.value.forEach(site => {
        if (site.chartExpanded) {
          fetchMetrics(site);
        }
      });
    };

    onMounted(() => {
      const savedLang = localStorage.getItem('statusLang');
      if (savedLang) {
        currentLang.value = savedLang;
      }
      
      updateCurrentTime();
      setInterval(updateCurrentTime, 1000);
      
      fetchStats();
      fetchUptimeHistory();
      
      updateInterval = setInterval(updateData, 30000);
    });

    onUnmounted(() => {
      if (updateInterval) {
        clearInterval(updateInterval);
      }
    });

    return {
      sites,
      currentTime,
      currentLang,
      translations,
      allOperational,
      toggleChart,
      changeTimeRange,
      handleMouseMove,
      handleMouseLeave,
      changeLang,
      showPieTooltip,
      hidePieTooltip
    };
  }
};
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

#app {
  background-color: #000000;
  color: white;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, sans-serif;
}

.container {
  max-width: 1000px;
  margin: 0 auto;
  padding: 0 1.5rem;
}

.main {
  flex: 1;
  padding: 3rem 0;
}

.language-selector {
  position: absolute;
  top: 1.5rem;
  right: 1.5rem;
  z-index: 10;
}

.lang-select {
  background-color: #1a1a1a;
  color: #fff;
  border: 1px solid #2a2a2a;
  padding: 0.5rem 2rem 0.5rem 1rem;
  border-radius: 0.375rem;
  font-size: 0.875rem;
  font-weight: 600;
  cursor: pointer;
  appearance: none;
  background-image: url("data:image/svg+xml;charset=UTF-8,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 24 24' fill='none' stroke='white' stroke-width='2' stroke-linecap='round' stroke-linejoin='round'%3e%3cpolyline points='6 9 12 15 18 9'%3e%3c/polyline%3e%3c/svg%3e");
  background-repeat: no-repeat;
  background-position: right 0.7rem center;
  background-size: 1em;
  outline: none;
  min-width: 80px;
  transition: all 0.2s;
}

.lang-select:hover {
  background-color: #2a2a2a;
  border-color: #5cdd8b;
}

.lang-select:focus {
  border-color: #5cdd8b;
}

.lang-select option {
  background-color: #1a1a1a;
  color: white;
}

.status-header {
  text-align: center;
  margin-bottom: 3rem;
}

.status-circle {
  width: 80px;
  height: 80px;
  border-radius: 50%;
  margin: 0 auto 1.5rem;
  animation: pulse 2s ease-in-out infinite;
}

@keyframes pulse {
  0%, 100% {
    transform: scale(1);
    opacity: 1;
  }
  50% {
    transform: scale(1.1);
    opacity: 0.8;
  }
}

.status-circle.operational {
  background-color: #5cdd8b;
}

.status-circle.degraded {
  background-color: #fbbf24;
}

.status-circle.down {
  background-color: #f87171;
}

.status-title {
  font-size: 1.75rem;
  font-weight: 700;
  margin-bottom: 0.5rem;
}

.status-time {
  color: #888;
  font-size: 0.875rem;
}

.services {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.service {
  background-color: #0a0a0a;
  border: 1px solid #1a1a1a;
  border-radius: 0.5rem;
  overflow: visible;
}

.service-header {
  padding: 1.25rem 1.5rem;
  cursor: pointer;
  transition: background-color 0.2s;
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;
  border-radius: 0.5rem 0.5rem 0 0;
}

.service-header:hover {
  background-color: #0f0f0f;
}

.service-info {
  flex: 1;
  min-width: 0;
}

.service-name {
  font-size: 1rem;
  font-weight: 600;
  margin-bottom: 0.25rem;
}

.service-url {
  font-size: 0.75rem;
  color: #666;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.service-meta {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.service-stats {
  display: flex;
  gap: 0.75rem;
  font-size: 0.875rem;
  color: #888;
}

.stat {
  white-space: nowrap;
}

.pie-chart-container {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
}

.pie-chart {
  display: block;
  filter: drop-shadow(0 0 4px rgba(92, 221, 139, 0.3));
  border-radius: 50%;
}

.pie-tooltip {
  position: absolute;
  background-color: rgba(0, 0, 0, 0.95);
  border: 1px solid #2a2a2a;
  border-radius: 0.375rem;
  padding: 0.5rem 0.75rem;
  pointer-events: none;
  z-index: 100;
  bottom: calc(100% + 10px);
  left: 50%;
  transform: translateX(-50%);
  font-size: 0.75rem;
  white-space: nowrap;
  min-width: 150px;
}

.pie-tooltip-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.25rem;
}

.pie-tooltip-item:last-child {
  margin-bottom: 0;
}

.pie-tooltip-color {
  width: 12px;
  height: 12px;
  border-radius: 2px;
  flex-shrink: 0;
}

.uptime-history {
  padding: 0.75rem 1.5rem 1rem;
  display: flex;
  align-items: center;
  gap: 0.75rem;
  border-radius: 0 0 0.5rem 0.5rem;
}

.history-label {
  font-size: 0.75rem;
  color: #666;
  min-width: 50px;
}

.history-label:first-child {
  text-align: right;
}

.history-label:last-child {
  text-align: left;
}

.history-bars {
  flex: 1;
  display: grid;
  grid-template-columns: repeat(90, 1fr);
  gap: 2px;
}

.history-bar {
  height: 32px;
  border-radius: 2px;
  cursor: pointer;
  transition: opacity 0.2s;
}

.history-bar:hover {
  opacity: 0.7;
}

.bar-operational {
  background-color: #5cdd8b;
}

.bar-degraded {
  background-color: #fbbf24;
}

.bar-down {
  background-color: #f87171;
}

.expand-enter-active, .expand-leave-active {
  transition: all 0.3s ease;
  max-height: 400px;
}

.expand-enter-from, .expand-leave-to {
  max-height: 0;
  opacity: 0;
}

.chart-container {
  border-top: 1px solid #1a1a1a;
  padding: 1.5rem;
  border-radius: 0 0 0.5rem 0.5rem;
}

.chart-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.chart-title {
  font-size: 0.875rem;
  font-weight: 600;
  color: #888;
}

.time-buttons {
  display: flex;
  gap: 0.5rem;
}

.time-btn {
  background-color: #1a1a1a;
  color: #888;
  border: none;
  padding: 0.375rem 0.75rem;
  border-radius: 0.25rem;
  font-size: 0.75rem;
  cursor: pointer;
  transition: all 0.2s;
}

.time-btn:hover {
  background-color: #2a2a2a;
}

.time-btn.active {
  background-color: #5cdd8b;
  color: black;
}

.chart-wrapper {
  position: relative;
  height: 220px;
  background-color: #0a0a0a;
  border-radius: 0.5rem;
  padding: 1rem;
}

.chart-wrapper canvas {
  width: 100% !important;
  height: 100% !important;
  cursor: crosshair;
}

.chart-tooltip {
  position: absolute;
  background-color: rgba(0, 0, 0, 0.95);
  border: 1px solid #2a2a2a;
  border-radius: 0.375rem;
  padding: 0.5rem 0.75rem;
  pointer-events: none;
  z-index: 10;
  transform: translateX(-50%);
  font-size: 0.75rem;
}

.tooltip-time {
  color: #888;
  margin-bottom: 0.25rem;
}

.tooltip-value {
  color: #5cdd8b;
  font-weight: 600;
}

.footer {
  border-top: 1px solid #1a1a1a;
  padding: 1.5rem 0;
  margin-top: auto;
}

.footer .container {
  display: flex;
  justify-content: center;
  align-items: center;
  font-size: 0.75rem;
}

.footer-copyright {
  color: #666;
  text-transform: lowercase;
}

@media (max-width: 768px) {
  .service-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.75rem;
  }

  .service-meta {
    width: 100%;
    justify-content: space-between;
  }

  .history-bars {
    grid-template-columns: repeat(45, 1fr);
  }

  .history-label {
    min-width: 40px;
    font-size: 0.625rem;
  }

  .history-bar {
    height: 28px;
  }

  .chart-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.75rem;
  }

  .time-buttons {
    width: 100%;
    justify-content: space-between;
  }

  .chart-wrapper {
    height: 180px;
  }
}
</style>
