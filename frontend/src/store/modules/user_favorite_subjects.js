import axios from "axios"

export default {
  actions: {
    async fetchUserFavoriteSubjects({commit}, id) {
      const response = await axios.get(`http://localhost:8080/api/v1/userFavoriteSubjects/${id}`)
      const subjects = response.data

      commit("updateUserFavoriteSubjects", subjects)
    }
  },
  mutations: {
    updateUserFavoriteSubjects(state, subjects) {
      state.userFavoriteSubjects = subjects
    }
  },
  state: {
    userFavoriteSubjects: []
  },
  getters: {
    getUserFavoriteSubjects: state => state.userFavoriteSubjects,
  }
}
