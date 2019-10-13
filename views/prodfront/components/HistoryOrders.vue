<template>
    <div v-resize="onResize" :class=classtable>
        <v-data-table
            :headers="headers"
            :items="orders"
            :pagination.sync="pagination"
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
                    ><v-icon small>mdi-chevron-down</v-icon>
                    {{header.text}}
                    </th>
                </tr>
            </template>
            <template slot="items" slot-scope="props">
                <td class="text-xs-center">{{ props.item.price }} {{props.item.buyCurrencySymbol}}</td>
                <td class="text-xs-center">{{ props.item.available }} {{props.item.sellCurrencySymbol}}</td>
                <td class="text-xs-center">{{ props.item.amount }} {{props.item.buyCurrencySymbol}}</td>
                <!--<td class="text-xs-center">Operation</td>-->
                <td class="text-xs-center">{{moment(new Date(props.item.createdAt * 1000)).format('MM.DD.YYYY — HH:mm')}}</td>
            </template>
            <template slot="no-data"><div></div></template>

        </v-data-table>
        <template v-if="!(orders.length > 0)">
            <v-layout theme--light align-center justify-center xs24 column alert alert-table>
                <i class="text-color body-10 icon-no_smart-contract"></i>
                <h3 class="text-color">No contracts</h3>
            </v-layout>
        </template>
    </div>
</template>
<script>
  import moment from 'moment'
  export default {
    props: ['orders', 'classtable'],
    data: () => ({
      pagination: {},
      heightComponent: 0,
      dialog: false,
      headers: [
        { text: 'Цена', value: 'price' },
        { text: 'Количество', value: 'available' },
        { text: 'Сумма', value: 'amount' },
       // { text: 'Операция', value: 'operation' },
        { text: 'Дата', value: 'date' }
      ]
    }),
    mounted () {
      this.$nextTick(() => {
        this.onResize()
        this.hideEmpty('tbody')
      })
    },
    computed: {
      formTitle () {
        return this.editedIndex === -1 ? 'New Item' : 'Edit Item'
      }
    },

    watch: {
      dialog (val) {
        val || this.close()
      }
    },
    methods: {
      hideEmpty (selector) {
        if (this.orders.length < 1) {
          if (this.$el) this.$el.querySelector(selector).style.display = 'none'
        } else {
          if (this.$el) this.$el.querySelector(selector).style.display = ''
        }
      },
      moment: function (data) {
        return moment(data)
      },
      onResize () {
        this.heightComponent = this.$el.clientHeight
        if (this.pagination.rowsPerPage) this.pagination.rowsPerPage = Math.floor((this.heightComponent - (28 + 30)) / 28)
      },
      changeSort (column) {
        if (this.pagination.sortBy === column) {
          this.pagination.descending = !this.pagination.descending
        } else {
          this.pagination.sortBy = column
          this.pagination.descending = false
        }
      }
    }
  }
</script>