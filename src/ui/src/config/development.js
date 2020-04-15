// you can set environment variable ( name: HOST_ENV ) to set your host
module.exports = {
  MOCK_HOST: 'http://localhost:5200/',
  HOST: 'http://localhost:5200/',
  ISTIO_LANG_KEY: 'en-US',
  WEBPACK_PROXY: {
    '/api': {
      target: 'http://www.naftis.com'
    }, // if your api server has been proxied by nginx or other web server, replace this host with your proxy configuration host.
    '/ws': {
      target: 'http://www.naftis.com',
      ws: true
    }, // if your api server has been proxied by nginx or other web server, replace this host with your proxy configuration host.
    '/prometheus': {
      target: 'http://www.naftis.com' // port forward your prometheus, and then replace this host with your exported prometheus's host.
    }
  }
}
