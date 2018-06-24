<template>
  <p v-if="isLoading">Carregando...</p>

  <harsh-screen v-else classy="sarradinhas-screen">
    <sarradinhas-header
      class="header"
      :quantity="+user.sarradinhas"
      :user-name="user.name"
      :user-avatar="user.avatar"
    />

    <sarradinhas-chart
      class="chart"
      :stats="stats"
    />

    <sarradinhas-friends
      :friends="friends"
    />
  </harsh-screen>
</template>

<script>
  import SarradinhasChart from '@/components/Sarradinhas/SarradinhasChart';
  import SarradinhasHeader from '@/components/Sarradinhas/SarradinhasHeader';
  import SarradinhasFriends from '@/components/Sarradinhas/SarradinhasFriends';
  import { getHistory } from '@/services/sarrada';

  export default {
    components: { SarradinhasChart, SarradinhasHeader, SarradinhasFriends },
    data () {
      return {
        isLoading: true,
        friends: [],
        stats: [],
        user: {},
      };
    },
    async mounted () {
      this.isLoading = true;
      const { friends, user, sarradas } = await getHistory();
      this.user = user
      this.friends = friends;
      this.stats = sarradas.map((_) => _.total);
      this.isLoading = false;
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .sarradinhas-screen
    padding: 0 30px

    > *
      margin-top: 25px
</style>
