module.exports = {
  mode: "jit",
  purge: [
    './public/**/*.html',
    './src/**/*.js',
    '../**/*.go',
  ],
  darkMode: false, // or 'media' or 'class'
  theme: {
    extend: {},
  },
  variants: {
    extend: {},
  },
  plugins: [],
}
