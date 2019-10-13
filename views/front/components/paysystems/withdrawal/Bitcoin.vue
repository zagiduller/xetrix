<template>
  <v-flex md12>
    <v-text-field
            label="Bitcoin address"
    ></v-text-field>
  </v-flex>
</template>

<script>
  export default {
    name: "Bitcoin",
    props: ['acc'],
    data () {
      return {
        amount: 0,
        sendingAddress: '',
        fullName: '',
        attributes: [],
      }
    },
    methods: {
      send () {
        this.$axios.post(process.env.apiUrl + '_v1/create_withdrawal')
      }
    },
  }
</script>

<style scoped>

</style>