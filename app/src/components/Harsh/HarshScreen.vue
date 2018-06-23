<template>
  <section class="harsh-screen">
    <slot />
  </section>
</template>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-screen
    width: 100%
    min-height: 100vh
    background-color: $color-dark-background
</style>
