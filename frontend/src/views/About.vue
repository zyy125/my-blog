<template>
  <div class="about-container">
    <el-card>
      <div class="profile">
          <el-avatar :size="100" src="/vite.png" />
          <h2>逐影</h2>
          <p>这是我的个人博客</p>
      </div>
      
      <el-divider content-position="left">关于我</el-divider>
      
      <div class="content">
          <p>
            一名UESTCer, 正在学习go后端
          </p>
          <p>其他：</p>
          <ul>
              <li>这个人很懒，什么都没留下</li>
          </ul>
      </div>
    </el-card>
  </div>
</template>

<style scoped>
.about-container {
    max-width: 800px;
    margin: 0 auto;
}
.profile {
    text-align: center;
    margin-bottom: 40px;
}
.content {
    line-height: 1.8;
    color: #555;
    padding: 0 20px;
}
</style>
