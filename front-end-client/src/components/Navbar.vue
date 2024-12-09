<template>
    <nav class="navbar">
      <router-link class="nav-link" to="/">首页</router-link>
      <RouterLink class="nav-link" to="/blog-create">新建</RouterLink>
      <RouterLink class="nav-link" to="/">友链</RouterLink>
      <RouterLink class="nav-link" to="/">登录</RouterLink>
      <RouterLink class="nav-link" to="/">注册</RouterLink>
    </nav>
</template>
  
<style scoped>
  .navbar{
    text-align: center;
    padding: 1rem 10rem;
    margin: 0 auto;
    background: rgba(200, 200, 200, 0.2); /* 半透明白色背景 */
    backdrop-filter: blur(4px); /* 背景模糊效果 */
    z-index: 1000; /* 确保导航栏在其他内容之上 */
  }

  .nav-link{
    font-weight: 600;
    text-decoration: none;
    color: #ffffffda;
    padding: 5px 15px;
    border-radius: 5px;
    transition: all 0.2s ease-in-out;
    margin-right: 1rem;
    &:hover{
      background: #ffffff48;
      font-weight: 900;
      color: #ffffff;
    }
  }

  .navbar a {
    color: white;
    text-decoration: none;
    margin: 0 1rem;
  }
</style>