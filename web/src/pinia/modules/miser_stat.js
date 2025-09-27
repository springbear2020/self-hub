import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'

export const useMiserStatStore = defineStore('miserStatStore', () => {
  /** ================================
   *  状态管理
   * ================================ */
  const startMonth = ref('')
  const endMonth = ref('')

  // 缓存：key -> { ts, data }
  const cache = reactive({})
  // 正在请求的任务：key -> Promise
  const inFlight = reactive({})
  // 订阅者：{ apiKey: Set(callback) }
  const subscribers = reactive({})

  /** ================================
   *  工具函数
   * ================================ */

  /**
   * 简单防抖函数
   * @param {Function} fn - 需要防抖的函数
   * @param {number} delay - 延迟时间 (默认 200ms)
   * @returns {Function} 防抖后的函数
   */
  const debounce = (fn, delay = 200) => {
    let timer = null
    return () => {
      clearTimeout(timer)
      timer = setTimeout(fn, delay)
    }
  }

  /**
   * 通知所有订阅者有数据更新
   */
  const notifySubscribers = () => {
    for (const apiKey in subscribers) {
      for (const cb of subscribers[apiKey]) {
        try {
          cb()
        } catch (err) {
          console.error('[miserStatStore] subscriber error:', err)
        }
      }
    }
  }
  const debouncedNotify = debounce(notifySubscribers)

  /**
   * 生成缓存 key
   */
  const getCacheKey = (apiKey, params = {}) =>
    JSON.stringify({
      apiKey,
      startMonth: startMonth.value,
      endMonth: endMonth.value,
      ...params
    })

  /** ================================
   *  业务逻辑
   * ================================ */

  /**
   * 设置统计时间区间
   * @param {string} start - 起始月份
   * @param {string} end - 结束月份
   */
  const setRange = (start, end) => {
    if (startMonth.value === start && endMonth.value === end) {
      return
    }

    startMonth.value = start
    endMonth.value = end

    // 重置缓存和正在请求的任务
    Object.keys(cache).forEach((k) => delete cache[k])
    Object.keys(inFlight).forEach((k) => delete inFlight[k])

    // 通知订阅者
    debouncedNotify()
  }

  /**
   * 获取数据（带缓存 & 并发控制）
   * @param {string} apiKey - API 唯一标识
   * @param {Function} apiFn - 请求函数
   * @param {Object} params - 额外参数
   * @param {Object} options - 配置项 { force: 是否强制刷新 }
   */
  const fetchData = async (
    apiKey,
    apiFn,
    params = {},
    { force = false } = {}
  ) => {
    const key = getCacheKey(apiKey, params)

    // 如果有缓存且不强制刷新，直接返回
    if (!force && cache[key]) {
      return cache[key].data
    }
    // 如果已有请求在进行中，复用它
    if (inFlight[key]) {
      return inFlight[key]
    }

    // 发起请求
    const promise = (async () => {
      try {
        const res = await apiFn({
          startMonth: startMonth.value,
          endMonth: endMonth.value,
          ...params
        })
        cache[key] = { ts: Date.now(), data: res }
        return res
      } finally {
        delete inFlight[key] // 请求结束后清理
      }
    })()

    inFlight[key] = promise
    return promise
  }

  /**
   * 订阅数据更新
   * @param {string} apiKey - API 唯一标识
   * @param {Function} cb - 回调函数
   */
  const subscribe = (apiKey, cb) => {
    if (!subscribers[apiKey]) {
      subscribers[apiKey] = new Set()
    }
    subscribers[apiKey].add(cb)
  }

  /**
   * 取消订阅
   * @param {string} apiKey - API 唯一标识
   * @param {Function} cb - 回调函数
   */
  const unsubscribe = (apiKey, cb) => {
    subscribers[apiKey]?.delete(cb)
  }

  /**
   * 清除所有缓存和请求
   */
  const clearCache = () => {
    Object.keys(cache).forEach((k) => delete cache[k])
    Object.keys(inFlight).forEach((k) => delete inFlight[k])
    debouncedNotify()
  }

  /** ================================
   *  导出
   * ================================ */
  return {
    startMonth,
    endMonth,
    setRange,
    fetchData,
    subscribe,
    unsubscribe,
    clearCache
  }
})
