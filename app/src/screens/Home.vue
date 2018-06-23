<template>
  <harsh-screen>
    <harsh-entry
      v-model="message"
      name="Mensagem"
    />

    <harsh-button>{{ message }}</harsh-button>
  </harsh-screen>
</template>

<script>
  export default {
    data: () => ({ message: 'Hello!' })
  };
</script>
