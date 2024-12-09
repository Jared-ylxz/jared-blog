<template>
  <router-link :to="`/articles/${article.id}`">
    <div class="article-card">
      <h3 class="article-title">{{ article.title }}</h3>
      <p class="article-message">{{ article.description }}</p>
    </div>
  </router-link>
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
    color: rgb(53, 53, 53);
    background-color: #fff;
    border: 1px solid #ddd;
    padding: 1rem;
    margin-bottom: 1rem;
    border-radius: 4px;
  }

  .article-card:hover{
      transform: translateY(-4px);
      background-color: #f9f9f9;
  }

  .article-title {
      color: #42b983;
      margin-bottom: 10px;
  }

  .article-message {
      font-size: 0.9em;
      color: #777;
      margin-bottom: 10px;
  }
  </style>