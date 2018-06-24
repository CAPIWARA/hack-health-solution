<template>
  <section class="sarradinhas-friends">
    <harsh-label class="label">Amigos sarradores</harsh-label>

    <ul class="friends">
      <li
        v-for="(friend, index) in friends"
        :key="index"
        class="friend"
      >
        <harsh-label class="number">#{{ index + 1 }}</harsh-label>

        <img :src="friend.avatar" class="avatar" />

        <p class="sarradinhas">
          <harsh-label class="name">{{ friend.name }}</harsh-label>
          <harsh-label class="quantity">{{ friend.quantity }} Sarradinhas</harsh-label>
        </p>
      </li>
    </ul>
  </section>
</template>

<script>
  export default {
    props: {
      friends: {
        type: Array,
        required: true
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .sarradinhas-friend
    display: flex
    align-items: center
    padding: 15px
    border-bottom: 1px solid $color-primary

    > .number
      min-width: 35px
      margin-right: 20px
      color: $color-primary
      font-size: 20pt
      font-weight: 500

    > .avatar
      width: 40px
      height: @width
      border-radius: 50%
      margin-right: 40px

    > .sarradinhas
      display: flex
      flex-direction: column

    > .sarradinhas > .name
      color: $color-primary
      font-size: 14pt
      font-weight: 500
      text-transform: uppercase
      overflow: hidden;
      white-space: nowrap;
      text-overflow: ellipsis;
      max-width: 190px;

    > .sarradinhas > .quantity
      color: $color-primary
      font-size: 12pt
      font-weight: 300

  .sarradinhas-friends
    width: 100%
    padding-bottom: 5px

    > .label
      font-size: 12pt
      font-weight: 500
      text-transform: uppercase
      text-align: center
      color: $color-primary

    > .friends > .friend
      @extends .sarradinhas-friend
</style>
