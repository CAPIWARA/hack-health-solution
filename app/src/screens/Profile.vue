<template>
  <p v-if="isLoading">Carregando...</p>

  <harsh-screen v-else classy="profile-screen">
    <img src="~@/assets/images/Avatar.png" alt="Avatar" />

    <harsh-entry
      v-model="name"
      type="text"
      name="Nome"
    />

    <harsh-entry
      v-model="email"
      type="email"
      name="E-Mail"
    />

    <harsh-entry
      v-model="birthdate"
      type="date"
      name="Data de Nascimento"
    />

    <harsh-select
      v-model="genderIndentity"
      label="Qual sua idêntidade de gênero?"
      :items="[
        { label: 'Mulher cis', value: 'Mulher cis' },
        { label: 'Homem cis', value: 'Homem cis' },
        { label: 'Mulher trans', value: 'Mulher trans' },
        { label: 'Homem trans', value: 'Homem trans' },
      ]"
    />

    <harsh-select
      v-model="orientation"
      label="Qual sua oriêntação sexual?"
      :items="[
        { label: 'Heterossexual', value: 'Hetero' },
        { label: 'Homossexual', value: 'Homo' },
        { label: 'Bissexual', value: 'Bi' },
        { label: 'Assexual', value: 'A' },
      ]"
    />

    <harsh-select
      v-model="ethnicity"
      label="Qual sua etnia?"
      :items="[
        { label: 'Negra', value: 'Negro'       },
        { label: 'Branca', value: 'Branco'     },
        { label: 'Indígena', value: 'Indígena' },
        { label: 'Asiática', value: 'Asiático' },
      ]"
    />

    <harsh-options
      v-model="disease"
      label="Possui alguma IST (HIV, Sífilis, etc)?"
      :options="[
        { label: 'Sim', value: true  },
        { label: 'Não', value: false }
      ]"
    />

    <harsh-button>Sarrar</harsh-button>
  </harsh-screen>
</template>

<script>
  import format from 'tiny-date-format';
  import { fetch } from '@/services/profile';

  export default {
    data () {
      return {
        name: '',
        email: '',
        birthdate: format(new Date(), 'YYYY-MM-DD'),
        genderIndentity: null,
        orientation: null,
        ethnicity: null,
        disease: null,
        isLoading: true
      };
    },
    async mounted () {
      const profile = await fetch();

      this.name = profile.name;
      this.email = profile.email;
      this.ethnicity = profile.ethnicity;
      this.orientation = profile.orientation;
      this.birthdate = format(new Date(profile.birthday), 'YYYY-MM-DD');
      this.genderIndentity = profile.genderIdentity;

      this.isLoading = false;
    }
  };
</script>

<style lang="stylus">
  .profile-screen
    padding: 15px 30px

    & > *
      margin-bottom: 25px
</style>
