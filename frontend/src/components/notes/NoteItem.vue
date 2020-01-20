<template>
  <v-container
      class="my-12">
    <v-row
        class="mx-md-6 px-2" justify="center" no-gutters>
      <v-dialog
          max-width="350" v-model="dialog">
        <v-card>
          <v-card-title class="headline">Вы действительно хотите удалить данную запись?</v-card-title>

          <v-card-text>
            После удаления доступ к данной записи пропадет, а также вы потеряете баллы рейтинга за данную работу.
          </v-card-text>
          <v-card-text>
            Последствия необратимы.
          </v-card-text>

          <v-card-actions>
            <v-spacer></v-spacer>

            <v-btn @click="dialog = !dialog" color="grey" text>
              Отмена
            </v-btn>

            <v-btn @click="confirmDeletion" color="red" text>
              Удалить
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <v-col :lg="6" :md="8" :sm="10" class="px-2">
        <div class="mb-3">
          <v-text-field
              class="display-2 display-3 note-title"
              id="note-title" outlined
              placeholder="Название"
              v-bind:class="{ 'd-none': !isEditing }" v-model="noteTitleInput"
          ></v-text-field>
          <span
              class="d-block display-2 display-sm-3"
              v-bind:class="{ 'd-none': isEditing }"
          >{{ getNote.title }}</span>
        </div>

        <div class="mb-2">
          <v-text-field
              class="headline note-description"
              id="note-description" outlined
              placeholder="Описание"
              v-bind:class="{ 'd-none': !isEditing }" v-model="noteDescriptionInput"
          ></v-text-field>
          <span class="headline font-weight-regular"
                v-bind:class="{ 'd-none': isEditing }"
          >{{getNote.description}}</span>
        </div>
        <div class="d-none d-md-block title mb-2">
          <a class="author font-weight-regular mr-2" href="">{{ getAuthor.name }},</a>
          <span class="font-weight-light">{{ getNote.posted_at }}</span>
        </div>
        <div class="d-block d-md-none title mb-2">
          <span class="d-block font-weight-regular mb-2">Автор:
            <a class="author" href="">{{ getAuthor.name }}</a>
          </span>
          <span class="d-block font-weight-light">{{ getNote.posted_at }}</span>
        </div>

        <div>
          <span class="title mb-2 d-block font-weight-regular">Предмет:
            <a class="subject" href="">{{subject.name}}</a>
          </span>
          <span class="title mb-2 d-block font-weight-regular">Преподаватель:
            <a class="teacher" href="">{{getTeacher.name}}</a>
          </span>
        </div>

        <div class="my-6">
          <v-btn class="d-sm-inline-block mr-1" color="primary" v-bind:disabled="!getNote.link" x-large>
            <a :href="getNote.link" style="color:white;text-decoration:none;">Отрыть</a>
          </v-btn>
          <div class="d-inline-block" v-if="!isEditing">
            <div class="d-inline-block mr-1">
              <v-btn @click="editNote" class="d-none d-sm-inline-block"
                     color="success" v-bind:disabled="!getNote.link" x-large>
                <a style="color:white;text-decoration:none;">Редактировать</a>
              </v-btn>
              <v-btn @click="editNote" class="d-sm-none my-2" color="success" style="min-width: 0;"
                     v-bind:disabled="!getNote.link" x-large>
                <a style="color:white;text-decoration:none;">
                  <v-icon dark>mdi-pencil</v-icon>
                </a>
              </v-btn>
            </div>
            <div class="d-inline-block mr-1">
              <v-btn @click="deleteNote" class="d-none d-sm-inline-block"
                     color="error" v-bind:disabled="!getNote.link" x-large>
                <a style="color:white;text-decoration:none;">Удалить</a>
              </v-btn>
              <v-btn @click="deleteNote" class="d-sm-none my-2" color="error" style="min-width: 0;"
                     v-bind:disabled="!getNote.link" x-large>
                <a style="color:white;text-decoration:none;">
                  <v-icon dark>mdi-delete</v-icon>
                </a>
              </v-btn>
            </div>
          </div>
          <div class="d-inline-block" v-else>
            <div class="d-inline-block mr-1">
              <v-btn @click="updateNote" class="d-none d-sm-inline-block mr-1" color="success"
                     v-bind:disabled="!getNote.link" x-large>
                <a style="color:white;text-decoration:none;">Сохранить</a>
              </v-btn>
              <v-btn @click="updateNote" class="d-sm-none my-2" color="success" style="min-width: 0;"
                     v-bind:disabled="!getNote.link" x-large>
                <a style="color:white;text-decoration:none;">
                  <v-icon dark>mdi-check</v-icon>
                </a>
              </v-btn>
            </div>
            <div class="d-inline-block mr-1">
              <v-btn @click="editNote" class="d-none d-sm-inline-block mr-1" v-bind:disabled="!getNote.link" x-large>
                <a style="color:inherit;text-decoration:none;">Отмена</a>
              </v-btn>
              <v-btn @click="editNote" class="d-sm-none my-2" style="min-width: 0;"
                     v-bind:disabled="!getNote.link" x-large>
                <a style="color:inherit;text-decoration:none;">
                  <v-icon dark>mdi-cancel</v-icon>
                </a>
              </v-btn>
            </div>
          </div>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
  import {mapActions, mapGetters} from "vuex";

  export default {
    name: "NoteItem",
    data() {
      return {
        subject: {
          name: "Неизвестный предмет"
        },
        teacher: {
          name: "Неизвестный преподаватель"
        },

        isEditing: false,
        noteTitleInput: '',
        noteDescriptionInput: '',
        dialog: false
      }
    },
    methods: {
      ...mapActions(["fetchNote"]),
      editNote: function () {
        this.isEditing = !this.isEditing
      },
      updateNote: function () {
        this.isEditing = !this.isEditing
      },
      deleteNote: function () {
        this.dialog = !this.dialog
      },
      confirmDeletion: function () {
        this.dialog = !this.dialog
      }
    },
    computed: {
      ...mapGetters(["getNote", "getAuthor", "getSubject", "getTeacher"]),
    },
    async mounted() {
      await this.fetchNote(this.$route.params.id)
    }
  }
</script>

<style scoped>
  .author, .subject, .teacher {
    text-decoration: none;
  }
  .author:hover, .subject:hover, .teacher:hover {
    text-decoration: underline;
  }
  .note-title, .note-description {
    padding: 0 !important;
    margin: 0 !important;
  }
</style>
<style>
  .note-title .v-input__slot, .note-description .v-input__slot {
    background: rgba(255, 255, 255, .5) !important;
    margin-bottom: 0;
  }
  .note-title .v-text-field__details, .note-description .v-text-field__details {
    display: none;
  }
  .note-title .v-text-field__slot input, .note-description .v-text-field__slot input {
    max-height: 100%;
  }
</style>