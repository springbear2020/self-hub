<script setup>
  import { formatBoolean } from '@/utils/format'

  defineProps({
    active: {
      type: Number,
      required: true,
      default: 0
    }
  })
</script>

<template>
  <div>
    <el-tag type="success" v-if="active === 1"
    >{{ formatBoolean(active) }}
    </el-tag>
    <el-tag type="danger" v-else-if="active === 0"
    >{{ formatBoolean(active) }}
    </el-tag>
  </div>
</template>
