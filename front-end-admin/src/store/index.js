// src/store/index.js
import { createStore } from 'vuex';

const store = createStore({
  state: {
    isLoggedIn: !!localStorage.getItem('token'), // 根据 token 判断初始登录状态
    user: {
      role: null, // 初始角色为 null
    },
  },
  mutations: {
    setLoginState(state, status) {
      state.isLoggedIn = status;
    },
    setUserRole(state, role) {
      state.user.role = role; // 设置用户角色
    },
  },
  actions: {
    login({ commit }, role) {
      commit('setLoginState', true); // 登录成功后修改登录状态
      commit('setUserRole', role); // 登录时设置用户角色
    },
    logout({ commit }) {
      commit('setLoginState', false); // 登出后修改登录状态
      commit('setUserRole', null); // 登出时重置用户角色
    },
  },
});

export default store;