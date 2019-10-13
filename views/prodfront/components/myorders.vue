<template>
    <div v-resize="onResize" :class=classtable>
        <v-data-table
                expand
                :headers="headers"
                :items="items"
                :disable-initial-sort=true
                :pagination.sync="pagination"
                :custom-sort="customSort"
                prev-icon="mdi-menu-left"
                next-icon="mdi-menu-right"
                sort-icon="mdi-chevron-down"
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
                        <template v-if="header.id == 'sell' && orders[0]">
                            {{scenario === 'sell' ? orders[0].buyCurrencySymbol : orders[0].sellCurrencySymbol}}
                        </template>
                        <template v-else-if="header.id == 'buy' && orders[0]">
                            Цена ({{scenario === 'sell' ? orders[0].sellCurrencySymbol : orders[0].buyCurrencySymbol}})
                        </template>
                        <template v-else="header.id">
                            {{header.text}}
                        </template>
                    </th>
                </tr>
            </template>
            <template slot="items" slot-scope="order">
                <tr @click="order.expanded = !order.expanded">
                    <td v-if="toRender('currency_pair')" class="text-xs-center">{{order.item.currency_pair}}</td>
                    <td v-if="toRender('amount')" class="text-xs-center">{{order.item.amount}}</td>
                    <td v-if="toRender('price')" class="text-xs-center">{{order.item.price}}</td>
                    <td v-if="toRender('sum')" class="text-xs-center">{{order.item.sum}}</td>
                    <td v-if="toRender('percentProgress')" class="text-xs-center" style="width: 100px">
                        <v-layout justify-center align-center>
                            <v-flex v-if="!mobile" class="theme--light progress-line">
                                <!--<v-progress-linear
                                        color="blue"
                                        height="15"
                                        :value="order.item.percentProgress">
                                </v-progress-linear>{{order.item.percentProgress}}-->
                            </v-flex>
                            <v-flex>
                                {{order.item.progress}}
                            </v-flex>
                        </v-layout>
                    </td>
                    <td v-if="toRender('operation')" class="text-xs-center">{{order.item.operation}}</td>
                    <td v-if="toRender('date_of_create')" class="text-xs-center">{{moment(order.item.date_of_create)}}</td>
                    <td v-if="toRender('date_of_close')" class="text-xs-center">{{order.item.date_of_close ? (moment(order.item.date_of_close)) : ''}}</td>
                    <td v-if="toRender('user')" class="text-xs-center">{{order.item.user}}</td>
                    <td v-if="toRender('status')" class="text-xs-center">{{order.item.status}}</td>
                    <td v-if="toRender('actions')" class="text-xs-center" style="width: 100px">
                        <v-btn :ripple="false" class="mx-0 my-0 v-btn-style btn-action-table" outline small @click="openDialog(order.item)" flat>
                            x
                        </v-btn>
                    </td>
                    <td v-if="toRender('simulation')" class="text-xs-center" style="width: 100px">
                        <v-btn :ripple="false" class="mx-0 my-0 v-btn-style btn-action-table" outline small  @click="openSimDialog(order.item)" flat>
                            Simulation
                        </v-btn>
                    </td>
                </tr>
            </template>
            <template slot="no-data"><div></div></template>
            <template v-if="mobile" slot="expand" slot-scope="order">
                <td v-if="toRenderInExpand('currency_pair')" class="text-xs-center">trading pair {{order.item.currency_pair}}</td>
                <td v-if="toRenderInExpand('amount')" class="text-xs-center">amount {{order.item.amount}}</td>
                <td v-if="toRenderInExpand('price')" class="text-xs-center">price {{order.item.price}}</td>
                <td v-if="toRenderInExpand('sum')" class="text-xs-center">sum {{order.item.sum}}</td>
                <td v-if="toRenderInExpand('percentProgress')" class="text-xs-center" style="width: 100px">
                    progress
                    <v-layout justify-center align-center>
                        <v-flex>
                            {{order.item.progress}}
                        </v-flex>
                    </v-layout>
                </td>
                <td v-if="toRenderInExpand('operation')" class="text-xs-center">operation {{order.item.operation}}</td>
                <td v-if="toRenderInExpand('date_of_create')" class="text-xs-center">date of create {{moment(order.item.date_of_create)}}</td>
                <td v-if="toRenderInExpand('date_of_close')" class="text-xs-center">date of close {{order.item.date_of_close ? (moment(order.item.date_of_close)) : ''}}</td>
                <td v-if="toRenderInExpand('user')" class="text-xs-center">user {{order.item.user}}</td>
                <td v-if="toRenderInExpand('status')" class="text-xs-center">status {{order.item.status}}</td>
                <td v-if="toRenderInExpand('actions')" class="text-xs-center" style="width: 100px">
                    <v-btn :ripple="false" class="mx-0 my-0 v-btn-style btn-action-table" outline small @click="openDialog(order.item)" flat>
                        Отмена
                    </v-btn>
                </td>
                <td v-if="toRenderInExpand('simulation')" class="text-xs-center" style="width: 100px">
                    <v-btn :ripple="false" class="mx-0 my-0 v-btn-style btn-action-table" outline small  @click="openSimDialog(order.item)" flat>
                        Simulation
                    </v-btn>
                </td>
            </template>
        </v-data-table>
        <template v-if="!(orders.length > 0)">
            <v-layout theme--light align-center justify-center xs24 column alert alert-table>
                <svg-icon name="no-orders"style="width: 40px; height: 40px;" class="mb-2"/>
                <h3 class="text-color">Нет контрактов</h3>
            </v-layout>
        </template>
        <v-dialog
                v-model="dialog"
                width="500">
            <v-card>
                <v-card-title
                        class="headline grey lighten-2"
                        primary-title>
                    Отменить ордер
                </v-card-title>
                <v-divider></v-divider>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn
                            color="primary"
                            flat
                            @click="cancelOrder(dialogValue)">
                        Да
                    </v-btn>
                    <v-btn
                            color="primary"
                            flat
                            @click="dialog = null">
                        Отмена
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <template v-if="((orders.length > pagination.rowsPerPage) && (scenario !== 'allOrders'))">
            <div class="pagination--table text-xs-center text-md-right">
                <pagination-for-table :pagination="pagination"></pagination-for-table>
            </div>
        </template>
    </div>
</template>

<style>
    tr.expand td {
        padding: 0 !important;
    }

    tr.expand .expansion-panel {
        box-shadow: none;
    }

    tr.expand .expansion-panel li {
        border: none;
    }
</style>

<script>
  import PaginationForTable from '~/components/PaginationForTable'
  import moment from 'moment'
  export default {
    props: [ 'orders', 'classtable', 'scenario', 'pair', 'userId' ],
    components: {
      PaginationForTable,
    },
    mounted () {
      this.$nextTick(() => {
        this.onResize()
      })
    },
    beforeUpdate () {
      this.$nextTick(() => {
        this.onResize()
      })
    },
    methods: {
      toRound (value, code) {
        const length = this.lengths[code] ? this.lengths[code] : 0
        const coef = Math.pow(10, length)
        return Math.round(value * coef) / coef
      },
      customSort (items, search, isDescending) {
        const onlyNumbers = /\d+/
        if (['currency_pair', 'operation', 'user', 'status'].indexOf(search) !== -1) {
          return items.sort((item1, item2) => (isDescending ? item1[search] > item2[search] : item1[search] < item2[search]))
        } else if (['amount', 'price', 'sum', 'percentProgress', 'date_of_create', 'date_of_close'].indexOf(search) !== -1) {
          return items.sort((item1, item2) => {
            const item1numbers = onlyNumbers.exec(item1[search])
            const item2numbers = onlyNumbers.exec(item2[search])
            if (item1numbers.length && item2numbers.length) {
              return isDescending ? Number(item1numbers[0]) > Number(item2numbers[0]) : Number(item1numbers[0]) < Number(item2numbers[0])
            } else {
              return isDescending ? item1[search] > item2[search] : item1[search] < item2[search]
            }
          })
        }
        return items
      },
      hideEmpty (selector) {
        if (this.orders.length < 1) {
          if (this.$el && this.$el.querySelector) this.$el.querySelector(selector).style.display = 'none'
        } else {
          if (this.$el && this.$el.querySelector) this.$el.querySelector(selector).style.display = ''
        }
      },
      onResize () {
        setTimeout(() => {
          if (this.scenario === 'allOrders') {
            this.hideEmpty('tbody')
            this.pagination.rowsPerPage = 2
            return
          }
          this.heightComponent = this.$el.clientHeight
          const newPagValue = Math.floor((this.heightComponent - (28 + 30)) / 28)
          if (newPagValue > 0) {
            this.pagination.rowsPerPage = this.mobile ? 7 : newPagValue
            this.hideEmpty('tbody')
          } else {
            setTimeout(() => this.onResize(), 2000)
          }
        }, 1)
      },
      touch (props, options) {
        if (this.mobile) {
          props.expanded = !props.expanded
        } else {
          // this.select(props.item, options)
        }
      },
      changeSort (column) {
        if (this.pagination.sortBy === column) {
          this.pagination.descending = !this.pagination.descending
        } else {
          this.pagination.sortBy = column
          this.pagination.descending = false
        }
      },
      cancelOrder (order) {
        this.dialog = null
        var that = this
        that.$store.dispatch('order/cancelOrder', order.id)
      },
      contractorId (local) {
        return this.userId === local.sellerId ? local.buyerId : local.sellerId
      },
      status (item) {
        return !item.status.status ? 'Создан' : 'Исполнен'
      },
      usersCount (item) {
        const length = this.contractsByOrder(item.id).length
        return length ? `${this.contractsByOrder(item.id).length} users` : (item.ownerId === this.userId ? '' : this.userId)
      },
      operation (item) {
        return this.routeIsBuy(item) ? 'buy' : 'sell'
      },
      onPrice (value, item) {
        return (Number(item.price) !== Number(item.frontMetaData.userPriceInput)) ? value / item.frontMetaData.userPriceInput : value * item.price
      },
      progress (item) {
        const amountOfContracts = this.contractsByOrder(item.id).reduce((accum, cur) => (accum + cur.amount), 0)
        if (this.routeIsBuy(item)) {
          return `${(this.toRound(this.onPrice((amountOfContracts || 0), item), item.buyCurrencySymbol))}/${this.toRound(this.onPrice(item.amount, item), item.buyCurrencySymbol)}  ${item.buyCurrencySymbol}`
        } else {
          return `${this.toRound(item.amount - (item.available || 0), item.sellCurrencySymbol)}/${this.toRound(item.amount, item.sellCurrencySymbol)}  ${item.sellCurrencySymbol}`
        }
      },
      percentProgress (item) {
        const amountOfContracts = this.contractsByOrder(item.id).reduce((accum, cur) => (accum + cur.amount), 0)
        return (amountOfContracts || 0) / item.amount * 100
      },
      price (item) {
        return this.routeIsBuy(item) ? `${this.toRound((Number(item.price) !== Number(item.frontMetaData.userPriceInput)) ? item.frontMetaData.userPriceInput : 1 / item.price, item.sellCurrencySymbol)}  ${item.sellCurrencySymbol}` : `${this.toRound(item.price, item.buyCurrencySymbol)}  ${item.buyCurrencySymbol}`
      },
      amount (item) {
        return this.routeIsBuy(item) ? `${this.toRound(this.onPrice(item.amount, item), item.buyCurrencySymbol)} ${item.buyCurrencySymbol}` : `${this.toRound(item.amount, item.sellCurrencySymbol)} ${item.sellCurrencySymbol}`
      },
      sum (item) {
        if (this.routeIsBuy(item)) {
          const value = ((Number(item.price) !== Number(item.frontMetaData.userPriceInput)) ? item.frontMetaData.userPriceInput : 1 / item.price) * this.onPrice(item.amount, item)
          return `${this.toRound(value, item.sellCurrencySymbol)}  ${item.sellCurrencySymbol}`
        } else {
          const value = Number(item.price) * Number(item.amount)
          return `${this.toRound(value, item.buyCurrencySymbol)}  ${item.buyCurrencySymbol}`
        }
      },
      amountLocal (item, local) {
        return this.routeIsBuy(item) ? `${this.onPrice(local.available, item)}  ${item.buyCurrencySymbol}` : `${local.available}  ${item.sellCurrencySymbol}`
      },
      currencyPair (item) {
        return this.routeIsBuy(item) ? `${item.buyCurrencySymbol} > ${item.sellCurrencySymbol}` : `${item.sellCurrencySymbol} > ${item.buyCurrencySymbol}`
      },
      moment: function (data) {
        return moment(data * 1000).format('MM.DD.YYYY — HH:mm')
      },
      select: function (order) {
        this.selected = order.id
      },
      contractsByOrder (orderId) {
        return this.userSellContracts.filter((i) => {
          return i.orderId === orderId
        })
      },
      toRender (code) {
        return Boolean(this.headers.find(item => (item.value === code)))
      },
      toRenderInExpand (code) {
        return Boolean(this.requidData.find(item => (item.value === code)) && !this.headers.find(item => (item.value === code)))
      },
      routeIsBuy (item) {
        console.log(this.pair)
        if (this.scenario === 'allOrders' || this.scenario === 'contracts') {
          return (Number(item.frontMetaData.userPriceInput) !== item.price)
        } else {
          return (this.pair.from.account.currency.symbol === item.buyCurrencySymbol && this.pair.to.account.currency.symbol === item.sellCurrencySymbol)
        }
      },
      openSimDialog (order) {
        this.simDialog = order
        this.simDialogValue = order
      },
      openDialog (item) {
        // this.$router.push(`/order?id=${item.id}`)
        this.dialog = item
        this.dialogValue = item
      }
    },
    computed: {
      accounts () {
        return this.$store.getters['account/symbolFilteredItems']
      },
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
      mobile: function () {
        return this.$vuetify.breakpoint.smAndDown
      },
      items () {
        return this.orders.map(item => (Object.assign({}, item, {
          currency_pair: this.currencyPair(item),
          amount: this.amount(item),
          price: this.price(item),
          sum: this.sum(item),
          percentProgress: this.percentProgress(item),
          progress: this.progress(item),
          operation: this.operation(item),
          date_of_create: item.createdAt,
          date_of_close: item.status.status ? item.status.createdAt : null,
          user: this.usersCount(item),
          status: this.status(item)
        })))
      },
      userSellContracts () {
        console.log(this.$store.getters['objects/items'])
        return this.$store.getters['objects/items'].filter((i) => {
          return (i.sellerId === this.$auth.user.id) || (i.buyerId === this.$auth.user.id)
        })
      },
      requidData () {
        return this.allHeaders.filter((item) => (item.filterOptions.all || item.filterOptions[this.scenario]))
      },
      headers () {
        const res = this.requidData
        if (this.mobile) {
          return res.slice(0, 3)
        } else {
          return res
        }
      }
    },
    watch: {
      orders (values) {
        this.$nextTick(() => {
          this.pagination.totalItems = this.orders.length
        })
      }
    },
    data () {
      return {
        pagination: {
          sortBy: 'available'
        },
        selected: '',
        dialog: null,
        simDialog: null,
        simDialogValue: null,
        dialogValue: null,
        expanded: {},
        expandedOrder: '',
        allHeaders: [
          { text: 'Валютная пара', value: 'currency_pair', align: 'center', filterOptions: { allOrders: true } },
          { text: 'Количество', value: 'amount', align: 'center', filterOptions: { all: true } },
          { text: 'Цена', value: 'price', align: 'center', filterOptions: { all: true } },
          { text: 'Сумма', value: 'sum', align: 'center', filterOptions: { tradingHistory: true } },
          { text: 'Прогресс', value: 'percentProgress', align: 'center', filterOptions: { myOrders: true, allOrders: true } },
          { text: 'Операция', value: 'operation', align: 'center', filterOptions: { allOrders: true, tradingHistory: true } },
          { text: 'Создана', value: 'date_of_create', align: 'center', filterOptions: { myOrders: true, allOrders: true, history: true, contracts: true } },
          { text: 'Закрыта', value: 'date_of_close', align: 'center', filterOptions: { history: true, tradingHistory: true } },
          { text: 'Статус', value: 'status', align: 'center', filterOptions: { myOrders: true, history: true } },
          { text: 'Операции', value: 'actions', align: 'center', filterOptions: { allOrders: true, myOrders: true } },
          { text: 'simulation', value: 'simulation', align: 'center', filterOptions: {} }
        ]
      }
    }
  }
</script>