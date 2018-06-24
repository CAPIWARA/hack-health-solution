<template>
  <p v-if="isLoading">Carregando...</p>

  <harsh-screen v-else classy="sarradinhas-details-screen">
    <header class="header sarradinhas-details-header">
      <harsh-title class="classification">{{ classification }}</harsh-title>
      <p class="points">
        <harsh-label class="quantity">{{ points }}</harsh-label>
        <harsh-label class="label">Sarradinhas</harsh-label>
      </p>
    </header>

    <!-- <harsh-label class="date">{{ date | toDate }}</harsh-label> -->

    <article class="sarradinhas-details-description">
      <harsh-label class="answer">Você usou camisinha? <strong>{{ camisinha | toAnswer }}</strong></harsh-label>
      <harsh-label class="answer">Até no sexo oral? <strong>{{ oral | toAnswer }}</strong></harsh-label>
      <harsh-label class="answer">Com quem você transou? <strong>{{ pessoa | toPessoa }}</strong></harsh-label>
      <harsh-label class="answer">Transou com quantas pessoas? <strong>{{ quantidade | toQuantidade }}</strong></harsh-label>
      <harsh-label class="answer">Usou drogas ou álcool <strong>{{ drogas | toAnswer }}</strong></harsh-label>
      <harsh-label class="answer">Ejaculou? <strong>{{ ejaculou | toAnswer }}</strong></harsh-label>
    </article>

    <harsh-button :to="{ name: 'Histórico de Sarradas' }">Okay</harsh-button>
  </harsh-screen>
</template>

<script>
  import * as types from '@/store/types'

  export default {
    data () {
      const DAY = 24 * 60 * 60 * 1000;
      return {
        isLoading: false,

        classification: null,
        date: null,
        points: null,
        camisinha: null,
        drogas: null,
        ejaculou: null,
        oral: null,
        pessoa: null,
        quantidade: null
      };
    },
    async mounted () {
      this.isLoading = true;
      const id = this.$route.params.id;
      await this.$store.dispatch(types.SARRADA, id);
      const sarrada = this.$store.getters[types.SARRADA] || {};

      this.classification = sarrada.message

      this.points = sarrada.total
      this.camisinha = sarrada.camisinha
      this.drogas = sarrada.drogas
      this.ejaculou = sarrada.ejaculou
      this.oral = sarrada.oral
      this.pessoa = sarrada.pessoa
      this.quantidade = sarrada.quantidade

      this.isLoading = false;
    },
    filters: {
      toPessoa: (value) => ({
                'FIXO': 'Parceira(o) Fixo',
           'UMA NOITE': 'Parceira(o) de uma única noite',
          'RECORRENTE': 'Parceira(o) recorrente',
        'PROFISSIONAL': 'Parceira(o) profissional',
      })[value] || null,
      toQuantidade: (value) => value === 4 ? 'Mais de 3' : value
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .sarradinhas-details-header
    display: flex
    align-items: center
    justify-content: space-between

    > .classification
      text-transform: uppercase

    > .points
      display: flex
      align-items: center
      flex-direction: column
      justify-content: center

      & > .quantity
        font-size: 26pt
        font-weight: 600

      & > .label
        font-size: 14pt

  .sarradinhas-details-screen
    padding: 15px 30px

    > .header
      margin-bottom: 15px

    > .date
      font-size: 12pt
      margin-bottom: 20px

  .sarradinhas-details-description
    width: 100%
    border-top: 1px solid $color-primary
    padding-top: 35px

    > .answer
      margin-bottom: 20px

    strong
      font-weight: 500
</style>


