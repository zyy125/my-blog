import request from '@/utils/request';
import adminRequest from '@/utils/admin-request';

// 公开接口
export function getArticles(params) {
  return request.get('/api/articles', { params });
}

export function getArticleDetail(id) {
  return request.get(`/api/articles/${id}`);
}

// 管理接口
export function createArticle(data) {
  return adminRequest.post('/api/admin/articles', data);
}

export function updateArticle(id, data) {
  return adminRequest.put(`/api/admin/articles/${id}`, data);
}

export function deleteArticle(id) {
  return adminRequest.delete(`/api/admin/articles/${id}`);
}
