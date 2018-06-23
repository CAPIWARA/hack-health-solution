<template>
  <fieldset class="harsh-entry">
    <input
      v-bind="$attrs"
      :name="name"
      :type="type"
      :value="value"
      :placeholder="name"
      class="entry"
      @input="input($event)"
    />
  </fieldset>
</template>

<script>
  export default {
    props: {
      value: String,
      type: {
        type: String,
        default: 'text',
        validator: (type) => [ 'text', 'email', 'password' ].includes(type)
      },
      name: {
        type: String,
        default: null
      }
    },
    methods: {
      input (event) {
        const value = event.target.value;
        this.$emit('input', value);
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-entry
    &
      display: flex
      border: 0
      width: 100%

    > .name
      { $typography-normal }
      display: block
      color: #FFFFFF

    > .entry
      display: block
      width: 100%
      padding: 10px
      border: 0
      border-bottom: 1px solid #FFFFFF
      background-color: transparent
      text-align: center
      outline: none

      &:focus
        border-bottom-color: $color-primary

      &
      &::placeholder
        { $typography-normal }
        color: #FFFFFF
</style>
