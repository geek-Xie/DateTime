<template>
  <div>
    <b-container>
      <b-col md="10" offset-md="2" lg="10" offset-lg="1">
        <b-table striped hover :items="items" :fields="fields"></b-table>
      </b-col>
    </b-container>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // Note `isActive` is left out and will not appear in the rendered table
      fields: ["Title", "Description", "Start_Date", "Start_Time", "L_T"],
      item: {
        Title: "",
        Describtion: "",
        Start_Date: "",
        Start_Time: "",
        L_T: "",
      },
      items: [],
    };
  },
  methods: {
    getEvents() {
      let api = "http://localhost:9090/getEvents";
      let useremail = {
        useremail: "",
      };
      // let useremail = JSON.parse(localStorage.getItem("userInfo"))["Email"];
      useremail.useremail = JSON.parse(localStorage.getItem("userInfo"))[
        "Email"
      ];
      this.axios.post(api, JSON.stringify(useremail)).then((res) => {
        for (var i = 0; i < res.data.data["events"].length; i++) {
          let item = {
            Title: res.data.data["events"][i]["Title"],
            Description: res.data.data["events"][i]["Description"],
            Start_Date: res.data.data["events"][i]["StartDate"],
            Start_Time: res.data.data["events"][i]["StartTime"],
            L_T: res.data.data["events"][i]["Description"],
          };
          this.items[i] = item;
        }
        localStorage.setItem("events", JSON.stringify(this.items));
      });
    },
  },
  created() {
    this.getEvents();
    var events = JSON.parse(localStorage.getItem("events"));
    for (var i = 0; i < events.length; i++) {
      let item = {
        Title: events[i]["Title"],
        Description: events[i]["Description"],
        Start_Date: events[i]["Start_Date"],
        Start_Time: events[i]["Start_Time"],
        L_T: events[i]["Description"],
      };
      this.items[i] = item;
    }
  },
};
</script>