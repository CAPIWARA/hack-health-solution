<template>
  <harsh-screen classy="create-sarrada-screen">
    <harsh-options
      v-model="params.camisinha"
      label="Você usou camisinha?"
      :options="[
        { label: 'Sim', value: true  },
        { label: 'Não', value: false }
      ]"
    />

    <harsh-options
      v-if="params.camisinha"
      v-model="params.oral"
      label="Até no sexo oral?"
      :options="[
        { label: 'Sim', value: true  },
        { label: 'Não', value: false }
      ]"
    />

    <harsh-select
      label="Quem você sarrou?"
      v-model="params.pessoa"
      :items="[
        { label: 'Parceira(o) Fixo', value: 'FIXO' },
        { label: 'Parceira(o) de uma única noite', value: 'UMA NOITE' },
        { label: 'Parceira(o) recorrente', value: 'RECORRENTE' },
        { label: 'Parceira(o) profissional', value: 'PROFISSIONAL' },
      ]"
    />

    <harsh-options
      v-model="params.quantidade"
      label="Transou com quantas pessoas?"
      :options="[
        { label: '1', value: 1 },
        { label: '2', value: 2 },
        { label: '3', value: 3 },
        { label: '3+', value: 4 }
      ]"
    />

    <harsh-options
      v-model="params.drogas"
      label="Estava sob efeito de drogas ou álcool?"
      :options="[
        { label: 'Sim', value: true  },
        { label: 'Não', value: false }
      ]"
    />

    <harsh-options
      v-model="params.ejaculou"
      label="Ejaculou?"
      :options="[
        { label: 'Sim', value: true  },
        { label: 'Não', value: false }
      ]"
    />

    <harsh-label v-if="error" class="error">{{ error }}</harsh-label>

    <harsh-button @click="onCreateSarrada()">Okay</harsh-button>
  </harsh-screen>
</template>

<script>
  import * as types from '@/store/types';

  export default {
    data () {
      return {
        error: '',
        params: {
          camisinha: null,
          drogas: null,
          ejaculou: null,
          oral: null,
          pessoa: null,
          quantidade: null
        }
      };
    },
    methods: {
      async onCreateSarrada () {
        try {
          await this.$store.dispatch(types.SARRADA_CREATE, {
            ...this.params,
            pessoa: this.params.pessoa && this.params.pessoa.value
          });
          const sarrada = this.$store.getters[types.SARRADA] || {};
          this.$router.push({ name: 'Detalhes da Sarrada', params: { id: sarrada.id } });
        } catch (error) {
          console.dir(error);
          this.error = 'Erro ao enviar a sarrada, confira os campos preenchidos ou sarre novamente.'
        }
      }
    }
  };
</script>

<style lang="stylus">
  @import '~@/assets/styles/theme'

  .create-sarrada-screen
    padding: 15px 30px

    & > .error
      color: $color-contrast

    & > *
      margin-bottom: 25px
</style>

