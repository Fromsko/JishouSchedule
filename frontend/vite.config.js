import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: '0.0.0.0',//ip地址
    port: 9090, // 设置服务启动端口号
    open: false, // 设置服务启动时是否自动打开浏览器
  }
})
