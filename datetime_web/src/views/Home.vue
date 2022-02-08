<template>
  <div>
    <Calendar></Calendar>
    <b-button @click="createEvent" variant="outline-secondary"
      >Create Event</b-button
    >
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
  methods: {
    createEvent() {
      this.$router.push("/createEvent");
      console.log("createEvent");
    },
  },
};
</script>
