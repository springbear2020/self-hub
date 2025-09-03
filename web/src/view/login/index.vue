<template>
  <div id="userLayout" class="user-layout">
    <div class="bg-shapes">
      <div class="shape shape-1"></div>
      <div class="shape shape-2"></div>
      <div class="shape shape-3"></div>
    </div>

    <div class="login-container">
      <div class="logo-area">
        <img class="logo-img" :src="$GIN_VUE_ADMIN.appLogo" alt="Logo" />
        <h1 class="app-name">{{ $GIN_VUE_ADMIN.appName }}</h1>
        <p class="app-desc">SelfHub - 安放您数字生活的私密一隅</p>
      </div>

      <el-form
        ref="loginForm"
        :model="loginFormData"
        :rules="rules"
        :validate-on-rule-change="false"
        @keyup.enter="submitForm"
        class="login-form"
      >
        <el-form-item prop="username" class="form-item">
          <el-input
            ref="usernameInputRef"
            v-model="loginFormData.username"
            size="large"
            placeholder="请输入用户名"
            suffix-icon="user"
          />
        </el-form-item>

        <el-form-item prop="password" class="form-item">
          <el-input
            v-model="loginFormData.password"
            show-password
            size="large"
            type="password"
            placeholder="请输入密码"
          />
        </el-form-item>

        <el-form-item
          v-if="loginFormData.openCaptcha"
          prop="captcha"
          class="form-item"
        >
          <div class="captcha-container">
            <el-input
              v-model="loginFormData.captcha"
              placeholder="请输入验证码"
              size="large"
              class="captcha-input"
            />
            <div class="captcha-image" @click="loginVerify()">
              <img
                v-if="picPath"
                :src="picPath"
                alt="请输入验证码"
              />
            </div>
          </div>
        </el-form-item>

        <el-form-item class="form-item">
          <el-button
            class="login-btn"
            type="primary"
            size="large"
            @click="submitForm"
          >登 录
          </el-button
          >
        </el-form-item>
      </el-form>

      <!--
      <div class="bottom-links">
        <a href="https://www.gin-vue-admin.com/" target="_blank">
          <img src="@/assets/docs.png" alt="文档" />
        </a>
        <a href="https://support.qq.com/product/371961" target="_blank">
          <img src="@/assets/kefu.png" alt="客服" />
        </a>
        <a
          href="https://github.com/flipped-aurora/gin-vue-admin"
          target="_blank"
        >
          <img src="@/assets/github.png" alt="github" />
        </a>
        <a href="https://space.bilibili.com/322210472" target="_blank">
          <img src="@/assets/video.png" alt="视频站" />
        </a>
      </div>
      -->

      <div class="copyright">
        © 2025 SelfHub · 基于 <a href="https://github.com/flipped-aurora/gin-vue-admin"
                                 target="_blank">Gin-Vue-Admin</a> 构建
      </div>
    </div>
  </div>
</template>

<script setup>
  import { captcha } from '@/api/user'
  import { checkDB } from '@/api/initdb'
  import BottomInfo from '@/components/bottomInfo/bottomInfo.vue'
  import { onMounted, reactive, ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { useRouter } from 'vue-router'
  import { useUserStore } from '@/pinia/modules/user'

  defineOptions({
    name: 'Login'
  })

  const router = useRouter()
  // 验证函数
  const checkUsername = (rule, value, callback) => {
    if (value.length < 5) {
      return callback(new Error('请输入正确的用户名'))
    } else {
      callback()
    }
  }
  const checkPassword = (rule, value, callback) => {
    if (value.length < 6) {
      return callback(new Error('请输入正确的密码'))
    } else {
      callback()
    }
  }

  // 获取验证码
  const loginVerify = async () => {
    const ele = await captcha()
    rules.captcha.push({
      max: ele.data.captchaLength,
      min: ele.data.captchaLength,
      message: `请输入${ele.data.captchaLength}位验证码`,
      trigger: 'blur'
    })
    picPath.value = ele.data.picPath
    loginFormData.captchaId = ele.data.captchaId
    loginFormData.openCaptcha = ele.data.openCaptcha
  }
  loginVerify()

  // 登录相关操作
  const usernameInputRef = ref()
  const loginForm = ref(null)
  const picPath = ref('')
  const loginFormData = reactive({
    username: '',
    password: '',
    captcha: '',
    captchaId: '',
    openCaptcha: false
  })
  const rules = reactive({
    username: [{ validator: checkUsername, trigger: 'blur' }],
    password: [{ validator: checkPassword, trigger: 'blur' }],
    captcha: [
      {
        message: '验证码格式不正确',
        trigger: 'blur'
      }
    ]
  })

  const userStore = useUserStore()
  const login = async () => {
    return await userStore.LoginIn(loginFormData)
  }
  const submitForm = () => {
    loginForm.value.validate(async (v) => {
      if (!v) {
        // 未通过前端静态验证
        ElMessage({
          type: 'error',
          message: '请正确填写登录信息',
          showClose: true
        })
        await loginVerify()
        return false
      }

      // 通过验证，请求登陆
      const flag = await login()

      // 登陆失败，刷新验证码
      if (!flag) {
        await loginVerify()
        return false
      }

      // 登陆成功
      return true
    })
  }

  // 跳转初始化
  const checkInit = async () => {
    const res = await checkDB()
    if (res.code === 0) {
      if (res.data?.needInit) {
        userStore.NeedInit()
        await router.push({ name: 'Init' })
      } else {
        ElMessage({
          type: 'info',
          message: '已配置数据库信息，无法初始化'
        })
      }
    }
  }

  onMounted(() => {
    usernameInputRef.value.focus()
  })
</script>

<style scoped lang="scss">
  .user-layout {
    min-height: 100vh;
    display: flex;
    justify-content: center;
    align-items: center;
    background: linear-gradient(135deg, #0a192f 0%, #0d7377 30%, #00ccc6 100%);
    position: relative;
    overflow: hidden;
    padding: 20px;
  }

  .bg-shapes {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 1;

    .shape {
      position: absolute;
      border-radius: 50%;
      opacity: 0.1;

      &.shape-1 {
        width: 500px;
        height: 500px;
        background: #00ccc6;
        top: -250px;
        left: -100px;
      }

      &.shape-2 {
        width: 400px;
        height: 400px;
        background: #32e0c4;
        bottom: -150px;
        right: -100px;
      }

      &.shape-3 {
        width: 300px;
        height: 300px;
        background: #7fffd4;
        top: 50%;
        left: 70%;
      }
    }
  }

  .login-container {
    position: relative;
    z-index: 2;
    background: rgba(255, 255, 255, 0.15);
    backdrop-filter: blur(10px);
    border-radius: 20px;
    padding: 40px;
    width: 100%;
    max-width: 450px;
    box-shadow: 0 15px 35px rgba(0, 0, 0, 0.2);
    border: 1px solid rgba(255, 255, 255, 0.1);
  }

  .logo-area {
    text-align: center;
    margin-bottom: 30px;
    display: flex;
    flex-direction: column;
    align-items: center;

    .logo-img {
      width: 96px;
      height: 96px;
      margin-bottom: 5px;
      display: block;
      margin-left: auto;
      margin-right: auto;
    }

    .app-name {
      color: white;
      font-size: 28px;
      font-weight: 600;
      margin-bottom: 5px;
    }

    .app-desc {
      color: rgba(255, 255, 255, 0.7);
      font-size: 14px;
    }
  }

  .login-form {
    .form-item {
      margin-bottom: 24px;
    }

    :deep(.el-input__wrapper) {
      background: rgba(255, 255, 255, 0.9);
      border-radius: 10px;
      padding: 4px 15px;
      box-shadow: 0 2px 6px rgba(0, 0, 0, 0.1);
    }

    :deep(.el-input__inner) {
      color: #333;
      font-size: 16px;
    }

    :deep(.el-input__suffix) {
      display: flex;
      align-items: center;
    }
  }

  .captcha-container {
    display: flex;
    gap: 12px;

    .captcha-input {
      flex: 1;
    }

    .captcha-image {
      flex-shrink: 0;
      width: 120px;
      height: 44px;
      background: #c3d4f2;
      border-radius: 10px;
      overflow: hidden;
      cursor: pointer;

      img {
        width: 100%;
        height: 100%;
        object-fit: cover;
      }
    }
  }

  .login-btn {
    width: 100%;
    height: 44px;
    background-color: #00ccc6;
    border: none;
    border-radius: 10px;
    font-size: 16px;
    font-weight: 600;
    box-shadow: 0 4px 10px rgba(0, 204, 198, 0.4);
    transition: all 0.3s ease;

    &:hover {
      background-color: #00b3ad;
      transform: translateY(-2px);
      box-shadow: 0 6px 14px rgba(0, 204, 198, 0.5);
    }
  }

  .bottom-links {
    display: flex;
    justify-content: center;
    gap: 20px;
    margin-top: 30px;

    a {
      display: flex;
      align-items: center;
      justify-content: center;
      width: 36px;
      height: 36px;
      border-radius: 50%;
      background: rgba(255, 255, 255, 0.1);
      transition: all 0.3s ease;

      &:hover {
        background: rgba(0, 204, 198, 0.3);
        transform: translateY(-3px);
      }

      img {
        width: 20px;
        height: 20px;
      }
    }
  }

  .copyright {
    text-align: center;
    margin-top: 20px;
    color: rgba(255, 255, 255, 0.6);
    font-size: 12px;
  }

  /* 响应式设计 */
  @media (max-width: 768px) {
    .login-container {
      padding: 30px 20px;
    }

    .logo-area .app-name {
      font-size: 24px;
    }

    .captcha-container {
      flex-direction: column;
    }

    .captcha-image {
      flex-shrink: 0;
      width: 120px;
      height: 44px;
      background: #b8f3f1;
      border-radius: 10px;
      overflow: hidden;
      cursor: pointer;
    }
  }
</style>