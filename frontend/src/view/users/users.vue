<template>
  <div class="content-box">
    <div class="container" style="min-width: 960px;">
      <el-button type="primary">
        <el-icon>
          <refresh/>
        </el-icon>
      </el-button>
      <el-button type="primary" @click="visibleUser=true">新建</el-button>
      <div style="display: flex; width:260px; align-items: center; float:right ;"> <!--搜索栏-->
        <el-input style="margin-right: 10px" placeholder="用户名 | 手机号 | 邮箱"></el-input>
        <el-button type="primary">
          <el-icon>
            <search/>
          </el-icon>
        </el-button>
      </div>

      <div class="table-container">
        <el-table :data="list" style="width: 100%;" @selection-change="handleSelectionChange" border>
          <el-table-column type="selection" align="center"></el-table-column>
          <el-table-column type="index" label="序号" align="center" width="60px"></el-table-column>
          <el-table-column label="用户名" align="center">
            <template #default="scope">{{ scope.row.username }}</template>
          </el-table-column>
          <el-table-column label="手机号" align="center">
            <template #default="scope">{{ scope.row.phone }}</template>
          </el-table-column>
          <el-table-column label="邮箱" align="center">
            <template #default="scope">{{ scope.row.email }}</template>
          </el-table-column>
          <el-table-column label="角色" align="center">
            <template #default="scope">
              <el-tag v-for="r in scope.row.roles" :key="r" style="margin: 0 5px;"> {{ r }}</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="MFA" width="100px" align="center">
            <template #default="scope">
              <el-tag type="success" v-if="scope.row.mfa==='enabled'">正常</el-tag>
              <el-tag type="info" v-else>禁用</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="状态" width="100px" align="center">
            <template #default="scope">
              <el-tag type="success" v-if="scope.row.status==='enabled'">正常</el-tag>
              <el-tag type="info" v-else>禁用</el-tag>
            </template>
          </el-table-column>
          <el-table-column label="操作" width="150" align="center" fixed="right">
            <template #default="scope">
              <el-button type="text" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
              <el-button type="text" v-if="scope.row.status==='disabled'"
                         @click="handleEnable(scope.$index, scope.row, true)">启用
              </el-button>
              <el-button type="text" v-else @click="handleEnable(scope.$index, scope.row, false)">禁用</el-button>
              <el-button type="text">更多</el-button>
            </template>
          </el-table-column>
        </el-table>
      </div>

      <div class="emptySpace"></div>
      <el-button type="primary" :disabled="selectedList.length===0">启用</el-button>
      <el-button type="primary" :disabled="selectedList.length===0">禁用</el-button>
      <el-button type="primary" :disabled="selectedList.length===0">启用MFA</el-button>
      <el-button type="primary" :disabled="selectedList.length===0">禁用MFA</el-button>
      <el-button type="primary" :disabled="selectedList.length===0">删除</el-button>
      <el-pagination style="float: right;"
                     layout="total,sizes,prev,pager,next,jumper"
                     :total="list.length"
                     :current-page="listQuery.pageNum"
                     :page-sizes="[10,20,50]"
                     :page-size="listQuery.pageSize"
                     @size-change="handleSizeChange"
                     @current-change="handleCurrentChange"
      >
      </el-pagination>
    </div>
    <el-dialog title="新建用户" v-model="visibleUser" width="600px">
      <div
          style="width: 100%; height: 100%; display: flex; align-items: center; flex-direction: column; padding:0 ; margin:0;">
        <div class='el-dialog-div-newuser'>
          <span>用户名：</span>
          <el-input class="el-input-newuser"></el-input>
        </div>
        <div class='el-dialog-div-newuser'>
          <span>手机号：</span>
          <el-input class="el-input-newuser"></el-input>
        </div>
        <div class='el-dialog-div-newuser'>
          <span>邮箱：</span>
          <el-input class="el-input-newuser"></el-input>
        </div>
        <div class='el-dialog-div-newuser'>
          <span>角色：</span>
          <el-select class="el-input-newuser"></el-select>
        </div>
        <div class='el-dialog-div-newuser'>
          <span>MFA：</span>
          <div class="el-input-newuser" style="display: flex; justify-content: space-between">
            <el-radio v-model="enableMFA" :lable=true size="large">启用</el-radio>
            <el-radio v-model="enableMFA" :label=false size="large">禁用</el-radio>
          </div>
        </div>
        <div style="margin-top:20px;">
          <el-button type="primary">添加</el-button>
          <el-button type="primary" @click="visibleUser=false">取消</el-button>
        </div>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import {Refresh} from "@element-plus/icons-vue";

const defaultListQuery = {
  pageNum: 1,
  pageSize: 10,
  keyword: ""
};

export default {
  components: {
    Refresh
  },
  data() {
    return {
      listQuery: Object.assign({}, defaultListQuery),
      list: [
        {
          username: 'admin',
          phone: '13344445555',
          email: 'abc@gg.com',
          roles: ['admin'],
          status: 'enabled',
        },
        {
          username: 'user',
          phone: '01088224393',
          email: '99show@hello.com',
          roles: ['admin', 'user'],
          status: 'disabled'
        }
      ],
      total: null,
      listLoading: true,
      selectedList: [],
      visibleUser: false,
      enableMFA: false,
    }
  },
  created() {
  },
  methods: {
    handleSelectionChange(val) {    // 批量选择行数
      this.selectedList = val;
      console.log(this.selectedList)
    },
    handleEdit(idx, row) {
      console.log('index:' + idx)
      console.log('row:' + row.username)
    },
    handleEnable(idx, row, enable) {
      row.status = enable === true ? 'enabled' : 'disabled'
    },
    handleSizeChange(val) {  // 改变列表显示条数
      this.listQuery.pageNum = 1;
      this.listQuery.pageSize = val;
    },
    handleCurrentChange(val) {  // 改变列表显示页数
      this.listQuery.pageNum = val;
    },
  }
}
</script>

<style scoped>
.emptySpace {
  height: 10px;
}

.el-input-newuser {
  width: 260px;
}

.el-dialog-div-newuser {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 360px;
  margin: 10px 0;
}

</style>
