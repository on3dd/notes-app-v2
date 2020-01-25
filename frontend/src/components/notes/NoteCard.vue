<template>
  <v-card class="pa-2 ma-2 ma-md-3" tile>
    <v-card-title class="text-capitalize">{{note.title || "Неизвестная тема"}}</v-card-title>

    <v-card-subtitle class="card-info">
      <a :to="'users/'+note.author_id">{{note.author_name || "Кирилл Петрунин"}}</a>,
      <span>{{note.posted_at}}</span>
    </v-card-subtitle>

    <v-card-text>
      <div class="card-info">
        <span>Предмет: </span>
        <a href="">{{note.subject_name || "Математический анализ"}}</a>
      </div>
      <div class="card-info">
        <span>Преподаватель: </span>
        <a href="">{{note.teacher_name || "Клевчихин Ю. А."}}</a>
      </div>
    </v-card-text>

    <v-card-actions>
      <v-btn text :to="'notes/'+note.id">Открыть</v-btn>

      <v-spacer/>

      <v-btn icon>
        <v-icon>mdi-heart</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-bookmark</v-icon>
      </v-btn>

      <v-btn icon>
        <v-icon>mdi-share-variant</v-icon>
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script>
  export default {
    name: "NoteCard",
    props: {
      note: {
        required: true
      },
    },
  }
</script>

<style scoped>
  .card-info a {
    text-decoration: none;
  }
  .card-info a:hover {
    text-decoration: underline;
  }
</style>