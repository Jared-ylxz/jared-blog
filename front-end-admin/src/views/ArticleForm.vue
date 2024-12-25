<template>
  <div class="article-form">
    <h1>{{ isEditing ? '编辑文章' : '新增文章' }}</h1>
    <form @submit.prevent="handleSubmit">
      <div>
        <label for="title">标题</label>
        <input id="title" v-model="form.title" required />
      </div>
      <div>
        <label for="content">内容</label>
        <div ref="quillEditor" class="quill-editor"></div>
      </div>
      <button type="submit">{{ isEditing ? '更新' : '创建' }}</button>
    </form>
  </div>
</template>

<script>
import { fetchArticleDetail, createArticle, updateArticle } from '../apis/article';
import Quill from 'quill';
import 'quill/dist/quill.snow.css';

export default {
  name: 'ArticleForm',
  data() {
    return {
      form: {
        title: '',
        content: '',
      },
      isEditing: false,
      quill: null,
    };
  },
  methods: {
    async loadArticle() {
      const id = this.$route.params.id;
      if (id) {
        this.isEditing = true;
        const response = await fetchArticleDetail(id);
        this.form.title = response.data.title;
        this.form.content = response.data.content;
        this.quill.root.innerHTML = this.form.content; // Set content in Quill
      }
    },
    async handleSubmit() {
      this.form.content = this.quill.root.innerHTML; // Get content from Quill
      if (this.isEditing) {
        const id = this.$route.params.id;
        await updateArticle(id, this.form);
        alert('文章更新成功！');
      } else {
        await createArticle(this.form);
        alert('文章创建成功！');
      }
      this.$router.push('/');
    },
  },
  mounted() {
    this.quill = new Quill(this.$refs.quillEditor, {
      theme: 'snow',
      modules: {
        toolbar: [
          ['bold', 'italic', 'underline'],
          [{ list: 'ordered' }, { list: 'bullet' }],
          ['link', 'image'],
          ['clean'], // Remove formatting button
        ],
      },
    });
    this.loadArticle();
  },
};
</script>

<style>
.article-form {
  padding: 20px;
}
form {
  display: flex;
  flex-direction: column;
  gap: 10px;
}
input {
  width: 100%;
  padding: 8px;
}
.quill-editor {
  height: 300px; /* Adjust height as needed */
}
</style>