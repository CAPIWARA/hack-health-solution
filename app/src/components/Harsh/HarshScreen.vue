<template>
  <main class="harsh-screen">
    <slot />
  </main>
</template>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-screen
    display: flex
    flex-direction: column
    width: 100%
    min-height: 100vh
</style>
