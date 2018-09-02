<template>
  <section class="container">
    <div>
      <h2 class="subtitle">
        title
      </h2>
        <img src="https://s3-ap-northeast-1.amazonaws.com/nagisa-intern/comic/1/thumb.jpeg">
      <div class="links">
         <a
          href="https://nuxtjs.org/"
          target="_blank"
          class="button--grey">NEXT</a>
        <a
          href="https://nuxtjs.org/"
          target="_blank"
          class="button--grey">PREV</a>
      </div>
      <div>
       <nuxt-link to="comment">comment</nuxt-link>
      </div>
      <div>
         <nuxt-link to="/">back</nuxt-link>
      </div>
      <br><br>
    </div>
  </section>
</template>

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

