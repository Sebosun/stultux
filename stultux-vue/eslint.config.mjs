import withNuxt from './.nuxt/eslint.config.mjs'

export default withNuxt(
  {
    rules: {
      '@typescript-eslint/no-explicit-any': 'off', // <-- Need this to not do type magic
    },
  },
)
