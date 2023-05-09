<template>
    <div>
        <!-- 面包屑导航区 -->
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>客户管理</el-breadcrumb-item>
            <el-breadcrumb-item>客户反馈</el-breadcrumb-item>
        </el-breadcrumb>

        <!-- 卡片视图区域 -->
        <el-card class="box-card">
            <el-table :data="feedback" stripe="true" border="true">
                <el-table-column type="index" align="center"></el-table-column>
                <el-table-column label="用户名" prop="username" align="center" width="150px"></el-table-column>
                <el-table-column label="用户反馈" prop="content" align="center">
                    <template slot-scope="scope">
                        {{ scope.row.content | addQuote }}
                    </template>
                </el-table-column>
                <el-table-column label="反馈时间" prop="CreatedAt" align="center" width="200px">
                    <template slot-scope="scope">
                        {{ scope.row.CreatedAt | dateChage }}
                    </template>
                </el-table-column>
            </el-table>
        </el-card>
    </div>
</template>
<script>

export default {
    filters:{
        addQuote(value){
            return`"${value}"`
        }
    },
    data() {
        return {
            feedback: []
        }
    },
    created() {
        this.getUserfeedback()
    },
    methods: {
        async getUserfeedback() {
            // console.log(this.$axios)
            await this.$axios.get('https://www.paper.matrix-studio.top/api/user/feedback/list', { params: this.queryInfo }).then(response => {
                console.log(response)
                this.feedback = response.detail
                console.log(11111111111)
            })
        }
    }
}
</script>