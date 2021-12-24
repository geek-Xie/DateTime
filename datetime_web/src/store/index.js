import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  // state: localStorage.getItem('state') ? JSON.parse(localStorage.getItem('state')) : {
  state: {
    isLogin: false,
    showDataTable: false,
    userInfo: {
      email: '',
      username: '',
      phone: '',
      token: ''
    }
  },
  mutations: {
    setUserInfo(state, user) {
      state.userInfo = user
      localStorage.setItem('userInfo', JSON.stringify(user))
      state.isLogin = true
    },
    setEmail(state, email) {
      state.userInfo.email = email
    },
    setUsername(state, username) {
      state.userInfo.username = username
    },
    setPhone(state, phone) {
      state.userInfo.phone = phone
    },
    setToken(state, token) {
      state.userInfo.token = token
    },
    setLogin(state, status) {
      state.isLogin = status
    },
    setDataTable(state, showDataTable) {
      state.showDataTable = showDataTable
      localStorage.setItem('showDataTable', showDataTable)
    },
  },
  getters: {
    getLogin: state => state.isLogin,
    getUserInfo: state => state.userInfo,
    getUserToken: state => state.userInfo['Token'],
    getDataTable: status => status.showDataTable
  },

  actions: {

  }
})
