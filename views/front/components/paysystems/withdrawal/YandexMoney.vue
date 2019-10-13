<template>
  <v-container>
    <v-flex md12>
      <v-text-field
              label="Количество"
              v-model="amount"
      ></v-text-field>
    </v-flex>
    <v-flex md12>
      <v-text-field
              label="Адрес счета"
              v-model="address"
      ></v-text-field>
    </v-flex>
    <v-flex md12>
      <v-text-field
              label="ФИО"
              v-model="fullName"
      ></v-text-field>
    </v-flex>
    <v-flex md12 class="text-md-right">
      <v-btn class="amber accend-3 ml-0 mt-3 mb-0" @click="send">Подать заявку</v-btn>
    </v-flex>
    <v-flex v-if="!!worder">
      <p class="title">
        Создана заявка:
      </p>
      <p>
        {{worder}}
      </p>
    </v-flex>
  </v-container>
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