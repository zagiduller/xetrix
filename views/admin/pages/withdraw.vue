<template>
  <v-container fluid class="pa-0">
    <v-layout row>
      <v-flex md12>
        <v-tabs
                slot="extension"
                v-model="tab"
                dark
        >
          <v-tabs-slider color="yellow"></v-tabs-slider>

          <v-tab
                v-for="ps in paysystems"
                :key="ps.Name"
                v-if="psWos(ps.Name).length > 0"
          >
            {{ ps.Name }} (<b>{{psWos(ps.Name).length}}</b>)
          </v-tab>
        </v-tabs>
      </v-flex>
    </v-layout>
    <v-layout row>
      <v-container fluid  class="pa-0">
        <v-tabs-items v-model="tab">
          <v-tab-item
                  v-for="ps in paysystems"
                  :key="ps.Name"
                  v-if="psWos(ps.Name).length > 0"
          >
            <v-layout row v-for="wo in psWos(ps.Name)" :key="wo.id">
              <Wdrwl :wo="wo"></Wdrwl>
            </v-layout>
          </v-tab-item>
        </v-tabs-items>
      </v-container>
    </v-layout>

  </v-container>
</template>

<script>
  import Wdrwl from '~/components/wdrwl'
  export default {
    name: "withdraw",
    components: {Wdrwl},
    fetch ({ store, app: {$axios}}) {
      return $axios.get(process.env.apiUrl + 'payments/systems')
        .then((res) => {
          store.commit('objects/setPaysystems', res.data)
        })
    },
    data () {
      return {
        tab: null,
      }
    },
    computed: {
      paysystems(){
        return this.$store.state.objects.paysystems
      },
      withdrawals(){
        return this.$store.state.objects.withdrawalOrders
      },
    },
    methods: {
      psWos(psName) {
        return this.withdrawals.filter((wo) => {
          return wo.paymentSystem === psName
        })
      },

    },
  }
</script>

<style scoped>

</style>