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
          <th class="action-column">操作</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="article in articles" :key="article.id">
          <td>{{ article.id }}</td>
          <td>{{ article.title }}</td>
          <td>
            <router-link :to="`/articles/${article.id}/edit`">
              <button :class="{ disabled: !isAdmin }" :disabled="!isAdmin">编辑</button>
            </router-link>
            <button 
              class="delete-button" 
              :class="{ disabled: !isAdmin }" 
              :disabled="!isAdmin" 
              @click="handleDelete(article.id)">
              删除
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script setup>
  import { ref, computed, onMounted } from 'vue';
  import { fetchArticles, deleteArticle } from '../apis/article';
  import { useStore } from 'vuex';

  const store = useStore();
  const articles = ref([]);

  // 计算属性判断用户是否为管理员
  const isAdmin = computed(() => store.state.user.role === 9);

  const loadArticles = async () => {
    const response = await fetchArticles();
    articles.value = response.data;
  };

  const handleDelete = async (id) => {
    if (confirm('确定删除该文章吗？')) {
      await deleteArticle(id);
      loadArticles(); // 重新加载文章列表
    }
  };

  onMounted(() => {
    loadArticles();
  });
</script>

<style scoped>
  .article-list {
    padding: 20px;
    background-color: #f9f9f9;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }

  .create-article {
    margin-bottom: 10px;
    background-color: #007bff;
    color: white;
    border: none;
    padding: 10px 15px;
    border-radius: 5px;
    cursor: pointer;
  }

  .create-article:hover {
    background-color: #0056b3;
  }

  table {
    width: 100%;
    border-collapse: collapse;
  }

  th, td {
    border: 1px solid #ddd;
    padding: 12px;
    text-align: left;
  }

  th {
    background-color: #f2f2f2;
  }

  .action-column {
    width: 122px; /* 设置操作列的宽度 */
  }

  button {
    padding: 8px 12px;
    border: none;
    border-radius: 5px;
    cursor: pointer;
    margin-right: 10px;
  }

  button.disabled {
    background-color: #ccc;
    cursor: not-allowed;
  }

  .delete-button {
    background-color: #dc3545;
    color: white;
  }

  .delete-button:hover:not(.disabled) {
    background-color: #c82333;
  }
</style>