/** @type {import('tailwindcss').Config} */
export default {
  content: ["./index.html", "./src/**/*.{js,jsx}"],
  theme: {
    extend: {
      fontFamily: {
        display: ["'Cormorant Garamond'", "serif"],
        body: ["'Jost'", "sans-serif"],
        fell: ["'IM Fell English'", "serif"],
      },
      colors: {
        lapis:    { DEFAULT: "#1a3a6e", light: "#2a5298", dark: "#0f2147" },
        saffron:  { DEFAULT: "#c8570a", light: "#e8761a", dark: "#8b3a06" },
        vermil:   { DEFAULT: "#c0392b", light: "#e74c3c", dark: "#922b21" },
        emerald:  { DEFAULT: "#1a5c3a", light: "#27ae60", dark: "#0f3d26" },
        gold:     { DEFAULT: "#c9a84c", light: "#f0c040", dark: "#8b6914" },
        ivory:    { DEFAULT: "#fdf6e3", dark: "#f5e6c8" },
        mughal:   { DEFAULT: "#2c1810", light: "#4a2c1a" },
      },
      backgroundImage: {
        "miniature-pattern": `
          radial-gradient(circle at 20% 20%, rgba(201,168,76,0.08) 0%, transparent 50%),
          radial-gradient(circle at 80% 80%, rgba(26,58,110,0.08) 0%, transparent 50%),
          radial-gradient(circle at 50% 50%, rgba(200,87,10,0.04) 0%, transparent 70%)
        `,
      },
      boxShadow: {
        "gold": "0 0 0 1px #c9a84c, 0 4px 24px rgba(201,168,76,0.15)",
        "mughal": "0 8px 40px rgba(44,24,16,0.18)",
      },
    },
  },
  plugins: [],
};