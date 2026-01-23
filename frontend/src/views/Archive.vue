<template>
  <div class="archive-page fade-in-up">
    <GlassCard class="archive-container">
        <div class="archive-header">
            <h2>归档 {{ queryDate ? `· ${queryDate}` : '' }}</h2>
            <el-button v-if="queryDate" link type="primary" @click="clearFilter">查看全部</el-button>
        </div>
        <el-timeline>
          <el-timeline-item
            v-for="article in filteredArticles"
            :key="article.id"
            :timestamp="formatDate(article.created_at)"
            placement="top"
            type="primary"
            size="large"
            hollow
          >
            <div class="archive-item" @click="goDetail(article.id)">
              <h4 class="article-title">{{ article.title }}</h4>
              <span class="read-more">阅读 <el-icon><ArrowRight /></el-icon></span>
            </div>
          </el-timeline-item>
        </el-timeline>
        
        <el-empty v-if="filteredArticles.length === 0" description="暂无文章" />
    </GlassCard>
  </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue';
import { getArticles } from '@/api/article';
import { useRouter, useRoute } from 'vue-router';
import dayjs from 'dayjs';
import GlassCard from '@/components/GlassCard.vue';
import { ArrowRight } from '@element-plus/icons-vue';

const router = useRouter();
const route = useRoute();
const articles = ref([]);
const queryDate = ref(route.query.date || '');

const filteredArticles = computed(() => {
    if (!queryDate.value) return articles.value;
    return articles.value.filter(a => {
        return dayjs(a.created_at).format('YYYY-MM') === queryDate.value;
    });
});

const fetchData = async () => {
    try {
        const res = await getArticles({
            page: 1,
            page_size: 1000,
            status: 1
        });
        articles.value = res.list;
    } catch(e) {
        console.error(e);
    }
};

const formatDate = (d) => dayjs(d).format('YYYY年MM月DD日');
const goDetail = (id) => router.push(`/article/${id}`);
const clearFilter = () => {
    router.push('/archive');
};

watch(() => route.query.date, (newVal) => {
    queryDate.value = newVal || '';
});

onMounted(() => {
    fetchData();
});
</script>

<style scoped>
.archive-page {
    max-width: 900px;
    margin: 20px auto;
}

.archive-container {
    padding: 40px;
    min-height: 600px;
}

.archive-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 40px;
    padding-bottom: 15px;
    border-bottom: 1px solid var(--border-color);
}

.archive-header h2 {
    margin: 0;
    font-size: 1.8rem;
    background: linear-gradient(120deg, var(--primary-color), var(--secondary-color));
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
}

.archive-item {
    cursor: pointer;
    padding: 15px 20px;
    background: rgba(255, 255, 255, 0.03);
    border-radius: 12px;
    border: 1px solid transparent;
    transition: all 0.3s ease;
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.archive-item:hover {
    background: rgba(255, 255, 255, 0.08);
    transform: translateX(5px);
    border-color: var(--primary-color);
}

.article-title {
    margin: 0;
    font-size: 1.1rem;
    font-weight: 500;
    color: var(--text-main);
}

.read-more {
    font-size: 0.85rem;
    color: var(--text-secondary);
    display: flex;
    align-items: center;
    gap: 4px;
    opacity: 0;
    transform: translateX(-10px);
    transition: all 0.3s ease;
}

.archive-item:hover .read-more {
    opacity: 1;
    transform: translateX(0);
    color: var(--primary-color);
}

/* Timeline Customization */
:deep(.el-timeline-item__timestamp) {
    color: var(--text-secondary);
    font-size: 0.9rem;
    margin-bottom: 8px;
}

:deep(.el-timeline-item__node--primary) {
    background-color: var(--primary-color);
}

:deep(.el-timeline-item__tail) {
    border-left: 2px solid rgba(255, 255, 255, 0.1);
}
</style>
