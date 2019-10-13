<template>
  
</template>

<script>
  import Order from '~/components/order'
  export default {
    name: "orders",
    components: [Order],
    data () {
      return {

      }
    },
    computed: {
      orders () {
        return this.$store.state.objects.orders
      }
    },
  }
</script>

<style scoped>

</style>