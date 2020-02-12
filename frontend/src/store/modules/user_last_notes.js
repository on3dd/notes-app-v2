import axios from "axios"

const options = {
  year: 'numeric',
  month: 'numeric',
  day: 'numeric',
  timezone: 'UTC',
  hour: 'numeric',
  minute: 'numeric'
};

export default {
  actions: {
    async fetchUserLastNotes({commit}, id) {
      const response = await axios.get(`http://localhost:8080/api/v1/userLastNotes/${id}`)
      const notes = response.data

      commit("updateUserLastNotes", notes)
    }
  },
  mutations: {
    updateUserLastNotes(state, notes) {
      notes.forEach(note => {
        const timestamp = Date.parse(note.posted_at)
        note.posted_at = new Date(timestamp).toLocaleString("ru", options)
      })

      state.userLastNotes = notes
    }
  },
  state: {
    userLastNotes: []
  },
  getters: {
    getUserLastNotes: state => state.userLastNotes,
  }
}
