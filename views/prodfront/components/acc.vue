<template>
  <v-layout wrap card-wallet>
    <v-flex xs24 row h-4 px-0 py-0>
      <h2><span class="text-uppercase">{{symbol}}</span> кошелек</h2>
    </v-flex>
    <v-flex xs24 class="card-wallet theme--light block-style-1 mx-0 my-0 px-0" overflow-hidden>
      <v-layout align-center justify-space-between row fill-height wrap px-0 class="header" style="width: max-content;min-width: 100%;">
        <v-flex style="flex: 0 0 30px;"><img :src="'/images/left/' + symbol.toLowerCase() + '.png'" width="48" alt=""></v-flex>
        <v-flex align-center fill-height py-0 px-0 style="width: calc(100% - 64px); min-width: 260px;">
          <v-layout align-center fill-height my-0>
            <v-flex xs24>
              <v-layout align-start justify-center column>
                <div>
                  <h1 style="line-height: 26px;" class="font-weight-regular text-uppercase">{{available}} {{symbol}}</h1>
                </div>
              </v-layout>
              <v-layout align-center fill-height wrap pt-1>
                <v-flex pl-0 py-0 style="flex: 0 0 40px;"><svg-icon name="wallet-small" style="width: 22px; height: 18px"/></v-flex>
                <v-flex>
                  <v-layout align-center wrap>
                    <v-flex xs24>
                      <v-layout align-center>
                        <template v-if="$vuetify.breakpoint.smAndDown">
                          <v-flex>
                            <v-layout align-space-between justify-space-between row fill-height mx-0 my-0 class="address_currency">
                              <div><span class="body-1 text-color font-weight-bold left">{{this.hashVisual(address, 'first', 5)}}</span></div>
                              <div><span class="body-1 text-color font-weight-bold right">{{this.hashVisual(address, 'last', 5)}}</span></div>
                            </v-layout>
                          </v-flex>
                        </template>
                        <template v-else>
                          <v-layout align-space-between justify-space-between row fill-height mx-0 my-0 class="address_currency">
                            <div><span class="body-2 text-color font-weight-bold left">{{this.hashVisual(address, 'first', 12)}}</span></div>
                            <div><span class="body-2 text-color font-weight-bold right">{{this.hashVisual(address, 'last', 12)}}</span></div>
                          </v-layout>
                        </template>
                        <v-flex text-xs-center  style="flex: 0 0 40px;"></v-flex>
                      </v-layout>
                    </v-flex>
                  </v-layout>
                </v-flex>
              </v-layout>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-layout>
      <v-layout align-center justify-space-between row footer px-xs-2 px-md-4 py-2>
        <v-flex>
          <span class="label body-5 pb-2">Доступный баланс:</span>&nbsp;<span class="body-5 data pb-3">{{available}}</span><br>
          <span class="label body-5 pb-2">Заблокировано:</span>&nbsp;<span class="body-5 data pb-3">{{locked ? locked : 0}}</span><br>
          <v-layout data theme--light table-style align-center justify-center row margin-minus>
            <v-flex xs12 mr-2>
              <v-btn block flat class="mb-0 v-btn-style v-btn__type3" :disabled="!ps.length" @click="paymentDialog = !paymentDialog">Пополнить</v-btn>
            </v-flex>
            <v-flex xs12 ml-2>
              <v-btn block flat class="mb-0 v-btn-style v-btn__type3" :disabled="!ps.length" @click="withdrawalDialog = !withdrawalDialog">Вывести</v-btn>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-layout>
      <div class="d-none">
        <v-dialog v-if="symbol === 'RUB'" v-model="paymentDialog" width="460" content-class="theme--light dialog--order" :fullscreen="$vuetify.breakpoint.xsOnly" transition="dialog-opacity-transition">
          <v-layout>
            <v-flex>
              <v-layout align-center justify-center column class="dialog--header">
                <h2 class="body-2">Платеж</h2>
                <a @click="paymentDialog = !paymentDialog" class="dialog--close-button"><v-icon small>mdi-close</v-icon></a>
              </v-layout>
              <v-layout justify-center px-md-5 px-xs-2 py-5 class="dialog--content" wrap >
                <v-flex xs24>
                  <v-layout align-center justify-space-between row>
                    <v-flex pl-0 style="flex: 0 0 40px;"><svg-icon name="wallet-small-invert" style="width: 30px; height: 24px;"/></v-flex>
                    <v-flex fill-height my-0 py-0 overflow-hidden>
                      <v-layout fill-height wrap>
                        <v-flex xs24 class="px-0 pb-0"><span class="body-3 gray gray-lighten-1 text-color">Баланс: </span><span class="font-weight-bold body-3">{{available}} {{symbol}}</span></v-flex>
                        <v-flex xs24 class="pt-0">
                          <v-layout>
                            <v-flex px-0 py-0>
                              <v-layout align-space-between justify-space-between row fill-height mx-0 my-0>
                                <span class="body-3 ">{{address}}</span>
                              </v-layout>
                            </v-flex>
                          </v-layout>
                        </v-flex>
                      </v-layout>
                    </v-flex>
                  </v-layout>
                </v-flex>
                <v-flex xs22 md24>
                  <v-layout column wrap>
                    <v-flex xs24 mb-2 pb-2 mt-2>
                      <iframe
                              :src="'https://money.yandex.ru/quickpay/shop-widget?writer=seller&targets='+address+'&label='+address+'&targets-hint=&default-sum=2&button-text=11&payment-type-choice=on&hint=&successURL=&quickpay=shop&account=410018657112160'"
                              width="100%" height="222"
                              frameborder="0"
                              allowtransparency="true"
                              scrolling="no"
                      >
                      </iframe>
                    </v-flex>
                  </v-layout>
                </v-flex>
              </v-layout>
            </v-flex>
          </v-layout>
        </v-dialog>
      </div>
    </v-flex>
    <v-dialog content-class="theme--light dialog--order" v-if="!!ps.length" v-model="withdrawalDialog" width="460" :fullscreen="$vuetify.breakpoint.xsOnly" transition="dialog-opacity-transition">
      <v-layout>
        <v-flex>
          <v-layout align-center justify-center column class="dialog--header">
            <h2 class="body-2">Заявка на вывод</h2>
            <a @click="withdrawalDialog = !withdrawalDialog" class="dialog--close-button"><v-icon small>mdi-close</v-icon></a>
          </v-layout>
          <v-layout justify-center px-md-5 px-xs-2 py-5 class="dialog--content" wrap >
            <v-flex xs24>
              <v-layout align-center justify-space-between row>
                <v-flex pl-0 style="flex: 0 0 40px;"><svg-icon name="wallet-small-invert" style="width: 30px; height: 24px;"/></v-flex>
                <v-flex fill-height my-0 py-0 overflow-hidden>
                  <v-layout fill-height wrap>
                    <v-flex xs24 class="px-0 pb-0"><span class="body-3 gray gray-lighten-1 text-color">Баланс: </span><span class="font-weight-bold body-3">{{available}} {{symbol}}</span></v-flex>
                    <v-flex xs24 class="pt-0">
                      <v-layout>
                        <v-flex px-0 py-0>
                          <v-layout align-space-between justify-space-between row fill-height mx-0 my-0>
                            <span class="body-3 ">{{address}}</span>
                          </v-layout>
                        </v-flex>
                      </v-layout>
                    </v-flex>
                  </v-layout>
                </v-flex>
              </v-layout>
            </v-flex>
            <v-flex xs22 md18>
              <v-layout column wrap>
                <v-flex xs24 mb-2>
                  <v-tabs fixed-tabs>
                    <v-tab
                            v-for="p in ps"
                            :key="p.Name"
                    >
                      {{ p.Name }}
                    </v-tab>
                    <v-tab-item
                            v-for="p in ps"
                            :key="p.Name"
                    >
                      <v-card flat>
                        <v-card-text>
                          <v-layout py-2>
                            <component :is="p.Name" :acc="acc"></component>
                          </v-layout>
                        </v-card-text>
                      </v-card>
                    </v-tab-item>
                  </v-tabs>
                </v-flex>
              </v-layout>
            </v-flex>
          </v-layout>
        </v-flex>
      </v-layout>
    </v-dialog>
  </v-layout>
</template>
<script>
  import Bitcoin from '~/components/paysystems/withdrawal/Bitcoin'
  import YandexMoney from '~/components/paysystems/withdrawal/YandexMoney'

  export default {
    name: "acc",
    props: ['acc', 'ps'],
    components: {YandexMoney, Bitcoin},
    computed: {
      currency() {
        return this.acc.account.currency.name
      },
      symbol() {
        return this.acc.account.currency.symbol
      },
      available() {
        return this.acc.balance.available > 0 ? this.toRound(this.acc.balance.available) : 0
      },
      locked() {
        return this.acc.balance.locked > 0 ? this.toRound(this.acc.balance.locked) : 0
      },
      address() {
        return this.acc.account.Address
      }
    },
    data () {
      return {
        paymentDialog: false,
        withdrawalDialog: false,
      }
    },
    methods: {
      toRound (value) {
        const length = this.acc.account.currency.decimal ? this.acc.account.currency.decimal : 0
        const coef = Math.pow(10, length)
        return Math.round(value * coef) / coef
      },
      hashVisual: function (number, type, countlast) {
        if (type === 'last') return number.substr(-countlast)
        if (type === 'first') return number.substr(0, countlast)
      },
    }
  }
</script>

<style scoped>

</style>