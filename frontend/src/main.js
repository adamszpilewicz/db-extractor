import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { store } from './store.js'; // Import the store

const app = createApp(App);
app.use(router);
app.use(store);  // Use the store
app.mount('#app');