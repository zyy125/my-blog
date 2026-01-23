<template>
  <GlassCard class="profile-card" hover-effect @click="goAbout">
    <div class="profile-header">
      <div class="avatar-wrapper">
        <img src="/vite.png" alt="Avatar" class="avatar-image">
      </div>
      <h3 class="profile-name">逐影</h3>

      <!-- Stats Row -->
      <div class="stats-row">
        <div class="stat-item" @click.stop="router.push('/archive')">
          <div class="stat-value">{{ articles }}</div>
          <div class="stat-label">文章</div>
        </div>
        <div class="stat-item" @click.stop="router.push('/category')">
          <div class="stat-value">{{ categories }}</div>
          <div class="stat-label">分类</div>
        </div>
        <div class="stat-item" @click.stop="router.push('/tag')">
          <div class="stat-value">{{ tags }}</div>
          <div class="stat-label">标签</div>
        </div>
      </div>
      
      <div class="social-links" @click.stop>
        <!-- GitHub -->
        <a href="https://github.com/zyy125" target="_blank" class="social-icon github">
            <svg viewBox="0 0 1024 1024" width="1.2em" height="1.2em" fill="currentColor">
                <path d="M511.6 76.3C264.3 76.2 64 276.4 64 523.5 64 718.9 189.3 885 363.8 946c23.5 5.9 19.9-10.8 19.9-22.2v-77.5c-135.7 15.9-141.2-73.9-150.3-88.9C215 726 171.5 718 184.5 703c16.9-1.5 61.4 26 81 77.5 30.1 62.1 89.4 51.5 108.2 38.2 2.9-31.4 14.6-55.8 26.6-69.1-141.4-15.3-259.3-107.5-259.3-300.4 0-66.8 19.4-124.8 61.6-172.6-21.7-72.3 3.3-138 6.8-146.5 45-12.2 136.1 52.8 136.1 52.8 38.2-11.8 86.2-16.6 128.1-16.4 42 0.2 90 4.7 128.1 16.4 0 0 91.1-65 136.1-52.8 3.5 8.5 28.5 74.3 6.8 146.5 42.2 47.7 61.6 105.8 61.6 172.6 0 193.6-118.3 285.3-260.2 300 22 23.4 34.6 63 34.6 110.8V926c0 10.2-2.3 26.3 21.6 20C836.3 882.9 960 717.6 960 523.5c0-247.1-200.3-447.3-448.4-447.2z"></path>
            </svg>
        </a>
        <!-- QQ -->
        <div class="social-icon qq" @click="showQQ = true">
            <svg viewBox="0 0 1024 1024" width="1.2em" height="1.2em" fill="currentColor">
                <path d="M824.8 613.2c-16-51.4-34.4-94.6-62.7-165.3C766.5 262.2 689.3 112 511.5 112 331.7 112 256.2 265.2 261 447.9c-28.4 70.8-46.7 113.7-62.7 165.3-34 109.5-23 154.8-14.6 155.8 18 2.2 70.1-82.4 70.1-82.4 0 49 25.2 112.9 79.8 159-26.4 8.1-85.7 29.9-71.6 53.8 11.4 19.3 196.2 12.3 249.5 6.3 53.3 6 238.1 13 249.5-6.3 14.1-23.8-45.3-45.7-71.6-53.8 54.6-46.2 79.8-110.1 79.8-159 0 0 52.1 84.6 70.1 82.4 8.5-1.1 19.5-46.4-14.5-155.8z" p-id="2676"></path>
            </svg>
        </div>
      </div>
      
      <!-- QQ Dialog -->
      <el-dialog v-model="showQQ" title="扫码添加QQ" width="300px" center append-to-body>
        <div style="text-align: center;">
            <img src="/qq.png" style="width: 100%; max-width: 250px; border-radius: 8px;" alt="QQ QRCode" />
        </div>
      </el-dialog>
      
    </div>
  </GlassCard>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import GlassCard from './GlassCard.vue';

const props = defineProps({
    articles: { type: Number, default: 0 },
    categories: { type: Number, default: 0 },
    tags: { type: Number, default: 0 }
});

const router = useRouter();
const showQQ = ref(false);

const goAbout = () => {
  router.push('/about');
};
</script>

<style scoped>
.profile-card {
  padding: 30px 20px;
  cursor: pointer;
  text-align: center;
  overflow: hidden;
  position: relative;
}

.avatar-wrapper {
  position: relative;
  display: inline-block;
  margin-bottom: 15px;
}

.avatar-image {
  width: 90px;
  height: 90px;
  border-radius: 50%;
  box-shadow: 0 10px 20px rgba(161, 196, 253, 0.4);
  transition: transform 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  object-fit: cover;
}

.profile-card:hover .avatar-image {
  transform: rotate(360deg) scale(1.05);
}

.profile-name {
  margin: 0 0 8px;
  font-size: 1.3rem;
  color: var(--text-main);
  font-weight: 700;
}

.profile-bio {
  color: var(--text-secondary);
  font-size: 0.9rem;
  margin: 0 0 20px;
  line-height: 1.5;
}

/* Stats Row */
.stats-row {
  display: flex;
  justify-content: space-around;
  margin-bottom: 25px;
  padding: 0 10px;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  transition: transform 0.2s;
}

.stat-item:hover {
    transform: translateY(-2px);
}

.stat-value {
  font-size: 1.2rem;
  font-weight: 700;
  color: var(--text-main);
}

.stat-label {
  font-size: 0.75rem;
  color: var(--text-secondary);
  margin-top: 4px;
}

/* Social Links */
.social-links {
    display: flex;
    justify-content: center;
    gap: 20px;
    opacity: 0.9;
    transform: translateY(0); /* Always visible now, or keep hover effect */
    transition: all 0.3s ease;
}

.social-icon {
    width: 36px;
    height: 36px;
    border-radius: 50%;
    background: rgba(0,0,0,0.03);
    display: flex;
    align-items: center;
    justify-content: center;
    color: var(--text-secondary);
    transition: all 0.3s;
    cursor: pointer;
}
:global(.dark) .social-icon {
    background: rgba(255,255,255,0.05);
}

.social-icon:hover {
    color: #fff !important;
    transform: translateY(-3px);
    background: var(--primary-color);
}

/* Removed specific brand colors to unify feedback effect as requested */
</style>
