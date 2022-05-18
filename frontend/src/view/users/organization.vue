<template>
  <div class="content-box">
    <div class="container">
      <vue3-tree-org
          :data="data"
          collapsable="collapsable"
          :default-expand-level="5"
          :node-draggable="false"
          :define-menus="menus"
          @on-node-blur="nodeBlur"
          @on-node-delete="nodeDelete"
      />
    </div>
  </div>
</template>

<script>
import {Vue3TreeOrg} from "vue3-tree-org";
import "vue3-tree-org/lib/vue3-tree-org.css";
import {fetch, post, put } from '@/util/http'

export default {
  name: "organization",
  components: {Vue3TreeOrg},
  data() {
    return {
      data: {id: 0, pid: 0, label: '默认组织'},
      menus: [
        {name: '复制组织名称', command: 'copy'},
        {name: '新增子组织', command: 'add'},
        {name: '编辑组织节点', command: 'edit'},
        {name: '删除组织节点', command: 'delete'}
      ]
    }
  },
  methods: {
    nodeDelete(data) {
      console.log("remove:",{data})
      put("/organization", data)
          .then(() => {
            this.$message.success('删除成功')
          })
          .catch(err => {
            this.$message.error(err)
          })
    },
    nodeBlur(e, data) {
      post("/organization", data)
          .then(() => {
            this.$message.success('更新成功')
          })
          .catch(err => {
            this.$message.error(err)
          })
    },
    getOrganization() {
      fetch("/organization")
          .then(data => {
            this.data = data
          })
          .catch(err => {
            this.$message.error(err)
          })
    }
  },
  created() {
    this.getOrganization()
  }
}
</script>

<style>
.content-box, .container {
  height: calc(100% - 30px);
  width: calc(100% - 40px);
}
</style>

