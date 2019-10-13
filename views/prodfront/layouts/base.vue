<template>
    <v-app light>
        <vue-snotify></vue-snotify>
        <v-dialog v-model="logoutModal" max-width="290" content-class="theme--light dialog-variant-2" transition="dialog-opacity-transition">
            <v-card>
                <v-card-title body-4>Вы действительно хотите выйти?</v-card-title>
                <v-card-actions>
                    <v-btn @click="logout" :ripple="false" color="" flat class="mb-0 v-btn-style v-btn__type3">Да</v-btn>
                    <v-spacer></v-spacer>
                    <v-btn @click="logoutModal = false" :ripple="false" color="" flat class="mb-0 v-btn-style v-btn__type3">Нет</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
        <v-toolbar class="toolbar-main toolbar-gradient" flat height="70">
            <img src="~assets/images/logo-xetrix.png" class="img-adaptive" alt="Xetrix" />
            <v-toolbar-items class="hidden-sm-and-down ml-md-4"></v-toolbar-items>
            <v-spacer></v-spacer>
            <v-toolbar-items class="mr-0" v-if="$vuetify.breakpoint.mdAndUp">
                <div v-show="$auth.$state.loggedIn" class="right-nav">
                    <v-btn flat @click="logoutModal = true">
                        <v-icon>person</v-icon> {{$auth.$state.loggedIn ? $auth.user.name : ''}}
                    </v-btn>
                </div>
                <div v-show="!$auth.$state.loggedIn" class="right-nav">
                    <v-btn :ripple="false" :class="['v-btn-style v-btn-toolbar-right', {'active' : $nuxt.$route.path === '/login'}]" @click="$router.push('/login')" flat >
                        Авторизация
                    </v-btn>
                    <v-btn :ripple="false" :class="['v-btn-style v-btn-toolbar-right', {'active' : $nuxt.$route.path === '/registration'}]" @click="$router.push('/registration')" flat>
                        Регистрация
                    </v-btn>
                </div>
            </v-toolbar-items>
            <v-menu offset-y content-class="main-navigation" full-width>
                <template #activator="data" v-slot:activator>
                    <v-toolbar-side-icon class="hidden-md-and-up burger" v-on="data.on"></v-toolbar-side-icon>
                </template>
                <v-layout px-3 py-3 class="mobile-navigation">
                    <v-flex xs24>
                        <button :class="['button-panel mobile', {'active' : menu.name === $route.name}]" @click="$router.push(menu.link)" v-for="menu in leftmenu" >
                            <v-layout align-center justify-space-between row px-0 py-0 >
                                <v-flex class="icon-nav">
                                    <svg-icon :name="menu.svg" class="main-menu"/>
                                    <div class="hover">
                                        <svg-icon :name="menu.svg+'_hover'" class="main-menu-hover"/>
                                    </div>
                                </v-flex>
                                <v-flex fill-height my-0 py-0 overflow-hidden text-xs-left pl-3 blue--text text--darken-1 body-4-5>
                                    <strong>{{menu.text}}</strong>
                                </v-flex>
                            </v-layout>
                        </button>
                        <v-btn v-show="$auth.$state.loggedIn" flat @click="logoutModal = true">
                            <v-icon>person</v-icon> {{$auth.$state.loggedIn ? $auth.user.name : ''}}
                        </v-btn>
                    </v-flex>
                </v-layout>
            </v-menu>
        </v-toolbar>
        <v-container grid-list-lg fluid content-exchange>
            <div class="left-panel" v-if="$vuetify.breakpoint.mdAndUp">
                <nav>
                    <button :class="['button-panel', {'active' : menu.name === $route.name}]" @click="$router.push(menu.link)" v-for="menu in leftmenu" >
                        <svg-icon :name="menu.svg" class="main-menu"/>
                        <div class="hover">
                            <svg-icon :name="menu.svg+'_hover'" class="main-menu-hover"/>
                        </div>
                        <span>{{menu.text}}</span>
                    </button>
                </nav>
            </div>
            <v-layout style="flex: 1;">
                <nuxt/>
            </v-layout>
        </v-container>
    </v-app>
</template>

<script>
  export default {
    name: "default",
    data () {
      return {
        drawer: false,
        logoutModal: false,
        leftmenu: [
          {text: 'Биржа', link: '/', svg: 'exchange', name: 'index'},
          {text: 'Кошелек', link: 'wallets', svg: 'wallets', name: 'wallets'},
          {text: 'Мои ордера', link: 'myorders', svg: 'myorders', name: 'myorders'},
          {text: 'История', link: 'history', svg: 'history', name: 'history'},
          {text: 'FAQ', link: 'faq', svg: 'faq', name: 'faq'},
          {text: 'Поддержка', link: 'support', svg: 'support', name: 'support'}
        ]
      }
    },
    methods: {
      logout () {
        this.logoutModal = false
        this.$auth.logout().then((e) => {
          console.log("logout yes");
          this.$router.push('/login')
        })
      },
      action (menuItem) {
        this.active = menuItem.route
        this.$router.push(this.active)
      },
      validate () {
        return this.$auth.$state.loggedIn
      }
    },
  }
</script>

<style scoped>

</style>