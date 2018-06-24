<template>
  <harsh-screen>
    <harsh-button>Okay!</harsh-button>
  </harsh-screen>
</template>

<script>
  import * as types from '@/store/types';

  export default {
    methods: {
      onCreateSarrada () {
        this.$store.dispatch(types.SARRADA_CREATE, {
          camisinha: params.camisinha,
          drogas: params.drogas,
          ejaculou: params.ejaculou,
          oral: params.oral,
          pessoa: params.pessoa,
          quantidade: params.quantidade,
        });
      }
    }
  };
</script>
