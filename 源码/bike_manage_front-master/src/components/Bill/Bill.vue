<template>
    <div>
        <!-- 面包屑导航区 -->
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>订单信息</el-breadcrumb-item>
            <!-- <el-breadcrumb-item>客户名单</el-breadcrumb-item> -->
        </el-breadcrumb>

        <!-- 卡片视图区域 -->
        <el-card class="box-card">
            
            <!-- 用户列表区域 -->
            <el-table :data="userlist" stripe="true" border="true">
                <el-table-column label="用户名" prop="username" width="90px" align="center"></el-table-column>
                <el-table-column label="车辆ID" prop="bike_id" align="center"></el-table-column>
                <el-table-column label="订单时间" prop="CreatedAt" align="center">
                    <template slot-scope="scope">
                        {{ scope.row.CreatedAt | dateChage }}
                    </template>
                </el-table-column>
                <el-table-column label="收益/元" prop="amount" align="center"></el-table-column>



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
            await this.$axios.get(' https://www.paper.matrix-studio.top/api/income/bill', { params: this.queryInfo }).then(response => {
                console.log(response)
                this.userlist = response.detail
            })
        },

    }
}
</script>