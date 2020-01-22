<template>
  <v-col
      :sm="10"
      :md="6"
  >
    <LoginForm/>
  </v-col>
</template>

<script>
  import LoginForm from "../components/LoginForm";

  export default {
    name: "Login",
    components: {
      LoginForm
    }
  }
</script>

<style scoped>

</style>