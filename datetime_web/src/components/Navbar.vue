<template>
  <div>
    <b-navbar toggleable="lg" type="dark" variant="dark" fixed="top">
      <b-container>
        <b-navbar-brand href="/">DateTime</b-navbar-brand>

        <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

        <b-collapse id="nav-collapse" is-nav>
          <!-- <b-navbar-nav>
            <b-nav-item href="#">Link</b-nav-item>
            <b-nav-item href="#" disabled>Disabled</b-nav-item>
          </b-navbar-nav> -->

          <!-- Right aligned nav items -->

          <b-navbar-nav class="ml-auto">
            <div>
              <!-- <b-nav-form>
                <b-form-input
                  size="sm"
                  class="mr-sm-2"
                  placeholder="Search"
                ></b-form-input>
                <b-button size="sm" class="my-2 my-sm-0" type="submit"
                  >Search</b-button
                >
              </b-nav-form> -->
            </div>

            <!-- <b-nav-item-dropdown text="Lang" right>
              <b-dropdown-item href="#">EN</b-dropdown-item>
              <b-dropdown-item href="#">ES</b-dropdown-item>
              <b-dropdown-item href="#">RU</b-dropdown-item>
              <b-dropdown-item href="#">FA</b-dropdown-item>
            </b-nav-item-dropdown> -->

            <!-- <b-nav-item-dropdown right> -->
            <!-- Using 'button-content' slot -->
            <!-- <template #button-content>
                <em>User</em>
              </template>
              <b-dropdown-item href="#">Profile</b-dropdown-item>
              <b-dropdown-item href="#">Sign Out</b-dropdown-item>
            </b-nav-item-dropdown> -->
            <b-nav-item @click="$router.push('/')">Home Page</b-nav-item>
            <b-nav-item
              v-if="!this.isLogin()"
              @click="$router.replace({ name: 'Register' })"
              >Register</b-nav-item
            >
            <b-nav-item
              v-if="!this.isLogin()"
              @click="$router.replace({ name: 'Login' })"
              >Login</b-nav-item
            >
            <b-nav-item
              v-if="this.isLogin()"
              @click="$router.replace({ name: 'Console' })"
              >Personal Console</b-nav-item
            >
            <b-nav-item v-if="this.isLogin()" @click="logout"
              >Logout</b-nav-item
            >
          </b-navbar-nav>
        </b-collapse>
      </b-container>
    </b-navbar>
  </div>
</template>

<script>
export default {
  methods: {
    isLogin() {
      return this.$store.getters.getLogin;
    },
    logout() {
      window.localStorage.clear();
      if (localStorage.getItem("userInfo") === null) {
        console.log("localstorage成功删除");
      } else {
        console.log("localstorage存在");
      }
      this.$router.replace({ name: "Login" });
      this.$router.go(0);
    },
  },
};
</script>

<style>
</style>