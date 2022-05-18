<template>

  <div class="loginContent">
    <div class="loginForm">
      <div class="loginHeader">
        <el-image class="logoImage" :src="require('../../assets/logo.png')"></el-image>
        <span class="logoFont">NickyShell</span>
      </div>

      <el-form @submit.enter.prevent>
        <el-form-item class="loginFormItem">
          <el-input v-model="username" placeholder="请输入用户名"></el-input>
        </el-form-item>
        <el-form-item class="loginFormItem">
          <el-input v-model="password" type="password" placeholder="请输入密码"></el-input>
        </el-form-item>
        <el-form-item class="loginFormItem">
          <el-button class="submit-button" type="primary" native-type="submit" @click="doLogin">登录</el-button>
        </el-form-item>
      </el-form>
    </div>

  </div>
</template>

<script>
import http from '../../util/http'
import router from '../../router'

export default {
  name: "login",
  data() {
    return {
      username: '',
      password: '',
    }
  },
  methods: {
    doLogin() {
      http.post('/user/admin/login', {
        username: this.username,
        password: this.password
      }).then((data) => {
        sessionStorage.setItem('token', data['token'])
        router.push('/')
      }).catch((err) => {
        this.$message(err)
      })
    }
  },
  mounted() {
    let token = sessionStorage.getItem('token')
    if (token !== "") {
      // router.push('/')
    }
  }
}
</script>

<style scoped>
.loginContent {
  width: 100%;
  height: 100%;
  background-color: #f0f2f5;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  flex-direction: column;
}

.loginForm {
  background-color: #fff;
  width: 400px;
  height: 300px;
  margin: 0;
  padding: 20px;
  border: 1px solid #fff;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.loginFormItem {
  width: 300px;
}

.submit-button {
  background-color: #ff7d0a;
  color: #ffffff;
  font-size: 16px;
  width: 100%;
  border: 1px solid #ff7d0a;
  border-radius: 4px;
}

.submit-button:hover {
  background-color: #ff8d0a;
}

.loginHeader {
  text-align: center;
  width: 100%;
  font-size: 44px;
  display: flex;
  justify-content: center;
  align-items: center;
  padding-bottom: 30px;
}

.logoImage {
  width: 60px;
}

.logoFont {
  position: relative;
  top: -2px;
}
</style>
