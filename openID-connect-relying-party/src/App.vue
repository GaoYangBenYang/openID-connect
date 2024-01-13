<script setup lang="ts">
import Login from './components/Login.vue'
</script>

<template>
  <div>
    <a href="https://vitejs.dev" target="_blank">
      <img src="/vite.svg" class="logo" alt="Vite logo" />
    </a>
  </div>
  <router-link to="/login">登录</router-link>
  <router-view>
    <Login/>
  </router-view>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
