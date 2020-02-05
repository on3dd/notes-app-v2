import axios from "axios";

export default {
  actions: {
    async fetchUsers({commit}) {
      const response = await axios.get("http://localhost:8080/api/v1/users")
      const users = response.data

      commit("updateUsers", users)
    }
  },
  mutations: {
    updateUsers(state, users) {
      state.users = users
    }
  },
  state: {
    users: []
  },
  getters: {
    getUsers: state => state.users,
  }
}