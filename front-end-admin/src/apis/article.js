import axios from 'axios';

// 创建 axios 实例
const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL, // 后端服务地址，使用 Vite 提供的环境变量
  timeout: 10000,
});

// 添加请求拦截器，在发送请求之前获取 token
apiClient.interceptors.request.use((config) => {
  const token = localStorage.getItem('token'); // 从 localStorage 获取 token
  if (token) {
    config.headers['Authorization'] = `${token}`; // 将 token 添加到请求头
  }
  return config;
}, (error) => {
  return Promise.reject(error);
});

// 获取文章列表
export const fetchArticles = () => apiClient.get('/articles/');

// 获取文章详情
export const fetchArticleDetail = (id) => apiClient.get(`/articles/${id}/`);

// 创建新文章
export const createArticle = (data) => apiClient.post('/articles/', data);

// 更新文章
export const updateArticle = (id, data) => apiClient.patch(`/articles/${id}/`, data);

// 删除文章
export const deleteArticle = (id) => apiClient.delete(`/articles/${id}/`);