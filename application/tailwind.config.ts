import type { Config } from 'tailwindcss'

const config: Config = {
  content: ['./src/**/*.{html,js}'],
  theme: {
    extend: {},
  },
  plugins: [],
} as const

export default config
