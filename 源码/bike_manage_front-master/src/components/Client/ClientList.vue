<template>
    <div>
      <!-- 面包屑导航区 -->
      <el-breadcrumb separator-class="el-icon-arrow-right">
        <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
        <el-breadcrumb-item>客户管理</el-breadcrumb-item>
        <el-breadcrumb-item>客户名单</el-breadcrumb-item>
      </el-breadcrumb>
  
      <!-- 卡片视图区域 -->
      <el-card class="box-card">

        <!-- 用户列表区域 -->
        <el-table :data="userlist" stripe="true" border="true">
          <el-table-column type="index" align="center"></el-table-column>
          <el-table-column label="用户名" prop="username" align="center" width="120px"></el-table-column>
          <!-- <el-table-column label="密码" prop="password" align="center"></el-table-column> -->
          <el-table-column label="创建时间" prop="CreatedAt" align="center">
            <template slot-scope="scope">
            <!-- addtime 是接口的字段 -->
            <!-- | 管道符 -->
            <!-- 设置全局过滤器 名为dateChage -->
            {{ scope.row.CreatedAt | dateChage }}
          </template>
          </el-table-column>
          <el-table-column label="订单" prop="bill" align="center">
            <template slot-scope="scope">
            <div v-if="scope.row.IsFault">请跳转订单页查询</div>
            <div v-if="!scope.row.IsFault">暂无订单</div>
          </template>
          </el-table-column>
          <el-table-column label="操作" width="200px" align="center">
          <template slot-scope="scope" >
            <!-- 删除按钮 -->
            <el-button type="danger" icon="el-icon-delete"></el-button>
          </template>
        </el-table-column>
        </el-table>

      </el-card>
    </div>
  </template>
  <script>
  
  export default {
    data() {
      return {
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
        await this.$axios.get('https://www.paper.matrix-studio.top/api/user/list', { params: this.queryInfo }).then(response => {
          console.log(response)
          this.userlist = response.detail
        })
      },
    }
  }
  </script>