<template>
  <div>
    <RegisterForm></RegisterForm>
  </div>
</template>

<script>
// @ is an alias to /src
import RegisterForm from "@/components/RegisterForm.vue";

export default {
  name: "Home",
  components: {
    RegisterForm,
  },
};
</script>
