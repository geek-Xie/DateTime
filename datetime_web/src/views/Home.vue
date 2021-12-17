<template>
  <div>
    <Calendar></Calendar>
  </div>
</template>

<script>
// @ is an alias to /src
// import RegisterForm from "@/components/RegisterForm.vue";
import Calendar from "@/components/Calendar.vue";

export default {
  name: "Home",
  components: { Calendar },
  created() {
    if (!this.$store.getters.getLogin) {
      this.$router.push("/login");
    }
  },
};
</script>
