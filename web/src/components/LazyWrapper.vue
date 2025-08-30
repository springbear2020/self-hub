<template>
  <div ref="container" style="min-height: 300px">
    <slot v-if="visible" />
    <el-skeleton v-else :rows="4" animated />
  </div>
</template>

<script setup>
  import { ref, onMounted, onBeforeUnmount } from 'vue'

  let observer
  const visible = ref(false)
  const container = ref(null)

  onMounted(() => {
    observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          visible.value = true
          observer.disconnect()
        }
      },
      {
        rootMargin: '100px' // 提前 100px 加载
      }
    )
    observer.observe(container.value)
  })

  onBeforeUnmount(() => {
    if (observer) {
      observer.disconnect()
    }
  })
</script>
