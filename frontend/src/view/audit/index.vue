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
  name: "audit",
  components: {SideMenu},
  data() {
    return {
      items: [{
        index: 'realtime',
        title: '实时会话',
        icon: 'edit',
      }, {
        index: 'history',
        title: '历史会话',
        icon: 'edit',
      },]
    }
  }
}
</script>

<style scoped>
/*.el-main {*/
/*  position: absolute;*/
/*  left: 250px;*/
/*  right: 0px;*/
/*  top: 70px;*/
/*  bottom: 0;*/
/*}*/
</style>
