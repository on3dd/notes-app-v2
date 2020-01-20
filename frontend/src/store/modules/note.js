import axios from "axios";

export default {
  actions: {
    async fetchNote({commit, dispatch}, id) {
      const response = await axios.get(`http://localhost:8080/api/v1/notes/${id}`)

      const note = response.data
      commit("updateNote", note)
      dispatch("fetchAuthor", note.author_id)
      dispatch("fetchCategory", note.author_id)
      dispatch("fetchTeacher", note.teacher_id)
    },
    async fetchAuthor({commit}, id) {
      const response = await axios.get(`http://localhost:8080/api/v1/users/${id}`)

      const author = response.data
      commit("updateAuthor", author)
    },
    async fetchCategory({commit, dispatch}, id) {
      const response = await axios.get(`http://localhost:8080/api/v1/categories/${id}`)

      const category = response.data
      commit("updateCategory", category)
      dispatch("fetchSubject", category.subject)
    },
    async fetchSubject({commit}, id) {
      const response = await axios.get(`http://localhost:8080/api/v1/subjects/${id}`)

      const subject = response.data
      commit("updateSubject", subject)
    },
    async fetchTeacher({commit}, id) {
      const response = await axios.get(`http://localhost:8080/api/v1/teachers/${id}`)

      const teacher = response.data
      commit("updateTeacher", teacher)
    },
  },
  mutations: {
    updateNote(state, note) {
      state.note = note
    },
    updateAuthor(state, author) {
      state.author = author
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
    note: {
      title: "Неизвестное название",
      description: "Неизвестное описание",
      posted_at: "1 января 2000 г., 00:00:00",
    },
    author: {
      name: "Неизвестный пользователь"
    },
    category: "",
    subject: {
      name: "Неизвестный предмет"
    },
    teacher: {
      name: "Неизвестный преподаватель"
    },
  },
  getters: {
    getNote: state => state.note,
    getAuthor: state => state.author,
    getCategory: state => state.category,
    getSubject: state => state.subjects,
    getTeacher: state => state.teacher
  }
}