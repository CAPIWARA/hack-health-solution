<template>
  <section class="harsh-screen" :class="{ '-is-FullScreen': isFullScreen }">
    <router-link v-if="!isFullScreen" class="link" :to="{ name: 'Home' }">
      <img class="logo" src="~@/assets/images/Logo.png" alt="Sarrada" />
    </router-link>

    <main :class="[ 'screen', classy ]">
      <slot />
    </main>

    <harsh-menu v-if="!isFullScreen" class="footer" />
  </section>
</template>

<script>
  export default {
    props: {
      isFullScreen: Boolean,
      classy: {}
    }
  };
</script>


<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-screen
    width: 100%
    height: 100%

    > .footer
      position: fixed
      left: 0
      bottom: 0

    > .screen
      display: flex
      align-items: center
      flex-direction: column
      height: calc(100vh - 95px)
      overflow-y: auto

    &.-is-FullScreen > .screen
      height: 100%


    > .link
      display: block
      text-align: center

    > .link > .logo
      width: auto
      height: 35px
      margin-top: 10px
</style>
