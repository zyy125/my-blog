<template>
  <el-container class="layout-container">
    <!-- Sidebar -->
    <el-aside :width="isCollapse ? 'var(--sidebar-width-collapsed)' : 'var(--sidebar-width)'" class="layout-sidebar">
      <div class="logo-container">
        <transition name="fade" mode="out-in">
          <span v-if="!isCollapse" class="logo-text">Blog Admin</span>
          <span v-else class="logo-text-mini">B</span>
        </transition>
      </div>
      
      <el-menu
        :default-active="activeMenu"
        class="sidebar-menu"
        :collapse="isCollapse"
        router
        :collapse-transition="false"
      >
        <el-menu-item index="/admin/dashboard">
          <el-icon><Odometer /></el-icon>
          <template #title>仪表盘</template>
        </el-menu-item>
        
        <el-menu-item index="/admin/articles">
          <el-icon><Document /></el-icon>
          <template #title>文章管理</template>
        </el-menu-item>
        
        <el-menu-item index="/admin/categories">
          <el-icon><Folder /></el-icon>
          <template #title>分类管理</template>
        </el-menu-item>
        
        <el-menu-item index="/admin/tags">
          <el-icon><Collection /></el-icon>
          <template #title>标签管理</template>
        </el-menu-item>
        
        <el-menu-item index="/admin/comments">
          <el-icon><ChatDotSquare /></el-icon>
          <template #title>评论管理</template>
        </el-menu-item>
        
        <el-menu-item index="/" @click="$router.push('/')">
          <el-icon><HomeFilled /></el-icon>
          <template #title>返回前台</template>
        </el-menu-item>
      </el-menu>
    </el-aside>


    <!-- Main Content -->
    <el-container class="is-vertical">
      <el-header class="layout-header">
        <div class="header-left">
          <el-button 
            type="text" 
            class="collapse-btn" 
            @click="toggleCollapse"
          >
            <el-icon :size="20">
              <Expand v-if="isCollapse" />
              <Fold v-else />
            </el-icon>
          </el-button>
          
          <el-breadcrumb separator="/">
            <el-breadcrumb-item :to="{ path: '/admin/dashboard' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>{{ currentRouteName }}</el-breadcrumb-item>
          </el-breadcrumb>
        </div>

        <div class="header-right">
            <el-tooltip content="设置前台背景" placement="bottom">
                <el-button circle text @click="bgDialogVisible = true">
                    <el-icon :size="18"><Picture /></el-icon>
                </el-button>
            </el-tooltip>
            
            <el-tooltip :content="isDark ? '切换亮色主题' : '切换深色主题'" placement="bottom">
                <el-button circle text @click="toggleTheme">
                    <el-icon :size="18">
                        <Moon v-if="isDark" />
                        <Sunny v-else />
                    </el-icon>
                </el-button>
            </el-tooltip>

            <el-dropdown trigger="click" @command="handleCommand">
                <div class="user-profile">
                    <el-avatar :size="32" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
                    <span class="username">Admin</span>
                    <el-icon><CaretBottom /></el-icon>
                </div>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item command="logout">退出登录</el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
        </div>
      </el-header>

      <el-main class="layout-main">
        <router-view v-slot="{ Component }">
          <transition name="fade-transform" mode="out-in">
            <component :is="Component" />
          </transition>
        </router-view>
      </el-main>
    </el-container>

    <!-- BG Setting Dialog -->
    <el-dialog v-model="bgDialogVisible" title="前台背景设置" width="400px" align-center>
        <el-form label-position="top">
            <el-form-item label="背景图片文件名 (public/images/)">
                <el-input v-model="bgImageName" placeholder="例如: starry-bg.jpg" />
            </el-form-item>
        </el-form>
        <template #footer>
            <div class="dialog-footer">
                <el-button @click="bgDialogVisible = false">取消</el-button>
                <el-button type="primary" @click="saveBgSetting">保存</el-button>
            </div>
        </template>
    </el-dialog>
  </el-container>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { 
  Odometer, Document, Folder, Collection, ChatDotSquare, 
  HomeFilled, Expand, Fold, Moon, Sunny, Picture, CaretBottom 
} from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';

const route = useRoute();
const router = useRouter();

const isCollapse = ref(false);
const isDark = ref(false);
const bgDialogVisible = ref(false);
const bgImageName = ref('starry-bg.jpg');

const activeMenu = computed(() => route.path);

const currentRouteName = computed(() => {
    const nameMap = {
        'dashboard': '仪表盘',
        'article-list': '文章管理',
        'article-create': '写文章',
        'article-edit': '编辑文章',
        'category-list': '分类管理',
        'tag-list': '标签管理',
        'comment-list': '评论管理'
    };
    // Try to match from route name or path
    for (const key in nameMap) {
        if (route.path.includes(key.replace('-list', 's'))) return nameMap[key];
    }
    return '控制台';
});

const toggleCollapse = () => {
    isCollapse.value = !isCollapse.value;
};

const toggleTheme = () => {
    isDark.value = !isDark.value;
    const html = document.documentElement;
    if (isDark.value) {
        html.classList.add('dark');
        localStorage.setItem('theme', 'dark');
    } else {
        html.classList.remove('dark');
        localStorage.setItem('theme', 'light');
    }
};

const handleCommand = (cmd) => {
    if (cmd === 'logout') {
        localStorage.removeItem('admin_token');
        router.push('/admin/login');
        ElMessage.success('已退出登录');
    }
};

const saveBgSetting = () => {
    if (bgImageName.value) {
        localStorage.setItem('bg_image', bgImageName.value);
        window.dispatchEvent(new Event('bg-image-changed'));
        ElMessage.success('背景设置已更新');
        bgDialogVisible.value = false;
    }
};

onMounted(() => {
    const theme = localStorage.getItem('theme');
    if (theme === 'dark') {
        isDark.value = true;
        document.documentElement.classList.add('dark');
    }
    
    const storedBg = localStorage.getItem('bg_image');
    if (storedBg) bgImageName.value = storedBg;
});
</script>

<style scoped>
.layout-container {
    height: 100vh;
    background-color: var(--bg-body);
}

.layout-sidebar {
    background-color: var(--bg-sidebar);
    border-right: 1px solid var(--border-light);
    transition: width 0.3s cubic-bezier(0.2, 0, 0, 1);
    display: flex;
    flex-direction: column;
    overflow: hidden;
}

.logo-container {
    height: var(--header-height);
    display: flex;
    align-items: center;
    justify-content: center;
    border-bottom: 1px solid var(--border-light);
}

.logo-text {
    font-size: 18px;
    font-weight: 700;
    color: var(--color-primary);
    letter-spacing: 1px;
}

.logo-text-mini {
    font-size: 20px;
    font-weight: 800;
    color: var(--color-primary);
}

.sidebar-menu {
    border-right: none;
    background: transparent;
    flex: 1;
}

.sidebar-footer {
    border-top: 1px solid var(--border-light);
    padding: 10px 0;
}

.layout-header {
    height: var(--header-height);
    background-color: var(--bg-header);
    border-bottom: 1px solid var(--border-light);
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 20px;
}

.header-left {
    display: flex;
    align-items: center;
    gap: 16px;
}

.collapse-btn {
    color: var(--text-regular);
}
.collapse-btn:hover {
    color: var(--color-primary);
}

.header-right {
    display: flex;
    align-items: center;
    gap: 16px;
}

.user-profile {
    display: flex;
    align-items: center;
    gap: 8px;
    cursor: pointer;
    color: var(--text-regular);
    padding: 4px 8px;
    border-radius: 4px;
    transition: background 0.2s;
}

.user-profile:hover {
    background: var(--bg-hover);
}

.username {
    font-size: 14px;
    font-weight: 500;
}

.layout-main {
    padding: 20px;
    overflow-y: auto;
}

/* Transitions */
.fade-transform-enter-active,
.fade-transform-leave-active {
  transition: all 0.3s;
}

.fade-transform-enter-from {
  opacity: 0;
  transform: translateX(-10px);
}

.fade-transform-leave-to {
  opacity: 0;
  transform: translateX(10px);
}
</style>