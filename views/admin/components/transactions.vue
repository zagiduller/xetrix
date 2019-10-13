<template>
  <div>
    <v-data-table
            :dark="!light"
            :headers="headers"
            :items="txs"
            class="elevation-1"
    >
      <template v-slot:items="props">
        <td class="text-xs-right">{{ props.item.id }}</td>
        <td class="text-xs-right" v-if="isBase( props.item.fromAddress )" >
          <span class="orange">( OUT )</span> <b>{{props.item.fromAddress}}</b>
        </td>
        <td class="text-xs-right" v-else>{{props.item.fromAddress}}</td>
        <td class="text-xs-right" v-if="isBase( props.item.toAddress )">
          <span class="green">( IN )</span> <b>{{props.item.toAddress}}</b>
        </td>
        <td class="text-xs-right" v-else>{{props.item.toAddress}}</td>
        <td class="text-xs-right">{{ toDate(props.item.createdAt) }}</td>
        <td class="text-xs-right">{{ props.item.amount }}</td>
        <td class="text-xs-right">{{ props.item.reason }}</td>
        <td class="text-xs-right">{{ props.item.status }}</td>
      </template>
    </v-data-table>

  </div>
</template>

<script>
  export default {
    name: "transactions",
    props: ['txs', 'base', 'perform', 'light'],
    computed: {
    },
    data () {
      return {
        headers: [
          { text: 'TXID', align: 'right', value: 'id'},
          { text: 'FROM', align: 'right', value: 'fromAddress'},
          { text: 'TO', align: 'right', value: 'toAddress'},
          { text: 'TIME', align: 'right', value: 'createdAt'},
          { text: 'AMOUNT', value: 'amount', align: 'right',},
          { text: 'REASON', value: 'reason', align: 'right',},
          { text: 'STATUS', value: 'status', align: 'right',},
        ],
      }
    },
    computed: {

    },
    methods: {
      isBase(addr) {
        return addr === this.base
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