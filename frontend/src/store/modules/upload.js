import axios from "axios";

export default {
  actions: {
    async fetchCategories({commit}) {
      const response = await axios.get('http://localhost:8080/api/v1/categories')

      const categories = response.data
      commit("updateCategories", categories)
    },
    async fetchData({commit, dispatch}, category) {
      commit("updateCategory", category)

      dispatch("fetchSubjects")
      dispatch("fetchTeachers")
    },
    async fetchSubjects({commit, getters}) {
      // commit("updateCategory", category)

      const categories = getters.allCategories
      const categoryIdx = getters.getCategoryIndex

      const response = await axios.get("http://localhost:8080/api/v1/subjects", {
        params: { category_id: categories[categoryIdx].id }
      })

      const subjects = response.data
      commit("updateSubjects", subjects)
    },
    async fetchTeachers({commit, getters}) {
      // commit("updateCategory", category)

      const categories = getters.allCategories
      const categoryIdx = getters.getCategoryIndex

      const response = await axios.get("http://localhost:8080/api/v1/teachers", {
        params: { category_id: categories[categoryIdx].id }
      })

      const teachers = response.data
      commit("updateTeachers", teachers)
    },
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
    allCategories: state => state.categories,
    allSubjects: state => state.subjects,
    allTeachers: state => state.teachers,

    getCategoryIndex(state) {
      return state.categories.indexOf(state.categories.find(el => el.name === state.category))
    },
    getSubjectIndex(state) {
      return state.subjects.indexOf(state.subjects.find(el => el.name === state.subject))
    },
    getTeacherIndex(state) {
      return state.teachers.indexOf(state.teachers.find(el => el.name === state.teacher))
    }
  }
}