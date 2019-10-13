export const state = () => ({
  pair: {},
  loading: false,
  items: [],
  history: {}
})

export const mutations = {
  InitPair (state, objects) {
    let from, to = {}
    if(objects.accounts !== undefined) {
      from = objects.accounts.find(item => (item.account.currency.symbol === 'BTC')) || objects.accounts[0]
      to = objects.accounts.find(item => (item.account.currency.symbol === 'USD')) || objects.accounts[1]
    } else {
      from = objects.WsAccounts.find(item => (item.account.currency.symbol === 'BTC')) || objects.WsAccounts[0]
      to = objects.WsAccounts.find(item => (item.account.currency.symbol === 'USD')) || objects.WsAccounts[1]
    }
    state.pair = {from: from, to: to}
  },
  setPair (state, pair) {
    state.pair = pair
  },
  updateItems (state, items) {
    state.items = items
  },
  setCoursesByAPI (state, data) {
    if (data && data.data && data.data.Data && data.data.Data.length) {
      state.history = Object.assign({}, state.history, {
        [`${data.exchange}_${data.element1}`]: data
      })
    }
  },
  setLoading (state, loading) {
    state.loading = loading
  }
}

export const getters = {
  items (state) {
    return state.items
  }
}
function getDataFn (exchange, currency1, currency2, callback) {
  return this.$axios.get(`https://min-api.cryptocompare.com/data/pricemulti?fsyms=${currency1.join(',')}&tsyms=${currency2.join(',')}&e=${exchange}`).then((res) => {
    const current = res.data
    if (current[currency1[0]]) {
      currency1.forEach(element1 => {
        currency2.forEach(element2 => {
          this.$axios.get(`https://min-api.cryptocompare.com/data/histohour?fsym=${element1}&tsym=${element2}&limit=24&aggregate=1&e=${exchange}`).then((res) => {
            const proportion = res.data
            callback(null, {
              timeOf: new Date(),
              element1,
              element2,
              current,
              data: proportion,
              exchange
            })
          })
        })
      })
    } else {
      console.warn('No data', exchange, currency1, currency2)
    }
  })
}

function getCoursesByAPI (commit) {
  const getData = getDataFn.bind(this)
  try {
    getData('Binance', ['BTC', 'ETH'], ['USDT'], (code, data) => {
      commit('setCoursesByAPI', data)
    })
    getData('Bittrex', ['BTC', 'ETH'], ['USD'], (code, data) => {
      commit('setCoursesByAPI', data)
    })
  } catch (e) {
    console.warn(e)
  }
}

export const actions = {

  loadItems ({commit}) {
    getCoursesByAPI.call(this, commit)
    return this.$axios.get(`/v1/get_currency`).then((resp) => {
      commit('updateItems', resp.data.items)
      commit('setLoading', false)
    }).catch((e) => {
      console.error(e)
      // commit('feedback/setError', e)
    })
  },
  pullingItems ({ commit }) {
    commit('setLoading', true)
    getCoursesByAPI.call(this, commit)
    return this.$axios.get(`/v1/get_currency`).then((resp) => {
      commit('updateItems', resp.data.items)
      commit('setLoading', false)
    }).catch((e) => {
      console.error(e)
      // commit('feedback/setError', e)
    })
  },
  setPair ({commit}, data) {
    commit('setPair', data)
  },
  InitPair ({commit}, data) {
    //if (!!data['pair']) {
      commit('InitPair', data)
    //}
  },
}
