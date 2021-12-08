import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
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
    }
  },
  getters: {
    getUserInfo(state) {
      return state.userInfo
    },
    getUserToken(state) {
      return state.userInfo['Token']
    }
  },

  actions: {

  }
})
