<template>
  <p class="harsh-label">
    <slot />
  </p>
</template>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-label
    { $typography-normal }

    &
    > a
      color: $color-primary
</style>
