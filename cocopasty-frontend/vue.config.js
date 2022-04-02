const { defineConfig } = require('@vue/cli-service')
const GoogleFontsPlugin = require("@beyonk/google-fonts-webpack-plugin")

// vue.config.js
module.exports = {
  }
module.exports = defineConfig({
  transpileDependencies: true,
  configureWebpack: {
    plugins: [
        new GoogleFontsPlugin({
            fonts: [
                { family: "Roboto Mono", variants: [ "500", "700" ] }
            ]
        })
    ]}
})
