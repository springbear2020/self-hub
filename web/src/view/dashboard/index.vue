<template>
  <div class="modern-dashboard">
    <el-row :gutter="24" class="stat-grid">
      <el-col
        v-for="(item, index) in statData"
        :key="index"
        :xs="24"
        :sm="12"
        :md="12"
        :lg="8"
        :xl="6"
        class="stat-col"
      >
        <div class="premium-card" :class="`theme-${index % 3}`">
          <div class="card-content">
            <div class="top-area">
              <div class="label-wrapper">
                <span class="dot"></span>
                <span class="label">{{ item.label }}</span>
              </div>
              <div class="icon-wrapper">
                <el-icon>
                  <Component :is="item.icon || 'DataLine'" />
                </el-icon>
              </div>
            </div>

            <div class="data-wrapper">
              <span class="number">{{ item.value }}</span>
              <span class="unit" v-if="item.unit">{{ item.unit }}</span>
            </div>
          </div>

          <div class="decoration-shape"></div>
          <div class="decoration-ring"></div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
  import { ref, onMounted } from 'vue'
  import { getDashboardStats } from '@/api/dashboard'

  const statData = ref([])
  const getStatData = async () => {
    const res = await getDashboardStats()
    statData.value = (res.data || []).sort((a, b) => a.sort - b.sort)
  }

  onMounted(getStatData)
</script>

<style lang="scss" scoped>
  .modern-dashboard {
    padding: 24px;
    background-color: #f3f6f9; /* 更清爽的背景色 */
    min-height: calc(100vh - 100px);
    font-family:
      -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue',
      Arial, sans-serif;

    .header-section {
      margin-bottom: 32px;
      padding-left: 8px;
      border-left: 4px solid #409eff; /* 增加左侧修饰线 */

      .title {
        margin: 0 0 6px 0;
        font-size: 22px;
        font-weight: 700;
        color: #1e293b;
        letter-spacing: 0.5px;
      }

      .subtitle {
        margin: 0;
        font-size: 13px;
        color: #64748b;
      }
    }

    .stat-grid {
      .stat-col {
        margin-bottom: 24px;
      }
    }

    .premium-card {
      position: relative;
      background: #ffffff;
      border-radius: 20px; /* 进一步圆润化 */
      padding: 24px;
      overflow: hidden;
      border: 1px solid rgba(255, 255, 255, 0.8);
      box-shadow:
        0 4px 12px -2px rgba(0, 0, 0, 0.03),
        0 2px 6px -1px rgba(0, 0, 0, 0.02);
      transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1); /* 加入果冻回弹 easing */
      cursor: pointer;

      /* 核心内容区 */
      .card-content {
        position: relative;
        z-index: 2;

        .top-area {
          display: flex;
          justify-content: space-between;
          align-items: center;
          margin-bottom: 20px;
        }

        .label-wrapper {
          display: flex;
          align-items: center;

          .dot {
            width: 8px;
            height: 8px;
            border-radius: 50%;
            margin-right: 10px;
            box-shadow: 0 0 8px currentColor; /* 让小圆点发光 */
          }

          .label {
            font-size: 15px;
            font-weight: 600;
            color: #475569;
          }
        }

        .icon-wrapper {
          display: flex;
          align-items: center;
          justify-content: center;
          width: 40px;
          height: 40px;
          border-radius: 12px;
          font-size: 20px;
          opacity: 0.85;
          transition: all 0.3s ease;
        }

        .data-wrapper {
          display: flex;
          align-items: baseline;
          margin-top: auto;

          .number {
            /* 推荐系统自带的无衬线等宽数字字体，防止数据跳动 */
            font-family:
              'DIN Alternate', 'Oswald', 'Helvetica Neue', Arial, sans-serif;
            font-size: 34px;
            font-weight: 700;
            color: #0f172a;
            line-height: 1.1;
            white-space: nowrap;
            letter-spacing: -0.5px;
          }

          .unit {
            margin-left: 8px;
            font-size: 14px;
            font-weight: 600;
            color: #94a3b8;
          }
        }
      }

      /* 底部光晕 */
      .decoration-shape {
        position: absolute;
        right: -30px;
        bottom: -30px;
        width: 120px;
        height: 120px;
        border-radius: 50%;
        filter: blur(30px);
        opacity: 0.15;
        transition: all 0.5s ease;
        z-index: 1;
      }

      /* 新增：右上角线框圆环辅助装饰 */
      .decoration-ring {
        position: absolute;
        right: -20px;
        top: -20px;
        width: 80px;
        height: 80px;
        border-radius: 50%;
        border: 2px solid transparent;
        opacity: 0.1;
        transition: all 0.5s ease;
        z-index: 1;
      }

      /* 通用悬浮态 */
      &:hover {
        transform: translateY(-8px);

        .icon-wrapper {
          transform: scale(1.1) rotate(5deg);
        }

        .decoration-shape {
          transform: scale(1.4);
          opacity: 0.25;
        }

        .decoration-ring {
          transform: scale(1.2);
          opacity: 0.2;
        }
      }
    }

    .theme-0 {
      &:hover {
        box-shadow: 0 16px 24px -6px rgba(103, 194, 58, 0.25);
      }

      .dot {
        background: linear-gradient(135deg, #67c23a 0%, #85ce61 100%);
        color: #67c23a;
      }

      .icon-wrapper {
        background: rgba(103, 194, 58, 0.1);
        color: #85ce61;
      }

      .decoration-shape {
        background: #67c23a;
      }

      .decoration-ring {
        border-color: #67c23a;
      }
    }

    .theme-1 {
      &:hover {
        box-shadow: 0 16px 24px -6px rgba(245, 108, 108, 0.25);
      }

      .dot {
        background: linear-gradient(135deg, #f56c6c 0%, #f78989 100%);
        color: #f56c6c;
      }

      .icon-wrapper {
        background: rgba(245, 108, 108, 0.1);
        color: #f78989;
      }

      .decoration-shape {
        background: #f56c6c;
      }

      .decoration-ring {
        border-color: #f56c6c;
      }
    }

    .theme-2 {
      &:hover {
        box-shadow: 0 16px 24px -6px rgba(0, 204, 198, 0.25);
      }

      .dot {
        background: linear-gradient(135deg, #00ccc6 0%, #33d9d5 100%);
        color: #00ccc6;
      }

      .icon-wrapper {
        background: rgba(0, 204, 198, 0.1);
        color: #33d9d5;
      }

      .decoration-shape {
        background: #00ccc6;
      }

      .decoration-ring {
        border-color: #00ccc6;
      }
    }
  }
</style>
