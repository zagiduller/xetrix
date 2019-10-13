import Vue from 'vue'
import VueNativeSock from 'vue-native-websocket'


export default ({ store, app  }, inject) => {
  Vue.use(VueNativeSock, process.env.wsServer, {
    connectManually: !app.$auth.loggedIn ,
    format: 'json',
    store: store,
    reconnection: true, // (Boolean) whether to reconnect automatically (false)
    reconnectionAttempts: 5, // (Number) number of reconnection attempts before giving up (Infinity),
    reconnectionDelay: 5000, // (Number) how long to initially wait before attempting a new (1000)
  })
}