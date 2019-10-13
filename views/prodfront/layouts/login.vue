<template>
    <v-app light>
        <v-toolbar class="toolbar-main toolbar-gradient" flat height="70">
            <img src="~assets/images/logo-xetrix.png" class="img-adaptive" alt="Xetrix" />
            <v-toolbar-items class="hidden-sm-and-down ml-md-4"></v-toolbar-items>
            <v-spacer></v-spacer>
            <v-toolbar-items class="mr-0" v-if="$vuetify.breakpoint.mdAndUp">
                <div v-show="$auth.$state.loggedIn" class="right-nav">
                    <v-btn flat @click="logoutModal = true">
                        {{$auth.$state.loggedIn ? $auth.user.name : ''}} <v-icon>person</v-icon>
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
        </v-toolbar>
        <v-container grid-list-lg fluid content-exchange>
            <v-layout style="flex: 1;" mx-0 my-0>
                <nuxt/>
            </v-layout>
        </v-container>
    </v-app>
</template>

<script>
  export default {
    name: "default",
    data () {
      return {}
    },
  }
</script>

<style scoped>

</style>