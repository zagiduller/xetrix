<template>
  <div>
    <v-dialog>

    </v-dialog>
    <v-card >
      <v-card-title
              class="headline grey lighten-2"
              primary-title
      >
        Управление выводом транзакций
      </v-card-title>
      <v-card-text>
        <p>Памятка:</p>
        <p>Пользователь пополнил свой счет. Нужно произвести вывод с него на мастер-счет на котором аккумулируются средства.</p>
        <p>Все транзакции требуют оплату в ETH. Но на токен счетах пользователей ETH нет, соответственно произвести вывод нельзя.</p>
        <p>Это окошко предлагает следующее решение:</p>
        <p>Проверяет баланс счета с которого нужно произвести вывод, если баланс меньше цены за транзакцию,
          появится кнопка создания пополняющей счет транзакции. При её нажатии со счета администратора будет создана
          транзакция на сумму которой хватит для оплаты комиссии, а транзакция в системе будет переведена в статус
          ожидания подтверждения.
        </p>
        <p>
          Если баланса для оплаты комиссии на счете достаточно, появится кнопка создания выходной транзакции на мастер-счет.
        </p>

      </v-card-text>
      <v-card-text v-if="tx">
        <p>
          Инфо:
        </p>
        <p  class="caption">
          Посмотреть баланс: <a target="_blank" v-bind:href="'https://api.etherscan.io/api?module=account&action=balance&tag=latest&address='+ address">{{address}}</a>
        </p>
        <p  class="caption">
          Входящая транзакция: <a target="_blank" v-bind:href="'https://etherscan.io/tx/' + psInTx">{{psInTx}}</a>
        </p>
        <p  class="caption">
          Транзакция подготовки: <a target="_blank" v-bind:href="'https://etherscan.io/tx/' + psPrepareTx">{{psPrepareTx}}</a>
        </p>
        <p  class="caption" >
          Транзакция на вывод: <a target="_blank" v-bind:href="'https://etherscan.io/tx/' + psOutTx">{{psOutTx}}</a>
        </p>
        <p class="caption">
          Текущий статус: <code>{{reasonStatus[tx.reason.status]}}</code>
        </p>
        <p class="caption">
          {{actions.current}}
        </p>
        <v-divider></v-divider>
        <v-layout wrap class="mt-5 mb-3" v-if="!psOutTx">
          <v-flex md12 class="pr-1" v-if="actions.needPrepare && !psPrepareTx">
            <v-btn block class="success" @click="needPrepare">Подготовить</v-btn></v-flex>
          <v-flex md12 class="pl-1" v-if="actions.prepared && !psOutTx" >
            <v-btn block class="warning" @click="prepared"> Вывести</v-btn></v-flex>
        </v-layout>
      </v-card-text>
      <v-card-text v-else>
        <v-progress-linear
                indeterminate
                class="mb-0"
        ></v-progress-linear>
      </v-card-text>

    </v-card>
  </div>
</template>

<script>
  export default {
    name: "transactions",
    props: ['tx'],
    data () {
      return {
        dialogTx: null,
        dialog: false,
        actionsMap: {
          'FUND_UNPERFORMED_TX': {current: 'Нужно подготовить счет к выводу', needPrepare: true, prepared: false },
          'FUND_WAIT_PREPARE_TX': {current: 'Транзакция отправлена, проверяйте счет. Если баланс есть, то жмите "Вывести"', needPrepare: true, prepared: true },
          'FUND_PREPARED_TX': {current: 'Вывод средств осуществлен системой. Ожидается перевод транзакции в статус "Исполнено"', needPrepare: false, prepared: true},
          'FUND_PERFORMED_TX': {current: 'Исполнено', needPrepare: false, prepared: false},
        },
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
        accBalance: false,
      }
    },
    computed: {
      psInTx () {
            return this.tx.reason.inPStxId ? this.tx.reason.inPStxId : false

      },
      psPrepareTx() {
        return this.tx.reason.preparePStxId ? this.tx.reason.preparePStxId : false
      },
      psOutTx() {
        return this.tx.reason.outPStxId ? this.tx.reason.outPStxId : false
      },
      address () {
        return this.tx.toAddress
      },
      actions () {
        return this.actionsMap[this.status]
      },
      status () {
        return this.reasonStatus[this.tx.reason.status]
      }
    },
    methods: {
      openDialog(){
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
      },
      balance() {
        if (!this.accountBalance) {
          this.$axios.get("https://api.etherscan.io/api?module=account&action=balanc&tag=latest&apikey=YourApiKeyTokene&address="+this.address)  .then(function (response) {
            // handle success
            console.log(response);
          })
        }
        return this.accountBalance
      },
      needPrepare() {
        this.$axios.post(process.env.apiUrl +"_v1-a/need-prepare-txs",{
          txsId: [
            this.tx.id,
          ]
        }).then((resp) => {
          console.log(resp)
        })
      },
      prepared () {
        this.$axios.post(process.env.apiUrl +"_v1-a/txs-prepared",{
          txsId: [
            this.tx.id,
          ]
        }).then((resp) => {
          console.log(resp)
        })
      }
    }
  }
</script>

<style scoped>

</style>