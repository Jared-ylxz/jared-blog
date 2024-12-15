# jaredBlog2
This is my blog


前端前台创建步骤：
cd <my-project-folder>
npm create vite@latest front-end-client --template vue
选择Vue和JavaScript
cd front-end-client
npm install

创建路由文件 src/router/index.js

npm install axios
npm install vue-router


前端后台创建步骤：
cd <my-project-folder>
npm create vite@latest front-end-admin --template vanilla


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