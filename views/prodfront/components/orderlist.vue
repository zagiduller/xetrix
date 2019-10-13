<template >
    <div v-resize="onResize" :class=classtable>
        <v-data-table
          expand
          :headers="headers"
          :headers-length="(headers.length)+1"
          :items="items"
          :disable-initial-sort=true
          :pagination.sync="pagination"
          prev-icon="mdi-menu-left"
          next-icon="mdi-menu-right"
          sort-icon="mdi-chevron-down"
          rows-per-page-text
        >
        <template slot="headers" slot-scope="props">
            <tr>
                <th
                    v-for="header in props.headers"
                    :key="header.text"
                    :class="['column sortable', pagination.descending ? 'desc' : 'asc', header.value === pagination.sortBy ? 'active' : '']"
                    @click="changeSort(header.value)"
                    >
                    <v-icon small>mdi-chevron-down</v-icon>
                    <template v-if="header.id == 'sell' ">
                       {{!orders[0]  ? (scenario === 'sell' ? pair.from.account.currency.symbol : pair.to.account.currency.symbol) : orders[0].sellCurrencySymbol}}
                   </template>
                    <template v-else-if="header.id == 'buy' && orders[0]">
                        Цена
                    </template>
                    <template v-else="header.id">
                        {{header.text}}
                    </template>
                </th>
            </tr>
        </template>
          <template slot="items" slot-scope="props">
            <tr :class="[props.item.id ? 'active' : '', props.expanded ? 'active-expanded' : ''] ">
              <td v-if="toRender('available')" class="text-xs-center"  @click="touch(props)">{{props.item.available}}</td>
              <td v-if="toRender('price')" class="text-xs-center"  @click="touch(props)">{{props.item.price}} {{props.item.buyCurrencySymbol}}</td>
              <td v-if="toRender('moment')" class="text-xs-center"  @click="touch(props)">{{props.item.moment}}</td>
            </tr>
          </template>
          <template slot="no-data"><table></table></template>
        </v-data-table>
        <template v-if="!(orders.length > 0)">
            <v-layout theme--light align-center justify-center xs24 column alert alert-table>
                <i class="text-color body-10 icon-myexpay icon-no_smart-contract"></i>
                <svg-icon name="no-orders"style="width: 40px; height: 40px;" class="mb-2"/>
                <h3 class="text-color">Ордеров в данный момент нет</h3>
            </v-layout>
        </template>

        <div class="pagination--table text-xs-center text-md-right" v-if="(orders.length > pagination.rowsPerPage)">
            <pagination-for-table :pagination="pagination"></pagination-for-table>
        </div>
        <v-dialog :fullscreen="$vuetify.breakpoint.xsOnly" v-model="dialog" width="460" content-class="theme--light dialog--order" transition="dialog-opacity-transition">
            <v-layout>
                <v-flex>
                    <v-layout align-center justify-center column class="dialog--header">
                        <h2 class="body-2">Покупка {{from}}</h2>
                        <a @click="dialog = null" class="dialog--close-button"><v-icon small>mdi-close</v-icon></a>
                    </v-layout>
                    <v-layout justify-center px-md-5 px-xs-2 py-5 class="dialog--content" wrap >
                        <v-flex xs24>
                            <v-layout align-center justify-space-between row>
                                <v-flex pl-0 style="flex: 0 0 40px;"><svg-icon name="wallet-small-invert" style="width: 30px; height: 24px;"/></v-flex>
                                <v-flex fill-height my-0 py-0 overflow-hidden>
                                    <v-layout fill-height wrap>
                                        <v-flex xs24 class="px-0 pb-0"><span class="body-3 text-color">Баланс: </span><span class="font-weight-bold body-3">{{availableBalance(pair.to.balance.available,pair.to.account.currency.symbol)}} {{pair.to.account.currency.symbol}}</span></v-flex>
                                        <v-flex xs24 class="pt-0">
                                            <v-layout>
                                                <v-flex px-0 py-0>
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
                        <v-flex xs22 md18>
                            <v-layout column wrap>
                                <v-flex xs24 mb-2>
                                    <span class="mt-3 label body-1 pb-2 d-block font-weight-bold">Количество {{from}}:</span>
                                    <v-text-field class="table-input" block single-line outline color="" v-model="amount"></v-text-field>
                                </v-flex>
                                <v-flex xs24 mb-2>
                                    <span class="label body-1 pb-2 d-block font-weight-bold">Итоговая цена</span>
                                    <span class="body-5 d-block data pb-3  d-block">{{amountBuy}} {{to}}</span>
                                </v-flex>
                                <v-flex xs24>
                                    <v-btn
                                            @click="buy"
                                            color="" block flat  class="mb-0 v-btn-style v-btn__type3">
                                        Купить
                                    </v-btn>
                                </v-flex>
                            </v-layout>
                        </v-flex>
                    </v-layout>
                </v-flex>
            </v-layout>
        </v-dialog>
    </div>
</template>
<script>
  import moment from 'moment'
  import PaginationForTable from '~/components/PaginationForTable'
  export default {
    props: ['orders', 'scenario', 'classtable', 'selected', 'accounts', 'pair'],
    components: {
      PaginationForTable
    },
    methods: {
      toRound (value, code) {
        const length = this.lengths[code] ? this.lengths[code] : 0
        const coef = Math.pow(10, length)
        return Math.round(value * coef) / coef
      },
      availableBalance(value,symbol) {
        return value > 0 ? this.toRound((value), symbol) : 0
      },
      roundedValue (value, code) {
        const decimal = this.accounts.find(item => (item.currency.symbol === code)).currency.decimal
        const coef = Math.pow(10, decimal)
        return Math.round(value * coef) / coef
      },
      hideEmpty (selector) {
        if (this.orders.length < 1) {
          if (this.$el) this.$el.querySelector(selector).style.display = 'none'
        } else {
          if (this.$el) this.$el.querySelector(selector).style.display = ''
        }
      },
      onResize () {
        setTimeout(() => {
          this.heightComponent = this.$el.clientHeight
          this.hideEmpty('tbody')
          const newPagValue = Math.floor((this.heightComponent - (28 + 30)) / 28)
          if (newPagValue > 0) {
            this.pagination.rowsPerPage = this.mobile ? 7 : newPagValue
          }
        }, 1)
      },
      query() {
        var q = {}
        q.orderId = this.id
        q.sendingAddress = this.send
        q.receiveAddress = this.receive
        q.buyerId = this.buyerId
        q.amount = this.amount
        if (!!q.amount <= this.available
          && !!q.orderId
          && !!q.sendingAddress
          && !!q.receiveAddress
          && !!q.orderId && !!q.buyerId) {
          return q
        }
        return false
      },
      price (item) {
        return item.price
      },
      moment: function (data) {
        return moment(data)
      },
      touch (props, options) {
        if(!this.isUser(props.item)) {
          this.dialog = true;
          this.select = props.item;
        }
      },
      buy() {
        if (this.query) {
          let url = process.env.apiUrl + '_v1/create_contract'
          let snotify = this.$snotify
          return this.$axios.post(url, this.query())
            .then(function(success) {
              success = success.data;
              console.log(success);
              snotify.success('Контракт успешно создан', 'Создание контракта', {
                timeout: 4000,
                showProgressBar: false,
                closeOnClick: false,
                pauseOnHover: true
              });
            })
            .catch(function(error) {
              console.log(error);
              snotify.error('При создании контракта что-то пошло не так', 'Ошибка', {
                timeout: 1000,
                showProgressBar: false,
                closeOnClick: false,
                pauseOnHover: true
              });
            });
        }
      },
      toRender (code) {
        return Boolean(this.headers.find(item => (item.value === code)))
      },
      changeSort (column) {
        if (this.pagination.sortBy === column) {
          this.pagination.descending = !this.pagination.descending
        } else {
          this.pagination.sortBy = column
          this.pagination.descending = false
        }
      },
      isUser (order) {
        return order.ownerId === this.$auth.user.id
      },
      filter: function (order) {
        if (this.scenario === 'sell') {
          this.$store.commit('objects/selectOrderSell', order)
        } else if  (this.scenario === 'buy') {
          this.$store.commit('objects/selectOrderBuy', order)
        }
      },
    },
    watch: {
      orders (values) {
        this.$nextTick(() => {
          this.pagination.totalItems = this.orders.length
          this.onResize()
          this.filter(this.orders)
        })
      },
      dialog () {
        if(this.dialog == false) this.select = {}
      }
    },
    computed: {
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
      amountBuy() {
        return this.toRound((this.amount * this.select.price), this.pair.to.account.currency.symbol)
      },
      wsaccs () {
        return this.$store.state.objects.WsAccounts
      },
      id () {
        return this.select.id
      },
      buyerId () {
        return this.$auth.user.id
      },
      send () {
        var acc = this.wsaccs.filter((a) => {
          return a.account.currency.symbol === this.to
        })[0]
        if (!!acc.account) {
          return acc.account.Address
        }
        return false
      },
      receive () {
        var acc = this.wsaccs.filter((a) => {
          return a.account.currency.symbol === this.from
        })[0]
        if (!!acc.account) {
          return acc.account.Address
        }
        return false
      },
      from () {
        return this.select.sellCurrencySymbol
      },
      to () {
        return this.select.buyCurrencySymbol
      },
      available () {
        return this.select.available
      },
      items: function () {
        return this.orders.map(item => (Object.assign({}, item, {
          //orderValue: item,
          //available: this.available(item),
          price: this.price(item),
          moment: moment(new Date(item.createdAt * 1000)).format('MM.DD.YYYY — HH:mm'),
          user: 'User'
        })))
      },
      mobile: function () {
        return this.$vuetify.breakpoint.smAndDown
      },
      pages () {
        if (this.pagination.rowsPerPage == null ||
          this.pagination.totalItems == null
        ) return 0

        return Math.ceil(this.pagination.totalItems / this.pagination.rowsPerPage)
      },
      headers () {
        return this.allHeaders.filter((item) => (!item.filterOptions.notMobile || (item.filterOptions.notMobile && !this.mobile)))
      }
    },
    mounted () {
      this.isHydrated = true
      /*this.$nextTick(() => {
        this.onResize()
        this.hideEmpty('tbody')
      })*/
    },
    data () {
      return {
        amount: 0,
        isHydrated: false,
        heightComponent: 0,
        pagination: {
          sortBy: 'available'
        },
        allHeaders: [
          { text: '', value: 'available', align: 'center', id: 'sell', filterOptions: {} },
          { text: `Цена`, value: 'price', align: 'center', id: 'buy', filterOptions: {} },
          { text: `Дата создания ордера`, value: 'moment', align: 'center', id: 'date', filterOptions: {} }
        ],
        dialog: false,
        select: {}
      }
    }
  }
</script>