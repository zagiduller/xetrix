<template>
  <div>
    <v-dialog
            v-if="symbol === 'RUB'"
            v-model="paymentDialog"
            width="500"
    >
      <v-card>
        <v-card-title>
          <span class="title">Платеж</span>
          <v-spacer></v-spacer>
          <v-icon @click="paymentDialog = !paymentDialog">close</v-icon>
        </v-card-title>
        <v-card-text>
          <iframe
                  :src="'https://money.yandex.ru/quickpay/shop-widget?writer=seller&targets='+address+'&label='+address+'&targets-hint=&default-sum=2&button-text=11&payment-type-choice=on&hint=&successURL=&quickpay=shop&account=410018657112160'"
                  width="423" height="222"
                  frameborder="0"
                  allowtransparency="true"
                  scrolling="no"
          >

          </iframe>
        </v-card-text>
      </v-card>
    </v-dialog>

    <v-dialog
            v-if="!!ps.length"
            v-model="withdrawalDialog"
            width="500"
    >
      <v-card>
        <v-card-title>
          <p class="title mb-0 pb-0">
            Заявка на вывод <br>
            <span class="caption">{{address}}</span>
          </p>
          <v-spacer></v-spacer>
          <v-icon @click="withdrawalDialog = !withdrawalDialog">close</v-icon>
        </v-card-title>
        <v-card-text>
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
                  <v-layout>
                    <component :is="p.Name" :acc="acc"></component>
                  </v-layout>
                </v-card-text>
              </v-card>
            </v-tab-item>
          </v-tabs>
        </v-card-text>
      </v-card>
    </v-dialog>

    <v-card>
      <v-card-title class="grey white--text">
        <v-flex md3>
          <span class="title">Счет</span>
        </v-flex>
        <v-spacer></v-spacer>
        <v-flex>
          <p class="text-xs-right ma-0">
            <span class="title"><b>{{currency}}</b></span>
          </p>
        </v-flex>
      </v-card-title>
      <v-card-text class="">
        <v-layout row wrap >
          <v-flex md2 d-flex>
            <p class="text-xs-left">{{symbol}}</p>
          </v-flex>
          <v-flex md3 d-flex>Доступно:</v-flex>
          <v-spacer></v-spacer>
          <v-flex md6 d-flex class="text-xs-right"><p class="text-xs-right"><b>{{available}}</b></p></v-flex>
        </v-layout>
        <v-layout row wrap class="mt-3">
          <v-flex md2 d-flex>
            <p class="text-xs-left">{{symbol}}</p>
          </v-flex>
          <v-flex md3 d-flex>Заблокировано:</v-flex>
          <v-spacer></v-spacer>
          <v-flex md6 d-flex><p class="text-xs-right"><b>{{locked}}</b></p></v-flex>
        </v-layout>
      </v-card-text>
      <v-card-actions class="grey lighten-2 pa-0">
        <template v-if="!!ps.length">
          <v-btn block flat small class="amber accent-3 ma-0" @click="paymentDialog = !paymentDialog">Пополнить</v-btn>
          <v-btn block flat small class="amber accent-3 ma-0" @click="withdrawalDialog = !withdrawalDialog">Вывести</v-btn>
        </template>
      </v-card-actions>
    </v-card>
  </div>

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
        return this.acc.balance.available > 0 ? this.acc.balance.available : 0
      },
      locked() {
        return this.acc.balance.locked > 0 ? this.acc.balance.locked : 0
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
    }
  }
</script>

<style scoped>

</style>