<template>
  <div class="category-page">
     <div class="page-header">
         <h2>分类: {{ categoryName }}</h2>
     </div>
     
     <div class="article-list" v-loading="loading">
        <el-card v-for="article in articles" :key="article.id" class="article-card" shadow="hover" @click="goDetail(article.id)">
          <div class="article-content">
            <h2 class="article-title">{{ article.title }}</h2>
             <div class="article-meta">
              <span class="meta-item"><el-icon><Calendar /></el-icon> {{ formatDate(article.created_at) }}</span>
              <span class="meta-item"><el-icon><View /></el-icon> {{ article.views }}</span>
            </div>
            <p class="article-summary">{{ article.summary }}</p>
          </div>
          <div class="article-cover" v-if="article.cover_img">
            <img :src="article.cover_img" alt="cover">
          </div>
        </el-card>

        <el-empty v-if="!loading && articles.length === 0" description="该分类下暂无文章" />

        <div class="pagination" v-if="total > 0">
          <el-pagination
            background
            layout="prev, pager, next"
            :total="total"
            :page-size="pageSize"
            v-model:current-page="currentPage"
            @current-change="fetchData"
          />
        </div>
      </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getArticles } from '@/api/article';
import { getCategories } from '@/api/category';
import dayjs from 'dayjs';
import { Calendar, View } from '@element-plus/icons-vue';

const route = useRoute();
const router = useRouter();
const articles = ref([]);
const loading = ref(false);
const total = ref(0);
const pageSize = ref(10);
const currentPage = ref(1);
const categoryName = ref('');

const fetchCategoryInfo = async () => {
    try {
        const id = parseInt(route.params.id);
        const res = await getCategories();
        const list = Array.isArray(res) ? res : res.list;
        const cat = list.find(c => c.id === id);
        if (cat) categoryName.value = cat.name;
    } catch(e) {}
};

const fetchData = async () => {
    loading.value = true;
    try {
        const res = await getArticles({
            page: currentPage.value,
            page_size: pageSize.value,
            status: 1,
            category_id: route.params.id
        });
        articles.value = res.list;
        total.value = res.total;
    } catch (e) {
        console.error(e);
    } finally {
        loading.value = false;
    }
};

const formatDate = (date) => dayjs(date).format('YYYY-MM-DD');
const goDetail = (id) => router.push(`/article/${id}`);

watch(() => route.params.id, () => {
    currentPage.value = 1;
    categoryName.value = '';
    fetchData();
    fetchCategoryInfo();
});

onMounted(() => {
    fetchData();
    fetchCategoryInfo();
});
</script>

<style scoped>
.page-header {
    margin-bottom: 30px;
    border-bottom: 2px solid #409eff;
    padding-bottom: 10px;
}
.article-card {
  margin-bottom: 20px;
  cursor: pointer;
  transition: all 0.3s;
}
.article-card:hover {
  transform: translateY(-2px);
}
:deep(.el-card__body) {
    display: flex;
    justify-content: space-between;
}
.article-content {
    flex: 1;
}
.article-title {
  margin: 0 0 10px 0;
  font-size: 20px;
  color: #333;
}
.article-meta {
  margin-bottom: 10px;
  color: #999;
  font-size: 13px;
  display: flex;
  align-items: center;
  gap: 15px;
}
.meta-item { display: flex; align-items: center; gap: 4px; }
.article-summary {
  color: #666;
  line-height: 1.6;
  margin: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
}
.article-cover {
    width: 150px;
    height: 100px;
    margin-left: 20px;
    flex-shrink: 0;
}
.article-cover img {
    width: 100%;
    height: 100%;
    object-fit: cover;
    border-radius: 4px;
}
.pagination {
  display: flex;
  justify-content: center;
  margin-top: 40px;
}
</style>
