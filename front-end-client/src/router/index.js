import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/Home.vue";
import ArticleDetailView from "../views/ArticleDetail.vue";

const routes = [
  { path: "/", name: "Home", component: HomeView },
  { path: "/articles/:id", name: "Article", component: ArticleDetailView },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;