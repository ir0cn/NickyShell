<template>
  <div class="sidebar">
    <!-- 左侧二级菜单栏的组件封装 -->
    <el-menu
        :class="isCollapsed ? 'sidebar-el-menu-collapse' : 'sidebar-el-menu'"
        :default-active="toIndex()"
        background-color="white"
        text-color="#7a8297"
        active-text-color="#2d8cf0"
        router>
      <el-menu-item v-for="item in items" :index="item.index" :key="item.index">
        <!-- 需要图标的在 item 对象中加上属性 icon -->
<!--        <i :class="item.icon"></i>-->
        <el-icon><component :is="item.icon"></component></el-icon>
        <template #title><span>{{ isCollapsed ? '' : item.title }}</span></template>
      </el-menu-item>

      <div @click="doCollapse" style="width: 100%; height: 16px"   class="collapseBtn">
        <el-icon v-if="isCollapsed"><d-arrow-right /></el-icon>
        <el-icon v-else><d-arrow-left /></el-icon>
<!--        <i :class="isCollapsed ? 'el-icon-d-arrow-right' : 'el-icon-d-arrow-left'"/>-->
      </div>
    </el-menu>
  </div>
</template>

<script>
export default {
  props: ['items'],
  data() {
    return {
      isCollapsed: false
    }
  },
  methods: {
    // 根据路径绑定到对应的二级菜单，防止页面刷新重新跳回第一个
    toIndex() {
      return this.$route.path.split('/')[2];
    },
    doCollapse() {
      this.isCollapsed = !this.isCollapsed
    }
  },
};
</script>

<style scoped>
/* 左侧菜单栏定位和位置大小设定 */
/*.sidebar {*/
/*display: block;*/
/*position: absolute;*/
/*left: 0;*/
/*top: 70px;*/
/*bottom: 0;*/
/*overflow-y: scroll;*/
/*height: 100%*/
/*}*/

.sidebar::-webkit-scrollbar {
  width: 0;
}

.sidebar-el-menu {
  width: 150px;
}

.sidebar-el-menu-collapse {
  width: 60px;
}

.sidebar > ul {
  height: 100%;
}

.sidebar {
  height: 100%;
}

/* 左侧二级菜单项的样式 */
.el-menu-item {
  font-size: 14px !important;
  padding-left: 15px !important;
  color: black !important;
}

/* 左侧二级菜单选中时的样式 */
.el-menu-item.is-active {
  color: white !important;
  background: #3989fa !important;
}

.el-menu-item, .el-submenu__title {
  height: 50px !important;
  line-height: 50px !important;
}

.collapseBtn {
  position: absolute;
  bottom: 10px;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
