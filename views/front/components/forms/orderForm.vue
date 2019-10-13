<template>
  <v-card>
    <v-card-title class="grey white--text pb-0">
      <v-layout row>
        <v-flex>
          <span class="title">Размещение ордера</span>
        </v-flex>
        <v-flex>
          <p class="text-xs-right">
            <v-btn class="ma-0" color="primary" :disabled="!query" small @click="open">Опубликовать</v-btn>
          </p>
        </v-flex>
      </v-layout>

    </v-card-title>
    <v-card-text class="pl-3 pr-3 pb-0">
      <v-layout row wrap justify-space-between>
        <v-flex md5>
          <p class="ma-0 mb-2">
            Что продаем?
          </p>
          <v-layout row wrap>
            <v-flex d-flex md6>
              <v-text-field
                      label="Кол-во"
                      v-model="amount"
                      solo
                      :disabled="!from"
              ></v-text-field>
            </v-flex>
            <v-flex d-flex md6>
              <v-select
                      class="elevation-0"
                      :items="currencies"
                      item-text="symbol"
                      item-value="symbol"
                      v-model="from"
                      label="Продам"
                      solo
              ></v-select>
            </v-flex>
          </v-layout>
        </v-flex>
        <v-flex md1>
          <p class="mt-5 mr-0 ml-0 text-xs-center">за</p>
        </v-flex>
        <v-flex md5>
          <p class="text-xs-right ma-0 mb-2">
            Почем?
          </p>
          <v-layout row wrap>
            <v-flex d-flex md6>
              <v-text-field
                      class="elevation-0"
                      label="Цена за ед."
                      solo
                      v-model="price"
                      :disabled="!to"
              ></v-text-field>
            </v-flex>
            <v-flex d-flex md6>
              <v-select
                      :items="toList"
                      item-text="symbol"
                      item-value="symbol"
                      v-model="to"
                      label="Куплю"
                      solo
                      :disabled="!from"
              ></v-select>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-layout>
      <v-layout row wrap justify-space-between>

      </v-layout>
    </v-card-text>
    <v-card-actions class="pt-0">
      <v-spacer></v-spacer>

    </v-card-actions>
  </v-card>
</template>

<script>
  export default {
    name: "orderForm",
    props: ['accs'],
    computed: {
      currencies() {
        var crs = []

        this.accs.forEach((a) => {
          var exist = crs.findIndex((c) => {return a.account.currency.symbol === c.symbol })
          if (exist < 0) {
            crs.push(a.account.currency)
          }
        })

        return crs
      },
      toList() {
        if (!!this.from) {
          return this.currencies.filter((c) => c.symbol !== this.from)
        }
        return this.currencies
      },
      send() {
        if (!!this.from) {
          var acc = this.accs.filter((a)=>{
            return a.account.currency.symbol === this.from
          })[0]
          if (!!acc.account){
            return acc.account.Address
          }

        }
        return false
      },
      receive() {
        if (!!this.to) {
          var acc = this.accs.filter((a)=>{
            return a.account.currency.symbol === this.to
          })[0]
          if (!!acc.account){
            return acc.account.Address
          }
        }
        return false
      },
      query() {
        var q = {}
        q.sendingAddress = this.send
        q.receiveAddress = this.receive
        q.amount = this.amount
        q.price = this.price
        console.log(q)
        if (q.sendingAddress && q.receiveAddress && q.price && q.amount ) {
          return q
        }
        return false
      }
    },
    data () {
      return {
        amount: 0,
        price: 0,
        from: '',
        to: '',
      }
    },
    methods: {
      open() {
        var q = this.query
        if (this.query) {
          var url = process.env.apiUrl + '_v1/create_order'
          this.$axios.post(url, q).then((resp)=>{
            console.log(resp)
          })
        }
      }
    }
  }
</script>

<style scoped>

</style>