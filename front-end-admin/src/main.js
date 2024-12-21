import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store'; // 引入 Vuex store

const app = createApp(App);

app.use(router);
app.use(store); // 将 Vuex store 添加到应用实例中
app.mount('#app');