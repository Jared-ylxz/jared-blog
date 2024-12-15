# jaredBlog2
This is my blog

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

# TODO
<!-- 前端添加文章发布功能
前端添加注册、登录功能
文章发布时，前后端校验是否已登录
前端增加管理页面，给用户删除、编辑文章
前端增加后台管理系统，用于管理员管理文章、评论等内容 -->

前端添加后台管理系统，用于对文章进行增、删、改
前端控制跳转到后台时，如果用户未登录，需要登录

后端添加评论功能
前端添加评论功能
文章详情返回创建时间、作者、分类、标签、评论等信息
管理后台添加评论管理功能