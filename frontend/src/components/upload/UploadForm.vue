<template>
  <v-container>
    <ValidationObserver ref="obs" v-slot="{ invalid, validated, passes, validate }">
      <v-form id="form" lazy-validation onsubmit="return false;">
        <span class="display-2 display-sm-3 text-center d-block mb-3">Загрузка работы</span>

        <VTextFieldWithValidation :counter="20" label="Название" rules="required|max:20" v-model="title"/>

        <v-textarea hint="Введите описание работы" label="Описание" name="descirption" v-model="description"/>

        <ValidationProvider rules="required" v-slot="{ errors }">
          <v-file-input accept="application/pdf" label="Файлы" name="file" v-model="file"/>
          <div class="v-text-field__details">
            <div class="v-messages theme--light error--text" role="alert">
              <div class="v-messages__wrapper">
                <div class="v-messages__message">{{ errors[0] }}</div>
              </div>
            </div>
          </div>
        </ValidationProvider>

        <VSelectWithValidation rules="required" :items="allCategories" @change="fetchData(category)"
          v-model="category" name="category" label="Категория" item-text="name"/>

        <VSelectWithValidation rules="required" :items="allSubjects" @change="updateSubject"
          v-model="subject" name="subject" label="Предмет" item-text="name"/>

        <VSelectWithValidation rules="required" :items="allTeachers" @change="updateTeacher"
          v-model="teacher" name="teacher" label="Преподаватель" item-text="name"/>

        <v-checkbox
            :rules="[v => !!v || 'Вы должны принять условия пользования!']"
            label="Я согласен с условиями пользования сервисом"
            required
            v-model="checkbox"/>

        <div class="d-flex justify-center">
          <v-btn
              :disabled="invalid" @click="submit" class="mx-4" color="success" large type="submit">
            Отправить
          </v-btn>
        </div>
      </v-form>
    </ValidationObserver>
  </v-container>
</template>

<script>
  import {ValidationObserver, ValidationProvider} from "vee-validate";
  import VTextFieldWithValidation from '../inputs/VTextFieldWithValidation';
  import VSelectWithValidation from '../inputs/VSelectWithValidation';

  import {mapActions, mapGetters} from 'vuex'

  export default {
    name: "UploadForm",
    components: {
      ValidationObserver,
      ValidationProvider,
      VTextFieldWithValidation,
      VSelectWithValidation
    },
    data: () => ({
      title: '',
      description: '',
      category: '',
      subject: '',
      teacher: '',

      file: null,
      fileRules: [
        v => !!v || "Поле должно быть заполнено"
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

        const response = await this.$http.post("http://localhost:8080/api/v1/notes", data)

        if (response.status === 200) {
          await this.$router.replace({path: `/notes/${response.data.id}`})
        }
      }
    },
    computed: {
      ...mapGetters(["allCategories", "allSubjects", "allTeachers", "getCategoryIndex", "getSubjectIndex", "getTeacherIndex"]),
    },
    async mounted() {
      await this.fetchCategories()
    }
  }

</script>

<style></style>