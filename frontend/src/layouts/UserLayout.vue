<template>
  <el-container class="layout-container">
    <el-header class="header">
      <div class="header-inner">
        <div class="logo">
           <span class="logo-text">My Blog</span>
        </div>
        <div class="nav-section">
          <el-menu 
            mode="horizontal" 
            router 
            :default-active="$route.path" 
            :ellipsis="false"
            class="main-menu"
          >
            <el-menu-item index="/">首页</el-menu-item>
            <el-menu-item index="/archive">归档</el-menu-item>
            <el-menu-item index="/about">关于</el-menu-item>
          </el-menu>
          
          <div class="actions">
            <div class="search-box">
                <el-input
                    v-model="keyword"
                    placeholder="搜索..."
                    @keyup.enter="handleSearch"
                    :prefix-icon="Search"
                    class="search-input"
                />
            </div>
            <el-button 
                circle 
                class="theme-toggle" 
                @click="toggleTheme"
            >
                <el-icon v-if="isDark"><Sunny /></el-icon>
                <el-icon v-else><Moon /></el-icon>
            </el-button>
          </div>
        </div>
      </div>
    </el-header>
    
    <el-main class="main-content">
      <router-view v-slot="{ Component }">
          <transition name="fade-transform" mode="out-in">
             <component :is="Component" />
          </transition>
      </router-view>
    </el-main>
    
    <el-footer class="footer">
      <div class="footer-content">
        <p>© 2026 My Blog. Built with Vue 3 & Go.</p>
      </div>
    </el-footer>
  </el-container>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useRouter } from 'vue-router';
import { Search, Moon, Sunny } from '@element-plus/icons-vue';

const router = useRouter();
const keyword = ref('');
const isDark = ref(false);

const handleSearch = () => {
  if (keyword.value.trim()) {
    router.push({ path: '/search', query: { keyword: keyword.value } });
  }
};

const toggleTheme = () => {
    isDark.value = !isDark.value;
    updateThemeClass();
};

const updateThemeClass = () => {
    const html = document.documentElement;
    if (isDark.value) {
        html.classList.add('dark');
        localStorage.setItem('theme', 'dark');
    } else {
        html.classList.remove('dark');
        localStorage.setItem('theme', 'light');
    }
};

onMounted(() => {
    // Initialize Theme
    const savedTheme = localStorage.getItem('theme');
    const systemDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
    
    if (savedTheme === 'dark' || (!savedTheme && systemDark)) {
        isDark.value = true;
    } else {
        isDark.value = false;
    }
    updateThemeClass();
});
</script>

<style scoped>
.layout-container {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
}

.header {
  position: sticky;
  top: 0;
  z-index: 1000;
  border-bottom: 1px solid var(--border-color);
  background: var(--card-bg);
  padding: 0;
  height: 60px;
  box-shadow: var(--shadow-sm);
  transition: background-color var(--transition-normal), border-color var(--transition-normal);
}

.header-inner {
    max-width: 1200px;
    margin: 0 auto;
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 20px;
}

.logo-text {
  font-size: 24px;
  font-weight: 700;
  background: linear-gradient(45deg, var(--primary-color), var(--secondary-color));
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  letter-spacing: -0.5px;
}

.nav-section {
    display: flex;
    align-items: center;
    gap: 30px;
}

.main-menu {
    border-bottom: none !important;
    background: transparent !important;
}

.actions {
    display: flex;
    align-items: center;
    gap: 15px;
}

.search-input {
    width: 200px;
    transition: width 0.3s;
}
.search-input:focus-within {
    width: 250px;
}

.main-content {
  flex: 1;
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 40px 20px;
}

.footer {
  text-align: center;
  color: var(--text-secondary);
  padding: 40px 0;
  background: var(--card-bg);
  border-top: 1px solid var(--border-color);
  margin-top: auto;
}

/* Page Transitions */
.fade-transform-enter-active,
.fade-transform-leave-active {
  transition: opacity 0.4s ease, transform 0.4s ease;
}

.fade-transform-enter-from {
  opacity: 0;
  transform: translateX(-20px);
}

.fade-transform-leave-to {
  opacity: 0;
  transform: translateX(20px);
}

@media (max-width: 768px) {
    .search-input {
        width: 140px;
    }
    .search-input:focus-within {
        width: 140px;
    }
}
</style>
