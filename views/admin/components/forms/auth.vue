<template>
  <div>
    <v-layout column>
      <v-flex xs6>
        <h3 class="mb-3">Вход в админ панель</h3>
        <v-text-field
                label="Email"
                append-icon="person"
                v-model="email"
                outline
        ></v-text-field>
      </v-flex>
      <v-flex xs6>
        <v-text-field
                label="Password"
                append-icon="vpn_key"
                v-model="password"
                type="password"
                outline
        ></v-text-field>
      </v-flex>
    </v-layout>
    <v-layout>
      <v-spacer></v-spacer>
        <v-btn color="primary"  @click="checkForm">Отправить</v-btn>
    </v-layout>
  </div>
</template>

<script>
  export default {
    name: "auth",
    data () {
      return {
        email: '',
        password: '',
        isNew: false,
        errors: []
      }
    },
    methods: {
      signup () {

      },
      login () {

      },
      checkForm(e) {
        this.errors = []
        if (this.email && this.password) {
          var url = process.env.apiUrl
          url += this.isNew ? 'v1/signup' : 'v1/start_session'

          var that = this;

          var authenticateData = {name: that.email, email: that.email, password: that.password}

          var login = () => {

            that.$auth.loginWith('local', {data: authenticateData }).then(() => {
              that.$connect();
              that.$router.push('/');
            }).catch((err)=>{
              console.log(err)
            })
          }

          if (this.isNew) {
            that.$axios.post(url, authenticateData).finally(() => {
              login()
            })
          } else {
            login()
          }



          return true;
        }

        if (!this.email) {
          this.errors.push('Требуется указать email.');
        }
        if (!this.password) {
          this.errors.push('Требуется указать пароль.');
        }
        e.preventDefault()
      }
    }
  }
</script>


<style>
    .m1 {
        margin-top: 10px;
    }
</style>