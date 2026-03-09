<template>
  <div>
    <el-upload
      :action="`${getBaseUrl()}/fileUploadAndDownload/upload`"
      :on-error="uploadError"
      :on-success="uploadSuccess"
      :show-file-list="false"
      :data="{'classId': props.classId}"
      :headers="{'x-token': token}"
      multiple
      class="upload-btn"
    >
      <el-button type="primary" :icon="Upload">普通上传</el-button>
    </el-upload>
  </div>
</template>

<script setup>
  import { ref } from 'vue'
  import { ElMessage } from 'element-plus'
  import { isVideoMime, isImageMime } from '@/utils/image'
  import { getBaseUrl } from '@/utils/format'
  import { Upload } from "@element-plus/icons-vue";
  import { useUserStore } from "@/pinia";

  defineOptions({
    name: 'UploadCommon'
  })

  const userStore = useUserStore()

  const token = userStore.token

  const props = defineProps({
    classId: {
      type: Number,
      default: 0
    }
  })

  const emit = defineEmits(['on-success'])

  const fullscreenLoading = ref(false)

  const uploadSuccess = (res) => {
    const { data } = res
    if (data.file) {
      emit('on-success', data.file.url)
    }
  }

  const uploadError = () => {
    ElMessage({
      type: 'error',
      message: '上传失败'
    })
    fullscreenLoading.value = false
  }
</script>
