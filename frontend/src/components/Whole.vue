<template>
  <div class="wrapper">
    <div class="header">
      <div class="logo" @click="this.$router.push('/')">
        <el-image class="logoImage" :src="require('../assets/logo.png')"></el-image>
        <span class="logoFont">NickyShell</span>
      </div>

      <!-- 水平一 级菜单 -->
      <div style="float:left;">
        <el-menu :default-active="toIndex()" mode="horizontal" @select="handleSelect">
          <el-menu-item v-for="item in items" :index="item.index" :key="item.index">
            <template #title>
              <span>{{ item.title }}</span>
            </template>
          </el-menu-item>
        </el-menu>
      </div>

      <div class="header-right">
        <div class="header-user-con">
          <!-- 用户头像 -->
          <div class="user-avator"><img src="../assets/img/avator.jpeg"/></div>
          <!-- 用户名下拉菜单 -->
          <el-dropdown class="user-name" trigger="click" @command="handleCommand">
            <span class="el-dropdown-link">
              {{ username }}
              <i class="el-icon-caret-bottom"></i>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item disabled>修改密码</el-dropdown-item>
                <el-dropdown-item command="loginout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </div>
    </div>

    <!-- 页面左侧二级菜单栏，和主体内容区域部分 -->
    <el-main class="whole-el-main">
      <router-view/>
    </el-main>
  </div>
</template>

<script>
export default {
  name: "Whole",
  data() {
    return {
      items: [    // 水平一级菜单栏的菜单
        {index: 'Home', title: '首页'},
        {index: 'users', title: '用户管理'},
        {index: 'assets', title: '资产管理'},
        {index: 'audit', title: '审计管理'},
        {index: 'setting', title: '系统设置'},
      ]
    }
  },
  computed: {
    username() {
      let username = localStorage.getItem('ms_username');
      return username ? username : 'admin';
    }
  },
  methods: {
    // 根据路径绑定到对应的一级菜单，防止页面刷新重新跳回第一个
    toIndex() {
      return this.$route.path.split('/')[1];
    },
    // 切换菜单栏
    handleSelect(index) {
      this.$router.push('/' + index);
    },
    // 用户名下拉菜单选择事件
    handleCommand(command) {
      if (command === 'loginout') {
        localStorage.removeItem('ms_username');
        this.$router.push('/login');
      }
    }
  }
}
</script>


<style scoped>
.wrapper {
  width: 100%;
  height: 100%;
  background: #f0f0f0;
}

.header {
  position: relative;
  box-sizing: border-box;
  width: 100%;
  height: 70px;
  font-size: 22px;
  min-width: 1124px;
}

.header .logo {
  float: left;
  margin-left: 12px;
  margin-top: 8px;
  height: 60px;
  width: 200px;
  vertical-align: middle;
}

.logo {
  display: flex;
  justify-content: center;
  align-items: center;
  cursor: pointer;
}

/* --------------- 用户头像区域的样式 ---------------- */
.header-right {
  float: right;
  padding-right: 50px;
}

.header-user-con {
  display: flex;
  height: 70px;
  align-items: center;
}

.user-avator {
  margin-left: 20px;
  margin-right: 10px;
}

.user-avator img {
  display: block;
  width: 40px;
  height: 40px;
  border-radius: 50%;
}

.user-name {
  margin-left: 10px;
  margin-right: 5px;
}

.el-dropdown-link {
  cursor: pointer;
}

.el-dropdown-menu__item {
  text-align: center;
}

/* --------------- 水平一级菜单栏的样式--------------------- */
.el-menu.el-menu--horizontal {
  border-bottom: none !important;
  float: left;
  margin-left: 50px;
}

.el-menu--horizontal > .el-menu-item.is-active {
  border-bottom: 2px solid #409eff;
  color: #3989fa;
  font-weight: 700;
}

.el-menu--horizontal > .el-menu-item {
  font-size: 16px;
  margin: 0 15px;
  color: black;
}

.whole-el-main {
  height: calc(100% - 70px);
  margin: 0;
  padding: 0;
}

.logoImage {
  height: 100%;
  position: relative;
}

.logoFont {

}
</style>
