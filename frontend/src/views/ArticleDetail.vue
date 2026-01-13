<template>
  <div v-if="article" class="article-detail">
    <div class="article-header">
       <h1 class="title">{{ article.title }}</h1>
       <div class="meta">
          <span class="meta-item">{{ formatDate(article.created_at) }}</span>
          <span class="meta-item">|</span>
          <span class="meta-item">阅读 {{ article.views }}</span>
          <span class="meta-item">|</span>
          <el-tag v-if="article.category" size="small">{{ article.category.name }}</el-tag>
       </div>
    </div>
    
    <div class="article-cover-lg" v-if="article.cover_img">
        <img :src="article.cover_img" />
    </div>

    <!-- Markdown Content -->
    <v-md-editor :model-value="article.content" mode="preview"></v-md-editor>

    <!-- Comments -->
    <div class="comments-section">
        <h3>评论 ({{ comments.length }})</h3>
        
        <div class="comment-form-card">
            <h4>发表评论</h4>
            <el-form :model="commentForm" ref="formRef" :rules="rules" label-position="top">
                <el-row :gutter="20">
                    <el-col :span="8">
                        <el-form-item prop="nickname" label="昵称">
                            <el-input v-model="commentForm.nickname" placeholder="必填" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="8">
                        <el-form-item prop="email" label="邮箱">
                            <el-input v-model="commentForm.email" placeholder="必填 (保密)" />
                        </el-form-item>
                    </el-col>
                    <el-col :span="8">
                        <el-form-item prop="website" label="网站">
                            <el-input v-model="commentForm.website" placeholder="选填" />
                        </el-form-item>
                    </el-col>
                </el-row>
                <el-form-item prop="content" label="内容">
                    <el-input v-model="commentForm.content" type="textarea" :rows="3" placeholder="写下你的评论..." />
                </el-form-item>
                <el-button type="primary" @click="submit(null)">提交评论</el-button>
            </el-form>
        </div>

        <div class="comment-list">
             <div v-for="item in comments" :key="item.id" class="comment-item">
                <div class="comment-avatar">{{ item.nickname.charAt(0).toUpperCase() }}</div>
                <div class="comment-body">
                    <div class="comment-info">
                        <span class="username">{{ item.nickname }}</span>
                        <span class="time">{{ formatDate(item.created_at) }}</span>
                    </div>
                    <div class="comment-text">{{ item.content }}</div>
                    <div class="comment-actions">
                        <el-button link type="primary" size="small" @click="handleSimpleReply(item)">回复</el-button>
                    </div>

                    <!-- Replies -->
                    <div v-if="item.replies && item.replies.length" class="replies">
                        <div v-for="reply in item.replies" :key="reply.id" class="reply-item">
                            <div class="comment-info">
                                <span class="username">{{ reply.nickname }}</span>
                                <span class="reply-target">回复 @{{ item.nickname }}</span>
                                <span class="time">{{ formatDate(reply.created_at) }}</span>
                            </div>
                            <div class="comment-text">{{ reply.content }}</div>
                        </div>
                    </div>

                    <!-- Inline Reply Form -->
                    <div v-if="replyId === item.id" class="inline-reply">
                        <el-form :model="replyForm" :rules="rules">
                           <el-row :gutter="10">
                              <el-col :span="8"><el-input v-model="replyForm.nickname" placeholder="昵称" size="small"/></el-col>
                              <el-col :span="8"><el-input v-model="replyForm.email" placeholder="邮箱" size="small"/></el-col>
                              <el-col :span="24" class="mt-2"><el-input v-model="replyForm.content" placeholder="内容" size="small" type="textarea" :rows="2"/></el-col>
                           </el-row>
                           <div class="mt-2">
                               <el-button type="primary" size="small" @click="submit(item.id)">提交回复</el-button>
                               <el-button size="small" @click="replyId = null">取消</el-button>
                           </div>
                        </el-form>
                    </div>
                </div>
             </div>
        </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue';
import { useRoute } from 'vue-router';
import { getArticleDetail } from '@/api/article';
import { getComments, submitComment } from '@/api/comment';
import dayjs from 'dayjs';
import { ElMessage } from 'element-plus';

const route = useRoute();
const article = ref(null);
const comments = ref([]);
const formRef = ref(null);

const commentForm = reactive({
    nickname: '',
    email: '',
    website: '',
    content: ''
});

const replyId = ref(null);
const replyForm = reactive({
    nickname: '',
    email: '',
    content: ''
});

const rules = {
    nickname: [{ required: true, message: '请输入昵称', trigger: 'blur' }],
    email: [{ required: true, message: '请输入邮箱', trigger: 'blur' }, { type: 'email', message: '邮箱格式不正确', trigger: 'blur' }],
    content: [{ required: true, message: '请输入内容', trigger: 'blur' }]
};

const formatDate = (date) => dayjs(date).format('YYYY-MM-DD HH:mm');

const fetchData = async () => {
    const id = route.params.id;
    try {
        const res = await getArticleDetail(id);
        article.value = res;
        const commentRes = await getComments(id);
        comments.value = commentRes;
    } catch (e) {
        console.error(e);
    }
};

const handleSimpleReply = (item) => {
    replyId.value = item.id;
    // Copy current user info to reply form if available
    replyForm.nickname = commentForm.nickname;
    replyForm.email = commentForm.email;
    replyForm.content = '';
};

const submit = async (parentId) => {
    const isReply = !!parentId;
    const form = isReply ? replyForm : commentForm;
    
    // Simple validation check
    if (!form.nickname || !form.email || !form.content) {
        ElMessage.warning('请填写完整信息');
        return;
    }

    try {
        const payload = {
            article_id: parseInt(route.params.id),
            nickname: form.nickname,
            email: form.email,
            content: form.content,
            parent_id: parentId ? parseInt(parentId) : undefined,
            website: isReply ? undefined : form.website
        };
        
        await submitComment(payload);
        ElMessage.success('评论提交成功，等待审核');
        
        if (isReply) {
            replyId.value = null;
            replyForm.content = '';
        } else {
            commentForm.content = '';
            formRef.value.resetFields();
        }
    } catch (e) {
        console.error(e);
    }
}

onMounted(() => {
    fetchData();
});
</script>

<style scoped>
.article-detail {
    background: white;
    padding: 30px;
    border-radius: 8px;
}
.article-header {
    text-align: center;
    margin-bottom: 30px;
}
.title {
    font-size: 32px;
    margin-bottom: 20px;
}
.meta {
    color: #999;
    font-size: 14px;
    display: flex;
    justify-content: center;
    gap: 15px;
    align-items: center;
}
.article-cover-lg {
    margin-bottom: 30px;
}
.article-cover-lg img {
    width: 100%;
    max-height: 400px;
    object-fit: cover;
    border-radius: 8px;
}
.comments-section {
    margin-top: 60px;
    border-top: 1px solid #eee;
    padding-top: 30px;
}
.comment-form-card {
    background: #f9f9f9;
    padding: 20px;
    border-radius: 8px;
    margin-bottom: 30px;
}
.comment-item {
    display: flex;
    margin-bottom: 20px;
    gap: 15px;
}
.comment-avatar {
    width: 40px;
    height: 40px;
    background: #e0e0e0;
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: bold;
    color: #666;
    flex-shrink: 0;
}
.comment-body {
    flex: 1;
}
.comment-info {
    margin-bottom: 5px;
}
.username {
    font-weight: bold;
    margin-right: 10px;
}
.time {
    color: #999;
    font-size: 12px;
}
.comment-text {
    line-height: 1.5;
    color: #333;
    margin-bottom: 8px;
}
.replies {
    background: #f9f9f9;
    padding: 10px;
    border-radius: 4px;
    margin-top: 10px;
}
.reply-item {
    margin-bottom: 10px;
    border-bottom: 1px dashed #eee;
    padding-bottom: 10px;
}
.reply-item:last-child {
    border-bottom: none;
    margin-bottom: 0;
    padding-bottom: 0;
}
.reply-target {
    color: #409eff;
    margin: 0 10px;
    font-size: 12px;
}
.inline-reply {
    margin-top: 10px;
    background: #f0f0f0;
    padding: 10px;
    border-radius: 4px;
}
.mt-2 { margin-top: 10px; }
.mb-2 { margin-bottom: 10px; }
</style>
