import axios from "axios"

export default {
  actions: {
    async fetchUser({commit}, id) {
      const response = await axios.get(`http://localhost:8080/api/v1/users/${id}`)
      const user = response.data

      commit("updateUser", user)
    }
  },
  mutations: {
    updateUser(state, user) {
      state.user = user
    }
  },
  state: {
    user: {}
  },
  getters: {
    getUser: state => state.user,
  }
}
