<template>
  <GlassCard 
    class="article-card" 
    :class="{ 'is-reversed': reversed }"
    @click="$emit('click')"
  >
    <div class="article-content">
      <div class="article-header">
           <el-tag v-if="article.category" size="small" effect="plain" class="category-tag">{{ article.category.name }}</el-tag>
           
           <div class="tags-group" v-if="article.tags && article.tags.length">
             <el-tag 
               v-for="tag in article.tags" 
               :key="tag.id" 
               size="small" 
               effect="plain" 
               type="info" 
               class="tag-pill"
             >
               #{{ tag.name }}
             </el-tag>
           </div>

           <span class="date">{{ formatDate(article.created_at) }}</span>
      </div>
      <h2 class="article-title">{{ article.title }}</h2>
      <p class="article-summary">{{ article.summary }}</p>
      
      <!-- Footer with conditional ordering -->
      <div class="article-footer">
          <!-- Normal Order: Image on Right -> Views Left, ReadMore Right -->
          <template v-if="!reversed">
              <span class="views"><el-icon><View /></el-icon> {{ article.views }} 阅读</span>
              <span class="read-more">阅读全文 <el-icon><ArrowRight /></el-icon></span>
          </template>
          
          <!-- Reversed Order: Image on Left -> ReadMore Left (near image), Views Right (far) -->
          <template v-else>
              <span class="read-more"><el-icon style="transform: rotate(180deg); margin-right: 4px;"><ArrowRight /></el-icon> 阅读全文</span>
              <span class="views">{{ article.views }} 阅读 <el-icon><View /></el-icon></span>
          </template>
      </div>
    </div>
    
    <div class="article-cover" v-if="article.cover_img">
      <el-image :src="article.cover_img" fit="cover" loading="lazy" class="cover-image">
          <template #error>
              <div class="image-slot">
                  <el-icon><Picture /></el-icon>
              </div>
          </template>
      </el-image>
    </div>
  </GlassCard>
</template>

<script setup>
import { ArrowRight, View, Picture } from '@element-plus/icons-vue';
import dayjs from 'dayjs';
import GlassCard from './GlassCard.vue';

defineProps({
  article: {
    type: Object,
    required: true
  },
  reversed: {
    type: Boolean,
    default: false
  }
});

defineEmits(['click']);

const formatDate = (date) => {
  return dayjs(date).format('MMM D, YYYY');
};
</script>

<style scoped>
.article-card {
  margin-bottom: 30px;
  cursor: pointer;
  overflow: hidden;
  position: relative;
  display: flex;
  padding: 35px;
  gap: 35px;
  transition: transform 0.3s cubic-bezier(0.2, 0.8, 0.2, 1), box-shadow 0.3s ease;
}

/* Reversed Layout Logic */
.article-card.is-reversed {
    flex-direction: row-reverse;
}

.article-card.is-reversed .article-content {
    text-align: right;
    align-items: flex-end;
}

.article-card.is-reversed .article-header {
    flex-direction: row-reverse;
}

/* 
   When reversed, the article content is aligned to the right (flex-end).
   The footer container (.article-footer) inherits this alignment context but uses space-between.
   
   However, in reversed mode (text-align: right), we want the footer elements to swap sides relative to the content block.
   
   Structure in DOM (Reversed Template):
   [Read More] ... [Views]
   
   Desired Visuals:
   [Read More (Left)] ........................ [Views (Right)]
   
   BUT: The parent .article-content is aligning items to flex-end (Right).
   We need to ensure .article-footer spans the full width and distributes items correctly.
*/
.article-card.is-reversed .article-footer {
    width: 100%;
    /* Ensure the footer justifies content across the full available width */
    justify-content: space-between;
    flex-direction: row; /* Keep row direction, don't reverse flex here, we reversed DOM order */
}

/* Ensure children don't inherit row-reverse which might mess up their internal icon spacing */
/* .article-card.is-reversed .article-footer > * {
    flex-direction: row;
} */
/* REMOVED: CSS-based reordering proved unreliable. Now handling via v-if in template. */

/* ... rest of styles ... */
.article-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.1);
}

.article-content {
    flex: 1;
    display: flex;
    flex-direction: column;
}

.article-header {
    display: flex;
    align-items: center;
    gap: 15px;
    margin-bottom: 12px;
}

.category-tag {
    border-radius: 20px;
    padding: 0 15px;
    font-weight: 600;
    border: 1px solid rgba(var(--primary-color-rgb), 0.3);
}

.date {
    font-size: 0.85rem;
    color: var(--text-secondary);
    letter-spacing: 0.5px;
    text-transform: uppercase;
    font-weight: 500;
}

.article-title {
  margin: 0 0 15px 0;
  font-size: 1.8rem;
  font-weight: 800;
  color: var(--text-main);
  line-height: 1.3;
  transition: color 0.3s;
  letter-spacing: -0.5px;
}

.article-card:hover .article-title {
    color: var(--primary-color);
}

.article-summary {
  color: var(--text-secondary);
  line-height: 1.7;
  font-size: 1.05rem;
  margin-bottom: 25px;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.article-footer {
    margin-top: auto;
    display: flex;
    align-items: center;
    justify-content: space-between;
}

.read-more {
    font-weight: 700;
    color: var(--primary-color);
    display: flex;
    align-items: center;
    gap: 8px;
    font-size: 0.95rem;
    transition: gap 0.3s;
}

.article-card:hover .read-more {
    gap: 12px;
}

.views {
    display: flex;
    align-items: center;
    gap: 6px;
    color: var(--text-secondary);
    font-size: 0.9rem;
}

.article-cover {
    width: 280px;
    height: 210px;
    border-radius: 16px;
    overflow: hidden;
    flex-shrink: 0;
    box-shadow: 0 8px 24px rgba(0,0,0,0.1);
}

.cover-image {
    width: 100%;
    height: 100%;
    transition: transform 0.6s cubic-bezier(0.2, 0.8, 0.2, 1);
}

.article-card:hover .cover-image {
    transform: scale(1.05);
}

@media (max-width: 900px) {
    .article-cover {
        width: 120px;
        height: 120px;
    }
}

@media (max-width: 768px) {
    .article-card {
        padding: 25px;
    }
    .article-cover {
        display: none;
    }
    .article-title {
        font-size: 1.5rem;
    }
}
</style>
