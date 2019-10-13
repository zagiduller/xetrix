<template>
    <div class="pag">
        <v-layout align-center justify-center>
            <v-flex no-gutters px-0>
                <span class="text-uppercase">{{pagination.page}} Of {{Math.ceil(pagination.totalItems / pagination.rowsPerPage)}}</span>
            </v-flex>
            <v-flex no-gutters px-0>
                <a @click="prev"><v-icon :disabled="!prev_page" medium>mdi-chevron-left</v-icon></a>
                <a @click="next"><v-icon :disabled="!next_page" medium>mdi-chevron-right</v-icon></a>
            </v-flex>
        </v-layout>
    </div>
</template>

<script>
  export default {
    props: ['pagination'],
    methods: {
      next (e) {
        if (this.next_page) {
          this.pagination.page = this.pagination.page + 1
        }
      },
      prev (e) {
        if (this.prev_page) {
          this.pagination.page = this.pagination.page - 1
        }
      }
    },
    computed: {
      prev_page () {
        return this.pagination.page > 1
      },
      next_page () {
        return this.pagination.page < Math.ceil(this.pagination.totalItems / this.pagination.rowsPerPage)
      }
    },
    watch: {
      pagination (value) {
        if (value.page > Math.ceil(value.totalItems / value.rowsPerPage) && (value.rowsPerPage > 0)) {
          value.page = Math.ceil(value.totalItems / value.rowsPerPage)
        }
      }
    },
    data () {
      return {
        amount: ''
      }
    }
  }
</script>