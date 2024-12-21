import { createApp } from 'vue';
import App from './App.vue';
import router from './router';
import store from './store'; // 引入 Vuex store

const app = createApp(App);

// 在应用启动时恢复用户角色
store.dispatch('restoreUserRole');

app.use(router);
app.use(store); // 将 Vuex store 添加到应用实例中
app.mount('#app');