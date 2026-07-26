<template>
  <div id="app">
    <router-view />
  </div>
</template>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

html, body {
  width: 100%;
  min-height: 100%;
  overflow-x: hidden;
}

#app {
  min-height: 100vh;
}

/* scrollbar arka planı da siyah olsun */
::-webkit-scrollbar {
  background: #000;
}
</style>