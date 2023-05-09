<template>
  <div>
    <!-- 面包屑导航区 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>收益管理</el-breadcrumb-item>
      <el-breadcrumb-item>日收益</el-breadcrumb-item>
    </el-breadcrumb>

    <el-card>
      <el-row :gutter="40">
        <el-col :span="9">
          <div>
            <el-statistic group-separator="," :precision="2" :value="value2" :title="title"></el-statistic>
          </div>
        </el-col>
        <el-col :span="6">
          <div>
            <el-statistic title="总收益/元" :value="total">
              <template slot="formatter">
              </template>
            </el-statistic>
          </div>
        </el-col>

        <el-col :span="9">
          <div>
            <el-statistic :value="totalFaultAmount" title="报废损失">
              <template slot="suffix">
                <span @click="like = !like" class="like">
                  <i class="el-icon-star-on" style="color:red" v-show="!!like"></i>
                  <i class="el-icon-star-off" v-show="!like"></i>
                </span>
              </template>
            </el-statistic>
          </div>
        </el-col>
      </el-row>
    </el-card>
    <el-card>
      <div ref="chart" :style="{ width: '100%', height: '400px' }" class="chart"></div>
    </el-card>
  </div>
</template>
  
<script>
import * as echarts from 'echarts';
import * as moment from 'moment';

export default {
  data() {
    return {
      like: true,
      // value1: 4154.564,
      value2: 2222,
      totalFaultAmount: 0,
      title: '储蓄基金',
      // title1: '相比昨日增长',
      chartData: [], // 图表数据
      earndata: [],
      information: {},
      income: {},
      total: 0,
      daily_income_data: []
    }
  },
  created() {
    this.getIncome()
  },
  mounted() {
    this.getEarnData()
  },
  methods: {
  async getEarnData() {
    const response = await this.$axios.get(' https://www.paper.matrix-studio.top/api/income/history',).then(response => {
      this.earndata = response.detail
      for (var i of this.earndata) {
        this.daily_income_data.push(i.amount)
      }
      this.renderChart()
     } )
  },
    renderChart() {
      const chart = echarts.init(this.$refs.chart)
      const currentDate = moment();
      const dateList = [];
      for (let i = 6; i >= 0; i--) {
        const date = moment(currentDate).subtract(i, 'days').format('MM-DD');
        dateList.push(date);
      }
      chart.setOption({
        title: {
          text: '每日收益',
        },
        tooltip: {
          trigger: 'axis',
          axisPointer: {
            type: 'cross',
            label: {
              backgroundColor: '#6a7985',
            },
          },
        },
        legend: {
          data: ['日收益/元'],
        },
        toolbox: {
          feature: {
            saveAsImage: {},
          },
        },
        grid: {
          left: '3%',
          right: '4%',
          bottom: '3%',
          containLabel: true,
        },
        xAxis: {
          type: 'category',
          boundaryGap: false,
          data: dateList,
        },
        yAxis: {
          type: 'value',
        },
        series: [
          {
            name: '日收益/元',
            type: 'line',
            stack: '总量',
            areaStyle: {},
            emphasis: {
              focus: 'series',
            },
            data: this.daily_income_data,
          },
        ],
      })
    },

  
  getIncome() {
    this.$axios.get(' https://www.paper.matrix-studio.top/api/income/information',).then(response => {
      this.value2 = response.totalBikeAmount
      this.totalFaultAmount = response.totalFaultAmount
      this.total = response.total
    })
  },
},
}

</script>
<style lang="scss">
.like {
  cursor: pointer;
  font-size: 25px;
  display: inline-block;
}

.chart {
  padding-top: 50px;
  padding-bottom: 50px;
}
</style>