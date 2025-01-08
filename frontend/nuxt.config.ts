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
      SHARE_SECRET_API_URL: process.env.SHARE_SECRET_API_URL,
    },
  },
})
