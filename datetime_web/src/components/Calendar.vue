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
      fields: ["first_name", "last_name", "age", "Address"],
      items: [
        {
          isActive: true,
          age: 40,
          first_name: "Dickerson",
          last_name: "Macdonald",
          Address: "123",
        },
        {
          isActive: false,
          age: 21,
          first_name: "Larsen",
          last_name: "Shaw",
          Address: "11111",
        },
        { isActive: false, age: 89, first_name: "Geneva", last_name: "Wilson" },
        { isActive: true, age: 38, first_name: "Jami", last_name: "Carney" },
      ],
    };
  },
};
</script>