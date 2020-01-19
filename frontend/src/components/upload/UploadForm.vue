<template>
  <v-container>
    <v-form
        id="form"
        lazy-validation
        onsubmit="return false;"
        v-model="valid"
    >
      <span class="display-2 display-sm-3 text-center d-block mb-3">Загрузка работы</span>
      <v-text-field
          :rules="titleRules" label="Название" name="title" required v-model="title"
      ></v-text-field>

      <v-textarea
          :rules="descriptionRules" hint="Введите описание работы" label="Описание" name="descirption"
          v-model="description"
      ></v-textarea>

      <v-file-input
          :rules="fileRules" accept="application/pdf" label="Файлы" name="file" required v-model="file"
      ></v-file-input>

      <v-select
          :items="categories" :rules="categoryRules" @change="getSubjects" item-text="name" label="Категория"
          name="category" required v-model="category"
      ></v-select>

      <v-select
          :items="subjects" :rules="subjectRules" @change="getTeachers" item-text="name" label="Предмет"
          name="subject" required v-model="subject"
      ></v-select>

      <v-select
          :items="teachers" :rules="teacherRules" item-text="name" label="Преподаватель" name="teacher"
          required v-model="teacher"
      ></v-select>

      <v-checkbox
          :rules="[v => !!v || 'Вы должны принять условия пользования!']"
          label="Я согласен с условиями пользования сервисом"
          required
          v-model="checkbox"
      ></v-checkbox>

      <div class="d-flex justify-center">
        <v-btn
            :disabled="!valid" @click="this.submit" class="mx-4" color="success" large type="submit"
        >
          Отправить
        </v-btn>
      </div>
    </v-form>
  </v-container>
</template>

<script>
  import axios from 'axios'

  export default {
    name: "UploadForm",
    data: () => ({
      valid: false,
      title: '',
      titleRules: [
        v => !!v || 'Поле должно быть заполнено',
        v => (v && v.length >= 10) || 'Название должно содержать как минимум 10 символов',
      ],

      file: null,
      fileRules: [
        v => !!v || "Поле должно быть заполнено"
      ],

      description: '',
      descriptionRules: [],

      category: '',
      categoryRules: [
        v => !!v || 'Поле должно быть заполнено',
      ],
      categories: ["Выберите категорию"],

      subject: '',
      subjectRules: [
        v => !!v || 'Поле должно быть заполнено',
      ],
      subjects: ["Выберите предмет"],

      teacher: '',
      teacherRules: [
        v => !!v || 'Поле должно быть заполнено',
      ],
      teachers: ["Выберите преподавателя"],

      checkbox: false,
    }),

    methods: {
      getCategories: async function () {
        const response = await axios.get('http://localhost:8080/api/v1/categories')

        this.categories = response.data

        this.subjects = []
        this.subject = ''

        this.teachers = []
        this.teacher = ''
      },
      getSubjects: async function () {
        let categoryIdx = this.categories.indexOf(this.categories.find(el => el.name == this.category))
        if (categoryIdx == -1) return false
        // console.log("subject id = ", this.categories[categoryIdx].subject)

        const response = await axios.get("http://localhost:8080/api/v1/subjects", {
          params: {
            id: this.categories[categoryIdx].subject
          }
        })

        this.subjects = response.data
        this.subject = ''

        this.teachers = []
        this.teacher = ''
      },
      getTeachers: async function () {
        let subjectIdx = this.subjects.indexOf(this.subjects.find(el => el.name == this.subject))
        if (subjectIdx == -1) return false
        // console.log("teacher id = ", this.subjects[subjectIdx].id)
        const response = await axios.get("http://localhost:8080/api/v1/teachers", {
          params: {
            id: this.subjects[subjectIdx].id
          }
        })

        this.teachers = response.data
        this.teacher = ''
      },
      submit: async function () {
        let categoryIdx = this.categories.indexOf(this.categories.find(el => el.name == this.category))
        if (categoryIdx == -1) return console.error("Error: Categories are empty!");

        let subjectIdx = this.subjects.indexOf(this.subjects.find(el => el.name == this.subject))
        if (subjectIdx == -1) return console.error("Error: Subjects are empty!");

        let data = new FormData()

        data.append("author", 1)
        data.append("category_id", this.categories[categoryIdx].id)
        data.append("teacher_id", this.subjects[subjectIdx].id)

        // data.append("posted_at", Date.now())
        data.append("title", this.title)
        data.append("description", this.description)
        data.append("file", this.file)

        const response = await axios.post("http://localhost:8080/api/v1/notes", data)

        if (response.status == 200) {
          this.$router.replace({path: `/notes/${response.data.id}`})
        }
      }
    },
    created() {
      this.getCategories()
    }
  }

</script>

<style></style>