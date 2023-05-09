<template>
  <div>
    <!-- 面包屑导航区 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>单车管理人员管理</el-breadcrumb-item>
      <el-breadcrumb-item>人员列表</el-breadcrumb-item>
    </el-breadcrumb>

    <!-- 卡片视图区域 -->
    <el-card class="box-card">
      <!-- 用户列表区域 -->
      <el-table :data="userlist" stripe="true" border="true">
        <el-table-column type="index"></el-table-column>
        <el-table-column label="姓名" prop="Username"></el-table-column>
        <el-table-column label="创建时间">
          <template slot-scope="scope">
            <!-- addtime 是接口的字段 -->
            <!-- | 管道符 -->
            <!-- 设置全局过滤器 名为dateChage -->
            {{ scope.row.CreatedAt | dateChage }}
          </template>
        </el-table-column>
        <el-table-column label="角色" width="350px">
          <template slot-scope="scope">
            <el-tag
              v-for="tag in scope.row.Roles"
              :key="tag.Name"
              :type="tag.success"
            >
              {{ tag.Name }}
            </el-tag>
          </template>
        </el-table-column>
        <!-- <el-table-column label="操作">
          <template slot-scope="scope">
            
            <el-button
              type="danger"
              icon="el-icon-delete"
              @click="handleDelete(scope.row)"
              style="margin-right: 10px"
            ></el-button>
     
            <el-dropdown
              @command="(command) => handleAllocatingUser(scope.row, command)"
            >
              <el-button type="primary">
                分配角色<i class="el-icon-arrow-down el-icon--right"></i>
              </el-button>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item
                  v-for="role in scope.row.Roles"
                  :key="role.ID"
                  :command="role.ID"
                  >{{ role.Name }}
                </el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </template>
        </el-table-column> -->
      </el-table>
      <!-- 分页区域 -->
      <el-pagination
        @size-change="addDialogVisible"
        @current-change="handleCurrentChange"
        :current-page="queryInfo.page"
        :page-sizes="[1, 3, 5, 10]"
        :page-size="queryInfo.pageSize"
        layout="total, sizes, prev, pager, next, jumper"
        :total="total_items"
      >
      </el-pagination>
    </el-card>
  </div>
</template>
<script>
// import { deleteUsers, getAllocatingUsers } from "@/api/user";
export default {
  data() {
    return {
      queryInfo: {
        // 查询关键字
        //  Username: '',
        // 第几页
        page: 1,
        // 一页几个
        pageSize: 10,
      },
      drownList: [],
      userlist: [],
      total_pages: 0,
      total_items: 0,
    };
  },
  created() {
    this.getUserList();
  },
  methods: {
    async handleDelete(Roles) {
      await deleteUsers(row.Roles[0].ID);
      this.getUserList();
    },
    async handleAllocatingUser(row, command) {
      await getAllocatingUsers(row.ID, command);
      this.getUserList();
    },
    handleClick() {
      alert("button click");
    },
    async getUserList() {
      // console.log(this.$axios)
      await this.$axios
        .get("https://www.paper.matrix-studio.top/api/rbac/user", {
          params: this.queryInfo,
        })
        .then((response) => {
          console.log(response);
          this.userlist = response.users;
          this.total_pages = response.total_pages;
          console.log(this.userlist);
          this.total_items = this.total_pages * this.queryInfo.pageSize;
        });
    },

    // 监听pagesize改变的事件
    handleSizeChange(newSize) {
      // console.log(newSize)
      this.queryInfo.pageSize = newSize;
      this.getUserList();
    },
    // 监听 页码值 改变的事件
    handleCurrentChange(newPage) {
      // console.log(newPage)
      this.queryInfo.page = newPage;
      this.getUserList();
    },
  },
};
</script>
<style scoped>
.el-dropdown {
  vertical-align: top;
}

.el-dropdown + .el-dropdown {
  margin-left: 15px;
}

.el-icon-arrow-down {
  font-size: 12px;
}
</style>
