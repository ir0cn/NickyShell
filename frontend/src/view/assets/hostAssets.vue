<template>
  <div class="content-box">
    <div class="container">
      <el-button type="primary">
        <el-icon>
          <refresh/>
        </el-icon>
      </el-button>
      <!--      <el-icon><refresh /></el-icon>-->
      <el-button type="primary" @click="visibleAdd = true">添加</el-button>
      <el-button type="primary" @click="visibleImport = true">导入</el-button>
      <el-container style="height:calc( 100% - 32px);">  <!-- 减掉一个button的高度 -->
        <el-aside class="left">
          <el-tree :data="assetGroup" :props="defaultProps" @node-click="handleGroupSelect">
            <template #default="{ node, data }">
              <div class="custom-tree-node">
                <div style="display: flex; justify-content: center; align-items: center">
                  <el-image v-if="data.iconUrl && node.isLeaf!==true" :src="data.iconUrl" class="imgLogo"
                            fit="scale-down"></el-image>
                  <el-image v-else-if="node.isLeaf!==true" :src="require('../../assets/img/default.png')"
                            class="imgLogo" fit="scale-down"/>
                  <span :style="node.isLeaf ? '' : 'font-weight: bold;'"> {{ node.label }} </span>
                </div>
                <div>
                  <el-icon v-if="data.iconAdd === true" class="dynamicIcon">
                    <plus/>
                  </el-icon>
                  <el-icon class="dynamicIcon">
                    <delete/>
                  </el-icon>
                </div>
              </div>
            </template>
          </el-tree>
        </el-aside>
        <el-main class="right">
          <el-table :data="assets">
            <el-table-column type="selection" align="center"></el-table-column>
            <el-table-column type="index" label="序号" align="center" width="60px"></el-table-column>
            <el-table-column label="资产名称" min-width="160px" align="center">
              <template #default="scope">{{ scope.row.assetName }}</template>
            </el-table-column>
            <el-table-column label="资产类型" min-width="82px" align="center">
              <template #default="scope">{{ scope.row.assetType }}</template>
            </el-table-column>
            <el-table-column label="状态" width="100px" align="center">
              <template #default="scope">
                <el-tag type="success" v-if="scope.row.status==='enabled'">正常</el-tag>
                <el-tag type="info" v-else>禁用</el-tag>
              </template>
            </el-table-column>
            <el-table-column label="操作" min-width="140px" align="center">
              <template #default="scope">
                <el-button type="text" @click="handleLogin(scope.row)">登录</el-button>
                <el-button type="text" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                <el-button type="text" v-if="scope.row.status==='disabled'"
                           @click="handleEnable(scope.$index, scope.row, true)">启用
                </el-button>
                <el-button type="text" v-else @click="handleEnable(scope.$index, scope.row, false)">禁用</el-button>
                <el-button type="text">更多</el-button>
              </template>
            </el-table-column>
          </el-table>
          <div class="emptySpace"></div>
          <el-pagination style="float: right"
                         layout="total,sizes,prev,pager,next,jumper"
                         :total="assets.length"
                         :current-page="listQuery.pageNum"
                         :page-sizes="[10,20,50]"
                         :page-size="listQuery.pageSize"
                         @size-change="handleSizeChange"
                         @current-change="handleCurrentChange"
          >
          </el-pagination>
        </el-main>
      </el-container>
    </div>
  </div>

  <el-dialog title="添加主机" v-model="visibleAdd">
    <div>
      <el-tree-select :data="assetGroup" style="margin-bottom: 10px; margin-right: 20px"/>
    </div>
    <el-button @click="doCommit" type="primary">确定</el-button>
    <el-button @click="visibleAdd=false" type="primary">取消</el-button>
  </el-dialog>

  <el-dialog title="导入云主机" v-model="visibleImport">
    <div style="display: flex; justify-content: space-between; margin: 5px 40px; padding: 0 40px">
      <div style="display: flex; justify-content: center; align-items: center; flex-direction: column">
        <el-image class="imgLogoImport" :src="require('../../assets/img/aliyun.png')"></el-image>
        <span style="font-size: 16px; margin-top:6px;">阿里云</span>
      </div>
      <div style="display: flex; justify-content: center; align-items: center; flex-direction: column">
        <el-image class="imgLogoImport" :src="require('../../assets/img/jdcloud.png')"></el-image>
        <span style="font-size: 16px; margin-top:6px;">京东云</span>
      </div>
      <div style="display: flex; justify-content: center; align-items: center; flex-direction: column">
        <el-image class="imgLogoImport" :src="require('../../assets/img/qcloud.png')"></el-image>
        <span style="font-size: 16px; margin-top:6px;">腾讯云</span>
      </div>
    </div>

    <div style="display: flex; justify-content: center">
    <el-button type="primary">AAA</el-button>
    <el-button @click="visibleImport=false" type="primary">取消</el-button>
    </div>
  </el-dialog>
</template>

<script>

import {Delete, Plus, Refresh} from "@element-plus/icons-vue";

export default {
  name: "hostAssets",
  components: {
    Refresh,
    Plus, Delete,
  },
  data() {
    return {
      listQuery: {
        pageNum: 1,
        pageSize: 10,
      },
      defaultProps: {
        children: 'children',
        label: 'label',
      },
      assetGroup: [{
        label: 'Default',
        iconAdd: true,
        children: [
          {label: 'vpc-100001', id: 'vpc-101'},
          {label: 'vpc-100002', id: 'vpc-102'},
          {label: 'vpc-100003', id: 'vpc-103'},
          {label: 'vpc-100004', id: 'vpc-104'},
        ]
      }, {
        label: '腾讯云',
        iconUrl: require('../../assets/img/qcloud.png'),
        children: [
          {label: 'vpc-002003', id: 'vpc-013'}
        ]
      }, {
        label: '京东云',
        iconUrl: require('../../assets/img/jdcloud.png'),
        children: [
          {label: 'vpc-000002', id: 'vpc-02'},
        ]
      }, {
        label: '阿里云',
        iconUrl: require('../../assets/img/aliyun.png'),
        children: [
          {label: 'vpc-000003', id: 'vpc-03'}
        ]
      }],
      assets: [
        {id: '1', assetId: 'id-1', assetName: '我的Linux服务器', assetType: 'linux', status: 'enabled', ip: '192.168.0.4'},
        {
          id: '1',
          assetId: 'id-2',
          assetName: '我的Window Server',
          assetType: 'windows',
          status: 'disabled',
          ip: '192.168.0.9'
        },
      ],
      visibleAdd: false,
      visibleImport: false,
    }
  },
  methods: {
    handleGroupSelect(data) {
      console.log(data.id)
    },
    handleSizeChange(val) {
      this.listQuery.pageNum = 1;
      this.listQuery.pageSize = val;
    },
    handleCurrentChange(val) {
      this.listQuery.pageNum = val;
    },
    handleEdit(idx, row) {
      console.log('index:' + idx)
      console.log('row:' + row.username)
    },
    handleEnable(idx, row, enable) {
      row.status = enable === true ? 'enabled' : 'disabled'
    },
    doCommit() {
      this.visibleAdd = false
    },
    handleLogin(row) {
      console.log(row)
      window.open("#/remote")
    }
  }
}
</script>

<style scoped>
.content-box {
  height: calc(100% - 20px); /* 减掉padding */
}

.container {
  height: calc(100% - 40px); /* 减掉padding */
}

.left {
  width: 200px;
  margin-top: 10px;
  height: 100%;
  border: 1px solid #f0f0f0;
  border-right: 0;
}

.right {
  width: auto;
  margin-top: 10px;
  height: 100%;
  padding: 0;
  border: 1px solid #f0f0f0;
}

.custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
}

.emptySpace {
  height: 10px;
}

.dynamicIcon {
  font-size: 16px;
  color: #909399;
  margin-right: 4px;
}

.dynamicIcon:hover {
  color: #20a0ff;
}

.imgLogo {
  height: 20px;
  width: 20px;
  margin-right: 6px;
}

.imgLogoImport {
  width: 80px;
  height: 60px;
}
</style>
