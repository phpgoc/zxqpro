import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";
import path from "path";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  build: {
    outDir: "../dist", // 输出目录设置为项目根目录的上一级目录的 dist 文件夹
  },
  // server: {
  //   proxy: {
  //     '/api': {
  //       target: 'http://localhost:8080',
  //       changeOrigin: true,
  //       rewrite: (path) => path.replace(/^\/api/, ''),
  //     },
  //   },
  // },
  resolve: {
    alias: {
      // 将 `~` 映射到 `src` 目录
      "~": path.resolve(__dirname, "./src"),
    },
  },
});
