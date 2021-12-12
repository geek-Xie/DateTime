import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  // state: localStorage.getItem('state') ? JSON.parse(localStorage.getItem('state')) : {
  state: {
    isLogin: false,
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
      localStorage.userInfo = user
      state.isLogin = true
    },
    setLogin(state, status) {
      state.isLogin = status
    }
  },
  getters: {
    getLogin: state => state.isLogin,
    getUserInfo: state => state.userInfo,
    getUserToken: state => state.userInfo['Token'],
  },

  actions: {

  }
})
