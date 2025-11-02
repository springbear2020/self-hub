import './style/element_visiable.scss'
import 'element-plus/theme-chalk/dark/css-vars.css'
import { createApp } from 'vue'
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import 'element-plus/dist/index.css'
// 引入gin-vue-admin前端初始化相关内容
import './core/gin-vue-admin'
// 引入封装的router
import router from '@/router/index'
import '@/permission'
import run from '@/core/gin-vue-admin.js'
import auth from '@/directive/auth'
import { store } from '@/pinia'
import App from './App.vue'
import dayjs from 'dayjs'
import 'dayjs/locale/zh-cn'

dayjs.locale('zh-cn')

const app = createApp(App)
app.config.productionTip = false

app
  .use(run)
  .use(ElementPlus, {
    locale: zhCn
  })
  .use(store)
  .use(auth)
  .use(router)
  .mount('#app')
export default app
