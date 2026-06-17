import { defineConfig } from "vite-plus";
import react from "@vitejs/plugin-react";
import tailwindcss from "@tailwindcss/vite";

export default defineConfig({
  plugins: [react(), tailwindcss()],
  // Emit the production build into the Go embed package so the binary can
  // serve the SPA. emptyOutDir is required because the dir is outside the
  // frontend project root.
  build: {
    outDir: "../internal/server/embed/dist",
    emptyOutDir: true,
  },
  resolve: {
    tsconfigPaths: true,
  },
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
