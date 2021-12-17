<template>
  <b-row>
    <b-calendar
      block
      v-model="value"
      @context="onContext"
      locale="en-US"
    ></b-calendar>
  </b-row>
</template>

<script>
export default {
  data() {
    return {
      api: "http://localhost:9090/console",
      value: "",
      context: null,
      userInfo: JSON.parse(localStorage.getItem("userInfo"))["Email"],
    };
  },
  methods: {
    onContext(ctx) {
      this.context = ctx;

      const CtxInfo = {
        activeYMD: this.context["activeYMD"],
        userEmail: this.userInfo,
      };
      console.log("CtxInfo", CtxInfo);
      // console.log(this.context["activeYMD"]);

      this.axios.post(this.api, CtxInfo).then((res) => {
        console.log("res", res);
      });
    },
  },
};
</script>