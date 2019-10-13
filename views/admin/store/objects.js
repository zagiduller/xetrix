import Vue from 'vue'

export const state = () => ({
  currencies: [],
  users: [],
  WsAccounts: [],
  orders: [],
  contracts: [],
  txs: [],
  paysystems: [],
  withdrawalOrders: [],
})

export const mutations = {
  initCurrencies(state, cs) {
    state.currencies = cs
  },
  initUsers (state, usrs) {
    state.users = usrs
  },
  initOrders (state, ords) {
    state.orders = ords
  },
  initContracts (state, cnts) {
    state.contracts = cnts
  },
  initTxs (state, txs) {
    state.txs = txs
  },
  initWithdrawal(state, wos) {
    state.withdrawalOrders = wos
  },
  addUser (state, user) {
    state.users.push(user)
  },
  addOrder (state, ord) {
    state.orders.push(ord)
  },
  addContract (state, contract) {
    state.contracts.push(contract)
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
  addNotWSAcc (state, acc) {
    state.WsAccounts.push({account: acc, balance: {accountId: acc.id}})
  },
  addTx (state, tx) {
    state.txs.push(tx)
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
  addWithdrawalOrder (state, wo) {
    state.withdrawalOrders.push(wo)
  },
  updateWithdrawalOrder (state, wo) {
    var woKey = state.withdrawalOrders.findIndex((el) => {
      return wo.id === el.id
    })
    console.log(woKey, Vue.set(state.withdrawalOrders, woKey, wo))
  },
}

export const actions = {
  InitCurrencies ({commit}, data) {
    if(!!data['currencies']) {
      commit('initCurrencies', data['currencies'])
    }
  },
  InitUsers ({commit}, data) {
    if (!!data['users']) {
      commit('initUsers', data['users'])
    }
  },
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
  InitContracts ({commit}, data) {
    if (!!data['contracts']) {
      commit('initContracts', data['contracts'])
    }
  },
  InitTxs ({commit}, data) {
    if (!!data['txs']) {
      commit('initTxs', data['txs'])
    }
  },
  InitWithdrawal ({commit}, data) {
    if (!!data['withdrawalOrders']) {
      commit('initWithdrawal', data['withdrawalOrders'])
    }
  },
  Event_NewUser ({commit}, {data}) {
    if (!!data) {
      commit('addUser', data.User)
    }
  },
  Event_NewAccount ({commit}, {data}) {
    // addNotWSAcc потому что новые акки появляются без balance части
    if (!!data['Account']) {
      commit('addNotWSAcc', data['Account'])
    }
  },
  Event_NewOrder ({commit}, {data}) {
    if (!!data['Order']) {
      commit('addOrder', data.Order)
    }
  },
  Event_NewContract ({commit}, {data}) {
    if (!!data['Contract']) {
      commit('addContract', data.Contract)
    }
  },
  Event_NewTransaction ({commit}, {data}) {
    if (!!data['Tx']) {
      commit('addTx', data.Tx)
    }
  },
  Event_OrderChange ({commit}, {data}) {
    if (!!data['Order']) {
      commit('updateOrder', data.Order)
    }
  },
  Event_BalanceChange({commit}, {data}) {
    commit('updateAccountBalance', data['AccountBalance'])
  },
  Event_NewWithdrawalOrder({commit}, {data}){
    commit('addWithdrawalOrder', data['WithdrawalOrder'])
  },
  Event_WithdrawalPerformed({commit}, {data}){
    commit('updateWithdrawalOrder', data['WithdrawalOrder'])
  },
  Event_TxConfirm({commit}, {data}){
    //Подтверждена транзакция
  },

}