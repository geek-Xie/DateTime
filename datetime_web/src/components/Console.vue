<template>
  <b-container>
    <b-row>
      <b-col md="10" offset-md="1" lg="10" offset-lg="1">
        <b-card title="Personal Console">
          <DateTable></DateTable>
        </b-card>
      </b-col>
    </b-row>
  </b-container>
</template>

<script>
import DateTable from "./DateTable.vue";
export default {
  name: "Console",
  components: { DateTable },
};
</script>

<style>
</style>