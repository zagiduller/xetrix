<template>
  <v-layout align-center wrap>
    <template v-for="(from, elem) in currencies" >
      <v-flex px-0 text-xs-center xs6>
        <v-menu offset-y content-class="menu-pairs-block" ><!--:offset-x="elem > 0 ? '-150' :0"-->
          <template #activator="data" v-slot:activator>
            <v-btn v-on="data.on" active-class block flat :class="['mb-0 v-btn-style v-btn__type2', {active: pair.from.account.currency.symbol == from.account.currency.symbol}]"><img :src="'/images/icons/' + from.account.currency.symbol.toLowerCase() + '.png'"  @error="setFallbackImageUrl" width="20" alt="">&nbsp;{{from.account.currency.symbol}}</v-btn>
          </template>
          <v-layout px-4 py-3>
            <v-flex md12 pl-2 pr-3 text-xs-center>
              <span class="body-1 gray-lighten-1">Пара валют</span>
              <hr class="my-1">
              <v-btn v-for="(to, index) in filtredItems(pairs, from.account.currency.symbol)" v-if="index < 3" :key="filtredItems(pairs, from.account.currency.symbol).from" block flat :class="['mb-0 mt-0 px-2 v-btn-style v-btn__type4',{active: (pair.from.account.currency.symbol === to.from_symbol && pair.to.account.currency.symbol === to.to_symbol)}]" @click="check(to.from_object,to.to_object)">{{to.from_symbol}} <span class="body-7 d-inline-block px-3" style="height: 16px; line-height: 10px;">&#8250</span> {{to.to_symbol}}</v-btn>
            </v-flex>
            <v-flex md12 pr-2 pl-3 text-xs-center>
              <span class="body-1 gray-lighten-1">Пара валют</span>
              <hr class="my-1">
              <v-btn v-for="(to, index) in filtredItems(pairs, from.account.currency.symbol)" v-if="index > 2" :key="filtredItems(pairs, from.account.currency.symbol).from" block flat :class="['mb-0 mt-0 px-2 v-btn-style v-btn__type4',{active: (pair.from.account.currency.symbol === to.from_symbol && pair.to.account.currency.symbol === to.to_symbol)}]" @click="check(to.from_object,to.to_object)">{{to.from_symbol}} <span class="body-7 d-inline-block px-3" style="height: 16px; line-height: 10px;">&#8250</span> {{to.to_symbol}}</v-btn>
            </v-flex>
          </v-layout>
        </v-menu>
      </v-flex>
    </template>
  </v-layout>
</template>

<script>
  import { mapMutations } from 'vuex'
  export default {
    mounted () {
      this.from = this.currencies.find(item => (item.account.currency.symbol === 'BTC')) || ''
      this.to = this.currencies.find(item => (item.account.currency.symbol === 'USD')) || ''
    },
    props: ['accs', 'pair'],
    methods: {
      setFallbackImageUrl(event) {
        event.target.src = "/images/xetrix_not_found.png"
      },
      filtredItems (items, filter) {
        return items ? items.filter((item) => (item.from_symbol === filter || item.to_symbol === filter)) : []
        //return items ? items.filter((item) => (item !== filter)) : []
      },
      check: function (from, to) {
        //if (!!from_symbol && !!to_symbol) {
          this.$store.dispatch('currency/setPair', {from: from, to: to})
        //}
      },
    },
    computed: {
      currencies() {
        let crs = []
        this.accs.forEach((a) => {
          let exist = crs.findIndex((c) => {return a.account.currency.symbol === c.account.currency.symbol })
          if (exist < 0) {
            crs.push(a)
          }
        })
        return crs
      },
      pairs() {
        let pair = []
        this.accs.forEach((a) => {
          let data = this.accs.filter((d) => {
            return d.account.currency.symbol !== a.account.currency.symbol
          })
          data.forEach((b) => {
            pair.push({
              from: a.account.currency.id,
              from_symbol: a.account.currency.symbol,
              to: b.account.currency.id,
              to_symbol: b.account.currency.symbol,
              from_object: a,
              to_object: b
            })
          })
        })
        return pair
      }
    },
    data () {
      return {
        from: '',
        to: ''
      }
    }
  }
</script>