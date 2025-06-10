import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
    plugins: [vue()],
    server: {
        proxy: {
            '/vote': {
                target: 'http://backend:8080',
                changeOrigin: true
            }
        }
    }
});
