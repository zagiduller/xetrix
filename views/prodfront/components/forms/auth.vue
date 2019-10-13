<template>
  <v-layout justify-center>
    <v-flex xs24 md22 theme--light auth-form>
      <template v-if="getScenario() == 'login'">
        <h1 class="header-title text-xs-center">Вход</h1>
        <h2 class="header-short-text  text-xs-center pb-3">Войдите, используя свой E-mail и пароль</h2>
        <v-layout wrap justify-center>
          <div class="form-wrapper">
            <div class="preloader-background" v-show="preloader">
              <div class="sk-fading-circle">
                <div class="sk-circle1 sk-circle"></div>
                <div class="sk-circle2 sk-circle"></div>
                <div class="sk-circle3 sk-circle"></div>
                <div class="sk-circle4 sk-circle"></div>
                <div class="sk-circle5 sk-circle"></div>
                <div class="sk-circle6 sk-circle"></div>
                <div class="sk-circle7 sk-circle"></div>
                <div class="sk-circle8 sk-circle"></div>
                <div class="sk-circle9 sk-circle"></div>
                <div class="sk-circle10 sk-circle"></div>
                <div class="sk-circle11 sk-circle"></div>
                <div class="sk-circle12 sk-circle"></div>
              </div>
            </div>
            <v-form ref="formauth" v-model="valid" lazy-validation>
              <v-flex>
                <v-text-field
                        v-model="email"
                        label="Email"
                        placeholder="Email"
                        :rules="emailRules"
                        class="input-style"
                        required
                ></v-text-field>
              </v-flex>
              <v-flex >
                <v-text-field
                        v-model="password"
                        :append-icon="e1 ? 'visibility' : 'visibility_off'"
                        @click:append="() => (e1 = !e1)"
                        :type="e1 ? 'password' : 'text'"
                        name="input-10-1"
                        :rules="passwordRules"
                        label="Пароль"
                        placeholder="Пароль"
                        class="input-style"
                        hint="Латинские буквы, не менее 7 символов"
                        min="8"
                        counter
                ></v-text-field>
              </v-flex>
              <v-flex text-xs-center>
                <v-btn block flat class="mb-0 v-btn-style v-btn__type3"  @click="checkForm">Войти</v-btn>
                <hr class="mt-3 mb-3">
                <span class="body-2 gray gray-lighten-1">Нет аккаунта? <a @click="$router.push('/registration')" class="blue-link">Зарегистрироваться</a></span>
              </v-flex>
            </v-form>
          </div>
        </v-layout>
      </template>
      <template v-else>
        <h1 class="header-title text-xs-center">Регистрация</h1>
        <v-layout wrap justify-center>
          <div class="form-wrapper">
            <div class="preloader-background" v-show="preloader">
              <div class="sk-fading-circle">
                <div class="sk-circle1 sk-circle"></div>
                <div class="sk-circle2 sk-circle"></div>
                <div class="sk-circle3 sk-circle"></div>
                <div class="sk-circle4 sk-circle"></div>
                <div class="sk-circle5 sk-circle"></div>
                <div class="sk-circle6 sk-circle"></div>
                <div class="sk-circle7 sk-circle"></div>
                <div class="sk-circle8 sk-circle"></div>
                <div class="sk-circle9 sk-circle"></div>
                <div class="sk-circle10 sk-circle"></div>
                <div class="sk-circle11 sk-circle"></div>
                <div class="sk-circle12 sk-circle"></div>
              </div>
            </div>
            <v-form ref="formauth" v-model="valid" lazy-validation>
              <v-flex>
                <v-text-field
                        v-model="email"
                        label="E-mail"
                        :rules="emailRules"
                        placeholder="E-mail"
                        class="input-style"
                        required
                ></v-text-field>
              </v-flex>
              <v-flex >
                <v-text-field
                        v-model="password"
                        :append-icon="e1 ? 'visibility' : 'visibility_off'"
                        @click:append="() => (e1 = !e1)"
                        :type="e1 ? 'password' : 'text'"
                        name="input-10-1"
                        label="Пароль"
                        placeholder="Пароль"
                        class="input-style"
                        hint="Латинские буквы, не менее 7 символов"
                        min="8"
                        counter
                ></v-text-field>
              </v-flex>
              <v-flex >
                <v-checkbox
                        v-model="agree"
                        :rules="[v => !!v || 'Вы должны согласиться продолжить!']"
                        color="grey"
                        label="">
                  <template slot="label">
                    <span class="body-2">Я принимаю&nbsp;<a href="/rules">условия использования</a></span>
                  </template>
                </v-checkbox>
              </v-flex>
              <v-flex text-xs-center>
                <v-btn block flat class="mb-0 v-btn-style v-btn__type3"  @click="checkForm">Зарегистрироваться</v-btn>
                <hr class="mt-3 mb-3">
                <span class="body-2 gray gray-lighten-1">Есть аккаунт? <a @click="$router.push('/login')" class="blue-link">Вход</a></span>
              </v-flex>
            </v-form>
          </div>
        </v-layout>
      </template>
    </v-flex>
  </v-layout>
</template>

<script>
  export default {
    props: ['scenario'],
    name: "auth",
    data () {
      return {
        valid: true,
        e1: true,
        e2: true,
        scen: '',
        preloader: false,
        password: '',
        passwordRules: [
          v => !!v || 'Необходим пароль',
        ],
        email: '',
        emailRules: [
          v => !!v || 'Требуется E-mail',
          v => /.+@.+/.test(v) || 'E-mail должен быть действительным'
        ],
        agree: false,
        isNew: false,
        errors: []
      }
    },
    mounted() {
      if(this.scenario == 'login') {
        return this.isNew = false
      } else {
        return this.isNew = true
      }
    },
    methods: {
      signup () {},
      login () {
      },
      getScenario () {
        return this.scen === '' ? this.scenario : this.scen
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