<template>
  <section class="sarradinhas-chart">
    <harsh-label class="label">Suas sarradas</harsh-label>
    <div class="chart" id="chart"></div>
  </section>
</template>

<script>
  import c3 from 'c3'

  export default {
    props: {
      stats: {
        type: Array,
        required: true
      }
    },
    mounted () {
      c3.generate({
        bindto: '#chart',
        data: {
          type: 'area-spline',
          columns: [
            ['Sarradinhas', ...this.stats]
          ],
          colors: { 'Sarradinhas': '#EFA118' }
        },
        legend: { show: false },
        axis: {
          y: { show: false },
          x: { show: false }
        }
      });
    }
  };
</script>

<style src="c3/c3.min.css"></style>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .sarradinhas-chart
    width: 100%
    border-bottom: 1px solid $color-primary
    padding-bottom: 5px

    > .label
      font-size: 12pt
      font-weight: 500
      text-transform: uppercase
      text-align: center
      color: $color-primary

    > .chart
      width: calc(100% - 30px)
      height: 125px
</style>


