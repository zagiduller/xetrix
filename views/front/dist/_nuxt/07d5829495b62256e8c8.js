(window.webpackJsonp=window.webpackJsonp||[]).push([[6],{322:function(t,e,r){"use strict";r(20);var n={name:"Bitcoin",props:["acc"],data:function(){return{amount:0,sendingAddress:"",fullName:"",attributes:[]}},methods:{send:function(){this.$axios.post("/_v1/create_withdrawal")}}},l=r(15),c=Object(l.a)(n,function(){var t=this.$createElement,e=this._self._c||t;return e("v-flex",{attrs:{md12:""}},[e("v-text-field",{attrs:{label:"Bitcoin address"}})],1)},[],!1,null,"2ba4bd12",null).exports,o={name:"YandexMoney",props:["acc"],data:function(){return{paymentSystem:"YandexMoney",amount:0,fullName:"",address:"",worder:""}},computed:{attributes:function(){return[{key:"fullName",value:this.fullName},{key:"YandexMoneyAddress",value:this.address}]}},methods:{send:function(){var t=this;this.$axios.post("/_v1/create_withdrawal",{paymentSystem:this.paymentSystem,amount:this.amount,sendingAddress:this.acc.account.Address,attributes:this.attributes}).then(function(e){console.log(e.data),t.worder=e.data,t.reset()})},reset:function(){this.amount=0,this.fullName="",this.address=""}}},d={name:"acc",props:["acc","ps"],components:{YandexMoney:Object(l.a)(o,function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("v-container",[r("v-flex",{attrs:{md12:""}},[r("v-text-field",{attrs:{label:"Количество"},model:{value:t.amount,callback:function(e){t.amount=e},expression:"amount"}})],1),t._v(" "),r("v-flex",{attrs:{md12:""}},[r("v-text-field",{attrs:{label:"Адрес счета"},model:{value:t.address,callback:function(e){t.address=e},expression:"address"}})],1),t._v(" "),r("v-flex",{attrs:{md12:""}},[r("v-text-field",{attrs:{label:"ФИО"},model:{value:t.fullName,callback:function(e){t.fullName=e},expression:"fullName"}})],1),t._v(" "),r("v-flex",{staticClass:"text-md-right",attrs:{md12:""}},[r("v-btn",{staticClass:"amber accend-3 ml-0 mt-3 mb-0",on:{click:t.send}},[t._v("Подать заявку")])],1),t._v(" "),t.worder?r("v-flex",[r("p",{staticClass:"title"},[t._v("\n      Создана заявка:\n    ")]),t._v(" "),r("p",[t._v("\n      "+t._s(t.worder)+"\n    ")])]):t._e()],1)},[],!1,null,"198bc34c",null).exports,Bitcoin:c},computed:{currency:function(){return this.acc.account.currency.name},symbol:function(){return this.acc.account.currency.symbol},available:function(){return this.acc.balance.available>0?this.acc.balance.available:0},locked:function(){return this.acc.balance.locked>0?this.acc.balance.locked:0},address:function(){return this.acc.account.Address}},data:function(){return{paymentDialog:!1,withdrawalDialog:!1}},methods:{hashVisual:function(t,e,r){return"last"===e?t.substr(-r):"first"===e?t.substr(0,r):void 0}}},v=Object(l.a)(d,function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("v-layout",{attrs:{wrap:"","card-wallet":""}},[r("v-flex",{attrs:{xs24:"",row:"","h-4":"","px-0":"","py-0":""}},[r("h3",[r("span",{staticClass:"text-uppercase"},[t._v(t._s(t.symbol))]),t._v(" кошелек")])]),t._v(" "),r("v-flex",{staticClass:"card-wallet theme--light block-style-1 mx-0 my-0",attrs:{xs24:""}},[r("v-layout",{staticClass:"header",attrs:{"align-center":"","justify-space-between":"",row:"","fill-height":"",wrap:"","px-3":""}},[r("v-flex",{staticStyle:{flex:"0 0 70px"}},[r("img",{attrs:{src:"/images/left/"+t.symbol.toLowerCase()+".png",width:"48",alt:""}})]),t._v(" "),r("v-flex",{attrs:{"align-center":"","fill-height":"","py-0":""}},[r("v-layout",{attrs:{"align-center":"","fill-height":"","my-0":""}},[r("v-flex",{attrs:{xs24:""}},[r("v-layout",{attrs:{"align-start":"","justify-center":"",column:""}},[r("div",[r("h1",{staticClass:"font-weight-regular text-uppercase",staticStyle:{"line-height":"26px"}},[t._v(t._s(t.available)+" "+t._s(t.symbol))])])]),t._v(" "),r("v-layout",{attrs:{"align-center":"","fill-height":"",wrap:""}},[r("v-flex",{staticStyle:{flex:"0 0 40px"},attrs:{"pl-0":""}},[r("svg-icon",{staticStyle:{width:"22px",height:"18px"},attrs:{name:"wallet-small"}})],1),t._v(" "),r("v-flex",[r("v-layout",{attrs:{"align-center":"",wrap:""}},[r("v-flex",{attrs:{xs24:""}},[r("v-layout",{attrs:{"align-center":""}},[t.$vuetify.breakpoint.smAndDown?[r("v-flex",[r("v-layout",{staticClass:"address_currency",attrs:{"align-space-between":"","justify-space-between":"",row:"","fill-height":"","mx-0":"","my-0":""}},[r("div",[r("span",{staticClass:"body-1 text-color font-weight-bold"},[t._v(t._s(this.hashVisual(t.address,"first",5)))])]),t._v(" "),r("div",[r("span",{staticClass:"body-1 text-color font-weight-bold"},[t._v(t._s(this.hashVisual(t.address,"last",5)))])])])],1)]:[r("v-layout",{staticClass:"address_currency",attrs:{"align-space-between":"","justify-space-between":"",row:"","fill-height":"","mx-0":"","my-0":""}},[r("div",[r("span",{staticClass:"body-1 text-color font-weight-bold"},[t._v(t._s(this.hashVisual(t.address,"first",14)))])]),t._v(" "),r("div",[r("span",{staticClass:"body-1 text-color font-weight-bold"},[t._v(t._s(this.hashVisual(t.address,"last",14)))])])])],t._v(" "),r("v-flex",{staticStyle:{flex:"0 0 40px"},attrs:{"text-xs-center":""}})],2)],1)],1)],1)],1)],1)],1)],1)],1),t._v(" "),r("v-layout",{attrs:{"align-center":"","justify-space-between":"",row:"",footer:"","px-4":"","py-2":""}},[r("v-flex",[r("span",{staticClass:"label body-5 pb-2"},[t._v("Доступный баланс:")]),t._v(" "),r("span",{staticClass:"body-5 data pb-3"},[t._v(t._s(t.available))]),r("br"),t._v(" "),r("span",{staticClass:"label body-5 pb-2"},[t._v("Заблокировано:")]),t._v(" "),r("span",{staticClass:"body-5 data pb-3"},[t._v(t._s(t.locked?t.locked:0))]),r("br"),t._v(" "),r("v-layout",{attrs:{data:"","theme--light":"","table-style":"","align-center":"","justify-center":"",row:""}},[r("v-flex",{attrs:{xs12:""}},[r("v-btn",{staticClass:"mb-0 v-btn-style v-btn__type3",attrs:{block:"",flat:""},on:{click:function(e){t.paymentDialog=!t.paymentDialog}}},[t._v("Пополнить")])],1),t._v(" "),r("v-flex",{attrs:{xs12:""}},[r("v-btn",{staticClass:"mb-0 v-btn-style v-btn__type3",attrs:{block:"",flat:""},on:{click:function(e){t.withdrawalDialog=!t.withdrawalDialog}}},[t._v("Вывести")])],1)],1)],1)],1),t._v(" "),"RUB"===t.symbol?r("v-dialog",{attrs:{width:"500"},model:{value:t.paymentDialog,callback:function(e){t.paymentDialog=e},expression:"paymentDialog"}},[r("v-card",[r("v-card-title",[r("span",{staticClass:"title"},[t._v("Платеж")]),t._v(" "),r("v-spacer"),t._v(" "),r("v-icon",{on:{click:function(e){t.paymentDialog=!t.paymentDialog}}},[t._v("close")])],1),t._v(" "),r("v-card-text",[r("iframe",{attrs:{src:"https://money.yandex.ru/quickpay/shop-widget?writer=seller&targets="+t.address+"&label="+t.address+"&targets-hint=&default-sum=2&button-text=11&payment-type-choice=on&hint=&successURL=&quickpay=shop&account=410018657112160",width:"423",height:"222",frameborder:"0",allowtransparency:"true",scrolling:"no"}})])],1)],1):t._e()],1),t._v(" "),t.ps.length?r("v-dialog",{attrs:{width:"500"},model:{value:t.withdrawalDialog,callback:function(e){t.withdrawalDialog=e},expression:"withdrawalDialog"}},[r("v-card",[r("v-card-title",[r("p",{staticClass:"title mb-0 pb-0"},[t._v("\n          Заявка на вывод "),r("br"),t._v(" "),r("span",{staticClass:"caption"},[t._v(t._s(t.address))])]),t._v(" "),r("v-spacer"),t._v(" "),r("v-icon",{on:{click:function(e){t.withdrawalDialog=!t.withdrawalDialog}}},[t._v("close")])],1),t._v(" "),r("v-card-text",[r("v-tabs",{attrs:{"fixed-tabs":""}},[t._l(t.ps,function(p){return r("v-tab",{key:p.Name},[t._v("\n            "+t._s(p.Name)+"\n          ")])}),t._v(" "),t._l(t.ps,function(p){return r("v-tab-item",{key:p.Name},[r("v-card",{attrs:{flat:""}},[r("v-card-text",[r("v-layout",[r(p.Name,{tag:"component",attrs:{acc:t.acc}})],1)],1)],1)],1)})],2)],1)],1)],1):t._e()],1)},[],!1,null,"9378b654",null);e.a=v.exports},323:function(t,e,r){"use strict";var n={name:"simple",props:["order"],computed:{wsaccs:function(){return this.$store.state.objects.WsAccounts},id:function(){return this.order.id},buyerId:function(){return this.$auth.user.id},send:function(){var t=this,e=this.wsaccs.filter(function(a){return a.account.currency.symbol===t.to})[0];return!!e.account&&e.account.Address},receive:function(){var t=this,e=this.wsaccs.filter(function(a){return a.account.currency.symbol===t.from})[0];return!!e.account&&e.account.Address},from:function(){return this.order.sellCurrencySymbol},to:function(){return this.order.buyCurrencySymbol},available:function(){return this.order.available},price:function(){return this.order.price},query:function(){var q={};return q.orderId=this.id,q.sendingAddress=this.send,q.receiveAddress=this.receive,q.buyerId=this.buyerId,q.amount=this.amount,!!(!!q.amount<=this.available&&q.orderId&&q.sendingAddress&&q.receiveAddress&&q.orderId&&q.buyerId)&&q}},data:function(){return{amount:0}},methods:{buy:function(){if(this.query){this.$axios.post("/_v1/create_contract",this.query).catch(function(t){console.log(t)})}}}},l=r(15),c={name:"order",props:["order"],components:{ContractSimpleForm:Object(l.a)(n,function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("v-container",{staticClass:"pa-0",attrs:{fluid:""}},[r("v-layout",{attrs:{row:"",wrap:""}},[r("v-flex",{attrs:{md6:""}},[r("v-text-field",{attrs:{solo:""},model:{value:t.amount,callback:function(e){t.amount=e},expression:"amount"}})],1),t._v(" "),r("v-flex",{attrs:{md4:""}},[r("v-btn",{staticStyle:{height:"30px",width:"30px"},attrs:{small:"",fab:"",color:"white"},on:{click:t.buy}},[r("v-icon",[t._v("add")])],1)],1)],1)],1)},[],!1,null,"1d1ab146",null).exports},computed:{isUser:function(){return this.order.ownerId===this.$auth.user.id},from:function(){return this.order.sellCurrencySymbol},to:function(){return this.order.buyCurrencySymbol},available:function(){return this.order.available},price:function(){return this.order.price}}},o=Object(l.a)(c,function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("v-card",[r("v-layout",{attrs:{row:"","justify-space-between":""}},[r("v-flex",{staticClass:"grey white--text",attrs:{md12:""}},[r("v-card-title",[r("span",{staticClass:"title"},[t._v("Ордер")])])],1),t._v(" "),r("v-flex",{attrs:{md24:""}},[r("v-card-title",[r("v-layout",{attrs:{row:"","justify-space-between":""}},[r("v-flex",{attrs:{md8:""}},[r("p",{staticClass:"ma-0 text-xs-center"},[r("span",{staticClass:"title"},[t._v(t._s(t.available))]),t._v(" "),r("span",{staticClass:"subheading"},[t._v(t._s(t.from))])])]),t._v(" "),r("v-flex",{attrs:{md4:""}},[r("p",{staticClass:"text-xs-center ma-0"},[r("span",{staticClass:"title"},[t._v("по")])])]),t._v(" "),r("v-flex",{attrs:{md8:""}},[r("p",{staticClass:"ma-0 text-xs-center"},[r("span",{staticClass:"title"},[t._v(t._s(t.price))]),t._v(" "),r("span",{staticClass:"subheading"},[t._v(t._s(t.to))])])])],1)],1)],1),t._v(" "),t.isUser?r("v-flex",{staticClass:"green white--text",attrs:{md8:""}},[r("v-card-title",{staticClass:"text-xs-right"},[r("p",{staticClass:"text-md-right ma-0"},[r("v-icon",[t._v("person")])],1)])],1):r("v-flex",{staticClass:"orange white--text",attrs:{md8:""}},[t.available>0?[r("v-card-title",{staticClass:"text-xs-right"},[r("p",{staticClass:"text-md-right ma-0"},[t._v("\n              Купить "+t._s(t.from)+"\n            ")])]),t._v(" "),r("v-card-title",[r("ContractSimpleForm",{attrs:{order:t.order}})],1)]:t._e()],2)],1)],1)},[],!1,null,"2f0c4b14",null);e.a=o.exports},324:function(t,e,r){"use strict";r(235),r(10);var n={name:"orderForm",props:["accs","pair"],created:function(){console.log("Значение props: "+JSON.stringify(this.pair))},computed:{currencies:function(){var t=[];return this.accs.forEach(function(a){t.findIndex(function(t){return a.account.currency.symbol===t.symbol})<0&&t.push(a.account.currency)}),t},toList:function(){var t=this;return this.from?this.currencies.filter(function(e){return e.symbol!==t.from}):this.currencies},send:function(){var t=this;if(this.from){var e=this.accs.filter(function(a){return a.account.currency.symbol===t.from})[0];if(e.account)return e.account.Address}return!1},receive:function(){var t=this;if(this.to){var e=this.accs.filter(function(a){return a.account.currency.symbol===t.to})[0];if(e.account)return e.account.Address}return!1},query:function(){var q={};return q.sendingAddress=this.send,q.receiveAddress=this.receive,q.amount=this.amount,q.price=this.price,console.log(q),!!(q.sendingAddress&&q.receiveAddress&&q.price&&q.amount)&&q}},data:function(){return{amount:0,price:0,from:"",to:""}},methods:{open:function(){var q=this.query;if(this.query){this.$axios.post("/_v1/create_order",q).then(function(t){console.log(t)})}},hashVisual:function(t,e){return"last"===e?t.substr(-10):"first"===e?t.substr(0,10):void 0}}},l=r(15),component=Object(l.a)(n,function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("v-layout",{staticClass:"data theme--light table-style style-1 buy-table top-panel-blocks block-style-1"},[r("v-flex",[0!==Object.keys(t.pair).length?r("v-layout",{attrs:{wrap:""}},[r("v-flex",{attrs:{xs24:""}},[r("h2",{staticClass:"text-xs-left heading"},[t._v("Кошельки")]),t._v(" "),r("v-layout",{attrs:{"align-center":"","justify-space-between":"",row:"","px-0":"","py-0":""}},[r("v-flex",{staticStyle:{flex:"0 0 40px"}},[r("svg-icon",{staticStyle:{width:"26px",height:"22px"},attrs:{name:"wallet-small-invert"}})],1),t._v(" "),r("v-flex",{attrs:{"fill-height":"","my-0":"","py-0":"","overflow-hidden":""}},[r("v-layout",{attrs:{"fill-height":"",wrap:""}},[r("v-flex",{staticClass:"px-0 pb-0",attrs:{xs24:""}},[r("span",{staticClass:"body-3 text-color"},[t._v("Баланс: ")]),r("span",{staticClass:"font-weight-bold body-3"},[t._v(t._s(t.pair.from.symbol))])]),t._v(" "),r("v-flex",{staticClass:"pt-0",attrs:{xs24:""}},[r("v-layout",[r("v-flex",[r("v-layout",{staticClass:"address_currency",attrs:{"align-space-between":"","justify-space-between":"",row:"","fill-height":"","mx-0":"","my-0":""}},[r("div",[r("span",{staticClass:"body-1 text-color font-weight-bold"},[t._v(t._s(t.hashVisual(t.pair.from.id,"first")))])]),t._v(" "),r("div",[r("span",{staticClass:"body-1 text-color font-weight-bold"},[t._v(t._s(t.hashVisual(t.pair.from.id,"last")))])])])],1)],1)],1)],1)],1)],1),t._v(" "),r("v-layout",{attrs:{"align-center":"","justify-space-between":"",row:"","px-0":"","py-0":""}},[r("v-flex",{staticStyle:{flex:"0 0 40px"}},[r("svg-icon",{staticStyle:{width:"26px",height:"22px"},attrs:{name:"wallet-small-invert"}})],1),t._v(" "),r("v-flex",{attrs:{"fill-height":"","my-0":"","py-0":"","overflow-hidden":""}},[r("v-layout",{attrs:{"fill-height":"",wrap:""}},[r("v-flex",{staticClass:"px-0 pb-0",attrs:{xs24:""}},[r("span",{staticClass:"body-3 text-color"},[t._v("Баланс: ")]),r("span",{staticClass:"font-weight-bold body-3"},[t._v(t._s(t.pair.to.symbol))])]),t._v(" "),r("v-flex",{staticClass:"pt-0",attrs:{xs24:""}},[r("v-layout",[r("v-flex",[r("v-layout",{staticClass:"address_currency",attrs:{"align-space-between":"","justify-space-between":"",row:"","fill-height":"","mx-0":"","my-0":""}},[r("div",[r("span",{staticClass:"body-1 text-color font-weight-bold"})]),t._v(" "),r("div",[r("span",{staticClass:"body-1 text-color font-weight-bold"})])])],1)],1)],1)],1)],1)],1)],1),t._v(" "),t.pair?r("v-flex",{attrs:{xs12:""}},[r("h2",{staticClass:"text-xs-left heading"},[t._v("Отдаете")]),t._v(" "),r("img",{attrs:{src:"/images/left/"+t.pair.from.symbol+".png",width:"18",alt:""}}),t._v(t._s(t.pair.from.name)+" ("+t._s(t.pair.from.symbol)+")\n        "),r("v-text-field",{attrs:{label:"Кол-во",solo:""},model:{value:t.amount,callback:function(e){t.amount=e},expression:"amount"}})],1):t._e(),t._v(" "),r("v-flex",{attrs:{xs12:""}},[r("h2",{staticClass:"text-xs-left heading"},[t._v("Получаете")]),t._v(" "),r("img",{attrs:{src:"/images/left/"+t.pair.to.symbol+".png",width:"18",alt:""}}),t._v(t._s(t.pair.to.name)+" ("+t._s(t.pair.to.symbol)+")\n        "),r("v-text-field",{staticClass:"elevation-0",attrs:{label:"Цена за ед.",solo:""},model:{value:t.price,callback:function(e){t.price=e},expression:"price"}})],1),t._v(" "),r("v-flex",{attrs:{xs12:""}},[r("span",{staticClass:"label body-1 pb-2 d-block"},[t._v("Комиссия")]),t._v(" "),r("span",{staticClass:"body-5 d-block data pb-3"}),t._v(" "),r("span",{staticClass:"label body-1 pb-2 d-block"},[t._v("Цена с комиссией")]),t._v(" "),r("span",{staticClass:"body-5 d-block data pb-3"})]),t._v(" "),r("v-flex",{attrs:{xs12:""}},[r("span",{staticClass:"label body-1 pb-2 d-block"},[t._v("Итоговая цена")]),t._v(" "),r("span",{staticClass:"body-5 d-block data pb-3  d-block"}),t._v(" "),r("v-btn",{staticClass:"mb-0 v-btn-style v-btn__type1",attrs:{color:"",flat:"",disabled:!t.query},on:{click:t.open}},[t._v("Создать ордер")])],1)],1):t._e()],1)],1)},[],!1,null,"0227e2f5",null);e.a=component.exports},340:function(t,e,r){"use strict";r.r(e);r(57);var n=r(322),l=r(324),c=r(323),o={name:"index",components:{Acc:n.a,OrderForm:l.a,Order:c.a},layout:"base",data:function(){return{isBordered:!0}},computed:{WsAccounts:function(){return this.$store.state.objects.WsAccounts},orders:function(){return this.$store.state.objects.orders},userOrders:function(){var t=this;return this.orders.filter(function(e){return e.ownerId===t.$auth.user.id})}},methods:{accPS:function(t){return this.$store.state.objects.paysystems.filter(function(e){return t===e.Symbol})}}},d=r(15),component=Object(d.a)(o,function(){var t=this.$createElement;return(this._self._c||t)("v-layout",{attrs:{wrap:"","justify-center":""}})},[],!1,null,null,null);e.default=component.exports}}]);