<template>
  <div>
    <h1 v-if="this.$store.getters.getLogin">这是Home， 您已登陆</h1>
    <h1 v-if="!this.$store.getters.getLogin">这是Home， 您未登陆</h1>
    <h1>这是Home</h1>
    <button @click="lo()">123123</button>
  </div>
</template>

<script>
// @ is an alias to /src
// import RegisterForm from "@/components/RegisterForm.vue";

export default {
  name: "Home",
  components: {
    // RegisterForm,
  },
  methods: {
    lo() {
      console.log(this.$store.getters.getLogin);
    },
  },
};
</script>
