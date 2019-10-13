export const state = () => ({
  timerOn: true
})

export const mutations = {
  off (state, items) {
    state.timerOn = false
  },
  on (state, items) {
    state.timerOn = true
  }
}

export const getters = {}
export const actions = {}
