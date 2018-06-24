<template>
  <div class="schedule-latest">
    <harsh-label class="label">Último teste</harsh-label>
    <harsh-title class="latest -xl">{{ latest }}</harsh-title>
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
      latest () {
        return format(this.date, 'DD/MM')
      }
    }
  }
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .schedule-latest
    width: 100%
    text-align: center
    border-bottom: 2px solid $color-primary

    > .label
      color: $color-primary

    > .latest
      margin-bottom: 30px
</style>


