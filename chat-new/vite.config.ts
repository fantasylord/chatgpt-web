import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  base: './',
  server: {
    port: 80,
    proxy: {
      '/completion': 'http://localhost:8080',
    },
  },
})
