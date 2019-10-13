export default ({ app, store }) => {
  setInterval(() => {
    if (store.state.auth.loggedIn && app.store.state.timer.timerOn) {
      app.store.dispatch('currency/pullingItems')
    }
  }, 50000)
}
