<template>
  <el-card>
     <div class="toolbar">
         <el-radio-group v-model="statusFilter" @change="fetchData">
             <el-radio-button :label="0">待审核</el-radio-button>
             <el-radio-button :label="1">已通过</el-radio-button>
             <el-radio-button :label="2">已拒绝</el-radio-button>
         </el-radio-group>
     </div>

     <el-table :data="comments" style="width: 100%; margin-top: 20px" v-loading="loading">
        <el-table-column prop="nickname" label="昵称" width="120" />
        <el-table-column prop="content" label="内容" min-width="200" show-overflow-tooltip/>
        <el-table-column label="文章" min-width="150" show-overflow-tooltip>
             <template #default="{ row }">
                 <el-link v-if="row.article" :href="`/article/${row.article.id}`" target="_blank">{{ row.article.title }}</el-link>
                 <span v-else>已删除文章</span>
             </template>
        </el-table-column>
        <el-table-column label="时间" width="160">
            <template #default="{ row }">{{ formatDate(row.created_at) }}</template>
        </el-table-column>
        <el-table-column label="操作" width="200" fixed="right">
            <template #default="{ row }">
                <el-button v-if="row.status === 0 || row.status === 2" size="small" type="success" @click="handleStatus(row, 'approve')">通过</el-button>
                <el-button v-if="row.status === 0 || row.status === 1" size="small" type="warning" @click="handleStatus(row, 'reject')">拒绝</el-button>
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
import { getAdminComments, approveComment, rejectComment, deleteComment } from '@/api/comment';
import dayjs from 'dayjs';
import { ElMessage } from 'element-plus';

const comments = ref([]);
const loading = ref(false);
const statusFilter = ref(0);
const page = ref(1);
const pageSize = ref(10);
const total = ref(0);

const fetchData = async () => {
    loading.value = true;
    try {
        const res = await getAdminComments({
            page: page.value,
            page_size: pageSize.value,
            status: statusFilter.value
        });
        comments.value = res.list;
        total.value = res.total;
    } catch (e) {
        console.error(e);
    } finally {
        loading.value = false;
    }
};

const handleStatus = async (row, action) => {
    try {
        if (action === 'approve') {
            await approveComment(row.id);
            ElMessage.success('已通过');
        } else {
            await rejectComment(row.id);
            ElMessage.success('已拒绝');
        }
        fetchData();
    } catch (e) {
        console.error(e);
    }
};

const handleDelete = async (row) => {
    try {
        await deleteComment(row.id);
        ElMessage.success('删除成功');
        fetchData();
    } catch(e) {
        console.error(e);
    }
};

const formatDate = (date) => dayjs(date).format('YYYY-MM-DD HH:mm');

onMounted(() => {
    fetchData();
});
</script>

<style scoped>
.toolbar {
    margin-bottom: 20px;
}
.pagination {
    margin-top: 20px;
    display: flex;
    justify-content: flex-end;
}
</style>
