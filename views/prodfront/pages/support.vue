<template>
    <v-layout justify-center mx-0 my-0 px-3 py-3 wrap>
        <v-flex xs24 md15>
            <v-layout wrap justify-xs-start  mt-md-5>
                <v-flex xs24 text-xs-center  mt-md-5>
                    <img src="/images/xetrix_support_logo.png" alt="">
                    <h1 class="body-9 font-weight-bold pb-5 pt-4">Свяжитесь с нами</h1>
                </v-flex>
                <v-flex xs24 md12>
                    <p class="body-5-md">Мы всегда готовы ответить на интересующие Вас <br>
                    вопросы, а также выслушать Ваши предложения<br>
                    по улучшению нашего сервиса.</p>
                    <span class="body-6 d-inline-block pb-3"><strong>E-mail</strong></span><br>
                    <svg-icon name="email-letter" style="width: 35px; height: 35px; vertical-align: middle" class="mr-4"/><a href="mailto:support@xet-rix.com" class="body-5-5"><strong>support@xet-rix.com</strong></a><br>
                    <span class="body-6 d-inline-block pb-3 pt-5"><strong>Мы в соц. сетях:</strong></span><br>
                    <div class="d-block text-xs-left">
                        <a class="body-4 links mr-3" target="_blank" href="https://www.facebook.com/xetrix"><v-icon class="body-7">mdi-facebook</v-icon></a>
                        <a class="body-4 links" target="_blank" href="https://twitter.com/xetrix"><v-icon class="body-7">mdi-twitter</v-icon></a>
                    </div>
                </v-flex>

                <v-flex xs24 md12>
                    <form>
                        <v-layout wrap justify-xs-start>
                            <v-flex xs24 md12>
                                <v-text-field
                                        v-model="name"
                                        class="table-input"
                                        single-line
                                        placeholder="Ваше имя"
                                        outline
                                        color=""
                                ></v-text-field>
                            </v-flex>
                            <v-flex xs24 md12>
                                <v-text-field
                                        v-model="email"
                                        placeholder="E-mail"
                                        class="table-input"
                                        single-line
                                        outline
                                        color=""
                                ></v-text-field>
                            </v-flex>
                            <v-flex xs24>
                                <v-textarea
                                        class="table-input"
                                        placeholder="Введите сообщение"
                                        single-line
                                        outline
                                        color=""
                                ></v-textarea>
                            </v-flex>
                            <v-flex xs24 offset-md12 md12>
                                <v-btn @click="submit" color="" block flat  class="mb-0 v-btn-style v-btn__type3">Отправить сообщение</v-btn>
                            </v-flex>
                        </v-layout>
                    </form>
                </v-flex>
            </v-layout>
        </v-flex>
    </v-layout>
</template>


<script>
  import { mapState } from 'vuex'

  export default {
    name: "support",
    layout: 'base',

  }
</script>