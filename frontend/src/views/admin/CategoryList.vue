<template>
  <el-card>
    <div class="toolbar">
      <el-button type="primary" :icon="Plus" @click="handleAdd">新建分类</el-button>
    </div>

    <el-table :data="categories" style="width: 100%; margin-top: 20px" v-loading="loading">
      <el-table-column prop="id" label="ID" width="80" />
      <el-table-column prop="name" label="名称" width="200" />
       <el-table-column prop="description" label="描述" />
      <el-table-column prop="article_count" label="文章数" width="100" />
      <el-table-column label="操作" width="180">
        <template #default="{ row }">
          <el-button size="small" type="primary" link @click="handleEdit(row)">编辑</el-button>
          <el-popconfirm title="确定删除吗? 删除前请确保该分类下无文章。" @confirm="handleDelete(row)">
            <template #reference>
              <el-button size="small" type="danger" link>删除</el-button>
            </template>
          </el-popconfirm>
        </template>
      </el-table-column>
    </el-table>

    <!-- Dialog -->
    <el-dialog v-model="dialogVisible" :title="isEdit ? '编辑分类' : '新建分类'" width="500px">
      <el-form :model="form" ref="formRef" :rules="rules" label-width="60px">
        <el-form-item label="名称" prop="name">
          <el-input v-model="form.name" placeholder="请输入分类名称" />
        </el-form-item>
        <el-form-item label="描述" prop="description">
          <el-input v-model="form.description" type="textarea" placeholder="请输入分类描述" />
        </el-form-item>
      </el-form>
      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogVisible = false">取消</el-button>
          <el-button type="primary" @click="submitForm" :loading="submitting">确定</el-button>
        </span>
      </template>
    </el-dialog>
  </el-card>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { getCategoriesStats, createCategory, updateCategory, deleteCategory } from '@/api/category';
import { Plus } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';

const categories = ref([]);
const loading = ref(false);
const dialogVisible = ref(false);
const isEdit = ref(false);
const submitting = ref(false);
const currentId = ref(null);

const form = reactive({
  name: '',
  description: ''
});
const formRef = ref(null);

const rules = {
  name: [{ required: true, message: '请输入名称', trigger: 'blur' }]
};

const fetchData = async () => {
    loading.value = true;
    try {
        const res = await getCategoriesStats();
        categories.value = Array.isArray(res) ? res : (res.list || []);
    } catch (e) {
        console.error(e);
    } finally {
        loading.value = false;
    }
};

const handleAdd = () => {
    isEdit.value = false;
    form.name = '';
    form.description = '';
    dialogVisible.value = true;
};

const handleEdit = (row) => {
    isEdit.value = true;
    currentId.value = row.id;
    form.name = row.name;
    form.description = row.description;
    dialogVisible.value = true;
};

const handleDelete = async (row) => {
    if (row.article_count > 0) {
        ElMessage.warning('该分类下有文章，无法删除');
        return;
    }
    try {
        await deleteCategory(row.id);
        ElMessage.success('删除成功');
        fetchData();
    } catch (e) {
    }
};

const submitForm = async () => {
    if (!formRef.value) return;
    await formRef.value.validate(async (valid) => {
        if (valid) {
            submitting.value = true;
            try {
                if (isEdit.value) {
                    await updateCategory(currentId.value, form);
                    ElMessage.success('更新成功');
                } else {
                    await createCategory(form);
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

onMounted(() => {
    fetchData();
});
</script>

<style scoped>
.toolbar {
    margin-bottom: 20px;
}
</style>
