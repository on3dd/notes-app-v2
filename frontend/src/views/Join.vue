<template>
  <v-col
      :sm="10"
      :md="6"
  >
    <JoinForm/>
  </v-col>
</template>

<script>
  import JoinForm from "../components/JoinForm";

  export default {
    name: "Join",
    components: {
      JoinForm
    }
  }
</script>

<style scoped>

</style>