import axios from 'axios';

// 创建 axios 实例
const apiClient = axios.create({
  baseURL: 'http://localhost:8081/api/v1', // 后端 API 的基础 URL
  timeout: 10000,
});

// 获取文章列表
export const fetchArticles = () => apiClient.get('/articles/');

// 获取文章详情
export const fetchArticleDetail = (id) => apiClient.get(`/articles/${id}`);

// 创建新文章
export const createArticle = (data) => apiClient.post('/articles', data);

// 更新文章
export const updateArticle = (id, data) => apiClient.put(`/articles/${id}`, data);

// 删除文章
export const deleteArticle = (id) => apiClient.delete(`/articles/${id}`);