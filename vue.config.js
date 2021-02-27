const path = require('path')

module.exports = {
  outputDir: 'view/dist',
  assetsDir: 'assets',
  runtimeCompiler: true,
  configureWebpack: {
    entry: {
      app: './view/src/main.ts'
    },
    plugins: []
  },
  chainWebpack: config => {
    config.resolve.alias.set('@', path.join(__dirname, '/view/src'))
  },
  devServer: {
    port: 30776
  },
  lintOnSave: false
}
