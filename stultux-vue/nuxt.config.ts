// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  modules: ['@nuxtjs/tailwindcss', '@nuxt/eslint', '@nuxtjs/color-mode'],
  devtools: { enabled: true },
  compatibilityDate: '2024-11-01',
  eslint: {
    config: {
      stylistic: true, // <---
    },
  },
})
