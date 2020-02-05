<template>
  <v-container class="grey lighten-5 my-3 mx-4 mx-sm-6 mx-md-8">
    <h1 class="display-2 display-sm-3 text-center">Пользователи</h1>
    <v-row no-gutters class="mt-3">
      <template v-for="(item, index) in getUsers">
        <v-col :key="index" class="col-12 col-sm-6 col-md-4">
          <UserCard :user="item"/>
        </v-col>
      </template>
    </v-row>
  </v-container>
</template>

<script>
  import UserCard from "./UserCard";
  import {mapActions, mapGetters} from "vuex";

  export default {
    name: "UserCards",
    components: {
      UserCard
    },
    methods: {
      ...mapActions(["fetchUsers"]),
    },
    computed: {
      ...mapGetters(["getUsers"]),
    },
    async mounted() {
      await this.fetchUsers()
      console.log(this.getUsers)
    }
  }
</script>

<style scoped>

</style>