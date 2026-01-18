<template>
  <div class="login-container">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <span>管理员认证</span>
        </div>
      </template>
      <el-form :model="form" @submit.prevent="handleLogin">
        <el-form-item>
          <el-input 
            v-model="form.secretKey" 
            placeholder="请输入管理密钥" 
            type="password" 
            show-password
            @keyup.enter="handleLogin"
          />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="handleLogin" :loading="loading" style="width: 100%">
            验证进入
          </el-button>
        </el-form-item>
      </el-form>
      <div v-if="errorMsg" class="error-msg">{{ errorMsg }}</div>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { verifySecretKey } from '../../api/auth';
import { ElMessage } from 'element-plus';

const router = useRouter();
const route = useRoute();
const form = ref({
  secretKey: ''
});
const loading = ref(false);
const errorMsg = ref('');

const handleLogin = async () => {
  if (!form.value.secretKey) {
    errorMsg.value = '请输入密钥';
    return;
  }

  loading.value = true;
  errorMsg.value = '';

  try {
    const res = await verifySecretKey(form.value.secretKey);
    // res 已经是 data 部分了，包含 token 和 expires_at
    
    if (res && res.token) {
      localStorage.setItem('admin_token', res.token);
      localStorage.setItem('token_expires_at', res.expires_at);
      
      ElMessage.success('验证成功');
      
      const redirect = route.query.redirect || '/admin/dashboard';
      router.push(redirect);
    } else {
      errorMsg.value = '验证失败';
    }
  } catch (error) {
    // 错误已经在拦截器中处理了，或者在这里额外显示
    // 拦截器会 reject error
    // 这里只设置 errorMsg
    if (error.response && error.response.data && error.response.data.msg) {
         errorMsg.value = error.response.data.msg;
    } else {
         errorMsg.value = '登录失败';
    }
    console.error(error);
  } finally {
    loading.value = false;
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f0f2f5;
}

.box-card {
  width: 400px;
}

.card-header {
  text-align: center;
  font-weight: bold;
}

.error-msg {
  color: #f56c6c;
  font-size: 14px;
  text-align: center;
  margin-top: 10px;
}
</style>
