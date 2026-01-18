<template>
  <el-row :gutter="20">
    <el-col :span="18" :xs="24">
      <div class="main-column">
        <!-- Hero / Welcome Section -->
        <div class="hero-section fade-in-up">
            <h1 class="hero-title">Welcome to My Blog</h1>
            <p class="hero-subtitle">Sharing thoughts on Tech, Life, and Code.</p>
        </div>

        <div class="article-list">
          <div class="list-container">
            <el-card 
              v-for="(article, index) in articles" 
              :key="article.id" 
              class="article-card scroll-reveal" 
              shadow="never" 
              @click="goDetail(article.id)"
              :ref="el => setArticleRef(el, index)"
            >
              <div class="article-content">
                <div class="article-header">
                     <el-tag v-if="article.category" size="small" effect="plain" class="category-tag">{{ article.category.name }}</el-tag>
                     <span class="date">{{ formatDate(article.created_at) }}</span>
                </div>
                <h2 class="article-title">{{ article.title }}</h2>
                <p class="article-summary">{{ article.summary }}</p>
                <div class="article-footer">
                    <span class="read-more">Read Article <el-icon><ArrowRight /></el-icon></span>
                    <span class="views"><el-icon><View /></el-icon> {{ article.views }} views</span>
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
            </el-card>
          </div>

          <el-empty v-if="articles.length === 0" description="No articles found" />

          <div class="pagination">
            <el-pagination
              background
              layout="prev, pager, next"
              :total="total"
              :page-size="pageSize"
              v-model:current-page="currentPage"
              @current-change="fetchArticles"
            />
          </div>
        </div>
      </div>
    </el-col>

    <el-col :span="6" :xs="0">
      <div class="sidebar">
          <!-- Profile Card -->
          <el-card class="sidebar-card profile-card" shadow="never">
              <div class="profile-header">
                  <div class="avatar-placeholder">Me</div>
                  <h3 class="profile-name">Zhuyin</h3>
                  <p class="profile-bio">Golang & Vue Developer. Minimalist.</p>
              </div>
          </el-card>

          <el-card class="sidebar-card" shadow="never">
            <template #header>
              <div class="card-header">
                <span>🗂 Categories</span>
              </div>
            </template>
            <div class="category-list">
              <div v-for="cat in categories" :key="cat.id" class="category-item" @click="goCategory(cat.id)">
                <span>{{ cat.name }}</span>
                <span class="count-badge">{{ cat.article_count }}</span>
              </div>
            </div>
          </el-card>

          <el-card class="sidebar-card mt-24" shadow="never">
            <template #header>
              <div class="card-header">
                <span>🏷 Tags</span>
              </div>
            </template>
            <div class="tag-cloud">
              <span
                v-for="tag in tags" 
                :key="tag.id" 
                class="tag-item" 
                @click="goTag(tag.id)"
              >
                #{{ tag.name }}
              </span>
            </div>
          </el-card>
      </div>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, onMounted, nextTick, onBeforeUnmount, watch } from 'vue';
import { useRouter } from 'vue-router';
import { getArticles } from '@/api/article';
import { getCategoriesStats } from '@/api/category';
import { getTagsStats } from '@/api/tag';
import dayjs from 'dayjs';
import { Calendar, View, ArrowRight, Picture } from '@element-plus/icons-vue';

const router = useRouter();
const articles = ref([]);
const total = ref(0);
const pageSize = ref(10);
const currentPage = ref(1);

const categories = ref([]);
const tags = ref([]);

// Animation Observer
const articleRefs = ref([]);
let observer = null;

const setArticleRef = (el, index) => {
    if (el) articleRefs.value[index] = el.$el; // el is component instance, $el is DOM
};

const initObserver = () => {
    if (observer) observer.disconnect();
    
    observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add('is-visible');
                observer.unobserve(entry.target); // Only animate once
            }
        });
    }, {
        threshold: 0.1, // Trigger when 10% visible
        rootMargin: '50px' // Trigger slightly before
    });

    articleRefs.value.forEach(el => {
        if (el) observer.observe(el);
    });
};

const fetchArticles = async () => {
  try {
    const res = await getArticles({
      page: currentPage.value,
      page_size: pageSize.value,
      status: 1
    });
    
    // Reset logic
    articleRefs.value = [];
    articles.value = res.list;
    total.value = res.total;
    
    // Wait for DOM update then observe
    nextTick(() => {
        initObserver();
    });

  } catch (error) {
    console.error(error);
  }
};

const fetchSidebar = async () => {
    try {
        const catRes = await getCategoriesStats();
        categories.value = catRes;
        
        const tagRes = await getTagsStats();
        tags.value = tagRes;
    } catch(e) {
        console.error(e);
    }
}

const formatDate = (date) => {
  return dayjs(date).format('MMM D, YYYY');
};

const goDetail = (id) => {
  router.push(`/article/${id}`);
};

const goCategory = (id) => {
    router.push(`/category/${id}`);
};
const goTag = (id) => {
    router.push(`/tag/${id}`);
};

onMounted(() => {
  fetchArticles();
  fetchSidebar();
});

onBeforeUnmount(() => {
    if (observer) observer.disconnect();
});
</script>

<style scoped>
.main-column {
    padding-top: 20px;
}

.hero-section {
    margin-bottom: 50px;
    padding: 60px 40px;
    text-align: center;
    background: linear-gradient(135deg, rgba(255,255,255,0.4), rgba(255,255,255,0.1));
    border-radius: 24px;
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.07);
    border: 1px solid rgba(255, 255, 255, 0.18);
}

:global(.dark) .hero-section {
    background: linear-gradient(135deg, rgba(30,30,30,0.4), rgba(30,30,30,0.1));
    border-color: rgba(255, 255, 255, 0.05);
}

.hero-title {
    font-size: 3.5rem;
    font-weight: 800;
    letter-spacing: -2px;
    margin: 0 0 15px 0;
    background: linear-gradient(120deg, var(--primary-color), var(--secondary-color));
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    text-shadow: 0 10px 30px rgba(0,0,0,0.1);
}

.hero-subtitle {
    font-size: 1.3rem;
    color: var(--text-secondary);
    margin: 0;
    font-weight: 300;
}

/* Scroll Reveal Animation Classes */
.scroll-reveal {
    opacity: 0;
    transform: translateY(60px) scale(0.9);
    transition: all 0.8s cubic-bezier(0.2, 0.8, 0.2, 1);
    top: 5em;
}

.scroll-reveal.is-visible {
    opacity: 1;
    transform: translateY(0) scale(1);
}

.article-card {
  margin-bottom: 30px;
  cursor: pointer;
  border: none;
  background: var(--card-bg); 
  backdrop-filter: blur(10px);
  border-radius: 20px;
  overflow: hidden;
  position: relative;
  /* Remove old animation props */
}

/* Border Gradient Effect on Hover */
.article-card::before {
    content: "";
    position: absolute;
    inset: 0;
    border-radius: 20px;
    padding: 2px; /* Border thickness */
    background: linear-gradient(45deg, transparent, transparent);
    -webkit-mask: linear-gradient(#fff 0 0) content-box, linear-gradient(#fff 0 0);
    -webkit-mask-composite: xor;
    mask-composite: exclude;
    transition: background 0.4s ease;
    pointer-events: none;
}

.article-card:hover::before {
    background: linear-gradient(45deg, var(--primary-color), var(--secondary-color));
}

.article-card:hover {
  transform: translateY(-8px) scale(1.02) !important; /* Slightly lift and scale */
  box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.15);
}

:deep(.el-card__body) {
    display: flex;
    padding: 35px;
    gap: 35px;
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
}

.date {
    font-size: 0.85rem;
    color: var(--text-secondary);
    letter-spacing: 0.5px;
    text-transform: uppercase;
}

.article-title {
  margin: 0 0 15px 0;
  font-size: 2rem;
  font-weight: 800;
  color: var(--text-main);
  line-height: 1.25;
  transition: color 0.3s;
  letter-spacing: -0.5px;
}

.article-card:hover .article-title {
    color: var(--primary-color);
}

.article-summary {
  color: var(--text-secondary);
  line-height: 1.8;
  font-size: 1.1rem;
  margin-bottom: 30px;
  display: -webkit-box;
  -webkit-line-clamp: 3; /* Show a bit more text */
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
    font-size: 1rem;
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
    width: 320px;
    height: 240px;
    border-radius: 16px;
    overflow: hidden;
    flex-shrink: 0;
    box-shadow: var(--shadow-md);
    position: relative;
}

.cover-image {
    width: 100%;
    height: 100%;
    transition: transform 0.8s cubic-bezier(0.2, 0.8, 0.2, 1);
}

.article-card:hover .cover-image {
    transform: scale(1.1);
}

/* Sidebar */
.sidebar {
    position: sticky;
    top: 90px;
}

.sidebar-card {
  margin-bottom: 24px;
  border-radius: 20px;
  border: none;
  background: var(--card-bg);
  backdrop-filter: blur(10px);
}

.profile-card {
    text-align: center;
    padding: 30px 0;
}

.avatar-placeholder {
    width: 100px;
    height: 100px;
    background: linear-gradient(135deg, #a1c4fd 0%, #c2e9fb 100%);
    border-radius: 50%;
    margin: 0 auto 20px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #fff;
    font-weight: 800;
    font-size: 1.5rem;
    box-shadow: 0 10px 20px rgba(161, 196, 253, 0.4);
}

.profile-name {
    margin: 0 0 5px;
    font-size: 1.4rem;
    color: var(--text-main);
    font-weight: 700;
}

.profile-bio {
    color: var(--text-secondary);
    font-size: 1rem;
    margin: 0;
}

.card-header {
    font-size: 1.1rem;
    font-weight: 700;
    color: var(--text-main);
    letter-spacing: -0.5px;
}

.category-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 10px;
  cursor: pointer;
  color: var(--text-secondary);
  transition: all 0.2s;
  border-radius: 8px;
}

.category-item:hover {
    color: var(--primary-color);
    background: rgba(var(--primary-color), 0.1); 
    /* Would need rgb var, but simple hover works */
    background: rgba(0,0,0,0.03);
}
:global(.dark) .category-item:hover {
    background: rgba(255,255,255,0.05);
}

.count-badge {
    background: var(--bg-color);
    padding: 4px 10px;
    border-radius: 12px;
    font-size: 0.75rem;
    font-weight: 600;
}

.tag-cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}

.tag-item {
    cursor: pointer;
    color: var(--text-secondary);
    background: var(--bg-color);
    padding: 6px 14px;
    border-radius: 20px;
    font-size: 0.9rem;
    transition: all 0.2s;
    font-weight: 500;
}
.tag-item:hover {
    color: #fff;
    background: var(--primary-color);
    transform: translateY(-2px);
    box-shadow: 0 5px 10px rgba(64, 158, 255, 0.3);
}

.mt-24 {
    margin-top: 24px;
}

@media (max-width: 900px) {
    .article-cover {
        width: 120px;
        height: 120px;
    }
}

@media (max-width: 768px) {
    .article-cover {
        display: none;
    }
    .hero-title {
        font-size: 2.5rem;
    }
}
</style>
