import axios from 'axios';
import { ElMessage } from 'element-plus';
import router from '../router';

const adminRequest = axios.create({
  baseURL: '',
  timeout: 10000
});

// 请求拦截器
adminRequest.interceptors.request.use(
  config => {
    // 从 localStorage 获取动态 token
    const token = localStorage.getItem('admin_token');
    
    // 检查 token 是否存在
    if (token) {
      config.headers['X-Admin-Token'] = token;
      
      // 检查是否过期 (可选，前端预检查)
      const expiresAt = localStorage.getItem('token_expires_at');
      if (expiresAt && parseInt(expiresAt) * 1000 < Date.now()) {
        // Token 过期，清除并跳转登录
        localStorage.removeItem('admin_token');
        localStorage.removeItem('token_expires_at');
        // 取消请求
         const controller = new AbortController();
         config.signal = controller.signal;
         controller.abort();
         
         ElMessage.warning('登录已过期，请重新验证');
         router.push('/admin/login');
         return config; // 返回config，雖然已被abort
      }
    } else {
      // 无 token，跳转登录
       // 这里可以做一些判断，如果是本来就是去login页面的请求就不管了，但这里是adminRequest，一般都是受保护的资源
       // 不过 router.push 更好在业务层做，或者在这里直接跳转
    }
    
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// 响应拦截器
adminRequest.interceptors.response.use(
  response => {
    const { code, msg, data } = response.data;
    if (code !== 200) {
      // 特定错误码处理，例如 token 失效
      if (code === 401 || msg === 'Token无效' || msg === 'Token无效或已过期') {
        localStorage.removeItem('admin_token');
        localStorage.removeItem('token_expires_at');
        router.push('/admin/login');
      }
      
      ElMessage.error(msg || 'Error');
      return Promise.reject(new Error(msg || 'Error'));
    }
    return data;
  },
  error => {
    console.error('Admin Request Error:', error);
    
    if (error.response) {
       if (error.response.status === 401 || (error.response.data && (error.response.data.msg === 'Token无效' || error.response.data.msg === 'Token无效或已过期'))) {
           localStorage.removeItem('admin_token');
           localStorage.removeItem('token_expires_at');
           router.push('/admin/login');
       }
        ElMessage.error(error.response.data.msg || 'Request failed');
    } else if (error.message !== 'canceled') { // 忽略取消的请求
        ElMessage.error(error.message || 'Request failed');
    }
    
    return Promise.reject(error);
  }
);

export default adminRequest;
