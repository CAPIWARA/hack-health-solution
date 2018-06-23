<template>
  <fieldset class="harsh-entry">
    <label class="name">{{ name }}</label>
    <input
      :name="name"
      :type="type"
      :value="value"
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

    > .name
      display: block
      { $typography-normal }

    > .entry
      display: block
      { $typography-normal }
</style>
