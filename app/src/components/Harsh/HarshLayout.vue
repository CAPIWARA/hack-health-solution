<template>
  <section class="harsh-layout">
    <transition name="fade">
      <slot />
    </transition>
  </section>
</template>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-layout
    height: 100vh
    background-color: $color-dark-background
    background-image: linear-gradient(320deg, #000000, #212121 50%, #000000 100%)
    background-position: center
    background-attachment: fixed

  .fade-enter-active
  .fade-leave-active
    transition: opacity .3s .3s ease-in

  .fade-enter
  .fade-leave-to
    opacity: 0
</style>
