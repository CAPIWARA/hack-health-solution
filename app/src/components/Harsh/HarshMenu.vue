<template>
  <nav class="harsh-menu">
    <harsh-menu-item
      v-for="item in items"
      :key="item.name"
      :icon="item.icon"
      :iconActive="item.iconActive"
      :name="item.name"
      :path="item.path"
      class="item"
    />
  </nav>
</template>

<script>
  import items from '@/services/items';
  import HarshMenuItem from '@/components/Harsh/HarshMenuItem'

  export default {
    components: { HarshMenuItem },
    data: () => ({ items })
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .harsh-menu
    display: flex
    align-items: center
    justify-content: space-around
    width: 100%
    height: 50px
    background-color: $color-medium-dark
</style>
