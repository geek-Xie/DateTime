<template>
  <div id="app">
    <Navbar />
    <router-view></router-view>
  </div>
</template>

<script>
// import HelloWorld from "./components/HelloWorld.vue";
import Navbar from "@/components/Navbar.vue";

export default {
  name: "app",
  components: {
    Navbar,
  },
  mounted() {
    window.addEventListener("unload", this.saveState);
  },
  methods: {
    saveState() {
      localStorage.setItem("state", JSON.stringify(this.$store.state));
    },
  },
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
