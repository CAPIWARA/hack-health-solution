<template>
  <section class="sarradinhas-header">
    <figure class="user">
      <img class="avatar" :src="userAvatar" :alt="userName" />
      <figcaption class="name">{{ userName }}</figcaption>
    </figure>

    <div class="quantity">
      <harsh-label class="label">você tem</harsh-label>
      <harsh-title class="label">{{ quantity }}</harsh-title>
      <harsh-label class="label">Sarradinhas</harsh-label>
    </div>
  </section>
</template>

<script>
  export default {
    props: {
      userName: {
        type: String,
        required: true
      },
      userAvatar: {
        type: String,
        required: true
      },
      quantity: {
        type: Number,
        required: true
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .sarradinhas-header-user
    display: flex
    align-items: center
    flex-direction: column

    > .avatar
      width: 98px
      height: @width
      border-radius: 50%
      margin-bottom: 6px

    > .name
      { $typography-normal }
      font-weight: 400
      color: $color-primary
      text-transform: uppercase

  .sarradinhas-header
    display: flex
    border-bottom: 1px solid $color-primary
    padding-bottom: 25px

    > .user
      @extends .sarradinhas-header-user

    > .quantity
      margin-left: 42px

      > .label
        color: $color-primary
        line-height: 1

      > .label:nth-child(1)
        font-size: 14pt
        font-weight: 500

      > .label:nth-child(2)
        font-size: 62pt
        font-weight: 500

      > .label:nth-child(3)
        font-size: 22pt

</style>

