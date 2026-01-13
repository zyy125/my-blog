import request from '@/utils/request';
import adminRequest from '@/utils/admin-request';

// 公开
export function getCategories() {
  return request.get('/api/categories');
}

export function getCategoriesStats() {
  return request.get('/api/categories/stats');
}

// 管理
export function createCategory(data) {
  return adminRequest.post('/api/admin/categories', data);
}

export function updateCategory(id, data) {
  return adminRequest.put(`/api/admin/categories/${id}`, data);
}

export function deleteCategory(id) {
  return adminRequest.delete(`/api/admin/categories/${id}`);
}
