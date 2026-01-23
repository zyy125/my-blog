<template>
  <div class="article-edit-page">
    <div class="edit-header">
      <div class="header-left">
        <el-input 
          v-model="form.title" 
          placeholder="请输入文章标题..." 
          class="title-input" 
          maxlength="100"
        />
      </div>
      <div class="header-right">
        <span class="status-badge" :class="form.status === 1 ? 'published' : 'draft'">
            {{ form.status === 1 ? '已发布' : '草稿' }}
        </span>
        <el-button @click="openSettings" :icon="Setting" circle />
        <el-button type="primary" @click="onSubmit" :loading="submitting">
          {{ isEdit ? '更新' : '发布' }}
        </el-button>
      </div>
    </div>

    <div class="edit-container">
      <v-md-editor 
        v-model="form.content"
        height="calc(100vh - 80px)"
        :disabled-menus="[]"
        @upload-image="handleEditorUpload"
        placeholder="开始您的创作..."
        class="immersive-editor"
      />
    </div>

    <!-- Settings Drawer -->
    <el-drawer
      v-model="drawerVisible"
      title="文章设置"
      direction="rtl"
      size="350px"
    >
      <el-form :model="form" ref="formRef" :rules="rules" label-position="top">
        <el-form-item label="分类" prop="category_id">
          <el-select v-model="form.category_id" placeholder="选择分类" style="width: 100%">
            <el-option v-for="item in categories" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </el-form-item>

        <el-form-item label="标签" prop="tag_ids">
          <el-select v-model="form.tag_ids" multiple placeholder="选择标签" style="width: 100%" filterable>
            <el-option v-for="item in tags" :key="item.id" :label="item.name" :value="item.id" />
          </el-select>
        </el-form-item>

        <el-form-item label="封面图" prop="cover_img">
            <el-upload
            class="avatar-uploader"
            action="#"
            :show-file-list="false"
            :http-request="handleCoverUpload"
            >
            <img v-if="form.cover_img" :src="form.cover_img" class="avatar" />
            <div v-else class="avatar-uploader-placeholder">
                <el-icon class="avatar-uploader-icon"><Plus /></el-icon>
                <span>点击上传封面</span>
            </div>
            </el-upload>
        </el-form-item>

        <el-form-item label="摘要" prop="summary">
          <el-input type="textarea" v-model="form.summary" :rows="4" placeholder="文章摘要（可选，默认截取前100字）" />
        </el-form-item>

        <el-form-item label="发布状态" prop="status">
            <el-radio-group v-model="form.status">
                <el-radio :label="0">草稿</el-radio>
                <el-radio :label="1">发布</el-radio>
            </el-radio-group>
        </el-form-item>

        <el-form-item label="其他设置">
            <el-checkbox v-model="form.is_top">置顶文章</el-checkbox>
        </el-form-item>
      </el-form>
      <template #footer>
        <div style="flex: auto">
          <el-button @click="drawerVisible = false">取消</el-button>
          <el-button type="primary" @click="drawerVisible = false">确认</el-button>
        </div>
      </template>
    </el-drawer>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getArticleDetail, createArticle, updateArticle } from '@/api/article';
import { getCategories } from '@/api/category';
import { getTags } from '@/api/tag';
import { uploadImage } from '@/api/upload';
import { ElMessage } from 'element-plus';
import { Plus, Setting } from '@element-plus/icons-vue';

const route = useRoute();
const router = useRouter();
const formRef = ref(null);
const submitting = ref(false);
const drawerVisible = ref(false);

const isEdit = computed(() => !!route.params.id);

const form = reactive({
    title: '',
    content: '',
    summary: '',
    cover_img: '',
    category_id: null,
    tag_ids: [],
    status: 1,
    is_top: false
});

const categories = ref([]);
const tags = ref([]);

const rules = {
    title: [{ required: true, message: '请输入标题', trigger: 'blur' }],
    // content: [{ required: true, message: '请输入内容', trigger: 'blur' }], // Content validation is manual often better
    category_id: [{ required: true, message: '请选择分类', trigger: 'change' }]
};

// Fetch dependencies
const initData = async () => {
    try {
        const [catRes, tagRes] = await Promise.all([getCategories(), getTags()]);
        categories.value = Array.isArray(catRes) ? catRes : (catRes.list || []);
        tags.value = Array.isArray(tagRes) ? tagRes : (tagRes.list || []);
        
        if (isEdit.value) {
            const id = route.params.id;
            const article = await getArticleDetail(id);
            Object.assign(form, {
                title: article.title,
                content: article.content,
                summary: article.summary,
                cover_img: article.cover_img,
                category_id: article.category?.id,
                tag_ids: article.tags ? article.tags.map(t => t.id) : [],
                status: article.status,
                is_top: article.is_top
            });
        }
    } catch (e) {
        console.error(e);
    }
};

const openSettings = () => {
    drawerVisible.value = true;
};

const handleCoverUpload = async (options) => {
    try {
        const formData = new FormData();
        formData.append('file', options.file);
        const res = await uploadImage(formData);
        form.cover_img = res.url; 
    } catch (e) {
        ElMessage.error('上传失败');
    }
};

const handleEditorUpload = async (event, insertImage, files) => {
  try {
      const formData = new FormData();
      formData.append('file', files[0]);
      const res = await uploadImage(formData);
      insertImage({
        url: res.url,
        desc: files[0].name,
      });
  } catch(e) {
      ElMessage.error('上传图片失败');
  }
};

const onSubmit = async () => {
    if (!form.title) {
        ElMessage.warning('请输入文章标题');
        return;
    }
    if (!form.content) {
        ElMessage.warning('请输入文章内容');
        return;
    }
    
    // We validate other fields (like category) only if drawer was opened or we force check
    // But usually for draft we might be lenient. Let's strict check category.
    if (!form.category_id) {
        ElMessage.warning('请在设置中选择文章分类');
        drawerVisible.value = true;
        return;
    }

    submitting.value = true;
    try {
        if (isEdit.value) {
            await updateArticle(route.params.id, form);
            ElMessage.success('更新成功');
        } else {
            await createArticle(form);
            ElMessage.success('发布成功');
        }
        router.push('/admin/articles');
    } catch (e) {
        console.error(e);
    } finally {
        submitting.value = false;
    }
};

onMounted(() => {
    initData();
});
</script>

<style scoped>
.article-edit-page {
  height: 100vh;
  display: flex;
  flex-direction: column;
  background-color: var(--bg-color);
}

.edit-header {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 20px;
  background-color: var(--card-bg);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid var(--border-color);
  z-index: 100;
}

.header-left {
  flex: 1;
  margin-right: 20px;
}

.title-input :deep(.el-input__wrapper) {
  background: transparent !important;
  box-shadow: none !important;
  font-size: 1.5rem;
  padding: 0;
}

.title-input :deep(.el-input__inner) {
  font-weight: 700;
  color: var(--text-main);
}

.header-right {
  display: flex;
  align-items: center;
  gap: 15px;
}

.status-badge {
    font-size: 0.85rem;
    padding: 2px 8px;
    border-radius: 4px;
    background: rgba(255,255,255,0.1);
    color: var(--text-secondary);
}
.status-badge.published { color: #67c23a; background: rgba(103, 194, 58, 0.1); }
.status-badge.draft { color: #e6a23c; background: rgba(230, 162, 60, 0.1); }


.edit-container {
  flex: 1;
  overflow: hidden;
}

.immersive-editor {
  border: none !important;
  box-shadow: none !important;
}

/* Upload Styles */
.avatar-uploader-placeholder {
  width: 100%;
  height: 150px;
  border: 1px dashed var(--border-color);
  border-radius: 8px;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  color: var(--text-secondary);
  cursor: pointer;
  transition: all 0.3s;
}

.avatar-uploader-placeholder:hover {
  border-color: var(--primary-color);
  color: var(--primary-color);
  background: rgba(255,255,255,0.02);
}

.avatar {
  width: 100%;
  height: 150px;
  object-fit: cover;
  border-radius: 8px;
}
</style>
