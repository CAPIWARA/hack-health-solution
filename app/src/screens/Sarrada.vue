<template>

</template>

<script>
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
