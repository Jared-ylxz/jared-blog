<template>
    <div class="article-list">
      <h1>文章管理</h1>
      <router-link to="/articles/new">
        <button class="create-article">新增文章</button>
      </router-link>
      <table>
        <thead>
          <tr>
            <th>ID</th>
            <th>标题</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="article in articles" :key="article.id">
            <td>{{ article.id }}</td>
            <td>{{ article.title }}</td>
            <td>
              <router-link :to="`/articles/${article.id}/edit`">
                <button>编辑</button>
              </router-link>
              <button @click="handleDelete(article.id)">删除</button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </template>
  
  <script>
  import { fetchArticles, deleteArticle } from '../apis/article';
  
  export default {
    name: 'ArticleList',
    data() {
      return {
        articles: [],
      };
    },
    methods: {
      async loadArticles() {
        const response = await fetchArticles();
        this.articles = response.data;
      },
      async handleDelete(id) {
        if (confirm('确定删除该文章吗？')) {
          await deleteArticle(id);
          this.loadArticles(); // 重新加载文章列表
        }
      },
    },
    mounted() {
      this.loadArticles();
    },
  };
  </script>
  
  <style>
  .article-list {
    padding: 20px;
  }
  .create-article {
    margin-bottom: 10px;
  }
  table {
    width: 100%;
    border-collapse: collapse;
  }
  th, td {
    border: 1px solid #ddd;
    padding: 8px;
  }
  </style>