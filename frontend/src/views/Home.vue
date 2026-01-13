<template>
  <el-row :gutter="20">
    <el-col :span="18" :xs="24">
      <div class="article-list">
        <el-card v-for="article in articles" :key="article.id" class="article-card" shadow="hover" @click="goDetail(article.id)">
          <div class="article-content">
            <h2 class="article-title">{{ article.title }}</h2>
             <div class="article-meta">
              <el-tag size="small" v-if="article.category">{{ article.category.name }}</el-tag>
              <span class="meta-item"><el-icon><Calendar /></el-icon> {{ formatDate(article.created_at) }}</span>
              <span class="meta-item"><el-icon><View /></el-icon> {{ article.views }}</span>
            </div>
            <p class="article-summary">{{ article.summary }}</p>
          </div>
          <div class="article-cover" v-if="article.cover_img">
            <img :src="article.cover_img" alt="cover">
          </div>
        </el-card>

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
    </el-col>

    <el-col :span="6" :xs="0">
      <el-card class="sidebar-card">
        <template #header>
          <div class="card-header">
            <span>分类</span>
          </div>
        </template>
        <div class="category-list">
          <div v-for="cat in categories" :key="cat.id" class="category-item" @click="goCategory(cat.id)">
            <span>{{ cat.name }}</span>
            <el-tag size="small" type="info" effect="plain">{{ cat.article_count }}</el-tag>
          </div>
        </div>
      </el-card>

      <el-card class="sidebar-card mt-20">
        <template #header>
          <div class="card-header">
            <span>标签</span>
          </div>
        </template>
        <div class="tag-cloud">
          <el-tag 
            v-for="tag in tags" 
            :key="tag.id" 
            class="tag-item" 
            @click="goTag(tag.id)"
            effect="plain"
          >
            {{ tag.name }} ({{ tag.article_count }})
          </el-tag>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { getArticles } from '@/api/article';
import { getCategoriesStats } from '@/api/category';
import { getTagsStats } from '@/api/tag';
import dayjs from 'dayjs';

const router = useRouter();
const articles = ref([]);
const total = ref(0);
const pageSize = ref(10);
const currentPage = ref(1);

const categories = ref([]);
const tags = ref([]);

const fetchArticles = async () => {
  try {
    const res = await getArticles({
      page: currentPage.value,
      page_size: pageSize.value,
      status: 1
    });
    articles.value = res.list;
    total.value = res.total;
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
  return dayjs(date).format('YYYY-MM-DD');
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
</script>

<style scoped>
.article-card {
  margin-bottom: 20px;
  cursor: pointer;
  transition: all 0.3s;
  display: flex;
}
.article-card:hover {
  transform: translateY(-2px);
}
:deep(.el-card__body) {
    display: flex;
    width: 100%;
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
.meta-item {
    display: flex;
    align-items: center;
    gap: 4px;
}
.article-summary {
  color: #666;
  line-height: 1.6;
  margin: 0;
  overflow: hidden;
  text-overflow: ellipsis;
  display: -webkit-box;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
}
.article-cover {
    width: 200px;
    height: 140px;
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

.sidebar-card {
  margin-bottom: 20px;
}
.category-item {
  display: flex;
  justify-content: space-between;
  padding: 10px 0;
  border-bottom: 1px solid #f0f0f0;
  cursor: pointer;
}
.category-item:hover {
    color: #409eff;
}
.category-item:last-child {
  border-bottom: none;
}
.tag-cloud {
  display: flex;
  flex-wrap: wrap;
  gap: 10px;
}
.tag-item {
    cursor: pointer;
}
.mt-20 {
    margin-top: 20px;
}

@media (max-width: 768px) {
    .article-cover {
        display: none;
    }
}
</style>
