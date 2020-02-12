import axios from "axios";

export default {
  actions: {
    async login({commit}, {username, password}) {
      return new Promise((resolve, reject) => {
        commit('auth_request')

        const data = new FormData()
        data.append("name", username)
        data.append("password", password)

        axios.post("http://localhost:8080/api/v1/login", data)
            .then(response => {
              const token = response.data.token
              const user = response.data.user
              console.log(user)

              localStorage.setItem('token', token)
              localStorage.setItem('user_name', user.name)
              localStorage.setItem('user_id', user.id)

              axios.defaults.headers.common['Authorization'] = token

              commit('auth_success', {token, user})
              resolve(response)
            })
            .catch(err => {
              commit('auth_error')
              localStorage.removeItem('token')
              reject(err)
            })
      })
    },
    register({commit}, {username, email, about, password}){
      return new Promise((resolve, reject) => {
        commit('auth_request')

        const data = new FormData()
        data.append("name", username)
        data.append("email", email)
        data.append("about", about)
        data.append("password", password)

        axios.post('http://localhost:8080/api/v1/register', data)
            .then(response => {
              const token = response.data.token
              const user = response.data.user

              localStorage.setItem('token', token)
              localStorage.setItem('user_name', user.name)
              localStorage.setItem('user_id', user.id)

              axios.defaults.headers.common['Authorization'] = token

              commit('auth_success', {token: token, user: user})
              resolve(response)
            })
            .catch(err => {
              commit('auth_error', err)
              localStorage.removeItem('token')
              reject(err)
            })
      })
    },
    logout({commit}){
      return new Promise((resolve, reject) => {
        commit('logout')

        localStorage.removeItem('token')
        localStorage.removeItem('user_name')
        localStorage.removeItem('user_id')

        delete axios.defaults.headers.common['Authorization']

        resolve()
      })
    }

  },
  mutations: {
    auth_request(state){
      state.status = 'loading'
    },
    auth_success(state, {token, user}){
      state.status = 'success'
      state.token = token
      state.user = user
    },
    auth_error(state){
      state.status = 'error'
    },
    logout(state){
      state.status = ''
      state.token = ''
    },
  },
  state: {
    status: '',
    token: localStorage.getItem('token') || '',
    user : {}
  },
  getters: {
    isLoggedIn: state => !!state.token,
    authStatus: state => state.status,
    getCurrentUser: state => state.user
  }
}