<template>
  <div class="category-list-page fade-in-up">
    <div class="page-header">
      <h2 class="title">文章分类</h2>
      <p class="subtitle">共 {{ categories.length }} 个分类</p>
    </div>

    <div class="categories-grid">
      <div 
        v-for="(cat, index) in categories" 
        :key="cat.id" 
        class="category-card-wrapper"
        :style="{ animationDelay: `${index * 0.1}s` }"
        @click="goCategory(cat.id)"
      >
        <div class="category-card">
          <div class="card-bg"></div>
          <div class="card-content">
            <!-- Icon Removed -->
            <h3 class="cat-name">{{ cat.name }}</h3>
            <div class="stat-row">
              <span class="article-count">{{ cat.article_count }}</span>
              <span class="label">篇文章</span>
            </div>
            <div class="action-arrow">
              <el-icon><ArrowRight /></el-icon>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getCategoriesStats } from '@/api/category';
import { ArrowRight } from '@element-plus/icons-vue';

const router = useRouter();
const categories = ref([]);

// Simple mapping for icons based on keywords, fallback to folder
const getIcon = (name) => {
  // Removed emoji icons as requested for better scalability and cleaner UI
  return ''; 
};

const fetchData = async () => {
  try {
    const res = await getCategoriesStats();
    categories.value = res;
  } catch (e) {
    console.error(e);
  }
};

const goCategory = (id) => {
  router.push(`/category/${id}`);
};

onMounted(() => {
  fetchData();
});
</script>

<style scoped>
.category-list-page {
  max-width: 1200px;
  margin: 40px auto;
  padding: 0 20px;
  min-height: 80vh;
}

.page-header {
  text-align: center;
  margin-bottom: 80px;
}

.title {
  font-size: 3rem;
  margin: 0 0 15px;
  font-weight: 800;
  background: linear-gradient(120deg, var(--primary-color), #a0cfff);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  letter-spacing: -1px;
}

.subtitle {
  color: var(--text-secondary);
  font-size: 1.2rem;
  opacity: 0.8;
}

.categories-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 30px;
}

.category-card-wrapper {
  /* opacity: 0; Removed to ensure visibility */
  animation: fadeInUp 0.6s ease both;
}

@keyframes fadeInUp {
  from { opacity: 0; transform: translateY(20px); }
  to { opacity: 1; transform: translateY(0); }
}

.category-card {
  position: relative;
  height: 160px;
  background: rgba(255, 255, 255, 0.03);
  border: 1px solid rgba(255, 255, 255, 0.08);
  border-radius: 16px;
  overflow: hidden;
  cursor: pointer;
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  backdrop-filter: blur(10px);
}

.category-card:hover {
  transform: translateY(-5px);
  border-color: var(--primary-color);
  box-shadow: 0 10px 25px -5px rgba(0, 0, 0, 0.3);
}

.card-bg {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 100%;
  background: linear-gradient(135deg, rgba(255,255,255,0.05) 0%, rgba(255,255,255,0) 100%);
  opacity: 0;
  transition: opacity 0.4s ease;
}

.category-card:hover .card-bg {
  opacity: 1;
}

.card-content {
  position: relative;
  z-index: 2;
  height: 100%;
  padding: 24px;
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  justify-content: center;
}

.cat-name {
  margin: 0 0 8px;
  font-size: 1.4rem;
  color: var(--text-main);
  font-weight: 600;
  letter-spacing: 0.5px;
}

.stat-row {
  display: flex;
  align-items: baseline;
  gap: 6px;
  color: var(--text-secondary);
  font-size: 0.95rem;
}

.article-count {
  font-size: 1.1rem;
  font-weight: 700;
  color: var(--primary-color);
}

.action-arrow {
  position: absolute;
  bottom: 24px;
  right: 24px;
  opacity: 0.5;
  transform: translateX(0);
  transition: all 0.3s ease;
  color: var(--text-secondary);
  font-size: 1.2rem;
}

.category-card:hover .action-arrow {
  opacity: 1;
  color: var(--primary-color);
  transform: translateX(5px);
}
</style>