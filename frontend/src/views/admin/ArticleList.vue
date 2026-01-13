<template>
  <el-card>
    <div class="toolbar">
        <div class="left">
             <el-button type="primary" :icon="Plus" @click="$router.push('/admin/articles/create')">写文章</el-button>
        </div>
        <div class="right">
             <el-select v-model="filterStatus" placeholder="状态" clearable style="width: 120px; margin-right: 10px" @change="fetchData">
                 <el-option label="已发布" :value="1" />
                 <el-option label="草稿" :value="0" />
             </el-select>
             <el-input v-model="searchKeyword" placeholder="搜索标题..." style="width: 200px;" @keyup.enter="fetchData" :prefix-icon="Search" />
        </div>
    </div>
    
    <el-table :data="articles" style="width: 100%; margin-top: 20px" v-loading="loading">
       <el-table-column prop="id" label="ID" width="60" />
       <el-table-column prop="title" label="标题" min-width="200" show-overflow-tooltip />
       <el-table-column label="分类" width="120">
           <template #default="{ row }">{{ row.category?.name || '-' }}</template>
       </el-table-column>
       <el-table-column label="标签" width="150" show-overflow-tooltip>
           <template #default="{ row }">
               <el-tag size="small" v-for="tag in row.tags" :key="tag.id" style="margin-right: 4px">{{ tag.name }}</el-tag>
           </template>
       </el-table-column>
       <el-table-column label="状态" width="100">
          <template #default="{ row }">
             <el-tag :type="row.status === 1 ? 'success' : 'info'">{{ row.status === 1 ? '已发布' : '草稿' }}</el-tag>
          </template>
       </el-table-column>
       <el-table-column prop="views" label="浏览" width="80" />
       <el-table-column label="发布时间" width="180">
           <template #default="{ row }">{{ formatDate(row.created_at) }}</template>
       </el-table-column>
       <el-table-column label="操作" width="180" fixed="right">
          <template #default="{ row }">
             <el-button size="small" type="primary" link @click="$router.push(`/admin/articles/edit/${row.id}`)">编辑</el-button>
             <el-popconfirm title="确定删除吗?" @confirm="handleDelete(row)">
                 <template #reference>
                     <el-button size="small" type="danger" link>删除</el-button>
                 </template>
             </el-popconfirm>
          </template>
       </el-table-column>
    </el-table>

    <div class="pagination">
      <el-pagination 
         background 
         layout="total, prev, pager, next" 
         :total="total" 
         v-model:current-page="page" 
         :page-size="pageSize"
         @current-change="fetchData"
      />
    </div>
  </el-card>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getArticles, deleteArticle } from '@/api/article';
import dayjs from 'dayjs';
import { Plus, Search } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';

const articles = ref([]);
const loading = ref(false);
const total = ref(0);
const page = ref(1);
const pageSize = ref(10);
const searchKeyword = ref('');
const filterStatus = ref('');

const fetchData = async () => {
    loading.value = true;
    try {
        const params = {
            page: page.value,
            page_size: pageSize.value,
            keyword: searchKeyword.value,
        };
        if (filterStatus.value !== '') {
            params.status = filterStatus.value;
        }
        const res = await getArticles(params);
        articles.value = res.list;
        total.value = res.total;
    } catch (e) {
        console.error(e);
    } finally {
        loading.value = false;
    }
};

const handleDelete = async (row) => {
    try {
        await deleteArticle(row.id);
        ElMessage.success('删除成功');
        fetchData();
    } catch (e) {
        console.error(e);
    }
};

const formatDate = (d) => dayjs(d).format('YYYY-MM-DD HH:mm');

onMounted(() => {
    fetchData();
});
</script>

<style scoped>
.toolbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.pagination {
    margin-top: 20px;
    display: flex;
    justify-content: flex-end;
}
</style>
