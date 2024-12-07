import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/Home.vue";
import ArticleView from "../views/Article.vue";

const routes = [
  { path: "/", name: "Home", component: HomeView },
  { path: "/articles/:id", name: "Article", component: ArticleView },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;