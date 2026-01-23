<template>
  <!-- Global Background Effect -->
  <BackgroundEffect />

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
            
            <!-- Article Dropdown Menu -->
            <el-sub-menu index="articles" popper-class="custom-submenu-popper">
              <template #title>文章</template>
              <el-menu-item index="/category">分类</el-menu-item>
              <el-menu-item index="/tag">标签</el-menu-item>
              <el-menu-item index="/archive">归档</el-menu-item>
            </el-sub-menu>

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
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { Search } from '@element-plus/icons-vue';
import BackgroundEffect from '@/components/BackgroundEffect.vue';

const router = useRouter();
const keyword = ref('');

const handleSearch = () => {
  if (keyword.value.trim()) {
    router.push({ path: '/search', query: { keyword: keyword.value } });
  }
};

onMounted(() => {
    // Force Dark Mode Class
    document.documentElement.classList.add('dark');
});
</script>

<style scoped>
.layout-container {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
    /* Ensure container is transparent so background shows through */
    background: transparent;
    position: relative;
    z-index: 1; /* Ensure content is above the background */
}

.header {
  position: sticky;
  top: 0;
  z-index: 1000;
  border-bottom: 1px solid var(--border-color);
  background: var(--card-bg); /* Glassmorphism */
  backdrop-filter: blur(12px);
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
  /* Use lighter gradient for dark theme */
  background: linear-gradient(45deg, #a1c4fd, #c2e9fb); 
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  letter-spacing: -0.5px;
  text-shadow: 0 0 10px rgba(161, 196, 253, 0.3);
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

/* Ensure menu items are white */
:deep(.el-menu-item), :deep(.el-sub-menu__title) {
    color: var(--text-secondary) !important;
}
:deep(.el-menu-item.is-active), :deep(.el-sub-menu.is-active .el-sub-menu__title) {
    color: var(--primary-color) !important;
    border-bottom-color: var(--primary-color) !important;
}
:deep(.el-menu-item:hover), :deep(.el-sub-menu__title:hover) {
    color: var(--text-main) !important;
    background-color: rgba(255,255,255,0.05) !important;
}

/* Custom Submenu Styles (since it's a popper, global styles might be needed in style.css, but we try deep first) */
:global(.custom-submenu-popper) {
    background: var(--card-bg) !important;
    backdrop-filter: blur(12px) !important;
    border: 1px solid var(--border-color) !important;
}
:global(.custom-submenu-popper .el-menu) {
    background: transparent !important;
}
:global(.custom-submenu-popper .el-menu-item) {
    color: var(--text-secondary) !important;
    background: transparent !important;
}
:global(.custom-submenu-popper .el-menu-item:hover) {
    color: var(--primary-color) !important;
    background-color: rgba(255,255,255,0.05) !important;
}
:global(.custom-submenu-popper .el-menu-item.is-active) {
    color: var(--primary-color) !important;
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
  backdrop-filter: blur(12px);
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
