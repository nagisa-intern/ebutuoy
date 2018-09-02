<template>
  <section class="container">
    <div>
      <h2 class="subtitle">
        {{title}}
      </h2>
      <img :src="`https://s3-ap-northeast-1.amazonaws.com/nagisa-intern/comic/${$route.params.id}/thumb.jpeg`">
      <h2 class="subtitle">
        概要
        {{summary}}
      </h2>
      <div>
        <nuxt-link to="viewer">READ</nuxt-link>
      </div>
      <div>
        <nuxt-link to="/">BACK</nuxt-link>
      </div>
      <br><br>
    </div>
  </section>
</template>

<script>
  import axios from 'axios'
   asyncData ({params}) {
      return axios.get(`http://54.168.127.220/${$route.params.id}`) 
   .then((res) => {
     return { title: res.data.title}
       })
     }
   }
 </script>

<script>
  import axios from 'axios'
   asyncData ({params}) {
      return axios.get(`http://54.168.127.220/${$route.params.id}`) 
   .then((res) => {
     return { summary: res.data.summary}
       })
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

