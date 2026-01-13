import request from '@/utils/request';
import adminRequest from '@/utils/admin-request';

// 公开
export function getTags() {
  return request.get('/api/tags');
}

export function getTagsStats() {
  return request.get('/api/tags/stats');
}

// 管理
export function createTag(data) {
  return adminRequest.post('/api/admin/tags', data);
}

export function updateTag(id, data) {
  return adminRequest.put(`/api/admin/tags/${id}`, data);
}

export function deleteTag(id) {
  return adminRequest.delete(`/api/admin/tags/${id}`);
}
