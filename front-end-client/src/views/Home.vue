<template>
  <div>
    <Navbar />
    <div class="container">
      <div v-if="loading">加载中...</div>
      <div v-else>
        <ArticleCard
          v-for="article in articles"
          :key="article.id"
          :article="article"
        />
      </div>
    </div>
  </div>
</template>
  
<script>
  import { ref, onMounted } from "vue";
  import { getArticles } from "../apis/article";
  import Navbar from "../components/Navbar.vue";
  import ArticleCard from "../components/ArticleCard.vue";
  
  export default {
    components: { Navbar, ArticleCard },
    setup() {
      const articles = ref([]);
      const loading = ref(true);
  
      onMounted(async () => {
        articles.value = await getArticles();
        loading.value = false;
      });
  
      return { articles, loading };
    },
  };
</script>
  
<style scoped>
  .container {
    max-width: 800px;
    margin: 0 auto;
    padding: 1rem;
  }
</style>