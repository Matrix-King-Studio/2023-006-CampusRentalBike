<!-- 去router导入home组件 -->

<template>
    <el-container class="home-container">
        <!-- 头部区域 -->
        <el-header class="title">
                <span >共享单车后台管理系统</span>
                <button type="info" @click="logout" >退出</button>
        </el-header>
        <!-- 页面主体区域 -->
        <el-container>
            <!-- 侧边栏 -->
            <el-aside :width="isCollapse ? '64px' : '200px'">
                <div class="toggle-buttom" @click="toggleCollapse">|||</div>
                <!-- 侧边栏菜单区 -->
                <el-menu background-color="#252526" text-color="#fff" active-text-color="#ffd04b" unique-opened="true"
                    :collapse="isCollapse" :collapse-transition="false" router :default-active="this.$router.path">
                    <!-- 首页 -->
                    <el-menu-item index="/person">
                        <i class="el-icon-menu"></i>
                        <span slot="title">首页</span>
                    </el-menu-item>

                
                    <!-- 车辆管理 /el-icon-bicycle-->
                    <el-submenu index="3">
                        <template slot="title">
                            <i class="el-icon-bicycle"></i>
                            <span>车辆管理</span>
                        </template>
                
                      
                        <el-menu-item index="/vehiclemap">
                            <template slot="title">
                                <i class="el-icon-bicycle"></i>
                                <span>车辆地图</span>
                            </template>
                        </el-menu-item>
                    
                        <el-menu-item index="/vehiclelist">
                            <template slot="title">
                                <i class="el-icon-bicycle"></i>
                                <span>车辆名单</span>
                            </template>
                        </el-menu-item>
                    </el-submenu>
                    <!-- 单车管理人员管理 /el-icon-s-custom-->
                    <el-submenu index="4">
                        <template slot="title">
                            <i class="el-icon-user"></i>
                            <span>单车管理人员管理</span>
                        </template>
                        <el-menu-item index="/managementlist">
                            <template slot="title">
                                <i class="el-icon-user"></i>
                                <span>人员列表</span>
                            </template>
                        </el-menu-item>

                    </el-submenu>
                    <!-- 限制管理/el-icon-map-location -->
                    <el-submenu index="5">
                        <template slot="title">
                            <i class="el-icon-map-location"></i>
                            <span>限制管理</span>
                        </template>
                        <!-- <el-menu-item index="/trafficcontrol">
                            <template slot="title">
                                <i class="el-icon-map-location"></i>
                                <span>单车停用限行</span>
                            </template>
                        </el-menu-item> -->
                        <el-menu-item index="/rail">
                            <template slot="title">
                                <i class="el-icon-map-location"></i>
                                <span>划定电子围栏</span>
                            </template>
                        </el-menu-item>
                    </el-submenu>
                    <!-- 状态查询 /el-icon-search-->
                    <!-- <el-submenu index="6">
                        <template slot="title">
                            <i class="el-icon-search"></i>
                            <span>状态查询</span>
                        </template>
                        <el-menu-item index="6-1-1">
                            <template slot="title">
                                <i class="el-icon-search"></i>
                                <span>车辆状态查询</span>
                            </template>
                        </el-menu-item>
                        <el-menu-item index="6-1-2">
                            <template slot="title">
                                <i class="el-icon-search"></i>
                                <span>人员状态查询</span>
                            </template>
                        </el-menu-item>
                    </el-submenu> -->
                    <!-- 客户管理 el-icon-s-order-->
                    <el-submenu index="7">
                        <template slot="title">
                            <i class="el-icon-s-order"></i>
                            <span>客户管理</span>
                        </template>
                        <el-menu-item index="/clientlist">
                            <template slot="title">
                                <i class="el-icon-s-order"></i>
                                <span>客户名单</span>
                            </template>
                        </el-menu-item>
                        <el-menu-item index="/feedback">
                            <template slot="title">
                                <i class="el-icon-s-order"></i>
                                <span>客户反馈</span>
                            </template>
                        </el-menu-item>
                        <el-menu-item index="/rechargeamount">
                            <template slot="title">
                                <i class="el-icon-s-order"></i>
                                <span>充值金额管理</span>
                            </template>
                        </el-menu-item>
                    
                    </el-submenu>
                    <!-- 收益管理/el-icon-s-data -->
                    <el-submenu index="8">
                        <template slot="title">
                            <i class="el-icon-s-data"></i>
                            <span>收益管理</span>
                        </template>
                 
                        <el-menu-item index="/earning">
                            <template slot="title">
                                <i class="el-icon-s-data"></i>
                                <span>日收益</span>
                            </template>
                        </el-menu-item>
                        <el-menu-item index="/pricecontrol">
                            <template slot="title">
                                <i class="el-icon-s-data"></i>
                                <span>价格管理</span>
                            </template>
                        </el-menu-item>
                 
                    </el-submenu>
                    <el-menu-item index="/bill">
                        <i class="el-icon-tickets"></i>
                        <span slot="title">订单信息</span>
                    </el-menu-item>
                </el-menu>
            </el-aside>
            <!-- 右侧主题内容 -->
            <el-main><router-view></router-view></el-main>
        </el-container>
    </el-container>
</template>

<script>
export default {

    data() {
        return {
            //是否折叠
            isCollapse: false
        }
    },
    methods: {
        //折叠菜单
        toggleCollapse() {
            this.isCollapse = !this.isCollapse
        },
        // 退出登录
        logout(){
            window.sessionStorage.clear()
            this.$router.push('/login')
        }
    }
}
</script>

<style lang="less"scoped>
.title {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
}

.home-container {
    height: 100%;
}

.el-header {
    background-color: #e9e9eb;
    display: flex;
    justify-content: space-between;
    padding-left: 0;
    align-items: center;
    font-size: 20px;
    font-style: italic;
    font-weight: bolder;
    padding-left: 30px;

    >div {
        display: flex;
        align-items: center;

        span {
            margin-left: 15px;
        }
    }
}

.el-aside {
    background-color: #252526;
}

.el-main {
    background-color: #f4f4f4;
}

.toggle-buttom {
    background-color: #252526;
    font-size: 10px;
    line-height: 24px;
    color: #fff;
    text-align: center;
    letter-spacing: 0.2em;
    cursor: pointer;
    ;
}
</style>