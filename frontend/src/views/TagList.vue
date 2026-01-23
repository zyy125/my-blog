<template>
  <div class="tag-list-page fade-in-up">
    <div class="page-header">
      <h2 class="title">文章标签</h2>
      <p class="subtitle">共 {{ tags.length }} 个标签</p>
    </div>

    <div class="tag-cloud-wrapper">
      <div class="tag-cloud">
        <span
          v-for="(tag, index) in tags" 
          :key="tag.id" 
          class="tag-item" 
          :class="getRandomClass(index)"
          :style="{ animationDelay: `${index * 0.05}s` }"
          @click="goTag(tag.id)"
        >
          <span class="hash">#</span>
          {{ tag.name }}
          <span class="tag-count" v-if="tag.article_count > 0">{{ tag.article_count }}</span>
        </span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getTagsStats } from '@/api/tag';

const router = useRouter();
const tags = ref([]);

const fetchData = async () => {
  try {
    const res = await getTagsStats();
    tags.value = res;
  } catch (e) {
    console.error(e);
  }
};

const goTag = (id) => {
  router.push(`/tag/${id}`);
};

const getRandomClass = (index) => {
  const sizes = ['size-s', 'size-m', 'size-l', 'size-xl'];
  // Deterministic "random" based on index/char code to keep layout stable
  return sizes[index % 4];
};

onMounted(() => {
  fetchData();
});
</script>

<style scoped>
.tag-list-page {
  max-width: 1000px;
  margin: 40px auto;
  padding: 0 20px;
  min-height: 80vh;
}

.page-header {
  text-align: center;
  margin-bottom: 60px;
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

.tag-cloud-wrapper {
  perspective: 1000px;
}

.tag-cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 15px;
  justify-content: center;
  align-items: center;
  padding: 40px;
}

.tag-item {
  position: relative;
  cursor: pointer;
  color: var(--text-secondary);
  background: rgba(255, 255, 255, 0.05);
  border: 1px solid rgba(255, 255, 255, 0.05);
  border-radius: 12px;
  transition: all 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
  display: inline-flex;
  align-items: center;
  gap: 8px;
  /* opacity: 0; Removed to ensure visibility safety */
  animation: zoomIn 0.5s ease both; /* 'both' handles initial hidden state during delay */
  backdrop-filter: blur(5px);
}

.tag-item:hover {
  color: #fff;
  background: var(--primary-color);
  transform: translateY(-5px) scale(1.05) !important;
  box-shadow: 0 10px 20px rgba(64, 158, 255, 0.3);
  border-color: transparent;
  z-index: 10;
}

.hash {
  opacity: 0.5;
  font-weight: 300;
}

.tag-count {
  font-size: 0.75em;
  background: rgba(0, 0, 0, 0.2);
  padding: 2px 6px;
  border-radius: 6px;
  min-width: 18px;
  text-align: center;
}

/* Size Variants */
.size-s {
  font-size: 0.9rem;
  padding: 6px 14px;
}

.size-m {
  font-size: 1.1rem;
  padding: 8px 18px;
}

.size-l {
  font-size: 1.3rem;
  padding: 10px 22px;
  font-weight: 500;
}

.size-xl {
  font-size: 1.6rem;
  padding: 14px 28px;
  font-weight: 700;
  background: rgba(255, 255, 255, 0.08);
  border-color: rgba(255, 255, 255, 0.1);
}

@keyframes zoomIn {
  from {
    opacity: 0;
    transform: scale(0.8);
  }
  to {
    opacity: 1;
    transform: scale(1);
  }
}
</style>