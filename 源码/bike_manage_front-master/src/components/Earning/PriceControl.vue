<template>
  <div>
    <!-- 面包屑导航区 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>收益管理</el-breadcrumb-item>
      <el-breadcrumb-item>价格管理</el-breadcrumb-item>
    </el-breadcrumb>
    <!--  -->
    <el-card>
      <h2>价格表单</h2>

      <el-table :data="products">
        <el-table-column type="index"></el-table-column>
        <!-- <el-table-column prop="name" label="车辆ID"></el-table-column> -->
        <el-table-column prop="name1" label="车辆类型"></el-table-column>
        <el-table-column label="单车押金/元">
          <template slot-scope="scope">
            <div v-if="!scope.row.isEdit">{{ scope.row.price }}</div>
            <div v-else>
              <el-input-number v-model="scope.row.price" controls-position="right" @blur="savePrice(scope.$index)"
                :step="0.1"></el-input-number>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="每小时/元">
          <template slot-scope="scope">
            <div v-if="!scope.row.isEdit">{{ scope.row.price1 }}</div>
            <div v-else>
              <el-input-number v-model="scope.row.price1" controls-position="right" @blur="savePrice(scope.$index)"
                :step="0.1"></el-input-number>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="每公里/元">
          <template slot-scope="scope">
            <div v-if="!scope.row.isEdit">{{ scope.row.price2 }}</div>
            <div v-else>
              <el-input-number v-model="scope.row.price2" controls-position="right" @blur="savePrice(scope.$index)"
                :step="0.1"></el-input-number>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="操作">
          <template slot-scope="scope">
            <el-button type="primary" v-if="!scope.row.isEdit" @click="editPrice(scope.$index)"
              icon="el-icon-edit"></el-button>
            <el-button v-else @click="savePrice(scope.$index)">保存</el-button>
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
      item: [],
      products: localStorage.getItem('products')
        ? JSON.parse(localStorage.getItem('products')) : [
          // { name: '001', name1: '自行车', price: 5.5, price1: 0.1, price2: 0.1, isEdit: false },

        ]
    }
  },
  created() {
    this.getitem()
  },
  methods: {
    async changeitem(dic) {
      await this.$axios.post('https://www.paper.matrix-studio.top/api/price/change', dic).then(response => {
        // 处理响应
      });
    },

    savePrice(index) {
      this.products[index].isEdit = false
      var dic = {}; // 初始化 dic 为空对象
      dic["price"] = this.products[index].price;
      dic["id"] = this.products[index].name;
      dic["type"] = this.products[index].name1;
      dic["start_price"] = this.products[index].price1;
      dic["km_price"] = this.products[index].price2;
      this.changeitem(dic);
    },
    editPrice(index) {
      this.products[index].isEdit = true;
    },
    async getitem() {
      await this.$axios.get(' https://www.paper.matrix-studio.top/api/price/list',).then(response => {
        for (var i of response.detail) {
          this.item.push({ name: i.id, name1: i.type, price: i.price, price1: i.start_price, price2: i.km_price, isEdit: false })
        }
        this.products = this.item
      })
    },

  }
}
</script>