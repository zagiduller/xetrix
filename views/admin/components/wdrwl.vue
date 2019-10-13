<template>
  <v-flex >
    <v-card class="pt-0 pb-5">
      <v-layout row wrap>
        <v-flex xs12>
          <v-card-title v-bind:class="{'amber': active, 'green': !active}" class="lighten-4 pt-3">
            <v-layout row justify-space-between class="">
              <v-flex xs11 >
                <p class="title">{{active ? 'Активная. ' : 'Исполненная. '}} Заявка на вывод. <br> <span class="subheading">{{ wo.id }}</span></p>
              </v-flex>
              <v-flex v-if="active">
                <v-btn @click="perform" color="success">Исполнить</v-btn>
              </v-flex>
          </v-layout>
          </v-card-title>
        </v-flex>
        <v-flex xs12>
          <v-card-text >
            <p>Сумма вывода: {{ wo.amount }} {{acc.currency.symbol}}</p>
            <hr>
            <p>
              Данные по заявке:
            </p>
            <ul>
              <li v-for="attr in wo.attributes" :key="attr.key">
                <b>{{attr.key}}</b> - {{attr.value}}
              </li>
            </ul>
          </v-card-text>
        </v-flex>
        <v-flex xs12>
          <v-card-text >
            <p>Связанная Транзакция:</p>
            <p><code>{{ wo.relatedTxId }}</code></p>
            <p><code>{{ tx}}</code></p>
          </v-card-text>
        </v-flex>
      </v-layout>
    </v-card>
  </v-flex>

</template>

<script>
  export default {
    name: "wdrwl",
    props: ['wo'],
    data () {
      return {
      }
    },
    computed: {
      tx () {
        return this.$store.state.objects.txs.find((e)=>{
          return e.id === this.wo.relatedTxId
        })
      },
      acc() {
        var acc = this.$store.state.objects.WsAccounts.filter((ws) => {
          return ws.account.Address === this.wo.sendingAddress
        })
        return acc[0].account
      },
      active() {
        return this.wo.status.status !== 1
      }
    },
    methods: {
      perform() {
        this.$axios.get(process.env.apiUrl + '_v1-a/confirm_tx/' + this.wo.relatedTxId).then((resp) => {
          console.log(resp.data)
        })
      },
    },
  }
</script>

<style scoped>

</style>