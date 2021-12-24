<template>
  <b-container>
    <b-row>
      <b-col md="10" offset-md="1" lg="10" offset-lg="1">
        <b-card v-if="!showDateTable" title="Personal Console">
          <b-calendar
            block
            v-model="value"
            @context="onContext"
            locale="en-US"
          ></b-calendar>
        </b-card>
        <b-table
          v-if="showDateTable"
          striped
          hover
          :items="items"
          :fields="fields"
        ></b-table>
        <b-button
          @click="clickBack"
          variant="outline-secondary"
          v-if="showDateTable"
          >Back</b-button
        >
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
// import DateTable from "./DateTable.vue";
export default {
  name: "Console",
  // components: { DateTable },
  data() {
    return {
      api: "http://localhost:9090/console",
      value: "",
      context: null,
      userInfo: JSON.parse(localStorage.getItem("userInfo"))["Email"],
      isFirstTime: false,

      showDateTable: false,
      fields: ["Title", "Description", "StartDate", "StartTime", "EndTime"],
      item: {
        Title: "",
        Description: "",
        StartDate: "",
        StartTime: "",
        EndTime: "",
      },
      items: [
        // {
        //   isActive: true,
        //   age: 40,
        //   first_name: "Dickerson",
        //   last_name: "Macdonald",
        // },
        // { isActive: false, age: 21, first_name: "Larsen", last_name: "Shaw" },
        // { isActive: false, age: 89, first_name: "Geneva", last_name: "Wilson" },
        // { isActive: true, age: 38, first_name: "Jami", last_name: "Carney" },
      ],
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
            // this.$router.push({
            //   path: "/eventitem",
            //   query: { data: res.data.data },
            // });
            for (var i = 0; i < res.data.data["eventItems"].length; i++) {
              let item = {
                Title: res.data.data["eventItems"][i]["Title"],
                Description: res.data.data["eventItems"][i]["Description"],
                StartDate: res.data.data["eventItems"][i]["StartDate"],
                StartTime: res.data.data["eventItems"][i]["StartTime"],
                EndTime: res.data.data["eventItems"][i]["EndTime"],
              };
              this.items[i] = item;
            }
            this.showDateTable = true;
            this.$store.commit("setDataTable", this.showDateTable);
            localStorage.setItem("items", JSON.stringify(this.items));
          } else {
            console.log(res.data.msg);
          }
        }
        this.isFirstTime = true;
      });
    },
    clickBack() {
      this.showDateTable = false;
      localStorage.removeItem("showDataTable");
      localStorage.removeItem("items");
    },
  },
  created() {
    if (localStorage.getItem("items") !== null) {
      for (
        var i = 0;
        i < JSON.parse(localStorage.getItem("items")).length;
        i++
      ) {
        let item = {
          Title: JSON.parse(localStorage.getItem("items"))[i]["Title"],
          Description: JSON.parse(localStorage.getItem("items"))[i][
            "Description"
          ],
          StartDate: JSON.parse(localStorage.getItem("items"))[i]["StartDate"],
          StartTime: JSON.parse(localStorage.getItem("items"))[i]["StartTime"],
          EndTime: JSON.parse(localStorage.getItem("items"))[i]["EndTime"],
        };
        this.items[i] = item;
      }
    }
    this.showDateTable = localStorage.getItem("showDataTable");
    // console.log("showDateTable", this.showDateTable);
  },
};
</script>

<style>
</style>