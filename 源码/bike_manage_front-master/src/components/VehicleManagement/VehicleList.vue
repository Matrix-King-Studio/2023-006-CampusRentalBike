<template>
  <div>
    <!-- 面包屑导航区 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>车辆管理</el-breadcrumb-item>
      <el-breadcrumb-item>车辆名单</el-breadcrumb-item>
    </el-breadcrumb>

    <!-- 卡片视图区域 -->
    <el-card class="box-card">
      <!-- 搜索与添加区域 -->
      <el-row :gutter="50">
        <el-col :span="8">
          <el-input placeholder="请输入车辆ID" type="text" v-model="detail.ID">
            <el-button slot="append" icon="el-icon-search" @click="searchBike"></el-button>
          </el-input>
        </el-col>
        <el-col :span="4">
          <el-button type="primary" @click="addDialogVisible = true">添加车辆</el-button>
        </el-col>
      </el-row>
      <!-- 用户列表区域 -->
      <el-table :data="bikelist" stripe="true" border="true">
        <el-table-column type="index"></el-table-column>
        <el-table-column label="车辆ID" prop="ID" align="center">
        </el-table-column>
        <el-table-column label="创建时间" align="center" width="300px">
          <template slot-scope="scope">
            {{ scope.row.CreatedAt | dateChage }}
          </template>
        </el-table-column>
        <el-table-column label="是否损坏" align="center">
          <template slot-scope="scope">
            <div v-if="scope.row.IsFault">是</div>
            <div v-if="!scope.row.IsFault">否</div>
          </template>
        </el-table-column>
        <el-table-column label="是否超出服务区" align="center">
          <template slot-scope="scope">
            <div v-if="scope.row.IsFault">是</div>
            <div v-if="!scope.row.IsFault">否</div>
          </template>
        </el-table-column>
        <el-table-column label="车辆位置" header-align="center">
          <el-table-column label="纬度" align="center">
            <template slot-scope="scope">
              <div v-for="Location in scope.row.Locations" v-bind:key="Location">
                <el-tag>{{ Location.latitude }}</el-tag>
              </div>
            </template>
          </el-table-column>
          <el-table-column label="经度" align="center">
            <template slot-scope="scope">
              <div v-for="Location in scope.row.Locations" v-bind:key="Location">
                <el-tag>{{ Location.longitude }}</el-tag>
              </div>
            </template>
          </el-table-column>
        </el-table-column>
        <el-table-column label="操作" width="300px" align="center">
          <template slot-scope="scope">
            <!-- 删除按钮 -->
            <el-button type="danger" icon="el-icon-delete" @click="deleteBike(scope.row.ID)"></el-button>
          </template>
        </el-table-column>
      </el-table>
      <!-- 分页区域 -->
      <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange" :current-page="queryInfo.page"
        :page-sizes="[1, 3, 5,10]" :page-size="queryInfo.limit" layout="total, sizes, prev, pager, next, jumper"
        :total="total_items">
      </el-pagination>
    </el-card>
    <!-- 添加用户的对话框 -->
    <el-dialog title="添加用户" :visible.sync="addDialogVisible" width="50%" @close="addDialogClosed">
      <!-- 内容主体区域 -->
      <el-form :model="addForm" :rules="addFormRules" ref="addFormRef" label-width="70px">
        <el-form-item label="经度" prop="latitude">
          <el-input v-model.number="addForm.latitude" type="number"></el-input>
        </el-form-item>
        <el-form-item label="纬度" prop="longitude">
          <el-input v-model.number="addForm.longitude" type="number"></el-input>
        </el-form-item>
      </el-form>
      <!-- 底部区域 -->
      <span slot="footer" class="dialog-footer">
        <el-button @click="addDialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="addBike">确 定</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>

export default {
  data() {
    return {
      addFormRules: {
        latitude: [ ],
      },
      queryInfo: {
        query: '',
        // 当前页数
        page: 1,
        // 当前每页显示多少条数据
        limit: 5,
      },
      bikelist: [],
      detail: {},
      total_pages: 0,
      //控制添加用户对话框的显示和隐藏
      addDialogVisible: false,
      modification: false,
      // 添加用户的表单数据
      addForm: {
        latitude: '',
        longitude: ''
      },
      total_items: 0,
    }
  },
  created() {
    this.getBikeList()
  },
  methods: {
    async getBikeList() {
      // console.log(this.$axios)
      await this.$axios.get('https://www.paper.matrix-studio.top/api/bike/bike_lst', { params: this.queryInfo }).then(response => {
        this.bikelist = response.bikes
        this.total_pages = response.total_pages
        this.total_items = this.total_pages * this.queryInfo.limit
      })
    },
    // 监听pagesize改变的事件
    handleSizeChange(newSize) {
      // console.log(newSize)
      this.queryInfo.limit = newSize
      this.getBikeList()
    },
    // 监听 页码值 改变的事件
    handleCurrentChange(newPage) {
      // console.log(newPage)
      this.queryInfo.page = newPage
      this.getBikeList()
    },
    // 监听添加用户对话框的关闭事件
    addDialogClosed() {
      this.$refs.addFormRef.resetFields()
    },
    addBike() {
      if (this.addForm.latitude != '') {
        this.addForm.latitude = parseFloat(this.addForm.latitude)
        this.addForm.longitude = parseFloat(this.addForm.longitude)
      }

      this.$refs.addFormRef.validate(async valid => {
        if (!valid) return
        const res = await this.$axios.post('https://www.paper.matrix-studio.top/api/bike/create', this.addForm)
        console.log(res, 'hhhh')
        if (res.status !== "success") {
          this.$message.error('添加车辆失败！')
        } else {
          this.$message.success('添加车辆成功！')
          this.addDialogVisible = false
          this.getBikeList()
        }

      })
    },
    async searchBike() {
      await this.$axios.get('https://www.paper.matrix-studio.top/api/bike/search_bike?bike_id=100000008',).then(response => {
        this.bikelist = []
        this.bikelist.push(response.detail)
        this.total_items = 1
      })
    },
    async deleteBike(id) {
      await this.$axios.get('https://www.paper.matrix-studio.top/api/bike/delete_bike?bike_id='+id,).then(response => { 
        location.reload()
      })
    },
  }
}
</script>