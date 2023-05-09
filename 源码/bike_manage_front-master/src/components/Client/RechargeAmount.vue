<template>
  <div>
    <!-- 面包屑导航区 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>客户管理</el-breadcrumb-item>
      <el-breadcrumb-item>充值金额管理</el-breadcrumb-item>
    </el-breadcrumb>

    <!-- 卡片视图区域 -->
    <el-card class="box-card">

      <!-- 用户列表区域 -->
      <el-table :data="userlist" stripe="true" border="true">
        <el-table-column label="用户ID" prop="ID" align="center">
        </el-table-column>
        <el-table-column label="用户名" prop="username" align="center"></el-table-column>
        <el-table-column label="创建时间" prop="CreatedAt" align="center">
          <template slot-scope="scope">
            {{ scope.row.CreatedAt | dateChage }}
          </template>
        </el-table-column>
        <el-table-column label="充值金额/元" prop="bill" align="center"></el-table-column>
      </el-table>

    </el-card>
  </div>
</template>
<script>

export default {
  data() {
    return {

      addFormRules: {},
      queryInfo: {
        // 第几页
        page: 1,
        // 一页几个
        pageSize: 10,
      },
      userlist: [],
      total_pages: 0,
    }
  },
  created() {
    this.getUserList()
  },
  methods: {
    async getUserList() {
      // console.log(this.$axios)
      await this.$axios.get(' https://www.paper.matrix-studio.top/api/user/bill/list', { params: this.queryInfo }).then(response => {
        console.log(response)
        this.userlist = response.detail


      })
    },

    addDialogClosed() {
      this.$refs.addFormRef.resetFields()
    },
    
  }
}
</script>