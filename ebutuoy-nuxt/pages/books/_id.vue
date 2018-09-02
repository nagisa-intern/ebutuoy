<template>
  <div>
    <!-- Header　-->
    <div>
      <AppHeader></AppHeader>
      <nuxt/>
    </div>

    <section class="container">
      <div>
        <h2 class="subtitle">

        </h2>
        <img :src="`https://s3-ap-northeast-1.amazonaws.com/nagisa-intern/comic/${$route.params.id}/thumb.jpeg`">
        <span>概要</span>
        <div>
          <nuxt-link to="viewer">READ</nuxt-link>
        </div>
        <div>
          <nuxt-link to="/">BACK</nuxt-link>
        </div>
        <br><br>
      </div>
    </section>

    <!-- Footer　-->
    <div class="books-footer">
      <AppFooter></AppFooter>
      <nuxt/>
    </div>

  </div>
</template>

<script>
import AppHeader from '@/components/Header.vue';
import AppFooter from '@/components/Footer.vue';
import axios from 'axios'

axios.defaults.withCredentials = true
//axios.defaults.baseURL = 'http://ebutuoy.to-hutohu.com/api'
axios.defaults.baseURL = 'http://54.248.63.189/api'

export default {
  data() {
    books: []
  },
  components: {
    AppHeader,
    AppFooter,
  },
  async created( ) {
    await axios.get('/me').catch(async () => {
      await axios.post('/me', {username: 'po', password: 'po'})
    })
    await axios
    .get(`/comics/${params.id}`)
    .then(response => (this.books = response.data))
    .catch(error => console.log(error))
    console.log(this.books);
  }
};

</script>

<style>
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}
.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}
.links {
  padding-top: 10px;
}
</style>
