import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../components/Login.vue'//注册登录组件
import Home from '../components/Home.vue'//注册登录组件
import Person from '../components/Person/Person.vue'
import Bikelogin from '../components/Person/Bikelogin.vue'
import ManagementList from '../components/Management/ManagementList.vue'
import VehicleList from '../components/VehicleManagement/VehicleList.vue'
import VehicleMap from '../components/VehicleManagement/VehicleMap.vue'
import Rail from '../components/LimitationManagement/Rail.vue'
import TrafficControl from '../components/LimitationManagement/TrafficControl'
import ClientList from '../components/Client/ClientList.vue'
import Earning from '../components/Earning/Earning'
import PriceControl from '../components/Earning/PriceControl'
import RechargeAmount from '../components/Client/RechargeAmount'
// import Collect from '../components/Earning/Collect'
import Bill from '../components/Bill/Bill'
import Feedback from '../components/Client/Feedback'

Vue.use(VueRouter)

const routes = [
  //注册路由
  { path: '/login', component: Login },
  {
    path: '/home', component: Home,
    redirect: '/person',
    children: [
      { path: '/managementlist', component: ManagementList, },
      { path: '/person', component: Person, },
      { path: '/bikelogin', component: Bikelogin, },
      { path: '/vehiclelist', component: VehicleList, },
      { path: '/vehiclemap', component: VehicleMap, },
      { path: '/rail', component: Rail, },
      { path: '/trafficcontrol', component: TrafficControl, },
      { path: '/clientlist', component: ClientList, },
      { path: '/earning', component: Earning, },
      { path: '/pricecontrol', component: PriceControl, },
      { path: '/rechargeamount', component: RechargeAmount, },
      // { path: '/collect', component: Collect, },
      { path: '/bill', component: Bill, },
      { path: '/feedback', component: Feedback, },

    ]
  },
  ,

  //重定向
  { path: '/', redirect: '/login' }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})
// 路由导航守卫
router.beforeEach((to, from, next) => {
  // to将要访问的路径
  // from 代表从哪个路径跳转而来
  // next 是一个函数，表示放行
  if (to.path === '/login') return next();
  const tokenSter = window.sessionStorage.getItem('token')
  if (!tokenSter) return next('/login')
  next()
})
export default router
