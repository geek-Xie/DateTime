import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import RegisterForm from '../components/auth/RegisterForm.vue'
import LoginForm from '../components/auth/LoginForm.vue'
import Console from '../components/Console.vue'
import EventItems from '../components/EventItems.vue'
import CreateEvent from '../components/CreateEvent.vue'

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
  },
  {
    path: '/eventitem',
    name: 'EventItems',
    component: EventItems
  },
  {
    path: '/createEvent',
    name: 'CreateEvent',
    component: CreateEvent
  }
]

const router = new VueRouter({
  routes
})


export default router
