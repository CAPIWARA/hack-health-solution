<template>
  <harsh-screen classy="sarradinhas-details-screen">
    <header class="header sarradinhas-details-header">
      <harsh-title class="classification">{{ classification }}</harsh-title>
      <p class="points">
        <harsh-label class="quantity">{{ points }}</harsh-label>
        <harsh-label class="label">Sarradinhas</harsh-label>
      </p>
    </header>

    <harsh-label class="date">{{ date | toDate }}</harsh-label>

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
  export default {
    data () {
      const DAY = 24 * 60 * 60 * 1000;
      return {
        classification: 'Deus da sarrada',
        date: new Date( Date.now() - 4 * DAY ),
        points: 35,
        camisinha: null,
        drogas: null,
        ejaculou: null,
        oral: null,
        pessoa: null,
        quantidade: null
      };
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


