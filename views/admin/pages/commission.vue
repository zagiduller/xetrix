<template>
  <v-container grid-list-lg class="pt-0">
    <v-layout row wrap >
      <v-flex class="mb-3" xs4 v-for="(a, i) in systemAccounts" :key="i">
        <Acc :acc="a"></Acc>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
  import Acc from '~/components/acc'
  export default {
    name: "commission",
    components: {Acc},
    data () {
      return {

      }
    },
    computed: {
      systemUserId() {
        return this.$auth.user.id
      },
      wsAccs() {
        return this.$store.state.objects.WsAccounts
      },
      systemAccounts() {
        return this.wsAccs.filter((a) => {
          return a.account.ownerId === this.systemUserId
        })
      }
    },
  }
</script>

<style scoped>

</style>