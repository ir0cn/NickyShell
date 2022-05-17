<template>
  <div class="content-box">
    <div class="container">
      <vue3-tree-org
          :data="data"
          collapsable="collapsable"
          :default-expand-level="level"
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
import {fetch, post} from '@/util/http'

export default {
  name: "organization",
  components: {Vue3TreeOrg},
  data() {
    return {
      level: 5,
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
      console.log("node delete:", data)
    },
    nodeBlur(e, data) {
      if (data.newNode === true) {
        console.log('new node:', data)
        post("/organization", data).then(res => {
          this.$message(res.data)
        }).catch(err => {
          this.$message(err)
        })
      } else {
        console.log("edit node:", data)
      }
    },
    getOrganization() {
      fetch("/organization")
          .then(data => {
            this.data = data.data
          })
          .catch(err => {
            this.$message(err)
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

