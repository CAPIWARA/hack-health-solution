<template>
  <harsh-screen classy="sarradinhas-history-screen">
    <header class="sarradinhas-history-header">
      <div class="quantity">
        <harsh-label class="label">Você tem</harsh-label>
        <harsh-title class="label">{{ quantity }}</harsh-title>
        <harsh-label class="label">Sarradinhas</harsh-label>
      </div>

      <harsh-label class="link">
        <router-link :to="{ name: 'Store' }">Trocar sarradinhas</router-link>
      </harsh-label>
    </header>

    <div class="sarradinhas-history-items">
      <router-link
        v-for="(sarrada, index) in sarradas"
        :key="index"
        :to="{ name: '' }"
        class="sarradinha-history-item"
      >
        <p class="sarrada">
          <harsh-label class="classification">{{ sarrada.classification }}</harsh-label>
          <harsh-label class="date">{{ sarrada.date | toDate }}</harsh-label>
        </p>

        <p class="points">
          <harsh-label class="quantity">{{ sarrada.points }}</harsh-label>
          <harsh-label class="label">Sarradinhas</harsh-label>
        </p>
      </router-link>
    </div>
  </harsh-screen>
</template>

<script>
  const DAY = 24 * 60 * 60 * 1000;

  export default {
    data () {
      return {
        quantity: 396,
        sarradas: [
          {
            classification: 'Sarrou gostoso',
            date: new Date( Date.now() - 2 * DAY ),
            points: 26
          },
          {
            classification: 'Sarrou mal',
            date: new Date( Date.now() - 4 * DAY ),
            points: 18
          },
          {
            classification: 'Deus da sarrada',
            date: new Date( Date.now() - 4 * DAY ),
            points: 35
          }
        ]
      };
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .sarradinhas-history-screen
    padding: 15px 30px

  .sarradinhas-history-header
    display: flex
    align-items: flex-end
    justify-content: space-between
    width: 100%
    margin-bottom: 30px

    & > .quantity
      text-decoration: none

    & > .quantity > .label
      color: $color-primary
      line-height: 1

    & > .quantity > .label:nth-child(1)
      font-size: 14pt
      font-weight: 500

    & > .quantity > .label:nth-child(2)
      font-size: 62pt
      font-weight: 500

    & > .quantity > .label:nth-child(3)
      font-size: 20pt

    & > .link
    & > .link > a
      font-size: 10pt
      text-transform: uppercase
      color: $color-primary

  .sarradinhas-history-items
    width: 100%
    border-top: 1px solid $color-primary

  .sarradinha-history-item
    display: flex
    justify-content: space-between
    width: 100%
    padding: 20px 0
    border-bottom: 1px solid $color-primary
    text-decoration: none

    > .sarrada
    > .points
      display: flex
      flex-direction: column

    > .sarrada > .classification
      font-size: 14pt
      font-weight: 500
      text-transform: uppercase

    > .sarrada > .date
      font-size: 12pt

    > .points
      align-items: center
      justify-content: center

    > .points > .quantity
      font-size: 24pt
      font-weight: 500

    > .points > .label
      font-size: 12pt
</style>
