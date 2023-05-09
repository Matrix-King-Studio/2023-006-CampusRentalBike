<!--
    【整体思路流程】
    1. 先通过后端接口获取appid等敏感数据，保存到data中。
    2. 将获取到的数据传递给组件，如果正确无误扫码登录框就自动出来了。
    3. 用户扫码登录后，会自动跳转到重定向的回调域地址(解析code)，也就是前端当前的页面。
    4. 解析成功或失败，前端就可以拿着code去找接口换数据了，进行您的逻辑即可。
    ———————————————————————————————————————————————————————————
    【注意事项】
    1. [非常重要] 前端和后端必须都在一个回调域下，比如前端是:www.abc.com/index，
    后端的是:www.abc.com/admin，这样就处于同一个【www.abc.com】回调域下。如果前端是abc，后端是def，则无法互通。
    2. 一定要等敏感数据(appid等)请求过来后，再传递给组件！我是用的if暂时不渲染！
    3. redirect_uri 是回调域，重定向的地址，必须是微信开放平台里设置好的回调域（找后端看）,咱们测试默认的localhost,什么192.168.xx 统统不行，
    所以每次调试必须部署到线上进行测试，非常麻烦。
    如果想要解决，请访问：https://wangjiabin.blog.csdn.net/article/details/127787561
-->

<template>
    <div class="background">
        <div class="border">
            <div class="title" style="padding-top: 42px;">
                <div class="title_item1" style="padding-left: 121px;">
                    欢迎进入共享单车管理系统
                </div>
                <div class="title_item" style="margin-top: 20px; margin-left: 5px;">
                    Welcome to the shared bike management system
                </div>
            </div>
            <!-- 微信扫码登录框 -->
            <!-- 配置说明请参考文章底部Props，或直接访问组件文档：https://github.com/toMatthew/vue-wxlogin -->

            <div v-if="sensitive != null" class="wxlogin">
                <wxlogin
                    href="data:text/css;base64,LmltcG93ZXJCb3ggLnFyY29kZSB7d2lkdGg6IDIwMHB4O30KLmltcG93ZXJCb3ggLnN0YXR1cyB7ZGlzcGxheTogbm9uZX0KLmltcG93ZXJCb3ggLnRpdGxlIHtkaXNwbGF5OiBub25lO30="
                    :appid="sensitive.appid" :scope="sensitive.scope" :redirect_uri="sensitive.redirect_uri" state="" />
            </div>
            <!-- END -->

            <!-- 数据还未请求过来 -->
            <div v-else>二维码加载中...</div>
            <!-- END -->
        </div>


    </div>
</template>

<script>
// 微信登录组件
import wxlogin from 'vue-wxlogin';
import api from '@/api/api'

export default {

    components: { wxlogin },
    data() {

        return {
            sensitive: {
                // 敏感数据(appid等)
                appid: 'wxc8ef5245335f581e',
                // 重定向的路径
                redirect_uri: 'http://www.paper.matrix-studio.top/',
                scope: 'snsapi_login'
            },

        }
    },

    mounted() {
        // 获取要重定向的路径(本页面url)
        this.getRedirect(),
            this.wx_login_interval = setInterval(() => {
                this.testing();
            }, 1000);
        // 获取敏感数据(appid等)
        // this.getSensitive()
    },

    // updated() {
    //     this.testing()

    // },

    methods: {
        /**
         * 获取要重定向的路径(本页面url)
         * @description 赋值到本地,传递给组件
         * @return void
         */
        getRedirect() {
            // 获取本页面的域名,也就是下面的代码(需要UrlEncode编码)
            this.redirect_uri = encodeURIComponent(window.location.href)
        },
        /**
         * 获取敏感数据(appid等)
         * @description 赋值到本地,传递给组件
         * @return void
         */
        // getSensitive() {
        //     // 请求接口
        //     this.$axios.post('/').then(res => {
        //         // console.log(res)
        //         this.sensitive = res.data.data
        //         // 检测是否扫码登录成功(因为用户扫码登录后,会刷新页面,必须进行检测)
        //         this.testing()
        //     })
        // },

        /**
         * 检测是否扫码登录成功
         * @description 判断检测url上有没有code拼接
         * @return void
         */
        testing() {

            // 获取URL上code
            const code = this.getUrlParam('code')
            console.log(code)
            // 判断是否存在code
            if (code == null || code == '' || code == undefined) {
                // code为空
                // 意味着没有扫码登录,您的逻辑...
            } else {
                // code存在,调用接口获取登录信息
                console.log('【当前code】', code)
                this.getInfo(code)
            }
        },

        /**
         * 获取登录信息(一定要后端去获取最终登录数据!)
         * @description 通过拿到的code,去接口换登录信息
         * @param {String} code - code
         * @return 
         */
        getInfo(code) {
            console.log("-------------------",code)
            // 请求接口
            api.wx_login(code).then(res => {
                // 判断code码
                // 获取失败
                if (res.status !== "success") {
                    // 要做的操作...
                    // 大部分是清空url上的code参数并跳转回本页面
                    this.delUrlParam(window.location.href)
                }
                // 获取成功
                else {
                    // 业务逻辑(res就是用户登录信息了,比如写入缓存)
                    console.log(res)
                    clearInterval(this.wx_login_interval);
                    // 回跳到原页面(去掉URL上的code参数)
                    this.delUrlParam(window.location.href)
                    window.sessionStorage.setItem("token", res.token);
                    this.$message({
                        message:'恭喜你，登录成功！',
                        type:'success'
                    })
                    this.$router.push('/home')
                }

            })
        },

        /**
         * 解析URL参数
         * @description 截取路由参数
         * @param {String} name - 要解析的路由参数
         * @return String
         */
        getUrlParam(name) {
            let reg = new RegExp('(^|&)' + name + '=([^&]*)(&|$)')
            let r = window.location.search.substr(1).match(reg)
            if (r != null) {// ok
                return unescape(r[2])
            }
            // false
            return null
        },

        /**
         * 删除url路径指定参数
         * @description 去除地址栏url上的code参数,回到原页面
         * @param {String} url - URL
         * @param {String} key - KEY
         * @return String
         */
        delUrlParam(url, key = 'code') {
            let baseUrl = url.split('?')[0] + '?';
            let query = url.split('?')[1];
            if (query.indexOf(key) > -1) {
                let obj = {};
                let arr = query.split('&');
                for (let i = 0; i < arr.length; i++) {
                    arr[i] = arr[i].split('=');
                    obj[arr[i][0]] = arr[i][1];
                }
                delete obj[key];
                let url =
                    baseUrl +
                    JSON.stringify(obj)
                        .replace(/[\"\{\}]/g, '')
                        .replace(/\:/g, '=')
                        .replace(/\,/g, '&');
                // return url;
                window.history.pushState({}, 0, url);//跳转页面
            } else {
                // return url;
                window.history.pushState({}, 0, url);//跳转页面
            }
        },
    }
}
</script>

<style lang="css" scoped>
.background {
    width: 100%;
    height: 100%;
    /* 加载背景图 */
    background-image: url('../static/login_photo.jpg');
    /* 背景图垂直、水平均居中 */
    background-position: center center;
    /* 背景图不平铺 */
    background-repeat: no-repeat;
    /* 当内容高度大于图片高度时，背景图像的位置相对于viewport固定 */
    background-attachment: fixed;
    /* 让背景图基于容器大小伸缩 */
    background-size: cover;
    /* 设置背景颜色，背景图加载过程中会显示背景色 */
}

.border {
    width: 500px;
    height: 420px;
    box-shadow: 0 5px 8px rgba(0, 0, 0, .12), 0 0 8px rgba(0, 0, 0, .06);
    left: 6%;
    top: 17%;
    position: absolute;
    /* transform: translate(-50%, -50%); */

}

.wxlogin {
    position: absolute;
    left: 50%;
    top: 79%;
    transform: translate(-50%, -50%);

}

.title {
    border-radius: 2px;
    font-size: larger;
    font-style: italic;
    padding: 15px;
}
</style>
