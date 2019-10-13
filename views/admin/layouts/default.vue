<template>
    <v-app light>
        <v-content class="blue-grey lighten-5">
            <v-layout>
                <v-flex v-if="$auth.loggedIn">
                    <v-navigation-drawer
                            fixed
                            app
                    >
                        <v-list :expand="true">
                            <v-list-tile to="/">
                                <v-list-tile-action>
                                    <v-icon>home</v-icon>
                                </v-list-tile-action>
                                <v-list-tile-title><h3>Администраторская</h3></v-list-tile-title>
                            </v-list-tile>

                            <v-list-group
                                    v-for="(menu, k) in menus"
                                    :key="k"
                                    value="true"
                            >
                                <v-list-tile slot="activator">
                                    <v-list-tile-title><h4>{{menu.category}}</h4></v-list-tile-title>
                                </v-list-tile>


                                <v-list-tile
                                        v-for="(list, i) in menu.nested"
                                        :key="i"
                                        @click=""
                                        :to="list.route"
                                >
                                  <v-list-tile-action v-if="!!list.icon">
                                    <v-icon v-text="list.icon"></v-icon>
                                  </v-list-tile-action>
                                    <v-list-tile-title v-text="list.name"></v-list-tile-title>

                                </v-list-tile>

                            </v-list-group>
                        </v-list>
                    </v-navigation-drawer>
                </v-flex>
                <v-flex md12>
                    <nuxt/>
                </v-flex>
            </v-layout>
        </v-content>
    </v-app>
</template>

<script>
  export default {
    name: "default",
    data () {
        return {
          menus: [
            {
              category: 'Общее',
              nested: [
                {
                  name: 'Валюты',
                  icon: 'list',
                  route: '/currencies',
                },
                {
                  name: 'Участники',
                  icon: 'list',
                  route: '/users',
                },
                {
                  name: 'Комиссия',
                  icon: 'money',
                  route: '/commission',
                },
                {
                  name: 'Вывод со счетов',
                  icon: 'call_received',
                  route: '/txasmaster',
                },
                {
                  name: 'Заявки на вывод',
                  icon: 'call_made',
                  route: '/withdraw',
                },
              ]
            },
            {
              category: 'Рынок',
              nested: [
                {
                  name: 'Ордера',
                  icon: 'work',
                  route: '/orders',
                },

              ]
            },

          ],

      }
    }
  }
</script>

<style scoped>

</style>