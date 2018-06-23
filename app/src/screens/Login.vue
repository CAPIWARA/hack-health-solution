<template>
  <harsh-screen class="login-screen">
    <form class="login-form" @submit.prevent="onLogin()">
      <img class="logo" src="~@/assets/images/Logo.png" alt="Sarrada" />

      <harsh-entry v-model="email" type="email" name="E-Mail" autocomplete="email" />

      <harsh-entry v-model="password" type="password" name="Senha" autocomplete="current-password" />

      <harsh-button type="submit">Entre</harsh-button>

      <harsh-label>
        ou <router-link class="link" :to="{ name: 'Subscribe' }">cadastre-se</router-link>
      </harsh-label>
    </form>
  </harsh-screen>
</template>

<script>
  export default {
    data () {
      return {
        email: '',
        password: ''
      };
    },
    methods: {
      onLogin () {
        this.$router.push({ name: 'Home' });
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .login-screen
    justify-content: center

  .login-form
    display: flex
    flex-direction: column
    justify-content: center
    align-items: center
    padding: 30px

    > *
      margin-bottom: 25px

    > .logo
      margin-bottom: 60px
</style>
