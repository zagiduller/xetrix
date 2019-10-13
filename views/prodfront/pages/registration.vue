<template>
    <v-layout mx-0 my-0 wrap column>
        <v-flex x24 d-flex>
            <v-layout align-center justify-center wrap  >
                <v-flex md6 mx-0 my-0>
                    <AuthForm scenario='registration'></AuthForm>
                </v-flex>
            </v-layout>
        </v-flex>
        <v-flex x24 text-xs-center style="height: 52px; flex: 0 1;">
            <span class="font-weight-bold">© 2019 Xetrix. Все права защищены.</span>
        </v-flex>
    </v-layout>
</template>
<script>
  import AuthForm from '~/components/forms/auth'
  export default {
    name: "login",
    layout: 'login',
    auth: false,
    components: {AuthForm},
    data() {
      return {}
    }
  }
</script>

<style scoped>
    .login-vh {
        margin-top: 15%;
    }
</style>