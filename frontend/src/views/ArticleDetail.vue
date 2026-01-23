<template>
  <el-row :gutter="20">
    <!-- Left Column: Content -->
    <el-col :span="18" :xs="24">
      <div v-if="article" class="main-column fade-in-up">
        
        <!-- Hero Header Card -->
        <GlassCard class="detail-card hero-header">
            <div class="hero-content">
                <div class="meta-tags">
                    <el-tag v-if="article.category" effect="dark" round class="category-pill">{{ article.category.name }}</el-tag>
                    <span class="date">{{ formatDate(article.created_at) }}</span>
                </div>
                <h1 class="article-title">{{ article.title }}</h1>
                <div class="meta-stats">
                    <span class="stat-item"><el-icon><View /></el-icon> {{ article.views }} 阅读</span>
                    <span class="stat-item"><el-icon><ChatLineRound /></el-icon> {{ comments.length }} 评论</span>
                </div>
            </div>
            <div class="hero-bg" v-if="article.cover_img" :style="{ backgroundImage: `url(${article.cover_img})` }"></div>
        </GlassCard>

        <!-- Markdown Content Card -->
        <GlassCard class="detail-card content-card">
            <v-md-editor :model-value="article.content" mode="preview"></v-md-editor>
        </GlassCard>

        <!-- Comments Section -->
        <GlassCard class="detail-card comments-card" id="comments">
            <div class="card-header-title">
                <el-icon><ChatDotRound /></el-icon> 评论 ({{ comments.length }})
            </div>
            
            <div class="comment-form-wrapper">
                <el-form :model="commentForm" ref="formRef" :rules="rules" label-position="top">
                    <el-row :gutter="20">
                        <el-col :span="24">
                            <el-form-item prop="nickname">
                                <el-input v-model="commentForm.nickname" placeholder="昵称 (选填，默认为随机昵称)" :prefix-icon="User" />
                            </el-form-item>
                        </el-col>
                    </el-row>
                    <el-form-item prop="content">
                        <el-input v-model="commentForm.content" type="textarea" :rows="2" placeholder="写下你的想法..." />
                    </el-form-item>
                    <div class="form-footer">
                        <el-button type="primary" round size="small" @click="submit(null)">发表评论</el-button>
                    </div>
                </el-form>
            </div>

            <div class="comment-list">
                 <div v-for="item in comments" :key="item.id" class="comment-item fade-in">
                    <div class="comment-avatar">{{ item.nickname.charAt(0).toUpperCase() }}</div>
                    <div class="comment-body">
                        <div class="comment-header">
                            <span class="username">{{ item.nickname }}</span>
                            <span class="time">{{ formatDate(item.created_at) }}</span>
                        </div>
                        <div class="comment-text">{{ item.content }}</div>
                        <div class="comment-actions">
                            <el-button link type="primary" size="small" @click="handleSimpleReply(item)">回复</el-button>
                        </div>

                        <!-- Replies -->
                        <div v-if="item.replies && item.replies.length" class="replies-list">
                            <div v-for="reply in item.replies" :key="reply.id" class="reply-item">
                                <div class="reply-header">
                                    <span class="username">{{ reply.nickname }}</span>
                                    <span class="reply-target">回复 @{{ item.nickname }}</span>
                                    <span class="time">{{ formatDate(reply.created_at) }}</span>
                                </div>
                                <div class="comment-text">{{ reply.content }}</div>
                            </div>
                        </div>

                        <!-- Inline Reply Form -->
                        <div v-if="replyId === item.id" class="inline-reply fade-in">
                            <el-form :model="replyForm" :rules="rules">
                               <el-row :gutter="10">
                                  <el-col :span="12"><el-input v-model="replyForm.nickname" placeholder="昵称 (选填)" size="small" :prefix-icon="User"/></el-col>
                                  <el-col :span="24" class="mt-2"><el-input v-model="replyForm.content" placeholder="内容" size="small" type="textarea" :rows="2"/></el-col>
                               </el-row>
                               <div class="mt-2 reply-actions">
                                   <el-button type="primary" size="small" round @click="submit(item.id)">提交</el-button>
                                   <el-button size="small" round @click="replyId = null">取消</el-button>
                               </div>
                            </el-form>
                        </div>
                    </div>
                 </div>
            </div>
        </GlassCard>
      </div>
    </el-col>

    <!-- Right Column: Sidebar -->
    <el-col :span="6" :xs="6" class="sidebar-col">
        <div class="sidebar-wrapper">
            <!-- Profile Card -->
            <ProfileCard 
                :articles="siteInfo.articleCount" 
                :categories="categories.length" 
                :tags="tags.length" 
            />

            <!-- Site Info Card -->
            <SiteInfoCard :info="siteInfo" />

            <!-- TOC Card (Manual Fixed) -->
            <div class="toc-placeholder" ref="tocWrapper" :style="{ height: isFixed ? tocHeight + 'px' : 'auto' }" v-if="toc.length > 0">
                <div :class="{ 'fixed-toc': isFixed }" :style="fixedStyle" ref="tocCardRef">
                    <GlassCard class="sidebar-card toc-card">
                        <div class="card-header">
                            <span><el-icon><List /></el-icon> 目录</span>
                        </div>
                        <div class="toc-list">
                            <div 
                                v-for="(item, index) in toc" 
                                :key="index" 
                                class="toc-item"
                                :class="{ 'active': activeToc === item.id, 'toc-h2': item.level === 2, 'toc-h3': item.level === 3 }"
                                @click="scrollToAnchor(item.id)"
                            >
                                {{ item.text }}
                            </div>
                        </div>
                    </GlassCard>
                </div>
            </div>
        </div>
    </el-col>
  </el-row>
</template>

<script setup>
import { ref, onMounted, reactive } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { getArticleDetail, getArticles } from '@/api/article';
import { getCategoriesStats } from '@/api/category';
import { getTagsStats } from '@/api/tag';
import { getComments, submitComment } from '@/api/comment';
import dayjs from 'dayjs';
import relativeTime from 'dayjs/plugin/relativeTime';
import { ElMessage } from 'element-plus';
import { 
    View, ChatLineRound, ChatDotRound, User, Link, List
} from '@element-plus/icons-vue';
import ProfileCard from '@/components/ProfileCard.vue';
import SiteInfoCard from '@/components/SiteInfoCard.vue';
import GlassCard from '@/components/GlassCard.vue';

dayjs.extend(relativeTime);

const route = useRoute();
const router = useRouter();
const article = ref(null);
const comments = ref([]);
const formRef = ref(null);
const toc = ref([]);
const activeToc = ref('');
const categories = ref([]);
const tags = ref([]);

// Manual Sticky Logic
const tocWrapper = ref(null);
const tocCardRef = ref(null);
const isFixed = ref(false);
const tocHeight = ref(0);
const fixedStyle = reactive({
    width: 'auto',
    top: '90px'
});

const handleScroll = () => {
    if (!tocWrapper.value) return;
    
    const rect = tocWrapper.value.getBoundingClientRect();
    const threshold = 90; // Top offset
    
    // When the top of the placeholder hits the threshold
    if (rect.top <= threshold) {
        if (!isFixed.value) {
            // Lock current width before fixing
            fixedStyle.width = `${tocWrapper.value.offsetWidth}px`;
            tocHeight.value = tocCardRef.value ? tocCardRef.value.offsetHeight : 0;
            isFixed.value = true;
        }
    } else {
        if (isFixed.value) {
            isFixed.value = false;
            fixedStyle.width = 'auto';
        }
    }
};

// Site Info Logic
const siteInfo = ref({
    articleCount: 0,
    runningTime: '',
    wordCount: '0',
    visitorCount: '12,345',
    totalViews: '0',
    lastUpdate: '-'
});

const commentForm = reactive({
    nickname: '',
    email: '',
    content: ''
});

const replyId = ref(null);
const replyForm = reactive({
    nickname: '',
    email: '',
    content: ''
});

const rules = {
    content: [{ required: true, message: '请输入内容', trigger: 'blur' }]
};

const formatDate = (date) => dayjs(date).format('YYYY年MM月DD日');

// TOC Logic
const generateTOC = (content) => {
    const lines = content.split('\n');
    const headers = [];
    const idMap = {};
    
    lines.forEach(line => {
        const match = line.match(/^(#{1,3})\s+(.+)$/);
        if (match) {
            const level = match[1].length;
            const text = match[2].trim();
            // Generate simple ID: "header-text"
            let id = text.toLowerCase().replace(/\s+/g, '-').replace(/[^\w\u4e00-\u9fa5-]/g, '');
            // Handle duplicate IDs
            if (idMap[id]) {
                idMap[id]++;
                id = `${id}-${idMap[id]}`;
            } else {
                idMap[id] = 1;
            }
            
            headers.push({ level, text, id });
        }
    });
    toc.value = headers;
};

const scrollToAnchor = (id) => {
    // Fallback: Find header by text
    const headers = document.querySelectorAll('.v-md-editor-preview h1, .v-md-editor-preview h2, .v-md-editor-preview h3');
    for (let h of headers) {
        if (h.innerText.trim() === toc.value.find(t => t.id === id)?.text) {
            h.scrollIntoView({ behavior: 'smooth', block: 'start' });
            activeToc.value = id;
            break;
        }
    }
};

const fetchData = async () => {
    const id = route.params.id;
    try {
        const res = await getArticleDetail(id);
        article.value = res;
        generateTOC(res.content);
        
        const commentRes = await getComments(id);
        comments.value = commentRes;
        
        // Fetch stats for ProfileCard
        const catRes = await getCategoriesStats();
        categories.value = catRes;
        const tagRes = await getTagsStats();
        tags.value = tagRes;

        // Fetch site info (reused from Home)
        fetchSiteInfo();
    } catch (e) {
        console.error(e);
    }
};

const fetchSiteInfo = async () => {
    try {
        const allRes = await getArticles({ page: 1, page_size: 1000, status: 1 });
        const list = allRes.list;
        
        siteInfo.value.articleCount = allRes.total;
        siteInfo.value.totalViews = list.reduce((acc, cur) => acc + (cur.views || 0), 0).toLocaleString();
        siteInfo.value.wordCount = (list.length * 1200).toLocaleString();
        
        const startDate = dayjs('2023-01-01');
        const now = dayjs();
        siteInfo.value.runningTime = now.diff(startDate, 'day') + ' 天';
        
        if (list.length > 0) {
            siteInfo.value.lastUpdate = dayjs(list[0].created_at).fromNow();
        }
    } catch(e) { console.error(e); }
};

const handleSimpleReply = (item) => {
    replyId.value = item.id;
    replyForm.nickname = commentForm.nickname;
    replyForm.email = '';
    replyForm.content = '';
};

const submit = async (parentId) => {
    const isReply = !!parentId;
    const form = isReply ? replyForm : commentForm;
    
    if (!form.content) {
        ElMessage.warning('请填写内容');
        return;
    }

    try {
        const payload = {
            article_id: parseInt(route.params.id),
            nickname: form.nickname,
            email: '', 
            content: form.content,
            parent_id: parentId ? parseInt(parentId) : undefined
        };
        
        await submitComment(payload);
        ElMessage.success('评论提交成功，等待审核');
        
        if (isReply) {
            replyId.value = null;
            replyForm.content = '';
        } else {
            commentForm.content = '';
        }
    } catch (e) {
        console.error(e);
    }
}

onMounted(() => {
    fetchData();
    window.addEventListener('scroll', handleScroll);
    window.addEventListener('resize', handleScroll);
});

import { onUnmounted } from 'vue';
onUnmounted(() => {
    window.removeEventListener('scroll', handleScroll);
    window.removeEventListener('resize', handleScroll);
});
</script>

<style scoped>
.main-column {
    padding-top: 20px;
}

/* Glassmorphism Cards */
.detail-card {
    min-height: auto; /* Allow auto height */
    overflow: visible; /* Fix overflow clipping if needed */
}

/* Hero Header */
.hero-header {
    position: relative;
    padding: 0;
    min-height: 200px;
    display: flex;
    align-items: center;
    justify-content: center;
    text-align: center;
    overflow: hidden;
    margin-bottom: 24px;
}

.hero-bg {
    position: absolute;
    inset: 0;
    background-size: cover;
    background-position: center;
    opacity: 0.15;
    z-index: 0;
    filter: blur(2px);
}
:deep(.hero-bg) { opacity: 0.3; } /* Remove global scope pollution */

.hero-content {
    position: relative;
    z-index: 1;
    padding: 40px;
    width: 100%;
}

.meta-tags {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 12px;
    margin-bottom: 10px;
    flex-wrap: wrap;
}

.detail-tags {
    display: flex;
    gap: 8px;
}

.detail-tag-pill {
    background: rgba(255,255,255,0.1);
    border: 1px solid rgba(255,255,255,0.2);
    color: var(--text-main);
    font-weight: 500;
}

.category-pill {
    margin-right: 12px;
    font-weight: bold;
    border: none;
}

.date {
    font-size: 0.9rem;
    color: var(--text-secondary);
    font-weight: 500;
}

.article-title {
    font-size: 2.5rem;
    margin: 15px 0;
    background: linear-gradient(120deg, var(--primary-color), var(--secondary-color));
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    font-weight: 800;
}

.meta-stats {
    display: flex;
    justify-content: center;
    gap: 20px;
    color: var(--text-secondary);
    font-size: 0.95rem;
}

.stat-item {
    display: flex;
    align-items: center;
    gap: 6px;
}

/* VMdEditor Override for Transparency - Increased Specificity */
.main-column :deep(.v-md-editor),
.main-column :deep(.v-md-editor-preview),
.main-column :deep(.github-markdown-body),
.main-column :deep(.v-md-editor__editor-wrapper) {
    background-color: transparent !important;
    background: transparent !important;
    border: none !important;
    box-shadow: none !important;
}

.main-column :deep(.v-md-editor-preview p),
.main-column :deep(.github-markdown-body p) {
    background-color: transparent !important;
}

/* Ensure glassmorphism applies nicely */
.content-card {
    padding: 20px;
    min-height: 400px;
    margin-bottom: 24px;
}

/* Comments */
.comments-card {
    padding: 30px;
    margin-bottom: 24px;
}

.card-header-title {
    font-size: 1.2rem;
    font-weight: 700;
    margin-bottom: 25px;
    display: flex;
    align-items: center;
    gap: 10px;
    color: var(--text-main);
}

.comment-form-wrapper {
    background: rgba(0,0,0,0.02);
    padding: 20px;
    border-radius: 16px;
    margin-bottom: 30px;
    border: 1px dashed var(--border-color);
}

.comment-item {
    display: flex;
    gap: 16px;
    margin-bottom: 24px;
}

.comment-avatar {
    width: 44px;
    height: 44px;
    background: linear-gradient(135deg, #e0c3fc 0%, #8ec5fc 100%);
    border-radius: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-weight: 800;
    color: white;
    font-size: 1.2rem;
    box-shadow: 0 4px 10px rgba(142, 197, 252, 0.4);
    flex-shrink: 0;
}

.comment-body {
    flex: 1;
}

/* Compact Comment List Styles */
.comment-list {
    margin-top: 24px;
}

.comment-item {
    display: flex;
    gap: 12px; /* Smaller gap */
    margin-bottom: 20px;
    padding-bottom: 20px;
    border-bottom: 1px solid rgba(255, 255, 255, 0.05);
}

.avatar-wrapper {
    flex-shrink: 0;
}

.avatar-wrapper :deep(.el-avatar) {
    width: 32px;
    height: 32px;
    font-size: 14px;
    border: 2px solid rgba(255, 255, 255, 0.1);
}

.comment-content {
    flex: 1;
}

.comment-header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    margin-bottom: 4px; /* Tighter spacing */
}

.user-info {
    display: flex;
    align-items: center;
    gap: 8px;
}

.nickname {
    font-weight: 600;
    color: var(--text-main);
    font-size: 0.9rem; /* Slightly smaller */
}

.comment-time {
    font-size: 0.75rem;
    color: var(--text-secondary);
}

.comment-text {
    font-size: 0.9rem;
    line-height: 1.5;
    color: var(--text-main);
    margin-bottom: 6px;
    padding: 8px 12px;
    background: rgba(255, 255, 255, 0.03);
    border-radius: 8px;
    display: inline-block;
}

.comment-actions {
    display: flex;
    gap: 12px;
}

.action-btn {
    padding: 0;
    height: auto;
    font-size: 0.75rem;
    color: var(--text-secondary);
}

.action-btn:hover {
    color: var(--primary-color);
    background: transparent;
}

/* Reply Input Compact */
.reply-input-wrapper {
    margin-top: 8px;
    display: flex;
    gap: 8px;
}

.reply-input-wrapper .el-input {
    flex: 1;
}

/* Nested Replies Compact */
.replies-list {
    margin-top: 12px;
    margin-left: 44px; /* Align with content */
    padding-left: 12px;
    border-left: 2px solid rgba(255, 255, 255, 0.05);
}

.reply-item {
    margin-top: 12px;
    padding-top: 12px;
    border-top: 1px solid rgba(255, 255, 255, 0.03);
}

.reply-item:first-child {
    margin-top: 0;
    padding-top: 0;
    border-top: none;
}

/* Sidebar & TOC */
.article-row {
    align-items: stretch;
    display: flex;
    flex-wrap: wrap;
    position: relative; /* Ensure row context */
}

/* Sidebar Column - Needs flex to stretch children */
.sidebar-col {
    display: flex !important;
    flex-direction: column;
    position: relative; /* Ensure column context */
}

/* Sidebar Wrapper - The main column container */
.sidebar-wrapper {
    /* flex: 1; Occupy full height */
    /* display: flex; */
    /* flex-direction: column; */
    position: relative; /* Ensure sticky context */
    min-height: 100%;
}

/* Regular sidebar content (Profile, Info) */
.sidebar-content {
    display: flex;
    flex-direction: column;
}

/* TOC Placeholder */
.toc-placeholder {
    margin-top: 40px;
    width: 100%;
}

.fixed-toc {
    position: fixed;
    z-index: 100;
}

.toc-card {
    max-height: calc(100vh - 120px);
    overflow: hidden;
    display: flex;
    flex-direction: column;
}

.toc-list {
    flex: 1;
    overflow-y: auto;
    padding-right: 5px;
    margin-top: 10px;
    /* Custom Scrollbar */
    scrollbar-width: thin;
    scrollbar-color: rgba(0,0,0,0.1) transparent;
}

.toc-item {
    padding: 8px 12px;
    font-size: 0.9rem;
    color: var(--text-secondary);
    cursor: pointer;
    transition: all 0.2s;
    border-radius: 6px;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    margin-bottom: 2px;
}

.toc-item:hover {
    color: var(--primary-color);
    background: rgba(var(--primary-color-rgb), 0.1);
}

.toc-item.active {
    color: var(--primary-color);
    background: rgba(var(--primary-color-rgb), 0.1);
    font-weight: 600;
}

.toc-h2 { 
    padding-left: 12px; 
}
.toc-h3 { 
    padding-left: 24px; 
    font-size: 0.85rem; 
}

/* Hide in-content TOC generated by plugins if any */
.v-md-editor-preview .toc,
.github-markdown-body .toc {
    display: none !important;
}

.sidebar-card {
    padding: 20px;
    margin-bottom: 24px;
}

.card-header {
    font-size: 1.1rem;
    font-weight: 700;
    color: var(--text-main);
    margin-bottom: 15px;
    padding-bottom: 10px;
    border-bottom: 1px solid rgba(255,255,255,0.1);
    display: flex;
    align-items: center;
    gap: 8px;
}

/* Animations */
.fade-in-up {
    animation: fadeInUp 0.8s ease-out forwards;
}
.fade-in {
    animation: fadeIn 0.5s ease-out forwards;
}

@keyframes fadeIn {
    from { opacity: 0; }
    to { opacity: 1; }
}
@keyframes fadeInUp {
    from { opacity: 0; transform: translateY(20px); }
    to { opacity: 1; transform: translateY(0); }
}
</style>