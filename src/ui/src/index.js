import React from 'react'
import ReactDOM from 'react-dom'
import { Provider } from 'react-redux'
import { addLocaleData, IntlProvider } from 'react-intl'
import App from './App'
import Login from './views/Login'
import configureStore from './redux/store'
import { getLocaleLanguage } from './commons/languages'
import './commons/common.scss'
import * as socketAction from './redux/actions/socket'

export const store = configureStore()

if (window.localStorage.getItem('isLogin')) {
  socketAction.connetctSocket()
}

if (module.hot) {
  module.hot.accept()
}

// initializes language components
const { appLocaleData, locale, messages } = getLocaleLanguage()
addLocaleData(appLocaleData)

ReactDOM.render((
  <Provider store={store}>
    <IntlProvider
      locale={locale}
      messages={messages}
    >
      {window.localStorage.getItem('isLogin') === 'true' ? <App /> : <Login />}
    </IntlProvider>
  </Provider>
), document.getElementById('app'))
