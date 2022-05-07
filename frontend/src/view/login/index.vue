<template>
  <div class="loginContent">
    <el-form class="loginForm">
      <el-form-item class="loginFormItem">
        <el-input v-model="username" placeholder="请输入用户名"></el-input>
      </el-form-item>
      <el-form-item class="loginFormItem">
        <el-input v-model="password" type="password" placeholder="请输入密码"></el-input>
      </el-form-item>
      <el-form-item class="loginFormItem">
        <el-button class="submit-button" type="primary" @click="doLogin">登录</el-button>
      </el-form-item>
    </el-form>
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
      }).then((res) => {
        sessionStorage.setItem('token', res.data['token'])
        router.push('/')
      }).catch((err) => {
        console.log(err)
        this.$message(err['code'] + ' ' + err['message'])
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
</style>
