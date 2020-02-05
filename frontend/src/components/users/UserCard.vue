<template>
  <v-card class="pa-2 ma-2 ma-md-3" tile>
    <v-list-item three-line>
      <v-list-item-avatar size="80" color="grey">
        <v-icon dark large>mdi-account-circle</v-icon>
      </v-list-item-avatar>

      <v-list-item-content>
        <div class="overline mb-1">Пользователь</div>
        <v-list-item-title class="headline mb-1 text-truncate">{{user.name}}</v-list-item-title>
        <v-list-item-subtitle>{{user.about}}</v-list-item-subtitle>
      </v-list-item-content>
    </v-list-item>

    <v-card-actions>
      <v-btn text :to="'users/'+user.id">Открыть</v-btn>
        <v-btn text>Отслеживать</v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
  export default {
    name: "UserCard",
    props: {
      user: {
        required: true
      },
    },
  }
</script>

<style scoped>

</style>