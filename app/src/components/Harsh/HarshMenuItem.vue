<template>
  <router-link
    :to="path"
    exact
    class="harsh-menu-item"
    active-class="-active"
    exact-active-class="-active"
  >
    <figure class="icon">
      <img class="normal" :src="icon" :alt="name" />
      <img class="active" :src="iconActive" :alt="name" />
    </figure>
    <span class="name">{{ name }}</span>
  </router-link>
</template>

<script>
export default {
  props: {
    icon: {
      type: String,
      required: true
    },
    iconActive: {
      type: String,
      required: true
    },
    name: {
      type: String,
      required: true
    },
    path: {
      type: [ String, Object ],
      required: true
    }
  }
};
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-menu-item-icon
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 15px;
    margin-bottom: 4px;

    > .normal
      max-height: 15px
      display: block

    > .active
      max-height: 15px
      display: none

  .harsh-menu-item-icon-active
    & > .normal
      display: none
    & > .active
      display: block

  .harsh-menu-item
    display: block
    text-decoration: none

    > .icon
      @extends .harsh-menu-item-icon

    &.-active > .icon
      @extends .harsh-menu-item-icon-active

    > .name
      { $typography-small }
      color: $color-primary

    &.-active > .name
      color: $color-contrast
</style>

