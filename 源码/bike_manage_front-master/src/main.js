import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/element.js'
//导入全局样式表
import './assets/css/global.css'
// 导入day.js文件
import dayjs from 'dayjs'
import axios from "./api/request.js";
import { Message } from 'element-ui';
//引入ECharts制作数据可视化
import * as echarts from 'echarts';
Vue.prototype.$echarts = echarts
Vue.prototype.$axios = axios;
Vue.prototype.$message = Message
Vue.config.productionTip = false

new Vue({
  router,
  store,
  echarts,
  render: h => h(App)
}).$mount('#app')

// 创建一个全局过滤器，把时间格式转化成YYYY-MM-DD
// 此处也可以用箭头函数
// Vue.filter('dateChage',time => {....代码还是如下 })
Vue.filter('dateChage', function (time) {
  //1.对拿到的time进行格式化处理，得到YYYY-MM-DD HH:mm:ss
  //2. 把格式化结果，return
  //直接调用dayjs()得到的是当前时间
  //dayjs(给定的时间)得到的是指定时间
  const dt = dayjs(time).format('YYYY-MM-DD HH:mm:ss')
  return dt
})

