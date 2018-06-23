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
</style>
