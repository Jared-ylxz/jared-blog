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
  
<script>
  export default {
    name: "App",
    data() {
      return {
        isLoggedIn: false, // 这里可以根据实际情况进行更新
      };
    },
    methods: {
      confirmLogout() {
        if (confirm("确定要退出登录吗？")) {
          localStorage.removeItem('token'); // 清除 token
          this.isLoggedIn = false; // 更新登录状态为未登录
          this.$router.push('/'); // 重定向到首页
        }
      },
    },
    mounted() {
      // 检查用户登录状态
      this.isLoggedIn = !!localStorage.getItem('token'); // 根据 token 是否存在判断登录状态
    },
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