<template>
  <div class="schedule-next">
    <harsh-label class="label -sm">Recomendamos um novo teste em</harsh-label>
    <harsh-title class="date -xl">{{ next.days }} dias</harsh-title>
    <harsh-label class="label -md">{{ next.date }}</harsh-label>
    <harsh-label class="label -xs">
      <a href="https://google.com.br/">Saiba mais</a>
    </harsh-label>
  </div>
</template>

<script>
  import format from 'tiny-date-format';

  export default {
    props: {
      date: {
        type: Date,
        required: true
      }
    },
    computed: {
      next () {
        const DAY = 24 * 60 * 60 * 1000;
        const date = this.date.getTime() + 90 * DAY
        return {
          date: format(new Date(date), 'DD/MM/YYYY'),
          days: ~~((date - Date.now()) / DAY)
        };
      }
    }
  }
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .schedule-next
    width: 100%
    text-align: center

    > .label
    > .label > a
      color: $color-primary

    > .label.-xs
      margin-top: 30px
      font-size: 10pt

    > .label.-sm
      margin-top: 30px
      font-size: 12pt

    > .label.-md
      font-size: 18pt

    > .date
      margin-top: 6px
</style>


