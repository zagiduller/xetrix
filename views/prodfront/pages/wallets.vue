<template>
    <v-layout align-content-start mx-0 my-0 px-3 py-3 wrap>
        <v-dialog max-width="290">
            <v-card>
                <v-card-title>Вывода нет</v-card-title>
                <v-card-actions>
                    <v-spacer></v-spacer>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <v-flex xs24 md24 lg20 pl-3>
            <v-layout wrap justify-xs-start>
                <v-flex xs24>
                    <h1>Выберите кошелек:</h1>
                </v-flex>
                <v-flex sm12 md6 lg6 mt-xs-4 mt-md-0 px-3 v-for="acc in WsAccounts" :key="acc.account.id">
                    <Acc :acc="acc" :ps="accPS(acc.account.currency.symbol)"></Acc>
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
    fetch ({ store, app: {$axios}}) {
      return $axios.get(process.env.apiUrl + 'payments/systems')
        .then((res) => {
          store.commit('objects/setPaysystems', res.data)
        })
    },
    layout: 'base',
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