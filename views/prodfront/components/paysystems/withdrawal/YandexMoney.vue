<template>
  <v-layout wrap>
    <v-flex xs24 mb-2>
      <v-text-field class="table-input"
                    single-line
                    outline
                    color=""
              label="Количество"
                    placeholder="Количество"
              v-model="amount"
      ></v-text-field>
    </v-flex>
    <v-flex xs24 mb-2>
      <v-text-field class="table-input"
                    single-line
                    outline
                    color=""
              label="Адрес счета"
                    placeholder="Адрес счета"
              v-model="address"
      ></v-text-field>
    </v-flex>
    <v-flex xs24 mb-2>
      <v-text-field class="table-input"
                    single-line
                    outline
                    color=""
                    placeholder="ФИО"
              label="ФИО"
              v-model="fullName"
      ></v-text-field>
    </v-flex>
    <v-flex xs24 mb-2 class="text-md-right">
      <v-btn block flat class="mb-0 v-btn-style v-btn__type3" @click="send">Подать заявку</v-btn>
    </v-flex>
    <v-flex v-if="!!worder">
      <p class="title">
        Создана заявка:
      </p>
      <p>
        {{worder}}
      </p>
    </v-flex>
  </v-layout>
</template>

<script>
  export default {
    name: "YandexMoney",
    props: ['acc'],
    data () {
      return {
        paymentSystem: 'YandexMoney',
        amount: 0,
        fullName: '',
        address: '',
        worder: '',
      }
    },
    computed: {
      attributes() {
        return [
          {
            key: 'fullName',
            value: this.fullName,
          },
          {
            key: 'YandexMoneyAddress',
            value: this.address,
          },
        ]
      }
    },
    methods: {
      send () {
        this.$axios.post(process.env.apiUrl + '_v1/create_withdrawal',{
          paymentSystem: this.paymentSystem,
          amount: this.amount,
          sendingAddress: this.acc.account.Address,
          attributes: this.attributes,
        }).then((resp) => {
          console.log(resp.data)
          this.worder = resp.data
          this.reset()
        })
      },
      reset() {
        this.amount = 0
        this.fullName = ''
        this.address = ''
      }
    },
  }
</script>

<style scoped>

</style>