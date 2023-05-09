<template>
  <div class="">
      <!-- 面包屑导航区 -->
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>限制管理</el-breadcrumb-item>
      <el-breadcrumb-item>单车停用限行</el-breadcrumb-item>
    </el-breadcrumb>

    <div id="mapFencecontainer" class="mapBox" style="width:85%;height:83%;position: absolute;">
      <div class="mapBtn">
        <!-- <el-button size="small" type="danger" @click="delpolygon()">删除多边形</el-button> -->
        <el-button size="small" type="danger" @click="clearOverlays()">清除</el-button>
        <el-button size="small" type="primary" @click="sendLngLat">确认</el-button>
        <el-button size="small" type="primary" @click="sendMsg">返回</el-button>
      </div>
    </div>
    <!-- <p>给地图绑定了点击事件，当点击地图时，会在点击位置显示Marker。</p> -->

  </div>
</template>

<script>
import { Notification } from 'element-ui';
let map, marker, polygon, drawingManager, lngLat; // 地图，坐标点，多边形初始化加载，多边形绘制，多边形每个点的集成数组
const markersArray = []; let overlaysArray = []; // 标注点集合,绘制多边形时的集合
let path4 = [];// 设置回显数据参数
// 默认展示假数据回显参数
let arr = [{ 'lat': 34.80506482194067, 'lng': 113.59760284423828 },
           { 'lat': 34.780394716350614, 'lng': 113.59365463256836 },
           { 'lat': 34.760794345237514, 'lng': 113.58541488647461 },
           { 'lat': 34.7330074293797, 'lng': 113.58455657958984 },
           { 'lat': 34.70281206762318, 'lng': 113.61854553222656 },
           { 'lat': 34.70027176815597, 'lng': 113.62884521484375 },
           { 'lat': 34.6987193245323, 'lng': 113.63296508789062 },
           { 'lat': 34.69914272113635, 'lng': 113.63880157470703 },
           { 'lat': 34.699001589175694, 'lng': 113.64498138427734 },
           { 'lat': 34.698578191849535, 'lng': 113.68309020996094 },
           { 'lat': 34.73907339121123, 'lng': 113.68206024169922 },
           { 'lat': 34.738932327388824, 'lng': 113.6920166015625 },
           { 'lat': 34.76319176276739, 'lng': 113.69218826293945 },
           { 'lat': 34.763050740134055, 'lng': 113.70386123657227 },
           { 'lat': 34.7864571976711, 'lng': 113.70214462280273 },
           { 'lat': 34.786739162702524, 'lng': 113.68206024169922 },
           { 'lat': 34.809997957307736, 'lng': 113.68223190307617 },
           { 'lat': 34.81098454894641, 'lng': 113.64789962768555 },
           { 'lat': 34.81366238099493, 'lng': 113.62936019897461 }]; // 用于数据回显的假数据  后期会调用接口返回数据
export default {
  props: ['parentmsg'],
  data() {
    return {
      longitude: '',
      latitude: '',
      lngLatData: [] // 绘制多边形坐标点
    };
  },
  mounted() {
    this.init();
  },
  methods: {

    // 子组件给父组件传值(返回按钮)
    sendMsg() {
        // func: 是父组件指定的传数据绑定的函数，this.msg:子组件给父组件传递的数据
        this.$emit('funcCancle', false);
    },
    // 子组件给父组件传值(确认按钮)
    sendLngLat() {
      
       if (lngLat == undefined) {
         lngLat = path4;
       }

       const h = this.$createElement;
       
    },
    // 地图初始化
   init() {
    map = new qq.maps.Map(document.getElementById('mapFencecontainer'), {
        center: new qq.maps.LatLng(34.764152, 113.667636), // 暂时默认郑州
        zoom: 13
    });
    // 标注的生成与回显
    this.addMarker(new qq.maps.LatLng(30.925788712587014, 103.8922119140625)); // 初始化回显标注
    this.markerOnly();
    // 多边形绘制及回显
    // 设置多边形路径以便回显
    // path4 = [];

    arr.forEach(item => {
      path4.push(new qq.maps.LatLng(item.lat, item.lng));
    });
   
    console.log(path4);
    this.showpolygon(path4);
    // 绘制
    this.addpolygon();
    },
    // 添加监听事件 获取鼠标单击事件（仅留一个标注点）
    markerOnly() {
      qq.maps.event.addListener(map, 'click', function(event) {
          this.addMarker(event.latLng);
          qq.maps.event.addListener(map, 'click', function(event) {
              this.deleteOverlays(); // 删除原有标注仅留下一个
              marker = new qq.maps.Marker({
                  position: event.latLng,
                  map: map
              });
          });
          var gps = event.latLng.getLat() + ',' + event.latLng.getLng(); // 解析出来的点方便给后端
          });
          
      },
    // 添加标记
    addMarker(location) {
      this.deleteOverlays();
      var marker = new qq.maps.Marker({
          position: location,
          map: map
      });
      markersArray.push(marker);
    },
    // 删除标记
    deleteOverlays() {
        if (markersArray) {
            // for (i in markersArray) {
            for (var i = 0;i < markersArray.length;i++) {
                markersArray[i].setMap(null);
            }
            markersArray.length = 0;
        }
        if (marker !== undefined) {
            marker.setMap(null);
        }
    },
    /**
 * 多边形绘制
*/
  // 绘制
  addpolygon() {
      drawingManager = new qq.maps.drawing.DrawingManager({
          drawingMode: qq.maps.drawing.OverlayType.POLYGON, // 默认选中多边形绘制
          drawingControl: true,
          drawingControlOptions: {
            position: qq.maps.ControlPosition.TOP_CENTER,
          //   drawingModes: [
          //     qq.maps.drawing.OverlayType.POLYGON
          //   ]
          // 表头显示区域
          drawingModes: [
              qq.maps.drawing.OverlayType.MARKER,
              qq.maps.drawing.OverlayType.CIRCLE,
              qq.maps.drawing.OverlayType.POLYGON,
              qq.maps.drawing.OverlayType.POLYLINE,
              qq.maps.drawing.OverlayType.RECTANGLE
              ]
          },
          markerOptions: {
            visible: false
          },
          // 多边形样式
          polygonOptions: {
            editable: true,
            strokeDashStyle:'dash',
            strokeColor: new qq.maps.Color(13, 188, 121, 0.8),
            fillColor: new qq.maps.Color(13, 188, 121, 0.2),
            clickable: false
          },
          // 圆形样式
          circleOptions: {
            fillColor: new qq.maps.Color(255, 208, 70, 0.3),
            strokeColor: new qq.maps.Color(88, 88, 88, 1),
            strokeWeight: 3,
            clickable: false
          }
      });

      drawingManager.setMap(map);
      qq.maps.event.addListener(drawingManager, 'overlaycomplete', function(event) {
          // clearOverlays(overlaysArray)
          lngLat = [];
          overlaysArray.push(event.overlay);
          for (const item of event.overlay.getPath().elems) {
            const lng = item.getLng();
            const lat = item.getLat();

            lngLat.push({
              lat: lat,
              lng: lng
            });
          }

          console.log(lngLat); // 获得相应的经纬度值
      });
  },
  // 初始化回显
  showpolygon(path3) {
    polygon = new qq.maps.Polygon({
        map: map,
        editable: true,
        strokeColor: new qq.maps.Color(202, 67, 58, 0.8),
        fillColor: new qq.maps.Color(202, 67, 58, 0.1)
        });
        polygon.setPath(path3);
    },
  // 删除初始化多边形
  delpolygon() {
      const a = [];
      polygon.setPath(a);
  },
  // 删除绘制的多边形
 clearOverlays() {
    arr = []; // 清除假数据
    path4 = [];
    polygon.setPath(path4);
    if (overlaysArray) { // 这个if判断在vue中会报错，不知道是里面的for in 被重写还是咋的，原因还不知道，如果报错可以不要这个if判断，删除标注那个地方跟这个原理是一样的，报错的话也不要，把长度重置为0
        // for (i in overlaysArray) {
        for (var i = 0;i < overlaysArray.length;i++) {
            overlaysArray[i].setMap(null);
        }
    }
    overlaysArray = []; // 需要重置为空，否则之前的数据还在这个数组里面
    lngLat = [];
    console.log(drawingManager);
    this.init();
}

  }
};
</script>
<style scoped lang="scss">
  .mapBox {
    position: relative;
    .mapBtn {
      position: absolute;
      bottom: 20px;
      right: 20px;
      z-index: 999;
    }

  }

</style>

