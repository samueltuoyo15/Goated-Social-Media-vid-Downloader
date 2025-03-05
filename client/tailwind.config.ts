import { Config } from 'tailwindcss'
const config: Config = {
  content: [
    './index.html',
    './src/**/*.{js,ts,jsx,tsx}',
  ],
  theme: {
    extend: {
      animation: {
        'bounce-sequential': 'bounce-sequential 1.5s infinite',
      },
      fontFamily:{
        sans: ['Inter', 'sans-serif'],
      },
      keyframes: {
        'bounce-sequential': {
          '0%, 80%, 100%': { transform: 'scale(0)' },
          '40%': { transform: 'scale(1)' },
        },
      },
    },
  },
  plugins: [],
}

export default config
