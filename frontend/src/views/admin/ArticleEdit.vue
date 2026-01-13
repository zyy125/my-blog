<template>
  <div class="article-edit">
    <el-card>
      <template #header>
          <div class="card-header">
              <span>{{ isEdit ? '编辑文章' : '写文章' }}</span>
          </div>
      </template>

      <el-form :model="form" ref="formRef" :rules="rules" label-width="80px">
        <el-row :gutter="20">
            <el-col :span="16">
                 <el-form-item label="标题" prop="title">
                    <el-input v-model="form.title" placeholder="请输入文章标题" />
                 </el-form-item>
            </el-col>
            <el-col :span="8">
                 <el-form-item label="分类" prop="category_id">
                    <el-select v-model="form.category_id" placeholder="选择分类" style="width: 100%">
                        <el-option v-for="item in categories" :key="item.id" :label="item.name" :value="item.id" />
                    </el-select>
                 </el-form-item>
            </el-col>
        </el-row>

        <el-row :gutter="20">
            <el-col :span="12">
                <el-form-item label="标签" prop="tag_ids">
                    <el-select v-model="form.tag_ids" multiple placeholder="选择标签" style="width: 100%" filterable>
                        <el-option v-for="item in tags" :key="item.id" :label="item.name" :value="item.id" />
                    </el-select>
                </el-form-item>
            </el-col>
             <el-col :span="12">
                <el-form-item label="置顶" prop="is_top">
                    <el-switch v-model="form.is_top" />
                </el-form-item>
            </el-col>
        </el-row>

        <el-form-item label="封面" prop="cover_img">
             <el-upload
                class="avatar-uploader"
                action="#"
                :show-file-list="false"
                :http-request="handleCoverUpload"
             >
                <img v-if="form.cover_img" :src="form.cover_img" class="avatar" />
                <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
             </el-upload>
             <div class="el-upload__tip">点击上传封面图，支持 jpg/png</div>
        </el-form-item>

        <el-form-item label="摘要" prop="summary">
            <el-input type="textarea" v-model="form.summary" :rows="2" placeholder="如果不填，将自动截取内容" />
        </el-form-item>

        <el-form-item label="状态" prop="status">
             <el-radio-group v-model="form.status">
                 <el-radio :label="0">草稿</el-radio>
                 <el-radio :label="1">发布</el-radio>
             </el-radio-group>
        </el-form-item>

        <el-form-item label="内容" prop="content">
             <v-md-editor 
                v-model="form.content"
                height="600px"
                :disabled-menus="[]"
                @upload-image="handleEditorUpload"
                placeholder="开始编辑..."
             />
        </el-form-item>

        <el-form-item>
             <el-button type="primary" @click="onSubmit" :loading="submitting">保存文章</el-button>
             <el-button @click="$router.back()">取消</el-button>
        </el-form-item>
      </el-form>
    </el-card>
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
import { Plus } from '@element-plus/icons-vue';

const route = useRoute();
const router = useRouter();
const formRef = ref(null);
const submitting = ref(false);

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
    content: [{ required: true, message: '请输入内容', trigger: 'blur' }],
    category_id: [{ required: true, message: '请选择分类', trigger: 'change' }]
};

// Fetch dependencies
const initData = async () => {
    try {
        const [catRes, tagRes] = await Promise.all([getCategories(), getTags()]);
        // getCategories -> {data: [...] } or just [...] depend on api wrapper?
        // My wrapper returns data part of response.
        // Doc: "data": [{"id": 1...}]
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

const handleCoverUpload = async (options) => {
    try {
        const formData = new FormData();
        formData.append('file', options.file);
        const res = await uploadImage(formData);
        // data.url
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
    if (!formRef.value) return;
    await formRef.value.validate(async (valid) => {
        if (valid) {
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
        }
    });
};

onMounted(() => {
    initData();
});
</script>

<style scoped>
.avatar-uploader .avatar {
  width: 178px;
  height: 178px;
  display: block;
  object-fit: cover;
}
.avatar-uploader .el-upload {
  border: 1px dashed var(--el-border-color);
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
  transition: var(--el-transition-duration-fast);
}
.avatar-uploader .el-upload:hover {
  border-color: var(--el-color-primary);
}
.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 178px;
  height: 178px;
  text-align: center;
  border: 1px dashed #d9d9d9;
  display: flex;
  align-items: center;
  justify-content: center;
}
</style>
