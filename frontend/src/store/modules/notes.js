import axios from "axios";

const options = {
  year: 'numeric',
  month: 'numeric',
  day: 'numeric',
  timezone: 'UTC',
  hour: 'numeric',
  minute: 'numeric',
  second: 'numeric'
};

export default {
  actions: {
    async fetchNotes({commit}) {
      const response = await axios.get("http://localhost:8080/api/v1/notes")
      const notes = response.data

      commit("updateNotes", notes)
    }
  },
  mutations: {
    updateNotes(state, notes) {
      notes.forEach(note => {
        const timestamp = Date.parse(note.posted_at)
        note.posted_at = new Date(timestamp).toLocaleString("ru", options)
      })

      state.notes = notes
    }
  },
  state: {
    notes: []
  },
  getters: {
    getNotes: state => state.notes,
  }
}