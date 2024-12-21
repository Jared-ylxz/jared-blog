// src/store/index.js
import { createStore } from 'vuex';

const store = createStore({
  state: {
    isLoggedIn: !!localStorage.getItem('token'), // 根据 token 判断初始登录状态
  },
  mutations: {
    setLoginState(state, status) {
      state.isLoggedIn = status;
    },
  },
  actions: {
    login({ commit }) {
      commit('setLoginState', true); // 登录成功后修改登录状态
    },
    logout({ commit }) {
      commit('setLoginState', false); // 登出后修改登录状态
    },
  },
});

export default store;