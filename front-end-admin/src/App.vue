<template>
  <div id="app">
    <!-- 页面头部导航栏 -->
    <header>
      <nav>
        <router-link to="/">文章列表</router-link>
        <router-link to="/articles/new">新增文章</router-link>
        <router-link class="right-nav" v-if="!isLoggedIn" to="/login">登录</router-link>
        <a class="right-nav" v-else @click="confirmLogout">登出</a>
      </nav>
    </header>

    <!-- 路由出口：显示当前路由对应的页面 -->
    <main>
      <router-view />
    </main>
  </div>
</template>
  
<script setup>
  import { useStore } from "vuex";
  import { useRouter } from "vue-router";
  import { computed } from "vue";

  // 获取 Vuex store 和 Vue Router 实例
  const store = useStore();
  const router = useRouter();

  // 映射 Vuex 的 isLoggedIn 状态
  const isLoggedIn = computed(() => store.state.isLoggedIn);

  // 定义登出方法
  const confirmLogout = () => {
    if (confirm("确定要退出登录吗？")) {
      localStorage.removeItem("token"); // 清除 token
      store.dispatch("logout"); // 调用 Vuex 的 logout action 更新状态
      router.push("/"); // 重定向到首页
    }
  };
</script>
  
<style scoped>
  /* 全局样式 */
  body {
    margin: 0;
    font-family: Arial, sans-serif;
    background-color: #f5f5f5;
    color: #333;
  }

  /* 导航栏样式 */
  header {
    padding: 10px 0;
    background-color: #007bff;
    color: white;
    margin-bottom: 20px;
  }

  nav {
    display: flex;
    gap: 10px;
  }

  nav a {
    color: white;
    text-decoration: none;
    padding: 5px 10px;
    border-radius: 5px;
    transition: background-color 0.3s;
  }

  nav a:hover {
    background-color: #0056b3;
  }

  /* 将登录按钮放到右边 */
  nav .right-nav {
    margin-left: auto;
  }

  /* 路由内容区域样式 */
  main {
    padding: 20px;
    background-color: #fff;
    border-radius: 5px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  }
</style>