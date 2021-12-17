import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import RegisterForm from '../components/auth/RegisterForm.vue'
import LoginForm from '../components/auth/LoginForm.vue'
import Console from '../components/Console.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterForm
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginForm
  },
  {
    path: '/console',
    name: 'Console',
    component: Console
  }
]

const router = new VueRouter({
  routes
})


export default router
