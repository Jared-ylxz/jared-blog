import { createRouter, createWebHistory } from 'vue-router';
import ArticleList from '../views/ArticleList.vue';
import ArticleForm from '../views/ArticleForm.vue';

const routes = [
  {path: '/', name: 'ArticleList', component: ArticleList },
  {path: '/articles/:id/edit', name: 'EditArticle', component: ArticleForm },
  {path: '/articles/new', name: 'NewArticle', component: ArticleForm },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;