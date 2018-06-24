<template>
  <div class="harsh-options">
    <label class="label">{{ label }}</label>
    <div class="options">
      <div
        v-for="option in options"
        :key="option.value"
        class="option"
        @click="$emit('input', option.value)"
      >
        <div :class="['checkbox', { '-is-checked': value === option.value }]" />
        <label class="label">{{ option.label }}</label>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    props: {
      value: {
        type: [ Number, String, Boolean ],
        default: null
      },
      options: {
        type: Array,
        required: true
      },
      label: {
        type: String,
        required: true
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-option
    display: flex
    flex: 1

    & > .label
      display: flex
      align-items: center
      color: $color-primary
      { $typography-normal }

    & > .checkbox
      display: block
      width: 27px
      height: @width
      border: 2px solid $color-primary
      border-radius: 10px

    & > .checkbox.-is-checked::before
      display: block
      width: 17px
      height: @width
      border-radius: 7px
      margin: 3px
      background-color: $color-contrast
      content: ''

    & > .label
      margin-left: 8px

  .harsh-options
    display: flex
    flex-direction: column
    width: 100%
    border: 0

    > .label
      { $typography-normal }
      color: $color-primary
      margin-bottom: 7px

    > .options
      display: flex
      align-items: flex-start
      justify-content: flex-start

    > .options > .option
      @extends .harsh-option
</style>
