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
    <v-flex wrap xs24 md8 lg8 pl-md-3 pr-md-3>
      <h1>Продавцы {{pair.from.account ? pair.from.account.currency.symbol : ''}}</h1>
     <order-list classtable="data theme--light table-style style-1 buy-table top-panel-blocks block-style-1" :pair="pair" scenario='sell' :orders="pairFilteredItems(pair)" ></order-list>

    </v-flex>
    <v-flex wrap xs24 md8 lg8 pl-md-3 pr-md-3>
      <h1>Покупатели {{pair.from.account ? pair.from.account.currency.symbol : ''}}</h1>
      <order-list classtable="data theme--light table-style style-1 buy-table top-panel-blocks block-style-1" :pair="pair" scenario='buy' :orders="invertedPairFilteredItems(pair)"></order-list>
   </v-flex>
    <v-flex wrap xs24 md8 lg8 pl-md-3 pr-md-3>
      <h1>Создать ордер</h1>
      <OrderForm :pair="pair" :accs="WsAccounts"></OrderForm>
    </v-flex>
  </v-layout>
</template>

<script>
import { mapState } from 'vuex'
import Acc from '~/components/acc'
import OrderForm from '~/components/forms/orderForm'
import Order from '~/components/order'
import OrderList from '~/components/orderlist'
import PairChecker from '~/components/PairChecker'
import InfoExchange from '~/components/InfoExchange'

export default {
    name: "index",
    components: { Acc, OrderForm, Order, OrderList, PairChecker, InfoExchange },
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
      pairFilteredItems () {
        const ordersForBuy = this.$store.getters['objects/ordersForBuy']
        if (ordersForBuy()) {
          return ordersForBuy
        } else {
          return this.$store.getters['objects/pairFilteredItems']
        }
      },
      invertedPairFilteredItems () {
        const ordersForSell = this.$store.getters['objects/ordersForSell']
        if (ordersForSell()) {
          return ordersForSell
        } else {
          return this.$store.getters['objects/invertedPairFilteredItems']
        }
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