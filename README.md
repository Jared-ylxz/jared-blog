# jaredBlog
This is my blog. It's built with Go and Vue.


# 创建项目过程的简单记录
1.创建后端项目：

2.创建前端前台项目：
cd <my-project-folder>
npm create vite@latest front-end-client --template vue
选择Vue和JavaScript
cd front-end-client
npm install
mkdir src/router/
touch src/router/index.js
npm install vue-router
npm install axios

3.创建前端后台项目：
cd <my-project-folder>
npm create vite@latest front-end-admin --template vue
选择Vue和JavaScript
cd front-end-admin
npm install
mkdir src/router/
touch src/router/index.js
npm install vue-router
npm install axios


# 启动项目：
cd <my-project-folder>
cd back-end
开发环境：RUNNING_ENV=development go run main.go
生产环境：RUNNING_ENV=production go run main.go
cd ../front-end-client
开发环境：npm run dev  # 会加载 .env.development 文件
生产环境：npm run build  # 会加载 .env.production 文件
cd ../front-end-admin
开发环境：npm run dev  # 会加载 .env.development 文件
生产环境：npm run build  # 会加载 .env.production 文件


# TODO
<!-- 前端添加文章发布功能
前端添加注册、登录功能
文章发布时，前后端校验是否已登录
前端增加管理页面，给用户删除、编辑文章
前端增加后台管理系统，用于管理员管理文章、评论等内容 -->

后端增加权限控制，只有admin用户才能管理文章（增加权限判断的中间件）
登录页面添加忘记密码功能
前端控制跳转到后台时，如果用户未登录，需要登录
将数据库、redis的配置信息添加到.env文件中，并防止密码泄露
跨域中的AllowOrigins使用环境变量配置

后端添加评论功能
前端添加评论功能
文章详情返回创建时间、作者、分类、标签、评论等信息
管理后台添加评论管理功能