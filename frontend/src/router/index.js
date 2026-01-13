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
  // Admin Routes
  {
    path: '/admin',
    component: AdminLayout,
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

export default router
