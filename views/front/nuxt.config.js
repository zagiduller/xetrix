var vm1 = false
var wsServer = vm1 ? 'wss://tubexhange.com/ws' : 'ws://localhost:8080/ws'

export default {
  mode: 'spa',
  modulesDir: ['../node_modules'],
  env: {
    wsServer:  process.env.NODE_ENV === 'production' ? wsServer :  'ws://localhost:8080/ws',
    apiUrl:    process.env.NODE_ENV === 'production' ? '/' :  '/api/'
  },
  plugins: [
    { src: '~/plugins/ws.js', ssr: false },
    { src: '~/plugins/vuetify.js',  }
  ],
  optimization: {
    splitChunks: {
      chunks: 'all',
      automaticNameDelimiter: '.',
      name: true,
      cacheGroups: {},
      minSize: 500000,
      maxSize: 500000
    }
  },
  maxChunkSize: 500000,
  extractCSS: true,
  filenames: {
    chunk: '[id].[chunkhash].js'
  },
  css: [
    // Загрузить модуль node.js
    'vuetify/dist/vuetify.min.css',
    'material-design-icons-iconfont/dist/material-design-icons.css',
  ],
  modules: [
    '@nuxtjs/axios',
    // '@nuxtjs/proxy',
    '@nuxtjs/auth'
  ],
  axios: {
    proxy: true,
  },
  proxy: {
    '/api/': { target: 'http://localhost:8080', ws: false,  pathRewrite: {'^/api/': ''},  },
  },
  router: {
    middleware: ['auth']
  },
  auth: {
    // Options
    redirect: {
      login: '/login',
      logout: '/',
      callback: '/login',
    },
    strategies: {
      local: {
        endpoints: {
          // Добавить префикс
          login: {
            url: process.env.NODE_ENV === 'production' ? '/v1/start_session' : '/api/v1/start_session',
            method: 'post',
            propertyName: 'Token'
          },
          logout: { url:'/auth/logout', method: 'post' },
          user: {
            url: process.env.NODE_ENV === 'production' ? '/_v1/get_info' : '/api/_v1/get_info',
            method: 'get',
            propertyName: ''
          }
        },
        // tokenRequired: true,
      }
    }
  }
}
