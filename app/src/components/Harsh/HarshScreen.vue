<template>
  <section class="harsh-screen" :class="{ '-is-FullScreen': isFullScreen }">
    <section class="content">
      <router-link v-if="!isFullScreen" class="link" :to="{ name: 'Home' }">
        <img class="logo" src="~@/assets/images/Logo.png" alt="Sarrada" />
      </router-link>

      <main :class="[ 'screen', classy ]">
        <slot />
      </main>
    </section>

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

    > .content
      height: calc(100% - 50px)
      overflow-y: auto

    > .content > .link
      display: block
      text-align: center

    > .content > .link > .logo
      width: auto
      height: 35px
      margin-top: 10px

    > .content > .screen
      display: flex
      align-items: center
      flex-direction: column
      min-height: calc(100% - 95px)

    &.-is-FullScreen > .content
      height: 100%

    &.-is-FullScreen > .content > .screen
      min-height: 100%
</style>
