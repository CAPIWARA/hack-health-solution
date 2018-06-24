<template>
  <h2 class="harsh-title">
    <slot />
  </h2>
</template>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-title
    { $typography-title }
    color: $color-primary

    &.-xl
      font-size: 48pt
</style>
