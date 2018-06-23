<template>
  <p class="harsh-label">
    <slot />
  </p>
</template>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-label
    color: #FFFFFF
    { $typography-normal }

    > a
      color: #FFFFFF
</style>
