<template>
  <v-container container fluid class="pa-0">
    <v-card v-for="(u, i) in users" :key="i" class="mb-3">
      <v-card-title class="grey white--text">
        <v-flex md3>
          <span class="title">Пользователь</span>
        </v-flex>
        <v-spacer></v-spacer>
        <v-flex>
          <p class="text-xs-right ma-0">
            <span class="title"><b>{{u.id}}</b></span>
          </p>
        </v-flex>
      </v-card-title>
      <v-card-text>
        <v-layout row>
          <v-flex md10>
            <ul>
              <li>
                Name: {{u.name}}
              </li>
              <li>
                Email: {{u.email}}
              </li>
            </ul>
          </v-flex>
          <v-flex md2 class="">
            <p class="text-xs-right">
              <v-btn color="primary" @click="loadDialog(u)">Открыть</v-btn>
            </p>
          </v-flex>
        </v-layout>
      </v-card-text>
    </v-card>

    <v-layout row justify-center>
      <v-dialog v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
        <v-card>
          <v-toolbar dark color="primary">
            <v-btn icon dark @click="dialog = false">
              <v-icon>close</v-icon>
            </v-btn>
            <v-toolbar-title>Характеристика</v-toolbar-title>
            <v-spacer></v-spacer>
            <v-toolbar-items>
              <v-btn dark flat @click="dialog = false">Применить</v-btn>
            </v-toolbar-items>
          </v-toolbar>
          <User :user="current"></User>
        </v-card>
      </v-dialog>
    </v-layout>
  </v-container>
</template>

<script>
  import User from '~/components/user'
  export default {
    name: "users",
    components: { User, },
    data () {
      return {
        dialog: false,
        current: null,
      }
    },
    computed: {
      users () {
        return this.$store.state.objects.users
      },
    },
    methods: {
      loadDialog (user) {
        this.current = null
        this.current = user
        this.dialog = true
      }
    },
  }
</script>

<style scoped>

</style>