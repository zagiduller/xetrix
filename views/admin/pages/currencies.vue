<template>
  <v-container>
    <v-dialog v-model="addDialog" width="800">
      <v-card>
        <v-card-title class="headline green lighten-2" primary-title>
          <span>Добавление контракта</span>
          <v-spacer></v-spacer>
          <span class="caption" ><v-checkbox color="red" v-model="contract.active" :label="contract.active ? 'Активен' : 'Не активен'"></v-checkbox></span>
        </v-card-title>
        <v-card-text>
          <v-list three-line>
            <v-list-tile >
              <v-list-tile-content>Адрес контракта</v-list-tile-content>
              <v-list-tile-content>
                <v-container grid-list-md>
                  <v-flex md12>
                    <v-text-field  v-model="contract.address" :error-messages="contract.errors.address"></v-text-field>
                  </v-flex>
                </v-container>
              </v-list-tile-content>
            </v-list-tile>
            <v-list-tile>
              <v-list-tile-content>Название валюты</v-list-tile-content>
              <v-list-tile-content>
                <v-container grid-list-md>
                  <v-flex md12>
                    <v-text-field v-model="contract.name" :error-messages="contract.errors.name"></v-text-field>
                  </v-flex>
                </v-container>
              </v-list-tile-content>
            </v-list-tile>
            <v-list-tile>
              <v-list-tile-content>Символ валюты</v-list-tile-content>
              <v-list-tile-content>
                <v-container grid-list-md>
                  <v-flex md12>
                    <v-text-field v-model="contract.symbol" :error-messages="contract.errors.symbol"></v-text-field>
                  </v-flex>
                </v-container>
              </v-list-tile-content>
            </v-list-tile>
            <v-list-tile>
              <v-list-tile-content>Разрядность</v-list-tile-content>
              <v-list-tile-content>
                <v-container grid-list-md>
                  <v-flex md12>
                    <v-text-field v-model="contract.decimal" :error-messages="contract.errors.decimal"></v-text-field>
                  </v-flex>
                </v-container>
              </v-list-tile-content>
            </v-list-tile>
            <v-list-tile >
              <v-list-tile-content>Подтвердите введенные данные</v-list-tile-content>
              <v-list-tile-avatar>
                <v-checkbox v-model="contract.accept" :error-messages="contract.errors.accept"></v-checkbox>
              </v-list-tile-avatar>
            </v-list-tile>
          </v-list>
        </v-card-text>
        <v-card-actions >
          <v-spacer></v-spacer>
          <v-btn flat color="orange" @click="addCurrency">Добавить</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
    <v-layout row>
      <v-flex md12 class="text-md-right">
        <v-btn class="ma-0" color="success" @click="addDialog = true">Добавить</v-btn>
      </v-flex>
    </v-layout>
    <v-layout wrap class="mt-5">
      <v-flex md4 v-for="c in currencies" :key="c.id" class="pr-1 pt-1">
        <Currency :currency="c"></Currency>
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
  import Currency from '~/components/currency'

  export default {
    name: "currency",
    components: {Currency},
    data () {
      return {
        addDialog: false,
        contract: {
          address: "0x0",
          name: "",
          symbol: "",
          decimal: 18,
          accept: false,
          active: true,
          errors: {
          },
        }
      }
    },
    computed: {
      currencies() {
        return this.$store.state.objects.currencies
      }
    },
    methods: {
      addCurrency () {
        this.contract.errors = {
          address: [],
          name: [],
          symbol: [],
          accept: [],
        }
        var errs = 0
        if (this.contract.address.length < 5) {
          this.contract.errors.address.push( "Задайте адрес контракта")
          errs += 1
        }
        if (this.contract.name.length < 3) {
          this.contract.errors.name.push( "Задайте имя контракта")
          errs += 1
        }
        if (this.contract.symbol.length < 2) {
          this.contract.errors.symbol.push( "Задайте символ валюты")
          errs += 1
        }

        if (errs == 0) {
          if (!this.contract.accept) {
            this.contract.errors.accept.push( "Подтвердите данные и добавление")
            return
          }


          var url = process.env.apiUrl + '_v1-a/create_currency'

          this.$axios.post(url,{
            object: {
              contractId: this.contract.address,
              name: this.contract.name,
              symbol: this.contract.symbol.toUpperCase(),
              type: 2, // ETH-Токен
              decimal: 18,
              active: this.contract.active,
            }
          }).then((resp)=>{
            console.log(resp.data)
          })
        }
        return
      }
    },
  }
</script>

<style scoped>

</style>