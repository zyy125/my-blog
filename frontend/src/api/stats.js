import adminRequest from '@/utils/admin-request';

export function getStats() {
  return adminRequest.get('/api/admin/stats');
}
