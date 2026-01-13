<template>
  <el-container>
    <el-header class="header">
      <div class="logo">My Blog</div>
      <div class="nav">
        <el-menu mode="horizontal" router :default-active="$route.path" :ellipsis="false">
          <el-menu-item index="/">首页</el-menu-item>
          <el-menu-item index="/archive">归档</el-menu-item>
          <el-menu-item index="/about">关于</el-menu-item>
        </el-menu>
        <div class="search-box">
             <el-input
                v-model="keyword"
                placeholder="搜索..."
                @keyup.enter="handleSearch"
                :prefix-icon="Search"
             />
        </div>
      </div>
    </el-header>
    <el-main class="main-content">
      <router-view />
    </el-main>
    <el-footer class="footer">
      © 2026 My Blog.
    </el-footer>
  </el-container>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { Search } from '@element-plus/icons-vue';

const router = useRouter();
const keyword = ref('');

const handleSearch = () => {
  if (keyword.value.trim()) {
    router.push({ path: '/search', query: { keyword: keyword.value } });
  }
};
</script>

<style scoped>
.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-bottom: 1px solid #dcdfe6;
  background: white;
  padding: 0 40px;
}
.logo {
  font-size: 24px;
  font-weight: bold;
  color: #333;
}
.nav {
    display: flex;
    align-items: center;
    gap: 20px;
}
.main-content {
  min-height: 80vh;
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;
  padding: 20px;
}
.footer {
  text-align: center;
  color: #909399;
  padding: 20px;
}
</style>
