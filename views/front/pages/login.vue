<template>
    <v-container fluid fill-height>
        <v-layout align-center justify-center>
            <v-flex md3>
                <AuthForm></AuthForm>
            </v-flex>
        </v-layout>
    </v-container>
</template>

<script>
  import AuthForm from '~/components/forms/auth'
  export default {
    name: "login",
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