<template>
  <section class="harsh-screen">
    <slot />
  </section>
</template>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-screen
    display: flex
    flex-direction: column
    width: 100%
    min-height: 100vh
    background-color: $color-dark-background
    background-image: linear-gradient(320deg, #000000, #212121 50%, #000000 100%)
</style>
