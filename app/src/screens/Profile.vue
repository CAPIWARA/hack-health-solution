<template>
  <harsh-screen classy="profile-screen">
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
        { label: 'Heterossexual', value: 'Heterossexual' },
        { label: 'Homossexual', value: 'Homossexual' },
        { label: 'Bissexual', value: 'Bissexual' },
        { label: 'Assexual', value: 'Assexual' },
      ]"
    />

    <harsh-select
      v-model="ethnicity"
      label="Qual sua etnia?"
      :items="[
        { label: 'Negra', value: 'Negra'       },
        { label: 'Branca', value: 'Branca'     },
        { label: 'Indígena', value: 'Indígena' },
        { label: 'Asiática', value: 'Asiática' },
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
  export default {
    data () {
      return {
        name: '',
        email: '',
        birthdate: new Date(),
        genderIndentity: null,
        orientation: null,
        ethnicity: null,
        disease: null
      };
    }
  };
</script>

<style lang="stylus">
  .profile-screen
    padding: 15px 30px

    & > *
      margin-bottom: 25px
</style>
