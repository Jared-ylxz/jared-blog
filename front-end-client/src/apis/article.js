import axios from "axios";

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL; // 后端服务地址，使用 Vite 提供的环境变量

// 获取文章列表
export const getArticles = async () => {
  const response = await axios.get(`${API_BASE_URL}/articles/`);
  return response.data;
};

// 获取文章详情
export const getArticleDetailById = async (id) => {
  const response = await axios.get(`${API_BASE_URL}/articles/${id}`);
  return response.data;
};