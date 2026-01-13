import adminRequest from '@/utils/admin-request';

export function uploadImage(formData) {
  return adminRequest.post('/api/admin/upload/image', formData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  });
}
