<template>
  <v-container class="grey lighten-5 my-3 mx-4 mx-sm-6 mx-md-8">
    <h1 class="display-2 display-sm-3 text-center">Все конспекты</h1>
    <v-row no-gutters class="mt-3">
      <template v-for="(item, index) in getNotes">
        <v-col :key="index" class="col-12 col-sm-6 col-md-4">
          <NoteCard :note="item"/>
        </v-col>
      </template>
    </v-row>
  </v-container>
</template>

<script>
  import {mapActions, mapGetters} from "vuex";
  import NoteCard from "./NoteCard";

  export default {
    name: "NoteCards",
    components: {
      NoteCard
    },
    methods: {
      ...mapActions(["fetchNotes"]),
    },
    computed: {
      ...mapGetters(["getNotes"]),
    },
    async mounted() {
      await this.fetchNotes()
      console.log(this.getNotes)
    }
  }
</script>

<style scoped>

</style>