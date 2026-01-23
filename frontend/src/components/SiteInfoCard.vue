<template>
  <GlassCard class="site-info-card mt-24" hover-effect>
    <div class="card-header">
      <el-icon><DataLine /></el-icon>
      <span>站点资讯</span>
    </div>
    <div class="site-info-list">
      <div class="info-item" v-for="(item, index) in infoItems" :key="index">
        <span class="label">{{ item.label }}</span>
        <span class="value">{{ item.value }}</span>
      </div>
    </div>
  </GlassCard>
</template>

<script setup>
import { computed } from 'vue';
import { DataLine } from '@element-plus/icons-vue';
import GlassCard from './GlassCard.vue';
import dayjs from 'dayjs';
import 'dayjs/locale/zh-cn';
import relativeTime from 'dayjs/plugin/relativeTime';

dayjs.locale('zh-cn');
dayjs.extend(relativeTime);

const props = defineProps({
  info: {
    type: Object,
    required: true,
    default: () => ({
        articleCount: 0,
        runningTime: '-',
        wordCount: '0',
        visitorCount: '-',
        totalViews: '0',
        lastUpdate: '-'
    })
  }
});

const infoItems = computed(() => [
  { label: '文章数目', value: props.info.articleCount },
  { label: '已运行时间', value: props.info.runningTime },
  { label: '本站文章总访问量', value: props.info.totalViews },
  { label: '最后更新时间', value: props.info.lastUpdate }
]);
</script>

<style scoped>
.mt-24 {
  margin-top: 24px;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 10px;
  font-weight: 700;
  font-size: 1.1rem;
  color: var(--text-main);
  margin-bottom: 20px;
  padding: 15px 20px 5px;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.site-info-list {
  padding: 0 20px 20px;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 10px; /* Increased padding for better interaction area */
  cursor: pointer;
  color: var(--text-secondary);
  transition: all 0.2s;
  border-radius: 8px;
  font-size: 0.9rem;
}

/* Add hover interaction similar to Category/Tags */
.info-item:hover {
    color: var(--primary-color);
    background: rgba(0,0,0,0.03);
}
:global(.dark) .info-item:hover {
    background: rgba(255,255,255,0.05);
}

.info-item .value {
  font-weight: 600;
  color: var(--text-main);
  /* Font family managed in style.css for consistency */
}
</style>
