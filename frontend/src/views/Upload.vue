<template>
  <v-col
      :sm="10"
      :md="6"
  >
    <UploadForm/>
  </v-col>
</template>

<script>
  import UploadForm from "../components/upload/UploadForm";

  export default {
    name: "Upload",
    components: {
      UploadForm
    }
  }
</script>

<style scoped>

</style>