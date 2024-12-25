<template>
  <div>
    <Navbar />
    <div class="container">
      <div v-if="loading">加载中...</div>
      <div v-else>
        <h1>{{ article.title }}</h1>
        <div v-html="article.content"></div> <!-- 使用 v-html 指令来渲染富文本内容 -->
      </div>
    </div>
  </div>
</template>
 
<script setup>
import { ref, onMounted } from "vue";
import { getArticleDetailById } from "../apis/article";
import Navbar from "../components/Navbar.vue";
import { useRoute } from "vue-router";
 
const route = useRoute();
const article = ref(null);
const loading = ref(true);
 
onMounted(async () => {
  const articleId = route.params.id;
  article.value = await getArticleDetailById(articleId);
  loading.value = false;
});
</script>
 
<style scoped>
.container {
  background-color: #fff;
  max-width: 800px;
  margin: 0 auto;
  padding: 1rem;
}
</style>