<template>
  <el-container style="width: 100%; height: 100%">
    <!-- 一级菜单下面所拥有的二级菜单 -->
    <el-aside style="width: auto">
      <SideMenu :items='items'></SideMenu>
    </el-aside>

    <!-- 以及二级菜单所对应的页面 -->
    <el-main style="padding:0;">
      <router-view></router-view>
    </el-main>
  </el-container>
</template>

<script>
import SideMenu from "@/components/SideMenu";

export default {
  name: "users",
  components: {SideMenu},
  data() {
    return {
      items: [{
        index: 'users',
        title: '用户管理',
        icon: 'user',
      }, {
        index: 'roles',
        title: '角色管理',
        icon: 'wallet',
      },{
        index: 'organization',
        title: '组织管理',
        icon: 'wallet'
      },{
        index: 'permission',
        title: '权限管理',
        icon: 'wallet'
      }]
    }
  }
}
</script>

<style scoped>
</style>
