<script setup>

</script>

<template>
  <div class="background-container">
    <router-view />
  </div>
</template>

<style scoped>
.background-container {
  background-image: url('./assets/site-background/wallhaven-yxdrex.png'); /* 背景图片 */
  background-size: cover;
  background-position: center;
  background-repeat: no-repeat;
  background-attachment: fixed;
  height: 100vh; /* 根据需要设置高度 */
}
</style>
