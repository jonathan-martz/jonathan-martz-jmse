import tailwindcss from "@tailwindcss/vite";
import { VitePWA } from 'vite-plugin-pwa'
export default defineNuxtConfig({
  compatibilityDate: "2025-05-05",
  devtools: { enabled: true },
  css: ['./tailwind.css'],

  vite: {
    plugins: [
      tailwindcss(),
      VitePWA({ registerType: 'autoUpdate' })
    ],
  },

  plausible: {
    // Prevent tracking on localhost
    ignoredHostnames: ['localhost'],
    apiHost: 'https://tracking.jmse.cloud',
    domain: 'jonathan-martz.de'
  },
  modules: ["@nuxtjs/plausible"],
});