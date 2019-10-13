<template>
  <v-container fluid>
    <v-dialog
            v-model="dialog"
            width="550"
    >
      <TxPerformProccess :tx="dialogTx"></TxPerformProccess>
    </v-dialog>
    <v-data-table
            :headers="headers"
            :items="notPerformedTxs"
            class="elevation-1"
    >
      <template v-slot:items="props">
        <td class="text-xs-right">{{ props.item.currencySymbol }}</td>
        <td class="text-xs-right">{{ props.item.id }}</td>
        <td class="text-xs-right">{{props.item.toAddress}}</td>
        <td class="text-xs-right">{{ toDate(props.item.createdAt) }}</td>
        <td class="text-xs-right">
          {{ reasonStatus[props.item.reason.status] }}
        </td>
        <td class="text-xs-right">{{ props.item.amount }}</td>
        <td class="text-xs-right">{{ txStatus[props.item.status] }}</td>
        <td class="text-xs-right">
          <v-btn color="primary" @click="openDialog(props.item)">Детали</v-btn>
        </td>
      </template>
    </v-data-table>
  </v-container>
</template>

<script>

  import TxPerformProccess from "~/components/forms/transaction/performProccess"
  export default {
    name: "txasmaster",
    components: {TxPerformProccess},
    computed: {
      notPerformedTxs() {
        return this.$store.state.objects.txs.filter(tx => tx.reason.status < 4)
      }
    },
    data () {
      return {
        dialogTx: null,
        dialog: false,
        headers: [
          { text: 'CURRENCY', align: 'right', value: 'currencySymbol'},
          { text: 'TXID', align: 'right', value: 'id'},
          { text: 'ADDRESS', align: 'right', value: 'toAddress'},
          { text: 'TIME', align: 'right', value: 'createdAt'},
          { text: 'REASON', value: 'reason', align: 'right',},
          { text: 'AMOUNT', value: 'amount', align: 'right',},
          { text: 'STATUS', value: 'status', align: 'right',},
          { text: 'ACTION', value: 'status', align: 'right',},
        ],
        reasonStatus: [
          'UNREASON_TX',
          'FUND_UNPERFORMED_TX',
          'FUND_WAIT_PREPARE_TX',
          'FUND_PREPARED_TX',
          'FUND_PERFORMED_TX',
          'WITHDRAW_TX',
          'SELLER_CONTRACT_TX',
          'BUYER_CONTRACT_TX',
          'CONTRACT_COMMISSION_TX',
        ],
        txStatus: [
          'UNCONFIRMED',
          'COLLATERALIZED',
          'SAVED',
          'CONFIRMED',
          'CANCELED',
        ],
      }
    },
    methods: {
      openDialog(tx){
        this.dialogTx = tx
        this.dialog = true
      },
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