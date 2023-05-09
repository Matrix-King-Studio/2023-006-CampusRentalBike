<template>
    <div class="flex">
        <div class="box">
            <el-card shadow="always" class="title">车辆情况汇总</el-card>
            <div class="echart">
                <el-card id="mychart1" :style="myChartStyle"></el-card>
                <el-card id="mychart2" :style="myChartStyle"></el-card>
                <el-card id="mychart3" :style="myChartStyle"></el-card>

            </div>
        </div>

        <!-- 显示地图的DOM节点 -->
        <div id="container" class="content"></div>
        <!-- END -->
    </div>
</template>
  
<script>
import * as echarts from "echarts";

export default {
    data() {
        return {
            // myChart1: ,
            detail: [],
            pieData: [],
            pieName: [],
            detail: {
                FaultBikes: 0,
                OutBikes: 0,
                TotalBikes: 0,
            },
            myChartStyle: { float: "left", width: "33.3%", height: "250px" } //图表样式
        };
    },
    created() {
        this.getPie(),
            this.fetchCoordsFromServer()
    },
    mounted() {
        // this.initDate(); //数据初始化
        this.initEcharts();
        this.loadMap();

    },
    methods: {
        getPie() {
            // console.log('---------getPie---------')
            this.$axios.get(' https://www.paper.matrix-studio.top/api/bike/summary',).then(response => {
                console.log(response.detail)
                this.detail = response.detail
                this.initEcharts()
            })
            // console.log('---------getPie---------')
        },
        /**
     * 初始化腾讯地图插件
     * @description 初始化完毕后执行业务操作
     * @return void
     */
        loadMap() {
            try {
                // 您申请的腾讯地图key
                const key = 'OUJBZ-MNIWJ-Y4DFO-XVRLH-RTAZ5-JKFI2'
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

        // initDate() {
        //     for (let i = 0; i < this.myChart1.length; i++) {
        //         // this.xData[i] = i;
        //         // this.yData =this.xData[i]*this.xData[i];
        //         this.pieName[i] = this.myChart1[i].name;
        //     }

        // },
        initEcharts() {
            // this.getPie();
            console.log('---------initEcharts---------')
            // console.log(this.detail)
            console.log('---------initEcharts---------')
            // 饼图
            const option1 = {

                title: {
                    // 设置饼图标题，位置设为顶部居中
                    text: "全部车辆",
                    top: "45%",
                    left: "center"
                },
                series: [
                    {
                        type: "pie",
                        label: {
                            show: true,
                            formatter: "{b} : {c} ({d}%)" // b代表名称，c代表对应值，d代表百分比
                        },
                        radius: ["50%", "60%"], //饼图半径
                        data: [
                            //     {
                            //     value: this.detail.OutBike,
                            //     name: "损坏车辆"
                            // },
                            {
                                value: this.detail.TotalBikes,
                                name: "车辆总数"
                            },
                        ]
                    }
                ]
            };
            const option2 = {

                title: {
                    // 设置饼图标题，位置设为顶部居中
                    text: "损坏车辆",
                    top: "45%",
                    left: "center"
                },
                series: [
                    {
                        type: "pie",
                        label: {
                            show: true,
                            formatter: "{b} : {c} ({d}%)" // b代表名称，c代表对应值，d代表百分比
                        },
                        radius: ["50%", "60%"], //饼图半径
                        data: [
                            { value: this.detail.FaultBikes, name: '损坏车辆' },
                            { value: this.detail.TotalBikes, name: '全部车辆' },

                        ]

                    }
                ]
            };
            const option3 = {

                title: {
                    // 设置饼图标题，位置设为顶部居中
                    text: "出界车辆",
                    top: "45%",
                    left: "center"
                },
                series: [
                    {
                        type: "pie",
                        label: {
                            show: true,
                            formatter: "{b} : {c} ({d}%)" // b代表名称，c代表对应值，d代表百分比
                        },
                        radius: ["50%", "60%"], //饼图半径
                        data: [
                            { value: this.detail.OutBikes, name: '出界车辆' },
                            { value: this.detail.TotalBikes, name: '全部车辆' },
                        ]

                    }
                ]
            };
            // console.log(this.seriesData);
            // const optionFree = {
            //     xAxis: {},
            //     yAxis: {},
            //     series: [
            //         {
            //             data: this.seriesData,
            //             type: "line",
            //             smooth: true
            //         }
            //     ]
            // };
            this.myChart1 = echarts.init(document.getElementById("mychart1"));
            this.myChart2 = echarts.init(document.getElementById("mychart2"));
            this.myChart3 = echarts.init(document.getElementById("mychart3"));
            this.myChart1.setOption(option1);
            this.myChart2.setOption(option2);
            this.myChart3.setOption(option3);
            //随着屏幕大小调节图表
            window.addEventListener("resize", () => {
                this.myChart1.resize();
            });
            window.addEventListener("resize", () => {
                this.myChart2.resize();
            });
            window.addEventListener("resize", () => {
                this.myChart3.resize();
            });
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
};
</script>

  
<style scoped>
/* 自定义宽高等样式 */

.content {
    height: 320px;
    width: 100%;
    margin-top: 15px;
}

.flex {
    margin: 0;
    padding: 0;
    display: flex;
    flex-direction: column;

}

.box {
    height: 320px;
    width: 100%;
}

.title {
    background-color: #e9e9eb;
    font-style: italic;
    font-size: larger;


}

.echart {
    height: 100%;
    display: flex;
    flex-direction: row;
}
</style>
  