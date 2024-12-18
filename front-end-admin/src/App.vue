<template>
  <div id="app">
    <!-- 页面头部导航栏 -->
    <header>
      <nav>
        <router-link to="/">文章列表</router-link>
        <router-link to="/articles/new">新增文章</router-link>
        <router-link to="/login" class="right">登录</router-link>
        <router-link to="/login" class="right">登出</router-link>
        <router-link to="/login" class="right">注册</router-link>
      </nav>
    </header>

    <!-- 路由出口：显示当前路由对应的页面 -->
    <main>
      <router-view />
    </main>
  </div>
</template>
  
  <script>
  export default {
    name: "App",
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
  
  #app {
    max-width: 1200px;
    margin: 0 auto;
    padding: 20px;
    background-color: #fff;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
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

  nav .right {
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