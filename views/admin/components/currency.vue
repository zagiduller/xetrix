<template>
  <v-card>
    <v-layout>
    <v-flex md12>
      <v-card-title>
        <span>
          {{c.name}}
        </span>
        <v-spacer></v-spacer>
        <span v-if="c.active" class="light-green--text darken-3">
        активно
      </span>
        <span v-else class="deep-orange--text darken-1">деактивировано</span>
      </v-card-title>
    </v-flex>

    </v-layout>
    <v-card-title>
      <v-spacer></v-spacer>
      <span>
        <v-chip small>{{c.type == 1 ? 'Фиат' : 'Криптовалюта'}}</v-chip>
        <v-chip small v-if="c.type == 2">Смартконтракт</v-chip>
      </span>
    </v-card-title>
    <v-card-text>
      <v-list dense>
        <v-list-tile>
          <v-list-tile-content>Активных счетов</v-list-tile-content>
          <v-list-tile-content class="align-end">{{accounts.length}}</v-list-tile-content>
        </v-list-tile>
      </v-list>
    </v-card-text>
  </v-card>
</template>

<script>
  export default {
    name: "currency",
    props: ['currency'],
    computed: {
      c() { return this.currency },
      accounts() {
        return this.$store.state.objects.WsAccounts.filter((wsa)=> wsa.account.currency.id === this.c.id)
      }
    }
  }
</script>

<style scoped>

</style>