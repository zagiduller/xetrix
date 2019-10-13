<template>
  <div class="text-xs-left">
    <v-menu
            v-model="menu"
            :close-on-content-click="false"
            :nudge-width="200"
            offset-x
    >
      <template #activator="data" >
        <v-btn
                class="ma-0"
                color="indigo"
                dark
                v-on="data.on"
        >
          Пополнение
        </v-btn>
      </template>

      <v-card>
        <v-list>
          <v-list-tile avatar>
            <v-list-tile-avatar>
              <v-icon>shopping_cart</v-icon>
            </v-list-tile-avatar>

            <v-list-tile-content>
              <v-list-tile-title>Давайте пополним счет!</v-list-tile-title>
              <v-list-tile-sub-title>{{account.Address}}</v-list-tile-sub-title>
            </v-list-tile-content>

            <v-list-tile-action>
              <v-btn
                      :class="fav ? 'red--text' : ''"
                      icon
                      @click="fav = !fav"
              >
                <v-icon>favorite</v-icon>
              </v-btn>
            </v-list-tile-action>
          </v-list-tile>
        </v-list>

        <v-divider></v-divider>

        <v-list>
          <v-list-tile>

            <v-list-tile-avatar>Сумма</v-list-tile-avatar>
            <v-spacer></v-spacer>
            <v-flex md8 class="text-xs-right">
              <v-text-field
                      :suffix="account.currency.symbol"
                      type="number"
                      label="Amount"
                      v-model="amount"
              ></v-text-field>
            </v-flex>
            <v-list-tile-action></v-list-tile-action>
          </v-list-tile>
        </v-list>
        <v-list v-if="proccLocked">
          <v-layout row>
            <v-progress-linear v-model="valueDeterminate"></v-progress-linear>
          </v-layout>
          <v-list-tile>
            <v-list-tile-content v-if="!!proccLocked">Понимание > </v-list-tile-content>
            <v-list-tile-content v-if="!!progress.rawtx">Сохранение > </v-list-tile-content>
            <v-list-tile-content v-if="!!progress.tx">Подтверждение > </v-list-tile-content>
            <v-list-tile-content v-if="!!progress.confirmed">Готово!</v-list-tile-content>
          </v-list-tile>
        </v-list>

        <v-card-actions>
          <v-spacer></v-spacer>

          <v-btn class="text-none" flat @click="doneOrCancel">Не надо</v-btn>
          <v-btn color="primary" flat @click="fund">Давай</v-btn>
        </v-card-actions>
      </v-card>
    </v-menu>
  </div>
</template>

<script>
  export default {
    name: "simpleFund",
    props: ['account'],
    data () {
      return {
        fav: true,
        menu: false,
        message: false,
        hints: true,
        amount: 0,
        proccLocked: false,
        progress: {
          rawtx: null,
          tx: null,
          confirmed: null,
        },
      }
    },
    computed: {
      valueDeterminate () {
        var value = 0
        if (!!this.progress.rawtx) {
          value += 35
        }
        if (!!this.progress.tx) {
          value += 35
        }
        if (!!this.progress.confirmed) {
          value += 30
        }
        return value
      }
    },
    methods: {
      fund() {
        if (this.amount > 0) {
          this.proccLocked = true
          this.$axios.post(process.env.apiUrl + '_v1-a/rawtx', {
            toAddress: this.account.Address,
            amount: this.amount,
          }).then((rawresp) => {
            var rawtx = rawresp.data.object
            console.log('rawtx',rawtx)
            this.progress.rawtx = rawtx
            this.$axios.post(process.env.apiUrl + '_v1-a/create_tx', rawtx).then((txresp)=>{
              var tx = txresp.data.object
              console.log('tx', tx)
              this.progress.tx = tx
              this.$axios.get(process.env.apiUrl + '_v1-a/confirm_tx/'+tx.id).then((confresp)=>{
                var conf = confresp.data.object
                console.log('confirmed', conf)
                this.progress.confirmed = conf

                setTimeout(this.doneOrCancel, 1000)
              })
            })
          })
        }
        // this.menu = false
      },
      doneOrCancel() {
        this.menu = false
        this.amount = 0
        this.proccLocked = false
        this.progress = {
          rawtx: null,
          tx: null,
          confirmed: null,
        }
      },
    }
  }
</script>

<style scoped>

</style>