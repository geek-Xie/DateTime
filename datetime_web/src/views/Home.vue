<template>
  <div>
    <Navbar></Navbar>
    <AuthForm></AuthForm>
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
import Navbar from "@/components/Navbar.vue";
import AuthForm from "@/components/AuthForm.vue";

export default {
  name: "Home",
  components: {
    Navbar,
    AuthForm,
  },
};
</script>
