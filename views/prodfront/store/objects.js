export const state = () => ({
  WsAccounts: [],
  orders: [],
  paysystems: [],
})
function selectOrder (state, order, type, options) {
  if (order) {
    if (options && options.force) {
      state[type] = [order]
    } else {
      const exist = state[type].findIndex(item => (item.id === order.id))
      if (exist !== -1) {
        if (!options || (options && !options.notUnselect)) {
          state[type] = state[type].slice(0, exist).concat(state[type].slice(exist + 1))
        }
      } else {
        state[type] = state[type].concat(order)
      }
    }
  }
}
export const mutations = {

  selectOrderBuy (state, data) {
    selectOrder(state, data.order, 'selectedBuy', data.options)
  },
  selectOrderSell (state, data) {
    selectOrder(state, data.order, 'selectedSell', data.options)
  },
  initOrders (state, ords) {
    state.orders = ords
  },
  addOrder (state, ord) {
    state.orders.push(ord)
    this._vm.$snotify.info(`Новый ордер ${ord.id} от ${ord.ownerId}`, `Новый ордер`, {
      timeout: 2000,
      showProgressBar: false,
      closeOnClick: false,
      pauseOnHover: true
    })
  },
  updateOrder (state, ord) {
    state.orders = state.orders.filter((o)=>{
      return o.id !== ord.id
    })
    state.orders.push(ord)
    state.orders.reverse()
  },
  addAccount (state, acc) {
    state.WsAccounts.push(acc)
  },
  addNotWSAcc (state, acc) {


    state.WsAccounts.push({account: acc, balance: {accountId: acc.id}})
    this.commit('currency/InitPair', state)
  },
  updateAccountBalance (state, balance) {
    var accKey = state.WsAccounts.findIndex((a) => {
      return a.account.id === balance.accountId
    })
    if (accKey >= 0) {
      var wsOld = state.WsAccounts.slice()
      var wsacc = wsOld[accKey]
      wsacc.balance = balance
      state.WsAccounts = state.WsAccounts.filter((a) => {
        return  a.account.id !== wsacc.account.id
      })
      state.WsAccounts.push(wsacc)
      state.WsAccounts.reverse()
    }

  },
  setPaysystems (state, systems) {
    state.paysystems = systems
  },
}

export const getters = {
  items (state) {
    return state.orders
  },
  allItems (state) {
    return state.orders
  },
  ordersForSell (state) {
    return () => state.ordersForSell
  },
  ordersForBuy (state) {
    return () => state.ordersForBuy
  },
  pairFilteredItems (state, getters, store) {
    if (store.auth !== undefined) {
      let user = store.auth.user
      return (pair) => {
        if (!state.orders || !pair.from.account || !pair.from.account.currency.symbol || !pair.to.account || !pair.to.account.currency.symbol) {
          return []
        }
        return state.orders.filter((i) => {
          return user.id !== i.ownerId && i.sellCurrencySymbol === pair.from.account.currency.symbol && i.buyCurrencySymbol === pair.to.account.currency.symbol
        })
      }
    }
  },
  invertedPairFilteredItems (state, getters, store) {
    if (store.auth !== undefined) {
      let user = store.auth.user
      return (pair) => {
        if (!state.orders || !pair.from.account || !pair.from.account.currency.symbol || !pair.to.account || !pair.to.account.currency.symbol) {
          return []
        }
        return state.orders.filter((i) => {
          return user.id !== i.ownerId && i.sellCurrencySymbol === pair.to.account.currency.symbol && i.buyCurrencySymbol === pair.from.account.currency.symbol
        })
      }
    }
  },
  userItems (state, getters, store) {
    if (store.auth !== undefined) {
      let user = store.auth.user
      return (pair) => {
        if (!state.orders || !pair.from.account || !pair.from.account.currency.symbol || !pair.to.account || !pair.to.account.currency.symbol) {
          return []
        }
        console.log(state.orders.filter((i) => {
          return user.id === i.ownerId && i.sellCurrencySymbol === pair.from.account.currency.symbol && i.buyCurrencySymbol === pair.to.account.currency.symbol
        }))
        return state.orders.filter((i) => {
          return user.id === i.ownerId && i.sellCurrencySymbol === pair.from.account.currency.symbol && i.buyCurrencySymbol === pair.to.account.currency.symbol
        })
      }
    }
  },
  userItemsHistory (state, getters, store) {
    if (store.auth !== undefined) {
      var user = store.auth.user
    }
    return (pair) => {
      if (!state.allItems || !pair.from || !pair.from.symbol || !pair.to || !pair.to.symbol) {
        return []
      }
      return state.allItems.filter((i) => {
        return user.id === i.ownerId && i.sellCurrencySymbol === pair.from.symbol && i.buyCurrencySymbol === pair.to.symbol && (!i.available || i.available === 0)
      })
    }
  },
  twoLastItems (state, getters, store) {
    if (store.auth !== undefined) {
      var user = store.auth.user
    }
    return () => {
      return state.items.filter((i) => {
        return user.id === i.ownerId
      }).reverse().slice(0, 2)
    }
  }
}


export const actions = {
  InitAccounts ({commit}, data) {
    if (!!data['accounts']) {
      data['accounts'].forEach((acc)=>{
        commit('addAccount', acc)
      })
    }
    this.commit('currency/InitPair', data)
  },
  InitOrders ({commit}, data) {
    if (!!data['orders']) {
      commit('initOrders', data['orders'])
    }
  },
  Event_NewOrder ({commit}, {data}) {
    if (!!data) {
      commit('addOrder', data.Order)

    }
  },

  Event_NewAccount ({commit}, {data}) {
    // addNotWSAcc потому что новые акки появляются без balance части
    if (!!data['Account']) {
      commit('addNotWSAcc', data['Account'])
    }
  },
  Event_OrderChange ({commit}, {data}) {
    if (!!data) {
      commit('updateOrder', data.Order)
    }
  },
  Event_BalanceChange({commit}, {data}) {
    commit('updateAccountBalance', data['AccountBalance'])
  },
}