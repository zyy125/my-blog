<template>
  <div class="page-container">
    <el-card class="box-card" shadow="never">
      <template #header>
        <div class="card-header">
          <div class="left">
            <h3 class="title">分类管理</h3>
          </div>
          <div class="right">
            <el-button type="primary" :icon="Plus" @click="handleCreate">新建分类</el-button>
          </div>
        </div>
      </template>

      <el-table :data="tableData" v-loading="loading" style="width: 100%">
        <el-table-column prop="id" label="ID" width="80" />
        <el-table-column prop="name" label="名称" width="200">
            <template #default="{ row }">
                <span class="category-name">{{ row.name }}</span>
            </template>
        </el-table-column>
        <el-table-column prop="article_count" label="文章数" width="120">
            <template #default="{ row }">
                <el-tag type="info" effect="plain" round>{{ row.article_count || 0 }}</el-tag>
            </template>
        </el-table-column>
        <el-table-column label="操作" min-width="150" align="right">
          <template #default="{ row }">
            <el-button size="small" :icon="Edit" @click="handleEdit(row)">编辑</el-button>
            <el-popconfirm title="确定删除该分类吗？" @confirm="handleDelete(row)">
              <template #reference>
                <el-button size="small" type="danger" :icon="Delete">删除</el-button>
              </template>
            </el-popconfirm>
          </template>
        </el-table-column>
      </el-table>
    </el-card>

    <!-- Dialog -->
    <el-dialog 
        v-model="dialogVisible" 
        :title="dialogTitle" 
        width="500px" 
        destroy-on-close
        @closed="resetForm"
    >
      <el-form 
        ref="formRef" 
        :model="form" 
        :rules="rules" 
        label-width="80px"
        label-position="top"
      >
        <el-form-item label="分类名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入分类名称" />
        </el-form-item>
        <el-form-item label="描述" prop="desc">
            <el-input v-model="form.desc" type="textarea" placeholder="可选" />
        </el-form-item>
      </el-form>
      
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitForm" :loading="submitting">确定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed, nextTick } from 'vue';
import { getCategoriesStats, createCategory, updateCategory, deleteCategory } from '@/api/category';
import { Plus, Edit, Delete } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';

// Data
const loading = ref(false);
const tableData = ref([]);
const dialogVisible = ref(false);
const submitting = ref(false);
const formRef = ref(null);
const isEdit = ref(false);
const currentId = ref(null);

const form = reactive({
    name: '',
    desc: ''
});

const rules = {
    name: [{ required: true, message: '请输入分类名称', trigger: 'blur' }]
};

const dialogTitle = computed(() => isEdit.value ? '编辑分类' : '新建分类');

// Methods
const fetchData = async () => {
    loading.value = true;
    try {
        const res = await getCategoriesStats();
        tableData.value = res;
    } catch (e) {
        console.error(e);
    } finally {
        loading.value = false;
    }
};

const handleCreate = () => {
    isEdit.value = false;
    currentId.value = null;
    form.name = '';
    form.desc = '';
    dialogVisible.value = true;
};

const handleEdit = (row) => {
    isEdit.value = true;
    currentId.value = row.id;
    form.name = row.name;
    form.desc = ''; // Backend doesn't return desc usually in list, ignore for now
    dialogVisible.value = true;
};

const handleDelete = async (row) => {
    try {
        await deleteCategory(row.id);
        ElMessage.success('删除成功');
        fetchData();
    } catch (e) {
        console.error(e);
    }
};

const submitForm = async () => {
    if (!formRef.value) return;
    
    await formRef.value.validate(async (valid) => {
        if (valid) {
            submitting.value = true;
            try {
                if (isEdit.value) {
                    await updateCategory(currentId.value, { name: form.name });
                    ElMessage.success('更新成功');
                } else {
                    await createCategory({ name: form.name });
                    ElMessage.success('创建成功');
                }
                dialogVisible.value = false;
                fetchData();
            } catch (e) {
                console.error(e);
            } finally {
                submitting.value = false;
            }
        }
    });
};

const resetForm = () => {
    if (formRef.value) formRef.value.resetFields();
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
.card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.title {
    margin: 0;
    font-size: 18px;
    font-weight: 600;
}
.category-name {
    font-weight: 500;
}
</style>