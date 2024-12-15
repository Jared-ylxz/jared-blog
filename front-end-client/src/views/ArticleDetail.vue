<template>
  <div>
    <Navbar />
    <div class="container">
      <div v-if="loading">加载中...</div>
      <div v-else>
        <h1>{{ article.title }}</h1>
        <p>{{ article.content }}</p>
      </div>
    </div>
  </div>
</template>
  
<script>
  import { ref, onMounted } from "vue";
  import { getArticleDetailById } from "../apis/article";
  import Navbar from "../components/Navbar.vue";
  import { useRoute } from "vue-router";
  
  export default {
    components: { Navbar },
    setup() {
      const route = useRoute();
      const article = ref(null);
      const loading = ref(true);
  
      onMounted(async () => {
        const articleId = route.params.id;
        article.value = await getArticleDetailById(articleId);
        loading.value = false;
      });
  
      return { article, loading };
    },
  };
</script>
  
<style scoped>
  .container {
    background-color: #fff;
    max-width: 800px;
    margin: 0 auto;
    padding: 1rem;
  }
</style>