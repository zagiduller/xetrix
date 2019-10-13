<template>
    <v-layout align-content-start mx-0 my-0 px-3 py-3 wrap>
        <v-flex wrap xs24 md8 lg8 pl-md-3 pr-md-3>
            <h1>Выберите валюту обмена</h1>
            <v-layout>
                <v-flex wrap xs24 md20 px-0>
                    <pair-checker :pair="pair" :accs="WsAccounts"></pair-checker>
                </v-flex>
            </v-layout>
        </v-flex>
        <v-flex wrap xs24 md8 lg8 pl-md-3 pr-md-3 v-if="$vuetify.breakpoint.mdAndUp">
            <info-exchange :history="historyBTC" classtable="data theme--light table-style style-1 info-exchange-table block-style-1"></info-exchange>
        </v-flex>
        <v-flex wrap xs24 md8 lg8 pl-md-3 pr-md-3 v-if="$vuetify.breakpoint.mdAndUp">
            <info-exchange :history="historyETH" classtable="data theme--light table-style style-1 info-exchange-table block-style-1"></info-exchange>
        </v-flex>
        <v-flex wrap xs24 md24 pl-md-3 pr-md-3>
            <h1>История</h1>
            <HistoryOrders classtable="data theme--light table-style style-1 buy-table top-panel-blocks block-style-1" scenario='sell' :orders="orders" ></HistoryOrders>
        </v-flex>
    </v-layout>
</template>

<script>
  import { mapState } from 'vuex'
  import Acc from '~/components/acc'
  import OrderForm from '~/components/forms/orderForm'
  import Order from '~/components/order'
  import PairChecker from '~/components/PairChecker'
  import InfoExchange from '~/components/InfoExchange'
  import HistoryOrders from '../components/HistoryOrders'

  export default {
    name: "index",
    components: {Acc, OrderForm, Order, HistoryOrders, PairChecker, InfoExchange },
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
      historyBTC () {
        return Object.values(this.$store.state.currency.history).filter(item => (item.element1 === 'BTC'))
      },
      historyETH () {
        return Object.values(this.$store.state.currency.history).filter(item => (item.element1 === 'ETH'))
      },
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
      },
      pair () {
        return this.$store.state.currency.pair
      },
    },
    methods: {
      accPS(accSymbol) {
        return this.$store.state.objects.paysystems.filter(ps => accSymbol === ps.Symbol)
      }
    },
    mounted: function () {
      this.$nextTick(() => {
        setTimeout(() => {
          this.$store.commit('timer/on')
        }, 1000)
      })
    },
    destroyed: function () {
      this.$store.commit('timer/off')
    },
  }
</script>

<style scoped>
    .m1 {
        margin-top: 20px;
    }
</style>