// src/store/index.js
import { createStore } from 'vuex';
import { jwtDecode } from 'jwt-decode';

const store = createStore({
  state: {
    isLoggedIn: !!localStorage.getItem('token'), // 页面加载时获取本地存储的 token
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
    restoreUserRole({ commit }) {
      const token = localStorage.getItem('token');
      if (token) {
        const role = getUserRoleFromToken(token);
        commit('setUserRole', role);
      }
    },
  },
});

function getUserRoleFromToken(token) {
    const decoded = jwtDecode(token);
    return decoded.role; // 角色存储在 token 的 payload 中
  }

export default store;