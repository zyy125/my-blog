import axios from 'axios';
import { ElMessage } from 'element-plus';

const ADMIN_TOKEN = 'sk_blog_admin_please_change_this_token';

const adminRequest = axios.create({
  baseURL: '',
  timeout: 10000,
  headers: {
    'X-Admin-Token': ADMIN_TOKEN
  }
});

adminRequest.interceptors.response.use(
  response => {
    const { code, msg, data } = response.data;
    if (code !== 200) {
      ElMessage.error(msg || 'Error');
      return Promise.reject(new Error(msg || 'Error'));
    }
    return data;
  },
  error => {
    console.error('Admin Request Error:', error);
    ElMessage.error(error.message || 'Request failed');
    return Promise.reject(error);
  }
);

export default adminRequest;
