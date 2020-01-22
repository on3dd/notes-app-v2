<template>
  <div></div>
</template>

<script>
  import axios from "axios"

  export default {
    name: "Notes",
    async mounted() {
      await axios.get("http://localhost:8080/api/v1/notes").then(response => console.log(response.data))
    }
  }
</script>

<style scoped>

</style>