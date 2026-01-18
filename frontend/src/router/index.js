import { createRouter, createWebHistory } from 'vue-router'
import UserLayout from '@/layouts/UserLayout.vue'
import AdminLayout from '@/layouts/AdminLayout.vue'

const routes = [
  // User Routes
  {
    path: '/',
    component: UserLayout,
    children: [
      { path: '', component: () => import('@/views/Home.vue') },
      { path: 'article/:id', component: () => import('@/views/ArticleDetail.vue') },
      { path: 'category/:id', component: () => import('@/views/Category.vue') },
      { path: 'tag/:id', component: () => import('@/views/Tag.vue') },
      { path: 'search', component: () => import('@/views/Search.vue') },
      { path: 'archive', component: () => import('@/views/Archive.vue') },
      { path: 'about', component: () => import('@/views/About.vue') }
    ]
  },
  // Admin Login
  {
    path: '/admin/login',
    component: () => import('@/views/admin/Login.vue')
  },
  // Admin Routes
  {
    path: '/admin',
    component: AdminLayout,
    meta: { requiresAuth: true }, // 需要认证
    children: [
      { path: '', redirect: '/admin/dashboard' },
      { path: 'dashboard', component: () => import('@/views/admin/Dashboard.vue') },
      { path: 'articles', component: () => import('@/views/admin/ArticleList.vue') },
      { path: 'articles/create', component: () => import('@/views/admin/ArticleEdit.vue') },
      { path: 'articles/edit/:id', component: () => import('@/views/admin/ArticleEdit.vue') },
      { path: 'categories', component: () => import('@/views/admin/CategoryList.vue') },
      { path: 'tags', component: () => import('@/views/admin/TagList.vue') },
      { path: 'comments', component: () => import('@/views/admin/CommentList.vue') }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth)) {
    const token = localStorage.getItem('admin_token');
    const expiresAt = localStorage.getItem('token_expires_at');
    
    if (!token) {
      next({
        path: '/admin/login',
        query: { redirect: to.fullPath }
      });
    } else if (expiresAt && parseInt(expiresAt) * 1000 < Date.now()) {
      // Token 过期
      localStorage.removeItem('admin_token');
      localStorage.removeItem('token_expires_at');
      next({
        path: '/admin/login',
        query: { redirect: to.fullPath }
      });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router
