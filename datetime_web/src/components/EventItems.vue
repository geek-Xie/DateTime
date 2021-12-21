<template>
  <div>
    <b-table striped hover :items="items" :fields="fields"></b-table>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // Note `isActive` is left out and will not appear in the rendered table
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
  created() {
    // for (let i = 0; i < this.$route.query.data.length; i++) {
    //   console.log(this.$route.query.data[i]);
    // }
    for (var i = 0; i < this.$route.query.data["eventItems"].length; i++) {
      let item = {
        Title: this.$route.query.data["eventItems"][i]["Title"],
        Description: this.$route.query.data["eventItems"][i]["Description"],
        StartDate: this.$route.query.data["eventItems"][i]["StartDate"],
        StartTime: this.$route.query.data["eventItems"][i]["StartTime"],
        EndTime: this.$route.query.data["eventItems"][i]["EndTime"],
      };
      this.items[i] = item;
    }
  },
};
</script>