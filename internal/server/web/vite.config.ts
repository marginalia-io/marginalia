import { defineConfig } from "vite-plus";

export default defineConfig({
  server: {
    proxy: {
      // Forward API calls to the Go backend (see `make dev`).
      "/api": "http://localhost:8090",
    },
  },
  staged: {
    "*": "vp check --fix",
  },
  fmt: {},
  lint: {
    jsPlugins: [{ name: "vite-plus", specifier: "vite-plus/oxlint-plugin" }],
    rules: { "vite-plus/prefer-vite-plus-imports": "error" },
    options: { typeAware: true, typeCheck: true },
  },
});
