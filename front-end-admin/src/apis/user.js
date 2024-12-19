import axios from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL; // 后端服务地址，使用 Vite 提供的环境变量

// 用户登录
export const login = async (credentials) => {
  const response = await axios.post(`${API_BASE_URL}/users/login/`, credentials);
  localStorage.setItem('token', response.data.token); // 保存 token 到 localStorage
  return response;
};

// 用户注册
export const register = async (data) => {
  const response = await axios.post(`${API_BASE_URL}/users/register/`, data);
  return response;
};