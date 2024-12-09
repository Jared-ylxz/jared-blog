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
    transition: transform 0.4s ease, background-color 0.4s ease;  /* 鼠标悬停时，元素上浮的速度 */
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
    
  .article-card a link{
    text-decoration: none; /* 移除下划线 */
  }
</style>