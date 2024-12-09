<template>
    <nav class="navbar">
      <router-link to="/">首页</router-link>
    </nav>
</template>
  
<style>
  .navbar{
    text-align: center;
    padding: 1rem 10rem;
    margin: 0 auto;
    background: rgba(200, 200, 200, 0.2); /* 半透明白色背景 */
    backdrop-filter: blur(4px); /* 背景模糊效果 */
    z-index: 1000; /* 确保导航栏在其他内容之上 */
  }
  .navbar a {
    color: white;
    text-decoration: none;
    margin: 0 1rem;
  }
</style>