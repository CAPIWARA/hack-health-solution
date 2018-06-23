<template>
  <harsh-screen>
    <form class="login-form" @submit.prevent="onLogin()">
      <harsh-entry
        v-model="email"
        name="E-Mail"
      />

      <harsh-entry
        v-model="email"
        name="Senha"
      />

      <harsh-button type="submit">Entre</harsh-button>

      <p class="paragraph">
        ou
        <router-link class="link" :to="{ name: 'Subscribe' }">cadastre-se</router-link>
      </p>
    </form>
  </harsh-screen>
</template>

<script>
  export default {
    methods: {
      onLogin () {
        this.$router.push('Home');
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .login-form
    display: flex
    flex-direction: column
    justify-content: center
    align-items: center
    margin-top: 35vh
    padding: 30px

    > *
      margin-bottom: 25px

    > .paragraph
      color: #FFFFFF
      { $typography-normal }

    > .paragraph > .link
      color: #FFFFFF
</style>


