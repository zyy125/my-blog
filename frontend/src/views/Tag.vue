<template>
  <div class="tag-page">
     <div class="page-header">
         <h2>标签: {{ tagName }}</h2>
     </div>
     
     <div class="article-list" v-loading="loading">
        <el-card v-for="article in articles" :key="article.id" class="article-card" shadow="hover" @click="goDetail(article.id)">
          <div class="article-content">
            <h2 class="article-title">{{ article.title }}</h2>
             <div class="article-meta">
              <span class="meta-item"><el-icon><Calendar /></el-icon> {{ formatDate(article.created_at) }}</span>
              <span class="meta-item"><el-icon><View /></el-icon> {{ article.views }}</span>
            </div>
          </div>
        </el-card>

        <el-empty v-if="!loading && articles.length === 0" description="该标签下暂无文章" />

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
import { getTags } from '@/api/tag';
import dayjs from 'dayjs';
import { Calendar, View } from '@element-plus/icons-vue';

const route = useRoute();
const router = useRouter();
const total = ref(0);
const pageSize = ref(10);
const currentPage = ref(1);
const tagName = ref('');
const loading = ref(false);
const articles = ref([]);

const fetchTagInfo = async () => {
    try {
        const id = parseInt(route.params.id);
        const res = await getTags();
        const list = Array.isArray(res) ? res : res.list;
        const tag = list.find(t => t.id === id);
        if (tag) tagName.value = tag.name;
    } catch(e) {}
};

const fetchData = async (page) => {
    if (typeof page === 'number') {
        window.scrollTo({ top: 0, behavior: 'smooth' });
    }
    loading.value = true;
    try {
        const res = await getArticles({
            page: currentPage.value,
            page_size: pageSize.value,
            status: 1,
            tag_id: parseInt(route.params.id)
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
    tagName.value = '';
    fetchData();
    fetchTagInfo();
});

onMounted(() => {
    fetchData();
    fetchTagInfo();
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
}
.article-title {
  font-size: 18px;
  margin-bottom: 10px;
}
.article-meta {
  color: #999;
  font-size: 13px;
  display: flex;
  gap: 15px;
}
.pagination {
  display: flex;
  justify-content: center;
  margin-top: 40px;
}
</style>
