<template>
  <div class="page-container">
    <el-card class="box-card" shadow="never">
      <div class="filter-container">
        <div class="left-actions">
          <el-input
            v-model="keyword"
            placeholder="搜索文章标题..."
            style="width: 240px"
            clearable
            @keyup.enter="handleFilter"
            @clear="handleFilter"
          >
            <template #prefix>
              <el-icon><Search /></el-icon>
            </template>
          </el-input>
          
          <el-select 
            v-model="statusFilter" 
            placeholder="发布状态" 
            clearable 
            style="width: 140px"
            @change="handleFilter"
          >
            <el-option label="已发布" :value="1" />
            <el-option label="草稿" :value="0" />
          </el-select>
          
          <el-button type="primary" @click="handleFilter">搜索</el-button>
        </div>
        
        <div class="right-actions">
          <el-button type="primary" :icon="Plus" @click="$router.push('/admin/articles/create')">
            写文章
          </el-button>
        </div>
      </div>

      <el-table 
        :data="tableData" 
        v-loading="loading" 
        style="width: 100%"
        stripe
      >
        <el-table-column prop="id" label="ID" width="80" align="center" />
        
        <el-table-column label="标题" min-width="250">
          <template #default="{ row }">
            <div class="article-title-cell">
              <span class="title-text">{{ row.title }}</span>
              <el-tag v-if="row.is_top" type="danger" size="small" effect="dark" class="top-tag">置顶</el-tag>
            </div>
          </template>
        </el-table-column>
        
        <el-table-column label="分类" width="140">
          <template #default="{ row }">
            <el-tag v-if="row.category" type="info" effect="plain">{{ row.category.name }}</el-tag>
            <span v-else class="text-placeholder">-</span>
          </template>
        </el-table-column>
        
        <el-table-column label="标签" width="200">
          <template #default="{ row }">
            <div class="tags-wrapper">
              <el-tag 
                v-for="tag in (row.tags || []).slice(0, 3)" 
                :key="tag.id" 
                size="small" 
                type="info" 
                effect="light"
                class="mr-1"
              >
                {{ tag.name }}
              </el-tag>
              <span v-if="(row.tags || []).length > 3" class="more-tags">...</span>
            </div>
          </template>
        </el-table-column>
        
        <el-table-column label="状态" width="100" align="center">
          <template #default="{ row }">
            <el-tag :type="row.status === 1 ? 'success' : 'warning'" effect="light" round>
              {{ row.status === 1 ? '已发布' : '草稿' }}
            </el-tag>
          </template>
        </el-table-column>
        
        <el-table-column prop="views" label="浏览量" width="100" align="center" sortable />
        
        <el-table-column label="发布时间" width="160" align="center">
          <template #default="{ row }">
            <span class="time-text">{{ formatDate(row.created_at) }}</span>
          </template>
        </el-table-column>
        
        <el-table-column label="操作" width="150" align="center" fixed="right">
          <template #default="{ row }">
            <el-tooltip content="编辑" placement="top">
              <el-button 
                type="primary" 
                link 
                :icon="Edit" 
                @click="$router.push(`/admin/articles/edit/${row.id}`)"
              />
            </el-tooltip>
            
            <el-popconfirm 
              title="确定要删除这篇文章吗？此操作不可恢复。" 
              confirm-button-text="删除" 
              cancel-button-text="取消"
              confirm-button-type="danger"
              @confirm="handleDelete(row)"
            >
              <template #reference>
                <el-button type="danger" link :icon="Delete" />
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>

      <div class="pagination-container">
        <el-pagination
          v-model:current-page="page"
          v-model:page-size="pageSize"
          :page-sizes="[10, 20, 50]"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          background
          @size-change="handleFilter"
          @current-change="handleFilter"
        />
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { getArticles, deleteArticle } from '@/api/article';
import { Search, Plus, Edit, Delete } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';
import dayjs from 'dayjs';

const loading = ref(false);
const tableData = ref([]);
const total = ref(0);
const page = ref(1);
const pageSize = ref(10);
const keyword = ref('');
const statusFilter = ref('');

const fetchData = async () => {
  loading.value = true;
  try {
    const params = {
      page: page.value,
      page_size: pageSize.value,
      keyword: keyword.value,
      status: statusFilter.value
    };
    const res = await getArticles(params);
    tableData.value = res.list || [];
    total.value = res.total || 0;
  } catch (e) {
    console.error(e);
  } finally {
    loading.value = false;
  }
};

const handleFilter = () => {
  page.value = 1; // Reset to first page on filter change
  fetchData();
};

const handleDelete = async (row) => {
  try {
    await deleteArticle(row.id);
    ElMessage.success('删除成功');
    // If last item on page, go back one page
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--;
    }
    fetchData();
  } catch (e) {
    console.error(e);
  }
};

const formatDate = (date) => {
  return dayjs(date).format('YYYY-MM-DD HH:mm');
};

onMounted(() => {
  fetchData();
});
</script>

<style scoped>
.page-container {
    padding: 0;
}
.box-card {
    border: none;
    box-shadow: none;
    background: transparent;
}
.filter-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
    flex-wrap: wrap;
    gap: 15px;
}
.left-actions {
    display: flex;
    align-items: center;
    gap: 12px;
}
.article-title-cell {
    display: flex;
    align-items: center;
    gap: 8px;
}
.title-text {
    font-weight: 500;
    color: var(--text-primary);
}
.top-tag {
    transform: scale(0.9);
}
.text-placeholder {
    color: var(--text-secondary);
}
.tags-wrapper {
    display: flex;
    flex-wrap: wrap;
    gap: 4px;
}
.time-text {
    color: var(--text-regular);
    font-size: 13px;
}
.pagination-container {
    margin-top: 24px;
    display: flex;
    justify-content: flex-end;
}
</style>