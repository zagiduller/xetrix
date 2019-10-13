<template>
  <v-container fluid class="pa-0">
    <v-layout row wrap>
      <v-flex md6>
        <v-text-field solo v-model="amount"></v-text-field>
      </v-flex>
      <v-flex md4>
        <v-btn
                @click="buy"
                small fab
                color="white"
                style="height: 30px; width: 30px">
          <v-icon>add</v-icon>
        </v-btn>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
  export default {
    name: "simple",
    props: ['order'],
    computed: {
      wsaccs () {
        return this.$store.state.objects.WsAccounts
      },
      id () {
        return this.order.id
      },
      buyerId () {
        return this.$auth.user.id
      },
      send () {
        var acc = this.wsaccs.filter((a) => {
          return a.account.currency.symbol === this.to
        })[0]
        if (!!acc.account) {
          return acc.account.Address
        }
        return false
      },
      receive () {
        var acc = this.wsaccs.filter((a) => {
          return a.account.currency.symbol === this.from
        })[0]
        if (!!acc.account) {
          return acc.account.Address
        }
        return false
      },
      from () {
        return this.order.sellCurrencySymbol
      },
      to () {
        return this.order.buyCurrencySymbol
      },
      available () {
        return this.order.available
      },
      price () {
        return this.order.price
      },
      query() {
        var q = {}
        q.orderId = this.id
        q.sendingAddress = this.send
        q.receiveAddress = this.receive
        q.buyerId = this.buyerId
        q.amount = this.amount
        if (!!q.amount <= this.available
          && !!q.orderId
          && !!q.sendingAddress
          && !!q.receiveAddress
          && !!q.orderId && !!q.buyerId) {
          return q
        }
        return false
      },
    },
    data() {
      return {
        amount: 0,
      }
    },
    methods: {
      buy() {
        if (this.query) {
          var url = process.env.apiUrl + '_v1/create_contract'
          this.$axios.post(url, this.query).catch((resp) => {
            console.log(resp)
          })
        }
      }
    },
  }
</script>

<style scoped>

</style>