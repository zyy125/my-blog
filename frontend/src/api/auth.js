import axios from 'axios';
import { ElMessage } from 'element-plus';

const authRequest = axios.create({
  baseURL: '',
  timeout: 10000
});

// 响应拦截器
authRequest.interceptors.response.use(
  response => {
    const { code, msg, data } = response.data;
    if (code !== 200) {
      ElMessage.error(msg || 'Error');
      return Promise.reject(new Error(msg || 'Error'));
    }
    return data;
  },
  error => {
    console.error('Auth Request Error:', error);
    ElMessage.error(error.message || 'Request failed');
    return Promise.reject(error);
  }
);

export const verifySecretKey = (secretKey) => {
  return authRequest.post('/api/admin/auth/verify', {
    secret_key: secretKey
  });
};
