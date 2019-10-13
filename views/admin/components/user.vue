<template>
  <div v-if="!!user" class="pa-2">
    <h3>Имя: {{user.name}}</h3>
    <h3>ИД: {{user.id}}</h3>
    <v-list>
      <v-subheader>Счета</v-subheader>
      <v-list-group
              v-for="acc in accounts" :key="acc.account.id"
              prepend-icon="settings_ethernet"
              no-action
      >
        <template v-slot:activator>
          <v-list-tile>
            <v-list-tile-content>
              <v-list-tile-title>
                <span>{{ acc.account.currency.name }}</span>
                <span class="pl-5" v-if="acc.balance.available  > 0 ">баланс: {{ acc.balance.available }}</span>
              </v-list-tile-title>
            </v-list-tile-content>
            <v-list-tile-avatar>
              <span v-if="acc.account.status == 1" class="green--text">Активен</span>
              <span v-else class="orange--text">Не активен</span>
            </v-list-tile-avatar>
          </v-list-tile>
        </template>
        <v-list-tile class="pa-3">
          <v-list-tile-content class="text-xs-left">
            <v-list-tile-title>Валюта </v-list-tile-title>
            <v-list-tile-sub-title>{{acc.account.currency.name}} ({{acc.account.currency.symbol}})</v-list-tile-sub-title>
          </v-list-tile-content>
          <v-list-tile-content >
            <v-list-tile-title>Зарегистрирован</v-list-tile-title>
            <v-list-tile-sub-title>{{toDate(acc.account.createdAt)}}</v-list-tile-sub-title>
          </v-list-tile-content>
          <v-list-tile-content >
            <v-list-tile-title>Адрес</v-list-tile-title>
            <v-list-tile-sub-title>{{acc.account.Address}}</v-list-tile-sub-title>
          </v-list-tile-content>
        </v-list-tile>

        <v-list-tile  class="pa-3">
          <v-list-tile-content >
            <v-list-tile-title>Баланс</v-list-tile-title>
            <v-list-tile-sub-title>По счету. Ед:</v-list-tile-sub-title>
          </v-list-tile-content>
          <v-list-tile-content >
            <v-list-tile-title >Доступно</v-list-tile-title>
            <v-list-tile-sub-title>{{acc.balance.available}}</v-list-tile-sub-title>
          </v-list-tile-content>
          <v-list-tile-content>
            <v-list-tile-title>Заблокировано</v-list-tile-title>
            <v-list-tile-sub-title>{{acc.balance.locked}}</v-list-tile-sub-title>
          </v-list-tile-content>
        </v-list-tile>
        <v-list-tile  class="pa-3">
          <v-layout row>
            <v-flex md3>
              <SimpleFund :account="acc.account"></SimpleFund>
            </v-flex>
          </v-layout>
        </v-list-tile>
        <Txs :txs="addrTxs(acc.account.Address)" :base="acc.account.Address"></Txs>
      </v-list-group>
    </v-list>

    <v-list three-line subheader>
      <v-subheader>Не используется</v-subheader>
      <v-list-tile avatar>
        <v-list-tile-action>
          <v-checkbox v-model="notifications"></v-checkbox>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Notifications</v-list-tile-title>
          <v-list-tile-sub-title>Notify me about updates to apps or games that I downloaded</v-list-tile-sub-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile avatar>
        <v-list-tile-action>
          <v-checkbox v-model="sound"></v-checkbox>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Sound</v-list-tile-title>
          <v-list-tile-sub-title>Auto-update apps at any time. Data charges may apply</v-list-tile-sub-title>
        </v-list-tile-content>
      </v-list-tile>
      <v-list-tile avatar>
        <v-list-tile-action>
          <v-checkbox v-model="widgets"></v-checkbox>
        </v-list-tile-action>
        <v-list-tile-content>
          <v-list-tile-title>Auto-add widgets</v-list-tile-title>
          <v-list-tile-sub-title>Automatically add home screen widgets</v-list-tile-sub-title>
        </v-list-tile-content>
      </v-list-tile>
    </v-list>
  </div>
</template>

<script>
  import Txs from '~/components/transactions'
  import SimpleFund from '~/components/forms/account/simpleFund'
  export default {
    name: "user",
    props: ['user'],
    components: {Txs,SimpleFund},
    computed: {
      accounts () {
        return this.$store.state.objects.WsAccounts.filter((ws) => ws.account.ownerId === this.user.id)
      },
    },
    data () {
      return {
        notifications: false,
        sound: true,
        widgets: false,
      }
    },
    methods: {
      addrTxs (addr) {
        return this.$store.state.objects.txs.filter((tx) => tx.fromAddress === addr || tx.toAddress === addr )
      },
      toDate(UNIX_timestamp) {
        var a = new Date(UNIX_timestamp * 1000);
        var months = ['Jan','Feb','Mar','Apr','May','Jun','Jul','Aug','Sep','Oct','Nov','Dec'];
        var year = a.getFullYear();
        var month = months[a.getMonth()];
        var date = a.getDate();
        var hour = a.getHours();
        var min = a.getMinutes();
        var sec = a.getSeconds();
        var time = date + ' ' + month + ' ' + year + ' ' + hour + ':' + min + ':' + sec ;
        return time;
      }
    }
  }
</script>

<style scoped>

</style>