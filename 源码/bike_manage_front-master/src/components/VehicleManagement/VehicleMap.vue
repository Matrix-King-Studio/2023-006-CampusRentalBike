<template>
    <div>
        <!-- 面包屑导航区 -->
        <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>车辆管理</el-breadcrumb-item>
            <el-breadcrumb-item>车辆地图</el-breadcrumb-item>
        </el-breadcrumb>
        <!-- 搜索与添加区域 -->
        <!-- <div class="search">
            <el-row :gutter="20">
                <el-col :span="8">
                    <el-input placeholder="查询车辆ID">
                        <el-button slot="append" icon="el-icon-search"></el-button>
                    </el-input>
                </el-col>
                <el-col :span="4">
                
                </el-col>
            </el-row>
        </div> -->

        <!-- 显示地图的DOM节点 -->
        <div id="container" class="content"></div>
        <!-- END -->
    </div>
</template>
  
<script type="text/javascript">

export default {
    data() {
        return {
            detail: []
        }
    },
    created() {
        this.fetchCoordsFromServer()
    },
    mounted() {
        this.loadMap()

    },

    methods: {

        /**
         * 初始化腾讯地图插件
         * @description 初始化完毕后执行业务操作
         * @return void
         */
        loadMap() {
            try {
                // 您申请的腾讯地图key
                const key = 'CWUBZ-JZ33G-SUXQV-QJ5GZ-UYMGE-3VFXW'
                // 动态创建script标签引入
                var script = document.createElement("script");
                script.type = "text/javascript";
                script.src = "https://map.qq.com/api/gljs?v=1.exp&callback=init&key=" + key;
                script.onload = script.onreadystatechange = () => {
                    if (!this.readyState || this.readyState === "loaded" || this.readyState === "complete") {
                        // 加载完毕,执行业务逻辑函数
                        this.actions();
                        script.onload = script.onreadystatechange = null;
                    }
                };
                document.body.appendChild(script);
            } catch (e) {
                //TODO handle the exception
            }
        },

        /**
         * 显示地图
         * @description 定位北京周边
         * @return void
         */
        actions() {

            // 地图中心点坐标(北京)
            // var center = new TMap.LatLng(39.90801507071309, 116.39093723423957);
            var center = new qq.maps.LatLng(39.089736, 117.704108);

            // 初始化地图(渲染到id=container DOM节点)
            var map = new qq.maps.Map("container", {
                rotation: 360, //设置地图旋转角度
                pitch: 0, //设置俯仰角度（0~45）
                zoom: 15, //设置地图缩放级别
                center: center, //设置地图中心点坐标
            })//初始化marker图
            for (var i = 0; i < this.detail.length; i++) {
                var latitude = this.detail[i].Locations[0].latitude;
                var longitude = this.detail[i].Locations[0].longitude;
                // console.log(666)
                // console.log(latitude, longitude)
                var marker = new qq.maps.Marker({
                    position: new qq.maps.LatLng(latitude, longitude),
                    map: map
                })
                console.log(111111111111)
                console.log(this.detail)
            }




        },
        async fetchCoordsFromServer() {
            // 使用 axios 或者其他库从后台获取坐标数组
            try {
                const response = await this.$axios.get(' https://www.paper.matrix-studio.top/api/bike/location',).then(response => {
                    this.detail = response.detail
                    // console.log(this.detail)
                    this.actions();
                })
            } catch (error) {
                console.log(error)
            }
        }

    }
}
</script>
  
<style scoped>
/* 自定义宽高等样式 */
.content {
    height: 630px;
    width: 100%;
}

.search {
    /* margin: 15px;
     */
    margin-top: 20px;
    margin-bottom: 20px;
}
</style>
  