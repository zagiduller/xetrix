<template>
  <v-card>
      <v-layout row justify-space-between>
        <v-flex md3 class="grey white--text">
          <v-card-title>
            <span class="title">Ордер</span>
          </v-card-title>
        </v-flex>
        <v-flex md5>
          <v-card-title>
            <v-layout row justify-space-between>
              <v-flex md2>
                <p class="ma-0 text-xs-center">
                  <span class="title">{{available}}</span>
                  <span class="subheading">{{from}}</span>
                </p>
              </v-flex>
              <v-flex md1>
                <p class="text-xs-center ma-0"><span class="title">по</span></p>
              </v-flex>
              <v-flex md2>
                <p class="ma-0 text-xs-center">
                  <span class="title">{{price}}</span>
                  <span class="subheading">{{to}}</span>
                </p>
              </v-flex>
            </v-layout>
          </v-card-title>
        </v-flex>
        <v-flex md2 class="orange white--text" v-if="!isUser">
          <template v-if="available > 0">
            <v-card-title class="text-xs-right">
              <p class="text-md-right ma-0">
                Купить {{from}}
              </p>
            </v-card-title>
            <v-card-title >
              <ContractSimpleForm :order="order"></ContractSimpleForm>
            </v-card-title>
          </template>
        </v-flex>
        <v-flex v-else md2 class="green white--text">
          <v-card-title class="text-xs-right">
            <p class="text-md-right ma-0">
              <v-icon>person</v-icon>
            </p>
          </v-card-title>
        </v-flex>
      </v-layout>
  </v-card>
</template>

<script>
  import ContractSimpleForm from '~/components/forms/contract/simple'
  export default {
    name: "order",
    props: ['order'],
    components: {ContractSimpleForm},
    computed: {
      isUser() {
        return this.order.ownerId === this.$auth.user.id
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
      }
    },

  }
</script>

<style scoped>

</style>