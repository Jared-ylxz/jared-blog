<template>
    <div class="login-form">
      <h1>登录</h1>
      <form @submit.prevent="handleLogin">
        <div>
          <label for="username">用户名</label>
          <input id="username" v-model="credentials.username" required />
        </div>
        <div>
          <label for="password">密码</label>
          <input type="password" id="password" v-model="credentials.password" required />
        </div>
        <button type="submit">登录</button>
        <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
      </form>
    </div>
  </template>
  
  <script>
  import { login } from '../apis/user'; // 假设你有一个用于登录的 API
  
  export default {
    name: 'Login',
    data() {
      return {
        credentials: {
          username: '',
          password: '',
        },
        errorMessage: '',
      };
    },
    methods: {
      async handleLogin() {
        try {
          await login(this.credentials); // 调用登录 API
          this.$router.push('/'); // 登录成功后重定向到首页
        } catch (error) {
          this.errorMessage = '登录失败，请检查用户名和密码。';
        }
      },
    },
  };
  </script>
  
  <style scoped>
  .login-form {
    padding: 20px;
    max-width: 400px;
    margin: 0 auto;
  }
  form {
    display: flex;
    flex-direction: column;
    gap: 10px;
  }
  input {
    width: 100%;
    padding: 8px;
  }
  .error {
    color: red;
  }
  </style>