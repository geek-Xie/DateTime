<template>
  <b-container fluid>
    <b-col md="10" offset-md="1" lg="10" offset-lg="1">
      <b-card title="Create New Event">
        <br />
        <b-form-group id="input-group-1" label="Tittle" label-for="input-1">
          <b-form-input
            id="input-1"
            v-model="newEvent.title"
            placeholder="Enter Title"
            required
          ></b-form-input>
        </b-form-group>

        <br />
        <b-form-group
          id="input-group-2"
          label="Describtion"
          label-for="input-2"
        >
          <b-form-input
            id="input-2"
            v-model="newEvent.description"
            placeholder="Enter Describtion"
            required
          ></b-form-input>
        </b-form-group>
        <br />
        <b-row>
          <b-col></b-col>
          <b-col></b-col>
          <b-col
            ><b-button @click="confirm" variant="outline-secondary"
              >Confirm</b-button
            ></b-col
          >
          <b-col
            ><b-button @click="back" variant="outline-secondary"
              >Back</b-button
            ></b-col
          >
          <b-col></b-col>
          <b-col></b-col>
        </b-row>
      </b-card>
    </b-col>
  </b-container>
</template>

<script>
export default {
  data() {
    return {
      newEvent: {
        useremail: "",
        title: "",
        description: "",
      },
    };
  },
  methods: {
    confirm() {
      let api = "http://localhost:9090/createEvent";
      this.newEvent.useremail = JSON.parse(localStorage.getItem("userInfo"))[
        "Email"
      ];
      this.axios.post(api, JSON.stringify(this.newEvent)).then((res) => {
        console.log(res);
      });
    },
    back() {
      this.$router.push("/");
    },
  },
};
</script>

<style>
</style>