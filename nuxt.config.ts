import tailwindcss from "@tailwindcss/vite";
export default defineNuxtConfig({
  compatibilityDate: "2025-05-05",
  devtools: { enabled: true },
  css: ['./tailwind.css'],

  vite: {
    plugins: [
      tailwindcss(),
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