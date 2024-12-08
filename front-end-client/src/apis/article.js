import axios from "axios";

const API_BASE_URL = "http://localhost:8081/api/v1"; // 后端服务地址

export const getArticles = async () => {
  const response = await axios.get(`${API_BASE_URL}/articles/0`);
  return response.data;
};

export const getArticleById = async (id) => {
  const response = await axios.get(`${API_BASE_URL}/articles/${id}`);
  return response.data;
};