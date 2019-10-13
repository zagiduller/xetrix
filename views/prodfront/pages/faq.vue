<template>
    <v-layout justify-center mx-0 my-0 px-3 py-3 wrap>
        <v-flex xs24 md15>
            <v-layout wrap justify-xs-start mt-md-5>
                <v-flex xs24 text-xs-center mt-md-5>
                    <img src="/images/xetrix_support_logo.png" alt="">
                    <h1 class="body-9 font-weight-bold pb-5 pt-4">Как мы можем вам помочь?</h1>
                </v-flex>
                <v-flex xs24>
                    <v-expansion-panel class="faq-panel">
                        <v-expansion-panel-content
                                v-for="(item,i) in 5"
                                :key="i"
                        >
                            <template v-slot:header>
                                <div>Item</div>
                            </template>
                            <v-card>
                                <v-card-text>Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.</v-card-text>
                            </v-card>
                        </v-expansion-panel-content>
                    </v-expansion-panel>
                    <!--<p class="body-5-md">Мы всегда готовы ответить на интересующие Вас <br>
                        вопросы, а также выслушать Ваши предложения<br>
                        по улучшению нашего сервиса.</p>
                    <span class="body-6 d-inline-block pb-3"><strong>E-mail</strong></span><br>
                    <svg-icon name="email-letter" style="width: 35px; height: 35px; vertical-align: middle" class="mr-4"/><a href="mailto:support@xet-rix.com" class="body-5-5"><strong>support@xet-rix.com</strong></a><br>
                    <span class="body-6 d-inline-block pb-3 pt-5"><strong>Мы в соц. сетях:</strong></span><br>
                    <div class="d-block text-xs-left">
                        <a class="body-4 links mr-3" target="_blank" href="https://www.facebook.com/xetrix"><v-icon class="body-7">mdi-facebook</v-icon></a>
                        <a class="body-4 links" target="_blank" href="https://twitter.com/xetrix"><v-icon class="body-7">mdi-twitter</v-icon></a>
                    </div>-->
                </v-flex>

            </v-layout>
        </v-flex>
    </v-layout>
</template>


<script>
  import { mapState } from 'vuex'
  import Acc from '~/components/acc'
  import OrderForm from '~/components/forms/orderForm'
  import Order from '~/components/order'

  export default {
    name: "faq",
    components: { Acc, OrderForm, Order },
    layout: 'base',
    data () {
      return {
        isBordered: true,
      }
    },
    computed: {
      WsAccounts () {
        return this.$store.state.objects.WsAccounts
      },
      orders () {
        return this.$store.state.objects.orders
      },
      userOrders() {
        return this.orders.filter((o)=>{
          return o.ownerId === this.$auth.user.id
        })
      }
    },
    methods: {
      accPS(accSymbol) {
        return this.$store.state.objects.paysystems.filter(ps => accSymbol === ps.Symbol)
      }
    },
  }
</script>