<template>
  <el-card>
    <template #header>
      <div>
        <span class="clearfix">Dashboard</span>
      </div>
    </template>
    <p>Dashboard</p>
  </el-card>
</template>

<script>
export default {
  name: "Dashboard"
}
</script>


<style scoped>
.el-row {
  margin-bottom: 20px;
}

/* 卡片标题字体的样式 */
.clearfix{
  padding-left: 5px;
  font-size: 22px;
  font-weight: 800;
}
/* 四个彩色卡片的相关样式 */
.colorCardsBody{
  padding: 18px 40px;
}
.colorCardsBody span{
  font-size: 16px;
  color: #515a6e;
}
.colorCardsBody p{
  font-family: Roboto-Medium !important;
  padding-top: 15px;
  font-size: 18px;
  font-weight: 800;
}
/* 数据总览区域的整体样式设置 */
.dataOverview{
  padding: 0 30px;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
/* 数据总览里的每一条信息项 */
.overview-list-item {
  padding: 0 25px;
  border-left: 1px solid rgba(220,222,226,.5);
  flex: 1;
}
.overview-list-item:first-child {
  border: none;
}
/* 数据总览里每条信息项的文本部分 */
.overview-list-text {
  color: #515a6e;
  font-size: 14px;
  margin-bottom: 24px;
}
.overview-list-count {
  color: #17233d;
  font-weight: 700;
  margin-bottom: 16px;
  font-size: 16px;
}
.overview-list-trend-up {
  font-size: 14px;
  color: blue;
}
.overview-list-trend-down {
  font-size: 14px;
  color: red;
}
.programmaticBatchBtn{    /* 程序化批量的按钮样式 */
  float: right;
  padding: 4px 11px 3px 11px;
  margin-left: 20px;
  margin-right: 50px;
  font-size: 14px;
  color: rgb(255, 255, 255);
  background-color: rgb(0, 182, 151);
}
.programmaticBatchBtn img{     /* 程序化批量按钮上的动图标 */
  width: 24px;
  display: inline-block;
  vertical-align: middle;    /* 让该图标在 Y 轴上居中 */
}
.search-form{      /* 日期筛选控件的样式 */
  display: inline;
  margin-left: 20px;
  margin-right: 8px;
  float: right;
}
.search-form span,.search-form-2 span{    /* 筛选控件的前面字体样式 */
  margin-right: 10px;
  font-weight: 700;
}
.search-form-2{      /* “效果总览” 卡片里的筛选控件定位设置 */
  margin-top:30px;
  margin-bottom: -40px;
  margin-left: 30px;
}
.blankContent{    /* TOP 创意卡片底部空白部分，为了和 TOP 广告底部对齐 */
  width: auto;
  height: 40px;
}
</style>
