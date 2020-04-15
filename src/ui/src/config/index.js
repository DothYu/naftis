/**
 * frontend config entry file
 */

const DEVELOPMET = require('./development')
const PRODUCTION = require('./production')

let CONFIG = {}
let nodeEnv = ''
let hostEnv = ''
try {
  nodeEnv = NODE_ENV
  hostEnv = HOST_ENV
} catch (error) {
  // console.log('err', error)
}

if (!nodeEnv || nodeEnv === 'development') {
  if (hostEnv && hostEnv !== 'undefined') {
    for (let key in DEVELOPMET.WEBPACK_PROXY) {
      DEVELOPMET.WEBPACK_PROXY[key] = hostEnv
    }
  }
  CONFIG = DEVELOPMET
} else if (nodeEnv === 'production') {
  if (hostEnv && hostEnv !== 'undefined') {
    PRODUCTION.HOST = hostEnv
  }
  CONFIG = PRODUCTION
}

CONFIG.LANGUAGES = {
  'zh-CN': '中文',
  'en-US': 'EN'
}

module.exports = {
  ...CONFIG
}
