<template>
  <div class="archive-page">
    <el-timeline>
      <el-timeline-item
        v-for="article in articles"
        :key="article.id"
        :timestamp="formatDate(article.created_at)"
        placement="top"
      >
        <el-card shadow="hover" @click="goDetail(article.id)" class="archive-card">
          <h4>{{ article.title }}</h4>
        </el-card>
      </el-timeline-item>
    </el-timeline>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getArticles } from '@/api/article';
import { useRouter } from 'vue-router';
import dayjs from 'dayjs';

const router = useRouter();
const articles = ref([]);

const fetchData = async () => {
    try {
        const res = await getArticles({
            page: 1,
            page_size: 1000, // Get all acts as archive
            status: 1
        });
        articles.value = res.list;
    } catch(e) {
        console.error(e);
    }
};

const formatDate = (d) => dayjs(d).format('YYYY-MM-DD');
const goDetail = (id) => router.push(`/article/${id}`);

onMounted(() => {
    fetchData();
});
</script>

<style scoped>
.archive-page {
    background: white;
    padding: 40px;
    border-radius: 8px;
}
.archive-card {
    cursor: pointer;
}
</style>
