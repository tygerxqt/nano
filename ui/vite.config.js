import { defineConfig } from 'vite'
import { svelte, vitePreprocess } from '@sveltejs/vite-plugin-svelte'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 3000
  },
  base: "./",
  build: {
    chunkSizeWarningLimit: 1000,
    reportCompressedSize: false
  },
  plugins: [svelte({
    preprocess: [vitePreprocess()]
  })],
  resolve: {
    alias: {
      "@": __dirname + "/src",
    }
  },
})
