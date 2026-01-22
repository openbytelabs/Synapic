const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  transpileDependencies: true,
  lintOnSave: false,
  devServer: {
    port: 8080,
    proxy: {
      '/api': {
        target: 'https://uptimeapi.synapic.com.tr',
        changeOrigin: true,
        secure: false
      }
    }
  }
})
