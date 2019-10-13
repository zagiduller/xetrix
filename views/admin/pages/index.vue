<template>
  <v-layout>
    <v-flex md10>
      <v-layout row wrap justify-space-between>
        <v-flex md2>
          Участники ({{users.length}})
        </v-flex>
        <v-flex md2>
          Ордера ({{orders.length}})
        </v-flex>
        <v-flex md2>
          Контракты ({{contracts.length}})
        </v-flex>
        <v-flex md2>
          Транзакции ({{txs.length}})
        </v-flex>
      </v-layout>
    </v-flex>
  </v-layout>
</template>

<script>
import { mapState } from 'vuex'
import Acc from '~/components/acc'
import OrderForm from '~/components/forms/orderForm'
import Order from '~/components/order'

export default {
    name: "index",
    components: { Acc, OrderForm, Order },
    data () {
      return {
        paysystems: []
      }
    },
    computed: {
      users() {
        return this.$store.state.objects.users
      },
      WsAccounts () {
        return this.$store.state.objects.WsAccounts
      },
      orders () {
        return this.$store.state.objects.orders
      },
      contracts () {
        return this.$store.state.objects.contracts
      },
      txs () {
        return this.$store.state.objects.txs
      },
    }
}
</script>

<style scoped>
    .m1 {
        margin-top: 20px;
    }
</style>