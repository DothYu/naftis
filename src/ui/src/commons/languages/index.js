import $$ from '../tools'
import CONFIG from '../../config'
import Cookies from 'js-cookie'
import { IntlProvider, addLocaleData } from 'react-intl'

const LANGUAGES = CONFIG.LANGUAGES

// get ISTIO_LANG_KEY from cookie
const getIstioLangFromCookie = () => {
  let ISTIO_LANG = $$.getCookie(CONFIG.ISTIO_LANG_KEY)
  if (!ISTIO_LANG) {
    ISTIO_LANG = 'en-US'
  }
  return ISTIO_LANG
}

const getLangFromCookie = () => {
  let ISTIO_LANG = $$.getCookie(CONFIG.ISTIO_LANG_KEY)
  if (!ISTIO_LANG) {
    return false
  }
  return true
}

const getLocaleLanguage = () => {
  const ISTIO_LANG = getIstioLangFromCookie()
  const language = require(`./lib/${ISTIO_LANG}`)
  return language && language.default
}

const setLanguageCookie = (lang) => {
  Cookies.set(CONFIG.ISTIO_LANG_KEY, lang)
}

const setDefaultLanguageCookie = () => {
  Cookies.set(CONFIG.ISTIO_LANG_KEY, getLocaleLanguage().locale)
}

const { appLocaleData, locale, messages } = getLocaleLanguage()
addLocaleData(appLocaleData)
const { intl } = new IntlProvider({ locale: locale, messages }, {}).getChildContext()
window.T = (id) => {
  return intl.formatMessage({ id })
}

export {
  getLangFromCookie,
  getIstioLangFromCookie,
  getLocaleLanguage,
  setLanguageCookie,
  setDefaultLanguageCookie,
  LANGUAGES
}
