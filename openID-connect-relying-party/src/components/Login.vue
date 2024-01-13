<script setup lang="ts">
  // import axios from "axios";
  // function login() {

  //   axios.get("http://rp.com:8081/login",
  //   {
  //     iss : "op.com",
  //   }
  //   ).then((_response:any) => {
  //     window.location.replace("http://op.com:8000/v1/authorize?redirect_uri=http://rp.com:8081/code_flow/oidc_op&scope=openid+profile&response_type=code&nonce=cdYrYNLv6wBHlBmZjWxvrQmmD&state=aHR0cDovL3JwLmNvbTo4MDgx&client_id=cnAuY29t");
  //   })
  // }
  
</script>

<template>
      <form action="/v1/account_verify" method="POST">
        <input type="text" placeholder="用户名" name="account" required />
        <input type="password" placeholder="密码" name="password" required />
        <button type="submit">登录</button>
      </form>
  <button @click="">OIDC登录</button>
</template>

<style scoped>

</style>