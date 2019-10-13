<template>
  <v-container fluid>
    <v-layout row wrap justify-space-between>
      <v-flex md2 sm6 xs12>
        <v-layout row class="mt-3" v-for="acc in WsAccounts" :key="acc.account.id">
          <v-flex md10 >
            <Acc :acc="acc" :ps="accPS(acc.account.currency.symbol)"></Acc>
          </v-flex>
        </v-layout>
      </v-flex>
      <v-flex md5>
        <v-layout row >
          <v-flex md12>

          </v-flex>
        </v-layout>
        <v-layout row class="mt-3" v-for="o in orders" :key="o.id" justify-end>
          <v-flex md10 >
            <Order :order="o"></Order>
          </v-flex>
        </v-layout>
      </v-flex>
      <v-flex md4>
        <v-layout column wrap class="mt-3" justify-end>
          <v-flex md10>
            <OrderForm :accs="WsAccounts"></OrderForm>
          </v-flex>
          <v-flex class="mt-3" v-for="o in userOrders" :key="'my' + o.id">
            <Order :order="o"></Order>
          </v-flex>
        </v-layout>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import { mapState } from 'vuex'
import Acc from '~/components/acc'
import OrderForm from '~/components/forms/orderForm'
import Order from '~/components/order'

export default {
    name: "index",
    components: { Acc, OrderForm, Order },
    fetch ({ store, app: {$axios}}) {
      return $axios.get(process.env.apiUrl + 'payments/systems')
        .then((res) => {
          store.commit('objects/setPaysystems', res.data)
        })
    },
    data () {
      return {
        isBordered: true,
      }
    },
    computed: {
      WsAccounts () {
        return this.$store.state.objects.WsAccounts
      },
      orders () {
        return this.$store.state.objects.orders
      },
      userOrders() {
        return this.orders.filter((o)=>{
          return o.ownerId === this.$auth.user.id
        })
      }
    },
    methods: {
      accPS(accSymbol) {
        return this.$store.state.objects.paysystems.filter(ps => accSymbol === ps.Symbol)
      }
    },
}
</script>

<style scoped>
    .m1 {
        margin-top: 20px;
    }
</style>