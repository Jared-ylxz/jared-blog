<template>
  <div class="login-container">
    <div class="login-form">
      <h1>登录</h1>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="username">用户名</label>
          <input id="username" v-model="credentials.username" placeholder="请输入用户名" required />
        </div>
        <div class="form-group">
          <label for="password">密码</label>
          <input type="password" id="password" v-model="credentials.password" placeholder="请输入密码" required />
        </div>
        <div class="form-actions">
          <button type="submit" class="btn btn-primary">登录</button>
          <button type="button" class="btn btn-secondary" @click="goToRegister">注册</button>
        </div>
        <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
      </form>
    </div>
  </div>
</template>

<script>
  import { login } from "../apis/user";
  import { mapActions } from "vuex";

  export default {
    name: "Login",
    data() {
      return {
        credentials: {
          username: "",
          password: "",
        },
        errorMessage: "",
      };
    },
    methods: {
      ...mapActions(["login"]), // 映射 Vuex 的 login action 到当前组件的 methods 中
      async handleLogin() {
        try {
          const response = await login(this.credentials); // 调用登录 API
          this.$router.push("/"); // 登录成功后重定向到首页
          // 更新 isLoggedIn 状态
          this.login(); // 调用 Vuex 的 login action 更新状态
          this.$router.push("/"); // 登录成功后重定向到首页
        } catch (error) {
          this.errorMessage = "登录失败，请检查用户名和密码。";
        }
      },
      goToRegister() {
        this.$router.push("/register"); // 跳转到注册页面
      },
    },
  };
</script>

<style scoped>
  .login-container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 86vh;  /* 86vh 即 86% 视窗高度 */
    background-color: #f5f5f5;
  }

  .login-form {
    background: #fff;
    padding: 30px;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    width: 100%;
    max-width: 400px;
  }

  h1 {
    text-align: center;
    margin-bottom: 20px;
    color: #333;
  }

  .form-group {
    margin-bottom: 15px;
  }

  label {
    display: block;
    margin-bottom: 5px;
    font-weight: bold;
    color: #555;
  }

  input {
    width: 94.5%;  /* 输入框宽度 */
    padding: 10px;
    border: 1px solid #ddd;
    border-radius: 6px;  /* 输入框圆角 */
    font-size: 14px;
  }

  input:focus {
    border-color: #007bff;
    outline: none;
  }

  .form-actions {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 20px;
  }

  .btn {
    padding: 10px 15px;
    border: none;
    border-radius: 4px;
    font-size: 14px;
    cursor: pointer;
  }

  .btn-primary {
    background-color: #007bff;
    color: white;
  }

  .btn-primary:hover {
    background-color: #0056b3;
  }

  .btn-secondary {
    background-color: #6c757d;
    color: white;
  }

  .btn-secondary:hover {
    background-color: #5a6268;
  }

  .error {
    color: red;
    margin-top: 10px;
    text-align: center;
  }
</style>