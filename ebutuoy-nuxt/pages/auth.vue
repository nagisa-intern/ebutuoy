<template>
  <section class="container">
    <div>
      <h2 class="subtitle">
        Login
      </h2>
      <form v-if="!$store.state.authUser" @submit.prevent="login">
        <p class="error" v-if="formError">{{ formError }}</p>
        <div>
          <input type="text" class="form-control" v-model="formUsername" name="username" placeholder="Username" />
          <input type="password" class="form-control" v-model="formPassword" name="password"  placeholder="Password" />
          <button type="submit" class="button--green block">Login</button>
        </div>
      </form>
      <div v-else>
        <h2>Hello {{ $store.state.authUser.username }}!</h2>
        <div class="links">
          <a href="/hello" class="button--green">go to hello page</a>
          <button class="button--grey" @click="logout">Logout</button>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
export default {
  data() {
    return {
      formError: null,
      formUsername: '',
      formPassword: ''
    }
  },
  methods: {
    async login() {
      try {
        await this.$store.dispatch('login', {
          username: this.formUsername,
          password: this.formPassword
        })
        this.formUsername = ''
        this.formPassword = ''
        this.formError = null
      } catch (e) {
        this.formError = e.message
      }
    },
    async logout() {
      try {
        await this.$store.dispatch('logout')
      } catch (e) {
        this.formError = e.message
      }
    }
  }
}
</script>
