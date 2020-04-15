import React from 'react'
import Icon from '@hi-ui/hiui/es/icon'
import { addLocaleData, IntlProvider } from 'react-intl'
import { getLocaleLanguage } from './../commons/languages'

const { appLocaleData, locale, messages } = getLocaleLanguage()

addLocaleData(appLocaleData)
const { intl } = new IntlProvider({ locale: locale, messages }, {}).getChildContext()
window.T = (id) => {
  return intl.formatMessage({ id })
}

// Two level menu is allowed at most.
export default [
  {
    key: 9,
    title: T('app.menu.worktop'),
    to: '',
    icon: <Icon name='refer' />,
    children: [
      {key: 21, title: T('app.menu.worktop.overview'), to: '/worktop/overview'}
    ]
  },
  {
    key: 10,
    title: T('app.menu.service'),
    to: '',
    icon: <Icon name='usergroup' />,
    children: [
      {key: 21, title: T('app.menu.service.manager'), to: '/service/serviceList'}
    ]
  },
  {
    key: 12,
    title: T('app.menu.istio'),
    to: '/istio',
    icon: <Icon name='tool' />
  }
]
