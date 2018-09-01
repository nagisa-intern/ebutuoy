<template>
  <section class="container">
 <div>
    <img :src="`https://github.com/${user.login}.png`" :alt="user.name">
    <h1>{{user.name}}</h1>
    <div class="links">
      <a href="${user.html_url}" target="_blank"
          class="button--green">Githubに飛ぶ</a>
      <a
        href="https://nuxtjs.org/"
        target="_blank"
        class="button--green">Documentation</a>
      <a
        href="https://github.com/nuxt/nuxt.js"
        target="_blank"
        class="button--grey">GitHub</a>
    </div>
    <div class="create_update">
        <h2>"create_at  : " {{user.created_at}}</h2>
        <h2>"updated_at : " {{user.updated_at}}</h2>
    </div>
    </div>
  </section>
</template>

<script>
import axios from 'axios'

export default {
  async asyncData ({ params }) {
    const { data: user } = await axios.get(`https://api.github.com/users/${params.id}`)
    return {
      user
    }
  }
}
</script>

<style>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  flex-direction: column;
}

.container img{
  margin-bottom: 20px;
  border-radius: 50%;
  overflow: hidden;
}

.create_update {
  padding-top:15px;
}
</style>
