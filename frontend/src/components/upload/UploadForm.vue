<template>
  <v-container>
    <v-form
        id="form"
        lazy-validation
        onsubmit="return false;"
        v-model="valid">
      <span class="display-2 display-sm-3 text-center d-block mb-3">Загрузка работы</span>
      <v-text-field
          :rules="titleRules" label="Название" name="title" required v-model="title" />
      <v-textarea
          :rules="descriptionRules" hint="Введите описание работы" label="Описание" name="descirption"
          v-model="description" />
      <v-file-input
          :rules="fileRules" accept="application/pdf" label="Файлы" name="file" required v-model="file" />
      <v-select
          :items="allCategories" :rules="categoryRules" @change="fetchData(category)" item-text="name" label="Категория"
          name="category" required v-model="category"/>
      <v-select
          :items="allSubjects" :rules="subjectRules" @change="updateSubject" item-text="name" label="Предмет"
          name="subject" required v-model="subject"/>
      <v-select
          :items="allTeachers" :rules="teacherRules" @change="updateTeacher" item-text="name" label="Преподаватель"
          name="teacher" required v-model="teacher" />
      <v-checkbox
          :rules="[v => !!v || 'Вы должны принять условия пользования!']"
          label="Я согласен с условиями пользования сервисом"
          required
          v-model="checkbox"/>
      <div class="d-flex justify-center">
        <v-btn
            :disabled="!valid" @click="submit" class="mx-4" color="success" large type="submit">
          Отправить
        </v-btn>
      </div>
    </v-form>
  </v-container>
</template>

<script>
  import axios from 'axios'
  import {mapActions, mapGetters, mapMutations} from 'vuex'

  export default {
    name: "UploadForm",
    data: () => ({
      title: '',
      description: '',
      category: '',
      subject: '',
      teacher: '',

      valid: false,
      titleRules: [
        v => !!v || 'Поле должно быть заполнено',
        v => (v && v.length >= 10) || 'Название должно содержать как минимум 10 символов',
      ],

      file: null,
      fileRules: [
        v => !!v || "Поле должно быть заполнено"
      ],

      descriptionRules: [],

      categoryRules: [
        v => !!v || 'Поле должно быть заполнено',
      ],

      subjectRules: [
        v => !!v || 'Поле должно быть заполнено',
      ],

      teacherRules: [
        v => !!v || 'Поле должно быть заполнено',
      ],

      checkbox: false,
    }),

    methods: {
      ...mapActions(["fetchCategories", "fetchData", "fetchSubjects", "fetchTeachers"]),

      updateSubject() {
        return this.$store.commit("updateSubject", this.subject)
      },
      updateTeacher() {
        return this.$store.commit("updateTeacher", this.teacher)
      },

      submit: async function () {
        const categoryIdx = this.getCategoryIndex
        if (categoryIdx === -1) return console.error('CategoryIdx is empty')

        const subjectIdx = this.getSubjectIndex
        if (subjectIdx === -1) return console.error('SubjectIdx is empty')

        const teacherIdx = this.getTeacherIndex
        if (teacherIdx === -1) return console.error('TeacherIdx is empty')

        const data = new FormData()

        const categories = this.allCategories
        const subjects = this.allSubjects
        const teachers = this.allTeachers

        data.append("author", 1)
        data.append("category_id", categories[categoryIdx].id)
        data.append("subject_id", subjects[subjectIdx].id)
        data.append("teacher_id", teachers[teacherIdx].id)

        data.append("title", this.title)
        data.append("description", this.description)
        data.append("file", this.file)

        const response = await axios.post("http://localhost:8080/api/v1/notes", data)

        if (response.status === 200) {
          await this.$router.replace({path: `/notes/${response.data.id}`})
        }
      }
    },
    computed: {
      ...mapGetters(["allCategories", "allSubjects", "allTeachers", "getCategoryIndex", "getSubjectIndex", "getTeacherIndex"]),
      // getCategoryIndex() {
      //   return this.$store.getters.getCategoryIndex(this.category)
      // },
      // getSubjectIndex() {
      //   return this.$store.getters.getSubjectIndex(this.subject)
      // },
      // getTeacherIndex() {
      //   return this.$store.getters.getTeacherIndex(this.teacher)
      // },
    },
    async mounted() {
      await this.fetchCategories()
    }
  }

</script>

<style></style>