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
      isFirstTime: false,
    };
  },
  methods: {
    onContext(ctx) {
      this.context = ctx;

      const CtxInfo = {
        activeYMD: this.context["activeYMD"],
        userEmail: this.userInfo,
      };
      this.axios.post(this.api, CtxInfo).then((res) => {
        if (this.isFirstTime) {
          if (res.data.code == 200) {
            this.$router.push({
              path: "/eventitem",
              query: { data: res.data.data },
            });
          } else {
            console.log(res.data.msg);
          }
        }
        this.isFirstTime = true;
      });
    },
  },
};
</script>