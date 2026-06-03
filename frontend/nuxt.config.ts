// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',

  devtools: { enabled: true },

  ssr: false,

  modules: [
    '@nuxt/ui',
    '@nuxtjs/tailwindcss',
    '@nuxtjs/color-mode',
    '@nuxt/icon',
    '@nuxt/image',
    '@nuxtjs/sitemap',
    '@nuxtjs/robots',
  ],

  colorMode: {
    preference: 'system',
    fallback: 'light'
  },

  robots: {
    blockAiBots: true
  },

  runtimeConfig: {
    public: {
      ghostSendApiUrl: '/api/v1',
    },
  },

  nitro: {
    preset: 'static',
    output: {
      // Output static files into frontend/dist so that frontend/fs.go
      // can embed them into the Go binary via //go:embed.
      publicDir: 'dist'
    },
    // In production the SPA is embedded in (and served by) the Go binary, so
    // `/api/*` is same-origin. During `nuxt dev` the SPA runs on :3000 while
    // the API runs on :8080, so proxy API calls to the backend.
    devProxy: {
      '/api': {
        target: 'http://localhost:8080/api',
        changeOrigin: true
      }
    }
  },

  vite: {
    build: {
      rollupOptions: {
        onwarn(warning, warn) {
          if (warning.code === 'INVALID_ANNOTATION') return
          warn(warning)
        }
      }
    }
  }
})
