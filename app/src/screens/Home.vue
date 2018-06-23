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
  import { requestGraphQL } from '@/helpers/request';

  export default {
    data: () => ({ message: 'Hello!' }),
    mounted () {
      requestGraphQL(`
        {
          user {
            name
          }
        }
      `).then((x) => {
        console.log(x);
      })
    }
  };
</script>
