export const state = () => ({
  WsAccounts: [],
  orders: [],
  paysystems: [],
})

export const mutations = {
  initOrders (state, ords) {
    state.orders = ords
  },
  addOrder (state, ord) {
    state.orders.push(ord)
  },
  updateOrder (state, ord) {
    // Найти другой способ менять данные в массиве
    state.orders = state.orders.filter((o)=>{
      return o.id !== ord.id
    })
    state.orders.push(ord)
    state.orders.reverse()
  },
  addAccount (state, acc) {
    state.WsAccounts.push(acc)
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

export const actions = {
  InitAccounts ({commit}, data) {
    if (!!data['accounts']) {
      data['accounts'].forEach((acc)=>{
        commit('addAccount', acc)
      })
    }
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
  Event_OrderChange ({commit}, {data}) {
    if (!!data) {
      commit('updateOrder', data.Order)
    }
  },
  Event_BalanceChange({commit}, {data}) {
    commit('updateAccountBalance', data['AccountBalance'])
  },
}