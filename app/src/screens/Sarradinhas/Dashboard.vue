<template>
  <harsh-screen classy="sarradinhas-screen">
    <sarradinhas-header
      class="header"
      :quantity="396"
      user-name="Vitor"
      user-avatar="https://s3-us-west-2.amazonaws.com/s.cdpn.io/78779/profile/profile-512.jpg"
    />

    <sarradinhas-chart
      class="chart"
      :stats="[ 32, 29, 39, 15, 14, 15 ]"
    />

    <sarradinhas-friends
      :friends="[
        {
          avatar: 'https://avatars3.githubusercontent.com/u/455023',
          name: 'Minko Gechev',
          quantity: 198
        },
        {
          avatar: 'https://avatars3.githubusercontent.com/u/455023',
          name: 'Minko Gechev',
          quantity: 211
        },
        {
          avatar: 'https://avatars3.githubusercontent.com/u/455023',
          name: 'Minko Gechev',
          quantity: 211
        }
      ]"
    />
  </harsh-screen>
</template>

<script>
  import SarradinhasChart from '@/components/Sarradinhas/SarradinhasChart';
  import SarradinhasHeader from '@/components/Sarradinhas/SarradinhasHeader';
  import SarradinhasFriends from '@/components/Sarradinhas/SarradinhasFriends';

  export default {
    components: { SarradinhasChart, SarradinhasHeader, SarradinhasFriends },
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .sarradinhas-screen
    padding: 0 30px

    > *
      margin-top: 25px
</style>
