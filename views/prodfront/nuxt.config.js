var vm1 = true
var domain = 'xetrix-platform.com'
var wsServer = vm1 ? 'wss://'+domain+'/ws' : 'ws://localhost:8080/ws'
var distpath = vm1 ? 'production' : 'dist'

export default {
  mode: 'spa',
  modulesDir: ['../node_modules'],
  env: {
    wsServer:  process.env.NODE_ENV === 'production' ? wsServer :  'ws://localhost:8080/ws',
    apiUrl:    process.env.NODE_ENV === 'production' ? '/' :  '/api/'
  },
  generate: {
    dir: distpath,
  },
  head: {
    title: 'Xetrix - Exchange service',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: 'Xetrix is a platform for buying, selling and exchanging money' }
    ],
    link: []

  },
  plugins: [
    { src: '~/plugins/ws.js', ssr: false },
    { src: '~/plugins/vuetify.js',  },
    { src: '~/plugins/timer.js',  },
    { src: '~/plugins/moment.js',  },
    { src: '~/plugins/snotify.js',  }
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
    '~/styles/app.styl',
    'material-design-icons-iconfont/dist/material-design-icons.css',
  ],
  build: {
    extend (config, { isDev, isClient }) {
      // ...
    }
  },
  modules: [
    '@nuxtjs/axios',
    '@nuxtjs/auth',
    '@nuxtjs/svg-sprite',
    ['nuxt-rfg-icon', { masterPicture: 'assets/favicon.png' }],
    '@nuxtjs/manifest'
  ],

  svgSprite: {
    input: '~/assets/icons/svg'
  },
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
