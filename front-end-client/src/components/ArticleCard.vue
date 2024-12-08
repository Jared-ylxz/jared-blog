<template>
    <div class="article-card">
      <h2>{{ article.title }}</h2>
      <p>{{ article.description }}</p>
      <router-link :to="`/articles/${article.id}`">阅读更多</router-link>
    </div>
  </template>
  
  <script>
  export default {
    props: {
      article: {
        type: Object,
        required: true,
      },
    },
  };
  </script>
  
  <style>
  .article-card {
    border: 1px solid #ddd;
    padding: 1rem;
    margin-bottom: 1rem;
    border-radius: 4px;
  }
  </style>