import request from '@/utils/request';
import adminRequest from '@/utils/admin-request';

// 公开
export function getComments(articleId) {
  return request.get(`/api/articles/${articleId}/comments`);
}

export function submitComment(data) {
  return request.post('/api/comments', data);
}

// 管理
export function getAdminComments(params) {
  return adminRequest.get('/api/admin/comments', { params });
}

export function approveComment(id) {
  return adminRequest.patch(`/api/admin/comments/${id}/approve`);
}

export function rejectComment(id) {
  return adminRequest.patch(`/api/admin/comments/${id}/reject`);
}

export function deleteComment(id) {
  return adminRequest.delete(`/api/admin/comments/${id}`);
}
