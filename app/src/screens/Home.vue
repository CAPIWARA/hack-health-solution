<template>
  <harsh-screen classy="home-screen">
    <img class="banner" src="~@/assets/images/Banner-Sarrador.png" alt="Sarrador" />

    <harsh-label class="label">SARROU?</harsh-label>
    <harsh-button class="button">Sarrei!</harsh-button>
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
    padding-top: 15px
    padding-bottom: @padding-top

    > .banner
      margin-bottom: auto
    > .label
      margin-bottom: 10px
</style>
