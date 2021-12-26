<template>
  <div>
    <Calendar></Calendar>
    <b-button @click="createItem" variant="outline-secondary"
      >Create Item</b-button
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
    createItem() {
      this.$router.push("/createItem");
      console.log("createItem");
    },
  },
};
</script>
