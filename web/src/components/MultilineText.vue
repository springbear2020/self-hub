<template>
  <span v-html="formattedText" />
</template>

<script setup>
  import { computed } from 'vue'

  defineOptions({
    name: 'MultilineText'
  })

  const props = defineProps({
    text: String
  })

  const formattedText = computed(() => {
    return props.text?.replace(/\n/g, '<br/>') ?? ''
  })
</script>
