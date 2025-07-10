<script setup lang="js">
import '../css/Loginstyle.css'
import '../css/boxicons.min.css'
import { container, toggleRegister, toggleLogin ,loginPost,registerPost,errorMessage,accountInput, passwordInput,telephoneInput } from '../js/login.js'
</script>

<template>
  <div class="container" ref="container">
    <div class="form-box login">
      <form @submit.prevent="loginPost">
        <h1>登录</h1>
        <div class="input-box">
          <input type="text" placeholder="账号" required v-model="accountInput">
          <i class="bx bxs-user"></i>
        </div>
        <div class="input-box">
          <input type="password" placeholder="密码" required v-model="passwordInput">
          <i class="bx bxs-lock-alt"></i>
        </div>
        <div v-if="errorMessage" class="error-message">
          {{ errorMessage }}
        </div>
        <div class="forgot-link">
          <a href="#">忘记密码?</a>
        </div>

        <button type="submit" class="btn">Login</button>
        <p>或者选用以下平台登录</p>
        <div class="social-icons">
          <a href="#"><i class="bx bxl-google"></i></a>
          <a href="#"><i class="bx bxl-facebook"></i></a>
          <a href="#"><i class="bx bxl-github"></i></a>
          <a href="#"><i class="bx bxl-linkedin"></i></a>
        </div>
      </form>
    </div>

    <div class="form-box register">
      <form @submit.prevent="registerPost">
        <h1>注册</h1>
        <div class="input-box">
          <input type="text" placeholder="Account" required v-model="accountInput">
          <i class="bx bxs-user"></i>
        </div>

        <div class="input-box">
          <input type="text" placeholder="Telephone" required v-model="telephoneInput">
          <i class="bx bxs-phone"></i>
        </div>
        <div class="input-box">
          <input type="password" placeholder="Password" required v-model="passwordInput">
          <i class="bx bxs-lock-alt"></i>
        </div>
        <div v-if="errorMessage" class="error-message">
          {{ errorMessage }}
        </div>
        <div class="forgot-link">
          <a href="#">忘记密码?</a>
        </div>
        <button type="submit" class="btn">Register</button>
        <p>或者选用以下平台登录</p>
        <div class="social-icons">
          <a href="#"><i class="bx bxl-google"></i></a>
          <a href="#"><i class="bx bxl-facebook"></i></a>
          <a href="#"><i class="bx bxl-github"></i></a>
          <a href="#"><i class="bx bxl-linkedin"></i></a>
        </div>
      </form>
    </div>

    <div class="toggle-box">
      <div class="toggle-panel toggle-left">
        <h1>Hello, Welcome!</h1>
        <p>还没有一个账户嘛?</p>
        <button class="btn register-btn" @click="toggleRegister">Register</button>
      </div>

      <div class="toggle-panel toggle-right">
        <h1>Welcome Back!</h1>
        <p>已经有一个账户啦?</p>
        <button class="btn login-btn" @click="toggleLogin">Login</button>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>