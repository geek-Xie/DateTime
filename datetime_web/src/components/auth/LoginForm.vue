<template>
  <div class="login">
    <b-row class="mt-50">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="Login">
          <b-form @submit="onSubmit" @reset="onReset" v-if="show">
            <b-form-group
              id="input-group-1"
              label="Email address:"
              label-for="input-1"
            >
              <b-form-input
                id="input-1"
                v-model="loginForm.email"
                type="email"
                placeholder="Enter email address"
                required
              ></b-form-input>
            </b-form-group>

            <b-form-group
              id="input-group-2"
              label="Your Password:"
              label-for="input-2"
            >
              <b-form-input
                id="input-2"
                v-model="loginForm.password"
                type="password"
                placeholder="Enter password"
                required
              ></b-form-input>
            </b-form-group>

            <!-- <b-button type="submit" variant="primary">Submit</b-button>
            <b-button type="reset" variant="danger">Reset</b-button> -->
            <br />

            <b-alert v-if="loginerror" show variant="primary">{{
              this.alertText
            }}</b-alert>
            <!-- <b-button @click="showAlert" variant="info" class="m-1">
              Show alert with count-down timer
            </b-button> -->

            <b-button
              variant="outline-secondary"
              @click="login"
              @keyup.enter="enterlogin"
              >Login</b-button
            >
          </b-form>
        </b-card>
      </b-col>
    </b-row>
    <!-- <b-card class="mt-3" header="Form Data Result">
      <pre class="m-0">{{ form }}</pre>
    </b-card> -->
  </div>
</template>

<script>
import { required } from "vuelidate/lib/validators";
export default {
  data() {
    return {
      loginForm: {
        email: "",
        password: "",
      },
      show: true,
      loginerror: false,
      alertText: "",
    };
  },
  validations: {
    loginForm: {
      email: {
        required,
      },
      password: {
        required,
      },
    },
  },
  created() {
    this.enterlogin();
  },
  methods: {
    countDownChanged(dismissCountDown) {
      this.dismissCountDown = dismissCountDown;
    },
    showAlert() {
      this.dismissCountDown = this.dismissSecs;
    },
    login() {
      this.$v.loginForm.$touch();
      if (this.$v.loginForm.$anyError) {
        return;
      }
      const api = "http://localhost:9090/auth/login";
      this.axios.post(api, this.loginForm).then((res) => {
        if (res.data.code == 200) {
          console.log(res.data.msg);
          this.$store.commit("setUserInfo", res.data.data["loginUser"]);
          this.$router.push("/");
        } else if (res.data.code == 400) {
          console.log(res.data.msg);
          this.alertText = res.data.msg;
          this.loginerror = true;
          this.$router.go(0);
        }
      });
    },
    enterlogin() {
      document.onkeydown = (e) => {
        if (e.keyCode == 13) {
          this.login();
        }
      };
    },
    onSubmit(event) {
      event.preventDefault();
      alert(JSON.stringify(this.loginForm));
    },
    onReset(event) {
      event.preventDefault();
      // Reset our form values
      this.loginForm.email = "";
      // Trick to reset/clear native browser form validation state
      this.show = false;
      this.$nextTick(() => {
        this.show = true;
      });
    },
  },
};
</script>

<style>
</style>