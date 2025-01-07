// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: [
    '@nuxt/ui',
    '@nuxtjs/tailwindcss',
    '@nuxtjs/color-mode',
    '@nuxt/icon',
    '@nuxt/image'
  ],
  colorMode: {
    preference: 'system',
    fallback: 'light'
  },
  runtimeConfig: {
    yourOrigin: "localhost",
    public: {
      hurdleBackendURL: process.env.HURDLE_API_URL,
    },
  },
})
