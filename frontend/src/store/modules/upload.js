import axios from "axios";

export default {
  actions: {
    async fetchCategories({commit}) {
      const response = await axios.get('http://localhost:8080/api/v1/categories')

      const categories = response.data
      commit("updateCategories", categories)
    },
    async fetchSubjects({commit, getters}, category) {
      commit("updateCategory", category)

      const categories = getters.allCategories
      const categoryIdx = getters.getCategoryIndex

      const response = await axios.get("http://localhost:8080/api/v1/subjects", {
        params: {
          id: categories[categoryIdx].subject
        }
      })

      const subjects = response.data
      commit("updateSubjects", subjects)
    },
    async fetchTeachers({commit, getters}, subject) {
      commit("updateSubject", subject)

      const subjects = getters.allSubjects
      const subjectIdx = getters.getSubjectIndex

      const response = await axios.get("http://localhost:8080/api/v1/teachers", {
        params: {
          id: subjects[subjectIdx].id
        }
      })

      const teachers = response.data
      commit("updateTeachers", teachers)
    }
  },
  mutations: {
    updateCategories(state, categories) {
      state.categories = categories
      state.subjects = state.teachers = []
      state.subject = state.teacher = ''
    },
    updateSubjects(state, subjects) {
      state.subjects = subjects
      state.teachers = []
      state.teacher = ''
    },
    updateTeachers(state, teachers) {
      state.teachers = teachers
    },
    updateCategory(state, category) {
      state.category = category
    },
    updateSubject(state, subject) {
      state.subject = subject
    },
    updateTeacher(state, teacher) {
      state.teacher = teacher
    },
  },
  state: {
    category: '',
    subject: '',
    teacher: '',

    categories: ["Выберите категорию"],
    subjects: ["Выберите предмет"],
    teachers: ["Выберите преподавателя"],

  },
  getters: {
    allCategories(state) {
      return state.categories
    },
    allSubjects(state) {
      return state.subjects
    },
    allTeachers(state) {
      return state.teachers
    },
    getCategoryIndex(state) {
      return state.categories.indexOf(state.categories.find(el => el.name === state.category))
    },
    getSubjectIndex(state) {
      return state.subjects.indexOf(state.subjects.find(el => el.name === state.subject))
    }
  }
}