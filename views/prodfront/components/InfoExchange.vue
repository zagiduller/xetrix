<template>
    <div :class=classtable>
        <v-data-table
            :headers="headers"
            :items="information ? information : informationDefaul"
            hide-actions
            prev-icon="mdi-menu-left"
            next-icon="mdi-menu-right"
            sort-icon="mdi-chevron-down"
        >
            {{history}}
            <template slot="headers" slot-scope="props">
                <tr>
                    <th v-for="header in props.headers" :key="header.text">
                        {{header.text}}
                    </th>
                </tr>
            </template>
            <template slot="items" slot-scope="props">
                <td class="text-xs-center">{{ props.item.exchange }}</td>
                <td class="text-xs-center">{{ props.item.pair }}</td>
                <td class="text-xs-center">{{ props.item.price }}</td>
                <td class="text-xs-center">{{ Math.abs(props.item.change) }}% <v-icon style="font-size: 20px;line-height: 14px;" small large color="blue darken-2">{{(props.item.change < 0) ? 'mdi-menu-down' : 'mdi-menu-up'}}</v-icon></td>
                <td class="text-xs-center">{{ props.item.marketcap }}</td>
            </template>
        </v-data-table>
    </div>
</template>
<script>
  export default {
    props: ['classtable', 'history'],
    created () {
      this.$store.dispatch('currency/pullingItems')
    },
    computed: {
      information () {
        if (this.history) {
          return this.history.map(item => {
            function toFormat (value) {
              return new Intl.NumberFormat('ru-RU').format(value).replace(',', '.').split(' ').join(',')
            }
            const marketcap = toFormat(Math.round(item.data.Data.reduce((accum, cur) => (accum + cur.volumeto), 0)))
            const price = toFormat(item.current[item.element1][item.element2])
            return {
              exchange: item.exchange,
              pair: `${item.element1}/${item.element2}`,
              price: `${price} ${item.element2}`,
              change: item.data.Data[1] ? (Math.round(((item.current[item.element1][item.element2] - item.data.Data[0].open) / item.current[item.element1][item.element2]) * 10000) / 100) : '',
              marketcap: item.data.Data[1] ? `${marketcap} ${item.element2}` : ''
            }
          })
        } else {
          return [{
            exchange: 'Bittrex',
            pair: 'USD/BTC',
            price: '',
            change: '',
            marketcap: ''
          }, {
            exchange: 'Binance',
            pair: 'USD/BTC',
            price: '',
            change: '',
            marketcap: ''
          }]
        }
      }
    },
    data: () => ({
      dialog: false,
      headers: [
        {text: 'Биржа', value: 'exchange'},
        {text: 'Пара', value: 'pair'},
        {text: 'Цена', value: 'price'},
        {text: 'Изменение (24ч)', value: 'change'},
        {text: 'Объем рынка (24ч)', value: 'marketcap'}
      ],
      informationDefaul: [{
        exchange: 'Bittrex',
        pair: 'USD/BTC',
        price: '',
        change: '',
        marketcap: ''
      }, {
        exchange: 'Binance',
        pair: 'USD/BTC',
        price: '',
        change: '',
        marketcap: ''
      }]
    })
  }
</script>