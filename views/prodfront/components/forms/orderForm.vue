<template>
  <v-layout class="data theme--light table-style style-1 buy-table top-panel-blocks block-style-1">
    <v-flex>
      <v-layout px-3 py-1 wrap v-if='Object.keys(pair).length !== 0'>
        <v-flex xs24>
          <h2 class="text-xs-left heading py-3">Кошельки</h2>
          <v-layout align-center justify-space-between row px-0 py-0>
            <v-flex pl-0  style="flex: 0 0 40px;"><svg-icon name="wallet-small-invert" style="width: 30px; height: 24px;"/></v-flex>
            <v-flex fill-height my-0 py-0 overflow-hidden>
              <v-layout fill-height wrap>
                <v-flex xs24 class="px-0 pb-0"><span class="body-3 text-color">Баланс: </span><span class="font-weight-bold body-3">{{available(pair.from.balance.available)}} {{pair.from.account.currency.symbol}}</span></v-flex>
                <v-flex xs24 class="pt-0">
                  <v-layout>
                    <v-flex  px-0 py-2>
                      <v-layout align-space-between justify-space-between row fill-height mx-0 my-0>
                        <span class="body-3 ">{{pair.from.account.Address}}</span>
                      </v-layout>
                    </v-flex>
                  </v-layout>
                </v-flex>
              </v-layout>
            </v-flex>
          </v-layout>
          <v-layout align-center justify-space-between row px-0 py-0>
            <v-flex pl-0 style="flex: 0 0 40px;"><svg-icon name="wallet-small-invert" style="width: 30px; height: 24px;"/></v-flex>
            <v-flex fill-height my-0 py-0 overflow-hidden>
              <v-layout fill-height wrap>
                <v-flex xs24 class="px-0 pb-0"><span class="body-3 text-color">Баланс: </span><span class="font-weight-bold body-3">{{available(pair.to.balance.available)}} {{pair.to.account.currency.symbol}}</span></v-flex>
                <v-flex xs24 class="pt-0">
                  <v-layout>
                    <v-flex px-0 py-2>
                      <v-layout align-space-between justify-space-between row fill-height mx-0 my-0>
                        <span class="body-3 ">{{pair.to.account.Address}}</span>
                      </v-layout>
                    </v-flex>
                  </v-layout>
                </v-flex>
              </v-layout>
            </v-flex>
          </v-layout>
        </v-flex>
        <v-flex xs12 v-if="pair">
          <h2 class="text-xs-left heading pb-3">Отдаете</h2>
          <img :src="'/images/' + pair.from.account.currency.symbol + '.png'" width="19" alt="" style="vertical-align: text-bottom"> <span class="body-4">{{pair.from.account.currency.name}} ({{pair.from.account.currency.symbol}})</span>
          <span class="mt-3 label body-1 pb-2 d-block font-weight-bold">Количество:</span>
          <v-text-field
                  class="table-input"
                  single-line
                  outline
                  color=""
                  v-model="amount"

          ></v-text-field>
        </v-flex>
        <v-flex xs12>
          <h2 class="text-xs-left heading pb-3">Получаете</h2>
          <img :src="'/images/' + pair.to.account.currency.symbol + '.png'" width="19" alt="" style="vertical-align: text-bottom"> <span class="body-4">{{pair.to.account.currency.name}} ({{pair.to.account.currency.symbol}})</span>
          <span class="mt-3 label body-1 pb-2 d-block font-weight-bold">Цена за единицу:</span>
          <v-text-field
                  class="table-input"
                  single-line
                  outline
                  color=""
                  v-model="price"></v-text-field>
        </v-flex>
        <v-flex xs12>
          <span class="label body-1 pb-2 d-block font-weight-bold">Комиссия</span>
          <span class="body-5 d-block data pb-3">{{commision}}%</span>
          <span class="label body-1 pb-2 d-block font-weight-bold">Цена с комиссией</span>
          <span class="body-5 d-block data pb-3" >{{actualAmount}} {{pair.to.account.currency.symbol}}</span>
        </v-flex>
        <v-flex xs12>
          <span class="label body-1 pb-2 d-block font-weight-bold">Итоговая цена</span>
          <span class="body-5 d-block data pb-3  d-block">{{Amount}} {{pair.to.account.currency.symbol}}</span>
          <v-btn color="" block flat :disabled="!query"  @click="open" class="mb-0 v-btn-style v-btn__type3">Создать ордер</v-btn>
        </v-flex>
      </v-layout>
    </v-flex>
  </v-layout>
</template>

<script>
  export default {
    name: "orderForm",
    props: ['accs', 'pair'],
    created: function () {
      this.from = this.pair.from.account.currency.symbol;
      this.to = this.pair.to.account.currency.symbol;
    },
    computed: {
      actualAmount() {
        return this.toRound((this.amount * this.price), this.pair.to.account.currency.symbol)
      },
      Amount () {
        return this.toRound((this.amount * this.price) * (1 - this.commision / 100), this.pair.to.account.currency.symbol)
      },
      lengths () {
        if (this.pair.to) {
          return {
            [this.pair.to.account.currency.symbol]: this.pair.to.account.currency.decimal,
            [this.pair.from.account.currency.symbol]: this.pair.from.account.currency.decimal
          }
        } else {
          return []
        }
      },
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
        if (q.sendingAddress && q.receiveAddress && q.price && q.amount ) {
          return q
        }
        return false
      },
    },
    watch: {
      pair() {
        this.$nextTick(() => {
          this.from = this.pair.from.account.currency.symbol;
          this.to = this.pair.to.account.currency.symbol;
        })
      }
    },
    data () {
      return {
        amount: 0,
        price: 0,
        from: '',
        to: '',
        dialog: false,
        commision: 1
      }
    },
    methods: {
      toRound (value, code) {
        const length = this.lengths[code] ? this.lengths[code] : 0
        const coef = Math.pow(10, length)
        return Math.round(value * coef) / coef
      },
      available(value) {
        return value > 0 ? value : 0
      },
      open() {
        let snotify = this.$snotify
        let pair = this.pair
        let q = this.query
        if (this.query) {
          var url = process.env.apiUrl + '_v1/create_order'
          return this.$axios.post(url, q)
            .then(function(success) {
              success = success.data;
              console.log(success);
              snotify.success('Ордер успешно создан', 'Создание ордера', {
                timeout: 4000,
                showProgressBar: false,
                closeOnClick: false,
                pauseOnHover: true
              });
            })
            .catch(function(error) {
              console.log(error);
              console.log(error.response);

              console.log(error.response.data.code);
              let message = ''

              if(error.response.data.code !== undefined && error.response.data.code == 2) {
                message = `Недостаточно средств на счету ${pair.from.account.currency.symbol}`
              } else {
                message = 'При создании ордера что-то пошло не так'
              }
              snotify.error(message, 'Ошибка', {
                timeout: 1000,
                showProgressBar: false,
                closeOnClick: false,
                pauseOnHover: true
              });
            });
        }
      },
      hashVisual: function (number, type) {
        let countlast = 10
        if (type === 'last') return number.substr(-countlast)
        if (type === 'first') return number.substr(0, countlast)
      },
    }
  }
</script>

<style scoped>

</style>