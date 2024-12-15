<template>
    <div class="article-form">
      <h1>{{ isEditing ? '编辑文章' : '新增文章' }}</h1>
      <form @submit.prevent="handleSubmit">
        <div>
          <label for="title">标题</label>
          <input id="title" v-model="form.title" required />
        </div>
        <div>
          <label for="content">内容</label>
          <textarea id="content" v-model="form.content" required></textarea>
        </div>
        <button type="submit">{{ isEditing ? '更新' : '创建' }}</button>
      </form>
    </div>
  </template>
  
  <script>
  import { fetchArticleDetail, createArticle, updateArticle } from '../apis/article';
  
  export default {
    name: 'ArticleForm',
    data() {
      return {
        form: {
          title: '',
          content: '',
        },
        isEditing: false, // 是否为编辑模式
      };
    },
    methods: {
      async loadArticle() {
        const id = this.$route.params.id;
        if (id) {
          this.isEditing = true;
          const response = await fetchArticleDetail(id);
          this.form = response.data;
        }
      },
      async handleSubmit() {
        if (this.isEditing) {
          const id = this.$route.params.id;
          await updateArticle(id, this.form);
          alert('文章更新成功！');
        } else {
          await createArticle(this.form);
          alert('文章创建成功！');
        }
        this.$router.push('/');
      },
    },
    mounted() {
      this.loadArticle();
    },
  };
  </script>
  
  <style>
  .article-form {
    padding: 20px;
  }
  form {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
  input, textarea {
    width: 100%;
    padding: 8px;
  }
  </style>