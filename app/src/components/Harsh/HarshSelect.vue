<template>
  <fieldset class="harsh-select">
    <label class="label">{{ label }}</label>
    <div class="selection">
      <p class="entry" @click="toggle()">{{ (value && value.label) || 'Selecione' }}</p>
      <img
        class="icon"
        src="~@/assets/icons/Arrow.png"
        :alt="isOpen ? 'Close' : 'Open'"
        :class="{ '-inverted': isOpen }"
        :title="isOpen ? 'Close' : 'Open'"
        @click="toggle()"
      />

      <transition name="open">
        <ul v-if="isOpen" class="harsh-select-items">
          <li
            class="harsh-select-item"
            v-for="item in items"
            :key="item.value"
            :class="{ '-selected': value === item }"
            @click="select(item)"
          >{{ item.label }}</li>
        </ul>
      </transition>
    </div>
  </fieldset>
</template>

<script>
  export default {
    data: () => ({ isOpen: false }),
    props: {
      value: {
        type: Object,
        default: null
      },
      label: {
        type: String,
        required: true
      },
      items: {
        type: Array,
        required: true
      }
    },
    methods: {
      select (value) {
        this.$emit('input', value);
        this.close();
      },
      toggle () {
        this.isOpen = !this.isOpen;
      },
      close () {
        this.isOpen = false;
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-select-selection
    &
      position: relative
      width: 100%
      padding: 6px 12px
      border-bottom: 2px solid $color-primary

    > .entry
      color: $color-primary
      { $typography-normal }
      font-size: 12pt

    > .icon
      position: absolute
      top: 50%
      right: 12px
      display: block
      width: 20px
      height: 13px
      transform: translateY(-50%)
      transition: transform 500ms ease

    > .icon.-inverted
      transform: translateY(-50%) rotate(180deg)

    > .harsh-select-items
      position: absolute
      top: 100%
      left: 0

  .harsh-select-items
    box-sizing: border-box
    width: 100%
    padding: 12px
    list-style: none
    background-color: $color-medium-dark
    z-index: 1

    > .harsh-select-item + .harsh-select-item
      margin-top: 12px

  .harsh-select-item
    { $typography-normal }
    font-size: 12pt
    color: #FFFFFF
    cursor: pointer
    transition: color 300ms ease

    &.-selected
      color: $color-primary + 25%
    &:hover
      color: $color-primary

  .harsh-select
    width: 100%
    border: 0

    > .selection
      @extends .harsh-select-selection

  .open-enter-active
  .open-leave-active
    transition: transform 500ms ease,
                opacity 500ms ease

  .open-enter
    transform: translateY(-50%) scaleY(0)
    opacity: 0

  .open-leave-to
    transform: translateY(-50%) scaleY(0)
    opacity: 0
</style>

