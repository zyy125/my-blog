import axios from 'axios';
import { ElMessage } from 'element-plus';

const request = axios.create({
  baseURL: '', 
  timeout: 10000
});

// 响应拦截器：统一处理错误
request.interceptors.response.use(
  response => {
    const { code, msg, data } = response.data;
    if (code !== 200) {
      ElMessage.error(msg || 'Error');
      return Promise.reject(new Error(msg || 'Error'));
    }
    return data;  // 直接返回 data 部分
  },
  error => {
    console.error('Request Error:', error);
    ElMessage.error(error.message || 'Request failed');
    return Promise.reject(error);
  }
);

export default request;
