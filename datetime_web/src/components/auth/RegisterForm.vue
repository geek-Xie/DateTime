<template>
  <div class="register">
    <b-row class="mt-50">
      <b-col md="8" offset-md="2" lg="6" offset-lg="3">
        <b-card title="Register">
          <b-form @submit="onSubmit" @reset="onReset" v-if="show">
            <b-form-group
              id="input-group-1"
              label="Email address:"
              label-for="input-1"
            >
              <b-form-input
                id="input-1"
                v-model="$v.registerForm.email.$model"
                type="email"
                placeholder="Enter email address"
                required
              ></b-form-input>
            </b-form-group>

            <b-form-group
              id="input-group-2"
              label="Your Name:"
              label-for="input-2"
            >
              <b-form-input
                id="input-2"
                v-model="$v.registerForm.username.$model"
                placeholder="Enter name"
                required
              ></b-form-input>
            </b-form-group>

            <b-form-group
              id="input-group-3"
              label="Your Phone Number:"
              label-for="input-3"
            >
              <b-form-input
                id="input-3"
                v-model="$v.registerForm.phone.$model"
                placeholder="Enter phone number"
                required
              ></b-form-input>
              <b-form-invalid-feedback :state="validateState('phone')">
                Your mobile phone number must be 11 digits
              </b-form-invalid-feedback>
            </b-form-group>

            <b-form-group
              id="input-group-4"
              label="Your Password:"
              label-for="input-4"
            >
              <b-form-input
                id="input-3"
                v-model="$v.registerForm.password.$model"
                type="password"
                placeholder="Enter password"
                required
              ></b-form-input>
            </b-form-group>

            <!-- <b-form-group id="input-group-3" label="Food:" label-for="input-3">
          <b-form-select
            id="input-3"
            v-model="form.food"
            :options="foods"
            required
          ></b-form-select>
        </b-form-group> -->

            <!-- <b-form-group id="input-group-4" v-slot="{ ariaDescribedby }">
          <b-form-checkbox-group
            v-model="form.checked"
            id="checkboxes-4"
            :aria-describedby="ariaDescribedby"
          >
            <b-form-checkbox value="me">Check me out</b-form-checkbox>
            <b-form-checkbox value="that">Check that out</b-form-checkbox>
          </b-form-checkbox-group>
        </b-form-group> -->
            <br />
            <b-alert v-if="registererror" show variant="primary">{{
              this.alertText
            }}</b-alert>

            <b-button
              variant="outline-secondary"
              @click="register"
              @keyup.enter="enterregister"
              >Register</b-button
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
import { required, minLength, maxLength } from "vuelidate/lib/validators";
export default {
  data() {
    return {
      registerForm: {
        email: "",
        username: "",
        phone: "",
        password: "",
      },
      show: true,
      alertText: "",
      registererror: false,
    };
  },
  validations: {
    registerForm: {
      email: {
        required,
      },
      username: {
        required,
      },
      phone: {
        required,
        minLength: minLength(11),
        maxLength: maxLength(11),
      },
      password: {
        required,
      },
    },
  },
  created() {
    this.enterregister();
  },
  methods: {
    validateState(name) {
      const { $dirty, $error } = this.$v.registerForm[name];
      return $dirty ? !$error : null;
    },

    register() {
      this.$v.registerForm.$touch();
      if (this.$v.registerForm.$anyError) {
        return;
      }
      const api = "http://localhost:9090/auth/register";
      this.axios.post(api, this.registerForm).then((res) => {
        if (res.data.code == 200) {
          this.$store.commit("setUserInfo", res.data.data["loginUser"]);
          this.$router.push("/");
        } else if (res.data.code == 400) {
          this.alertText = res.data.msg;
          this.registererror = true;
          this.this.$router.go(0);
        }
      });
    },
    enterregister() {
      document.onkeydown = (e) => {
        if (e.keyCode == 13) {
          this.register();
        }
      };
    },
    onSubmit(event) {
      event.preventDefault();
      alert(JSON.stringify(this.registerForm));
    },
    onReset(event) {
      event.preventDefault();
      // Reset our form values
      this.registerForm.email = "";
      this.registerForm.username = "";
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