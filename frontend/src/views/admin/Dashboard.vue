<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="6" v-for="(value, key) in stats" :key="key">
        <el-card shadow="hover" class="stat-card">
          <div class="stat-value">{{ value }}</div>
          <div class="stat-label">{{ labels[key] || key }}</div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getStats } from '@/api/stats';

const stats = ref({});
const labels = {
  article_count: '总文章数',
  published_count: '已发布',
  category_count: '分类数',
  tag_count: '标签数',
  comment_count: '评论数',
  pending_comment_count: '待审核评论',
  total_views: '总浏览量'
};

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
.stat-card {
  text-align: center;
  margin-bottom: 20px;
}
.stat-value {
  font-size: 28px;
  font-weight: bold;
  color: #409eff;
}
.stat-label {
  color: #666;
  margin-top: 5px;
}
</style>
