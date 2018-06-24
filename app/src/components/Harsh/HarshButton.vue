<template>
  <component
    v-bind="$attrs"
    v-on="$listeners"
    :class="classes"
    :is="to ? 'router-link' : 'button'"
    :to="to ? to : null"
  >
    <figure class="background"></figure>
    <figure class="background"></figure>

    <span class="content">
      <slot />
    </span>
  </component>
</template>

<script>
  export default {
    props: {
      to: {
        type: [ String, Object ],
        default: null
      },
      isAlternative: {
        type: Boolean,
        default: false
      },
      size: {
        type: String,
        default: 'md',
        validator: (size) => ['md'].includes(size)
      }
    },
    computed: {
      classes () {
        return [
          'harsh-button',
          '-' + this.size, {
            '-is-link': this.to,
            '-is-button': !this.to,
            '-is-alternative': this.isAlternative
          }
        ];
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  $harsh-button-size = 48px

  .harsh-button
    position: relative
    z-index: 0
    height: $harsh-button-size
    min-height: $harsh-button-size
    border: 0
    padding: 0
    padding-left: 40px
    padding-right: @padding-left
    outline: none
    cursor: pointer

    > .background
      position: absolute
      top: 0
      left: 0
      z-index: -1
      display: block
      width: 100%
      height: $harsh-button-size
      border-radius: @height
      background-color: $color-primary

    > .background:nth-child(1)
      background-color: $color-contrast
      transform: translate(3px, 3px)
      transition: transform .5s ease

    &:active > .background:nth-child(1)
      transform: translate(-2px, -2px)
      transition: transform .1s ease

    &.-is-button
      background-image: none
      background-color: transparent

    &.-is-link
      text-decoration: none
      line-height: $harsh-button-size

    > .content
      { $typography-normal }
      text-transform: uppercase
      color: $color-dark-typography
</style>
