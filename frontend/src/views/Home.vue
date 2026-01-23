<template>
  <el-row :gutter="20">
    <el-col :span="18" :xs="24">
      <div class="main-column">
        <!-- Hero / Welcome Section -->
        <div class="hero-section fade-in-up">
            <h1 class="hero-title">欢迎来到我的博客</h1>
        </div>

        <div class="article-list">
          <div class="list-container">
            <ArticleCard 
              v-for="(article, index) in articles" 
              :key="article.id" 
              :article="article"
              :reversed="index % 2 !== 0"
              class="anim-scale-in" 
              :style="{ animationDelay: `${index * 0.2}s` }"
              @click="goDetail(article.id)"
            />
          </div>

          <el-empty v-if="articles.length === 0" description="暂无文章" />

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

    <el-col :span="6" :xs="6">
      <div class="sidebar">
        <!-- Profile Card -->
        <ProfileCard 
            class="sidebar-card anim-scale-in" 
            style="animation-delay: 0.2s"
            :articles="total" 
            :categories="categories.length" 
            :tags="tags.length"
        />

        <!-- Latest Articles Card -->
        <GlassCard class="sidebar-card anim-scale-in" style="animation-delay: 0.4s" hover-effect v-if="latestArticles.length > 0">
          <div class="card-header">
            <span><el-icon><Timer /></el-icon> 最新文章</span>
          </div>
          <div class="latest-list">
            <div v-for="article in latestArticles" :key="article.id" class="latest-item" @click="goDetail(article.id)">
              <div class="latest-title">{{ article.title }}</div>
              <div class="latest-date">{{ formatDate(article.created_at) }}</div>
            </div>
          </div>
        </GlassCard>

        <!-- Site Info Card -->
        <SiteInfoCard :info="siteInfo" class="sidebar-card anim-scale-in" style="animation-delay: 0.3s" />

        <GlassCard class="sidebar-card anim-scale-in" style="animation-delay: 0.8s" hover-effect>
          <div class="card-header">
            <span>🗂 分类</span>
          </div>
          <div class="category-list">
            <div v-for="cat in categories" :key="cat.id" class="category-item" @click="goCategory(cat.id)">
              <span>{{ cat.name }}</span>
              <span class="count-badge">{{ cat.article_count }}</span>
            </div>
          </div>
        </GlassCard>

        <GlassCard class="sidebar-card anim-scale-in" style="animation-delay: 1s" hover-effect>
          <div class="card-header">
            <span>🏷 标签</span>
          </div>
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
        </GlassCard>

        <!-- Archives Card -->
        <GlassCard class="sidebar-card anim-scale-in" style="animation-delay: 1.2s" hover-effect>
          <div class="card-header">
            <span><el-icon><Collection /></el-icon> 归档</span>
          </div>
          <div class="category-list">
            <div v-for="arc in archives" :key="arc.date" class="category-item" @click="goArchive(arc.rawDate)">
              <span>{{ arc.date }}</span>
              <span class="count-badge">{{ arc.count }}</span>
            </div>
          </div>
        </GlassCard>
      </div>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, onMounted, nextTick, onBeforeUnmount } from 'vue';
import { useRouter } from 'vue-router';
import { getArticles } from '@/api/article';
import { getCategoriesStats } from '@/api/category';
import { getTagsStats } from '@/api/tag';
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';
import { Collection, Timer } from '@element-plus/icons-vue';
import ArticleCard from '@/components/ArticleCard.vue';
import ProfileCard from '@/components/ProfileCard.vue';
import SiteInfoCard from '@/components/SiteInfoCard.vue';
import GlassCard from '@/components/GlassCard.vue';

dayjs.extend(relativeTime);

const router = useRouter();
const articles = ref([]);
const latestArticles = ref([]);
const total = ref(0);
const pageSize = ref(8);
const currentPage = ref(1);

const categories = ref([]);
const tags = ref([]);
const archives = ref([]);
const siteInfo = ref({
    articleCount: 0,
    runningTime: '',
    wordCount: '0',
    visitorCount: '12,345',
    totalViews: '0',
    lastUpdate: '-'
});

// Animation Observer
const articleRefs = ref([]);
let observer = null;

const setArticleRef = (el, index) => {
    // ArticleCard component root element
    if (el) articleRefs.value[index] = el.$el; 
};

const initObserver = () => {
    if (observer) observer.disconnect();
    
    observer = new IntersectionObserver((entries) => {
        entries.forEach(entry => {
            if (entry.isIntersecting) {
                entry.target.classList.add('is-visible');
                observer.unobserve(entry.target); 
            }
        });
    }, {
        threshold: 0.1, 
        rootMargin: '50px' 
    });

    articleRefs.value.forEach(el => {
        if (el) observer.observe(el);
    });
};

const fetchArticles = async (page) => {
  if (typeof page === 'number') {
    window.scrollTo({ top: 0, behavior: 'smooth' });
  }
  try {
    const res = await getArticles({
      page: currentPage.value,
      page_size: pageSize.value,
      status: 1
    });
    
    articleRefs.value = [];
    articles.value = res.list;
    total.value = res.total;
    
    nextTick(() => {
        initObserver();
    });

  } catch (error) {
    console.error(error);
  }
};

const fetchSidebar = async () => {
    try {
        // Fetch Latest Articles (Limit 5)
        const latestRes = await getArticles({ page: 1, page_size: 5, status: 1 });
        latestArticles.value = latestRes.list;

        const catRes = await getCategoriesStats();
        categories.value = catRes;
        
        const tagRes = await getTagsStats();
        tags.value = tagRes;

        const allRes = await getArticles({ page: 1, page_size: 1000, status: 1 });
        const list = allRes.list;

        const groups = {};
        list.forEach(article => {
            const dateObj = dayjs(article.created_at);
            const dateStr = dateObj.format('YYYY年MM月');
            const rawDate = dateObj.format('YYYY-MM');
            
            if (!groups[dateStr]) {
                groups[dateStr] = { count: 0, rawDate: rawDate };
            }
            groups[dateStr].count++;
        });
        archives.value = Object.keys(groups).map(date => ({
            date,
            count: groups[date].count,
            rawDate: groups[date].rawDate
        }));

        siteInfo.value.articleCount = allRes.total;
        siteInfo.value.totalViews = list.reduce((acc, cur) => acc + (cur.views || 0), 0).toLocaleString();
        siteInfo.value.wordCount = (list.length * 1200).toLocaleString(); 
        
        const startDate = dayjs('2023-01-01');
        const now = dayjs();
        siteInfo.value.runningTime = now.diff(startDate, 'day') + ' 天';
        
        if (list.length > 0) {
            siteInfo.value.lastUpdate = dayjs(list[0].created_at).fromNow();
        }

    } catch(e) {
        console.error(e);
    }
}

const goDetail = (id) => {
  router.push(`/article/${id}`);
};

const goCategory = (id) => {
    router.push(`/category/${id}`);
};
const goTag = (id) => {
    router.push(`/tag/${id}`);
};

const goArchive = (date) => {
    router.push({ path: '/archive', query: { date } });
};

const formatDate = (date) => {
  return dayjs(date).format('YYYY-MM-DD');
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

/* Scroll Reveal Animation Classes - Deprecated / Replaced by global anim-scale-in */
/* .scroll-reveal logic removed as we use anim-scale-in now */

/* Zigzag Layout for Articles */
/* Remove old :deep CSS logic as it is now handled by ArticleCard props */
/* Clean up unused selectors */
:deep(.article-card:nth-child(even)),
:deep(.article-card:nth-child(even) .article-content),
:deep(.article-card:nth-child(even) .article-header),
:deep(.article-card:nth-child(even) .article-footer) {
    /* Styles handled by prop */
}

/* Remove any order overrides to ensure natural flex flow */
:deep(.article-card .article-footer .views),
:deep(.article-card .article-footer .read-more) {
    order: unset;
}

/* Sidebar */
.sidebar {
    position: sticky;
    top: 90px;
}

.sidebar-card {
    padding: 20px;
    margin-bottom: 0; /* Handled by flex gap now */
    /* opacity: 0; REMOVED to ensure visibility if animation fails */
    animation: fadeInUp 1s ease-out forwards;
}

@keyframes fadeInUp {
    from { opacity: 0; transform: translateY(20px); }
    to { opacity: 1; transform: translateY(0); }
}

/* Staggered animation delays for sidebar items */
.sidebar > *:nth-child(1) { animation-delay: 0.2s; }
.sidebar > *:nth-child(2) { animation-delay: 0.4s; }
.sidebar > *:nth-child(3) { animation-delay: 0.6s; }
.sidebar > *:nth-child(4) { animation-delay: 0.8s; }
.sidebar > *:nth-child(5) { animation-delay: 1.0s; }

/* Flex container for sidebar to manage spacing consistently */
.sidebar {
    display: flex;
    flex-direction: column;
    gap: 40px; /* Increased consistent spacing */
}

/* Remove manual margins since we use gap */
.mt-24 {
    margin-top: 0 !important;
}

.card-header {
    font-size: 1.1rem;
    font-weight: 700;
    color: var(--text-main);
    margin-bottom: 15px;
    padding-bottom: 10px;
    border-bottom: 1px solid rgba(255,255,255,0.1);
    display: flex;
    align-items: center;
    gap: 8px;
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

/* Latest Articles List */
.latest-list {
  display: flex;
  flex-direction: column;
  gap: 8px; /* Slightly tighter gap */
}

.latest-item {
  display: flex;
  flex-direction: column;
  padding: 10px;
  cursor: pointer;
  border-radius: 8px;
  transition: all 0.2s;
}

.latest-item:hover {
  background: rgba(0,0,0,0.03);
}
:global(.dark) .latest-item:hover {
  background: rgba(255,255,255,0.05);
}

.latest-title {
  font-size: 0.95rem;
  color: var(--text-main);
  margin-bottom: 6px;
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2; /* Allow 2 lines for title */
  -webkit-box-orient: vertical;
  overflow: hidden;
  transition: color 0.2s;
  font-weight: 500;
}

.latest-item:hover .latest-title {
  color: var(--primary-color);
}

.latest-date {
  font-size: 0.8rem;
  color: var(--text-secondary);
}
</style>
