<template>
  <div id="app">
    <Navbar />
    <br />
    <br />
    <br />
    <br />
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
  created() {
    // window.addEventListener("unload", this.saveState);
    if (localStorage.getItem("userInfo") === null) {
      // this.$store.commit.setLogin();
      this.$store.commit("setLogin", false);
      console.log("isLogin", this.$store.getters.getLogin);
    } else {
      this.$store.commit("setLogin", true);
      this.$store.commit(
        "setEmail",
        JSON.parse(localStorage.getItem("userInfo"))["Email"]
      );
      this.$store.commit(
        "setUsername",
        JSON.parse(localStorage.getItem("userInfo"))["Username"]
      );
      this.$store.commit(
        "setPhone",
        JSON.parse(localStorage.getItem("userInfo"))["Phone"]
      );
      this.$store.commit(
        "setToken",
        JSON.parse(localStorage.getItem("userInfo"))["Token"]
      );
      console.log("isLogin", this.$store.getters.getLogin);
    }
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
