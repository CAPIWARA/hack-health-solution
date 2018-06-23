<template>
  <component
    v-bind="$attrs"
    v-on="$listeners"
    :class="classes"
    :is="to ? 'router-link' : 'button'"
    :to="to ? to : null"
  >
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

  harsh-pill($size, $background)
    display: block
    height: $size
    border-radius: $size
    background-color: $background

  .harsh-button
    harsh-pill(48px, $color-primary)
    position: relative
    padding: 0 40px
    border: 0
    outline: none
    cursor: pointer

    &::before
      harsh-pill(48px, $color-contrast)
      position: absolute
      top: 0
      left: 0
      width: 100%
      z-index: -1
      content: ''
      transform: translate(3px, 3px)
      transition: transform .5s ease

    &:active::before
      transform: translate(-2px, -2px)
      transition: transform .1s ease

    &.-is-button
      background-image: none

    &.-is-link
      text-decoration: none

    > .content
      { $typography-normal }
      text-transform: uppercase
      color: $color-dark-typography
</style>
