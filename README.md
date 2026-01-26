# 个人博客系统 

这是一个用ai开发的基于 **Go (Gin)** 和 **Vue 3**的个人博客系统。

**在线网址：[zhu-ying.online](zhu-ying.online)**

## 功能特性

- **文章管理**: 支持 Markdown 编辑器，实时预览，代码高亮。
- **分类与标签**: 灵活的内容组织方式。
- **评论系统**: 访客可以发表评论（支持嵌套回复）。
- **响应式设计**: 适配桌面端和移动端（保持一致的布局）。
- **后台管理**: 包含文章发布、编辑、删除等管理功能。
- **Docker 部署**: 一键容器化部署，开箱即用。

## 技术栈

### 后端 (Backend)
- **语言**: Go (Golang)
- **Web 框架**: Gin
- **ORM**: GORM
- **数据库**: MySQL
- **配置管理**: Viper

### 前端 (Frontend)
- **核心框架**: Vue 3, Vite
- **UI 组件库**: Element Plus
- **状态管理**: Pinia
- **路由**: Vue Router
- **Markdown**: v-md-editor, marked, highlight.js
- **工具**: Axios, Dayjs

### 部署 (Deployment)
- **容器化**: Docker, Docker Compose

##  项目结构
my-blog/
├── backend/            # 后端项目代码 (Go)
│   ├── cmd/            # 入口文件
│   ├── config/         # 配置文件
│   ├── internal/       # 业务逻辑
│   ├── uploads/        # 上传文件存储
│   └── Dockerfile      # 后端构建文件
├── frontend/           # 前端项目代码 (Vue 3)
│   ├── src/            # 源代码
│   └── Dockerfile      # 前端构建文件
├── docker-compose.yml  # Docker 编排文件
└── README.md           # 项目说明文档

## 快速开始 (Docker 部署)

确保本地已安装 [Docker](https://www.docker.com/) 和 [Docker Compose](https://docs.docker.com/compose/)。

### 1. 克隆项目

```bash
git clone https://github.com/zyy125/my-blog.git
cd my-blog
```

### 2. 配置文件准备

**后端配置**:
检查 `backend/config/` 目录，确保存在 `config.yaml`。如果不存在，请复制示例配置：

```bash
cp backend/config/config.example.yaml backend/config/config.yaml
```

**数据库配置**:
在 `docker-compose.yml` 中，您可以修改 MySQL 的环境变量（如密码）：

```yaml
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: my_blog_db
      MYSQL_USER: user
      MYSQL_PASSWORD: password
```

### 3. 启动服务

在项目根目录下运行：

```bash
docker compose up -d --build
```

该命令将自动构建前端和后端镜像，并启动 MySQL、后端和前端容器。

### 4. 访问应用

- **前台页面**: http://localhost
- **后端接口**: http://localhost:8080

## 本地开发指南

如果您需要在本地进行开发调试：

### 数据库准备
1. 启动一个本地 MySQL 实例。
2. 创建数据库 `my_blog_db`。
3. 修改 `backend/config/config.yaml` 中的数据库连接信息（Host 改为 `localhost` 或 `127.0.0.1`）。

### 后端启动
```bash
cd backend
go mod download
go run cmd/main.go
```

### 前端启动
```bash
cd frontend
npm install
npm run dev
```

## 开源协议
MIT License