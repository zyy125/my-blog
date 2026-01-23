<template>
  <div class="page-container">
    <div class="welcome-section">
      <h2 class="welcome-title">
        欢迎回来，管理员
        <span class="welcome-emoji">👋</span>
      </h2>
      <p class="welcome-subtitle">这里是您的博客数据概览</p>
    </div>

    <div class="stats-grid">
      <div 
        v-for="item in statItems" 
        :key="item.key" 
        class="stat-card"
        :style="{ '--card-accent': item.color }"
      >
        <div class="stat-icon-wrapper">
          <el-icon><component :is="item.icon" /></el-icon>
        </div>
        <div class="stat-info">
          <div class="stat-value">{{ stats[item.key] || 0 }}</div>
          <div class="stat-label">{{ item.label }}</div>
        </div>
        <div class="stat-bg-icon">
          <el-icon><component :is="item.icon" /></el-icon>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getStats } from '@/api/stats';
import { Document, Folder, Collection, ChatDotSquare, View, Timer } from '@element-plus/icons-vue';

const stats = ref({});

const statItems = [
  { key: 'article_count', label: '总文章数', icon: Document, color: '#409eff' },
  { key: 'published_count', label: '已发布', icon: Document, color: '#67c23a' },
  { key: 'category_count', label: '分类数', icon: Folder, color: '#e6a23c' },
  { key: 'tag_count', label: '标签数', icon: Collection, color: '#f56c6c' },
  { key: 'comment_count', label: '评论总数', icon: ChatDotSquare, color: '#909399' },
  { key: 'total_views', label: '总浏览量', icon: View, color: '#b37feb' }
];

onMounted(async () => {
  try {
    const res = await getStats();
    stats.value = res;
  } catch (e) {
    console.error(e);
  }
});
</script>

<style scoped>
.page-container {
    padding: 0;
}

.welcome-section {
    margin-bottom: 32px;
}

.welcome-title {
    font-size: 24px;
    font-weight: 700;
    margin: 0 0 8px 0;
    color: var(--text-primary);
    display: flex;
    align-items: center;
    gap: 8px;
}

.welcome-emoji {
    animation: wave 2.5s infinite;
    transform-origin: 70% 70%;
    display: inline-block;
}

.welcome-subtitle {
    margin: 0;
    color: var(--text-secondary);
    font-size: 14px;
}

.stats-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
    gap: 24px;
}

.stat-card {
    background: var(--bg-container);
    border-radius: 12px;
    padding: 24px;
    display: flex;
    align-items: center;
    position: relative;
    overflow: hidden;
    box-shadow: var(--shadow-base);
    transition: var(--transition-all);
    border: 1px solid var(--border-lighter);
}

.stat-card:hover {
    transform: translateY(-4px);
    box-shadow: var(--shadow-hover);
    border-color: var(--card-accent);
}

.stat-icon-wrapper {
    width: 48px;
    height: 48px;
    border-radius: 12px;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 24px;
    margin-right: 16px;
    background: color-mix(in srgb, var(--card-accent), transparent 90%);
    color: var(--card-accent);
}

.stat-info {
    position: relative;
    z-index: 2;
}

.stat-value {
    font-size: 28px;
    font-weight: 700;
    line-height: 1.2;
    color: var(--text-primary);
    margin-bottom: 4px;
}

.stat-label {
    font-size: 13px;
    color: var(--text-secondary);
}

.stat-bg-icon {
    position: absolute;
    right: -10px;
    bottom: -10px;
    font-size: 100px;
    opacity: 0.05;
    color: var(--text-primary);
    transform: rotate(-15deg);
    pointer-events: none;
    z-index: 1;
}

@keyframes wave {
    0% { transform: rotate(0deg); }
    10% { transform: rotate(14deg); }
    20% { transform: rotate(-8deg); }
    30% { transform: rotate(14deg); }
    40% { transform: rotate(-4deg); }
    50% { transform: rotate(10deg); }
    60% { transform: rotate(0deg); }
    100% { transform: rotate(0deg); }
}
</style>