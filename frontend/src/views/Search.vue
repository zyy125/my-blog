<template>
  <div class="search-page">
     <div class="page-header">
         <h2>搜索结果: "{{ keyword }}"</h2>
     </div>
     
     <div class="article-list" v-loading="loading">
        <el-card v-for="article in articles" :key="article.id" class="article-card" shadow="hover" @click="goDetail(article.id)">
            <!-- Simple List Item -->
            <h3 class="article-title" v-html="highlight(article.title)"></h3>
            <p class="article-summary" v-html="highlight(article.summary)"></p>
        </el-card>

        <el-empty v-if="!loading && articles.length === 0" description="未找到相关文章" />
      </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getArticles } from '@/api/article';

const route = useRoute();
const router = useRouter();
const articles = ref([]);
const loading = ref(false);
const keyword = ref('');

const fetchData = async () => {
    const kw = route.query.keyword;
    if (!kw) return;
    
    keyword.value = kw;
    loading.value = true;
    try {
        const res = await getArticles({
            page: 1,
            page_size: 50, // Get more for search
            status: 1,
            keyword: kw
        });
        articles.value = res.list;
    } catch (e) {
        console.error(e);
    } finally {
        loading.value = false;
    }
};

const highlight = (text) => {
    if (!text) return '';
    if (!keyword.value) return text;
    const reg = new RegExp(`(${keyword.value})`, 'gi');
    return text.replace(reg, '<span style="color: red; font-weight: bold;">$1</span>');
};

const goDetail = (id) => router.push(`/article/${id}`);

watch(() => route.query.keyword, () => {
    fetchData();
});

onMounted(() => {
    fetchData();
});
</script>

<style scoped>
.page-header {
    margin-bottom: 20px;
}
.article-card {
    margin-bottom: 15px;
    cursor: pointer;
}
.article-title {
    margin: 0 0 10px 0;
}
.article-summary {
    color: #666;
    margin: 0;
    font-size: 14px;
}
</style>
