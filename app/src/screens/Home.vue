<template>
  <harsh-screen class="home-screen">
    <img class="logo" src="~@/assets/images/Logo.png" alt="Sarrada" />

    <img class="banner" src="~@/assets/images/Banner-Sarrador.png" alt="Sarrador" />

    <harsh-label>SARROU?</harsh-label>
    <harsh-button>Sarrei!</harsh-button>
  </harsh-screen>
</template>

<script>
  import { requestGraphQL } from '@/helpers/request';

  export default {
    data: () => ({ message: 'Hello!' }),
    async mounted () {
      const response = await requestGraphQL(`
        {
          user {
            name
            gender
          }
        }
      `);
      const user = response.data.user;
      console.dir({ user });
    }
  };
</script>

<style lang="stylus">
  .home-screen
    align-items: center

    > .logo
      width: auto
      height: 35px
      margin-top: 10px

    > .banner
      margin-top: 20px
</style>

