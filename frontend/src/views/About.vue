<template>
  <div class="about-container">
    <el-card>
      <div class="profile">
          <el-avatar :size="100" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
          <h2>My Blog</h2>
          <p>这是基于 Go (Gin) + Vue 3 的个人博客系统。</p>
      </div>
      
      <el-divider content-position="left">关于本站</el-divider>
      
      <div class="content">
          <p>
            欢迎来到我的博客！这里是我记录技术学习、生活感悟的地方。
          </p>
          <p>技术栈：</p>
          <ul>
              <li>后端：Go, Gin, GORM, MySQL</li>
              <li>前端：Vue 3, Vite, Element Plus, Pinia</li>
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
