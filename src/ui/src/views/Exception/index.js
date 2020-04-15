import React from 'react'
import { setBreadCrumbs } from '../../redux/actions/global'
import './index.scss'

class Exception extends React.Component {
  componentDidMount () {
    const crumbsItems = [
      { title: T('app.common.home'), to: '/' }
    ]
    setBreadCrumbs(crumbsItems)
  }

  setText (status) {
    let text = ''
    switch (status) {
      case '403': text = T('app.common.err403')
        break
      case '404': text = T('app.common.err404')
        break
      case '500': text = T('app.common.err500')
        break
      default: text = T('app.common.err403')
    }
    return text
  }

  // exception page
  render () {
    const location = this.props.location
    let status = location ? location.pathname : '/404'
    status = status.slice(1)
    let text = this.setText(status)

    return (
      <div className='exception'>
        <div className='exception__content'>
          <div className='status'>{status}</div>
          <div className='text'>{text}</div>
        </div>
      </div>
    )
  }
}

export default Exception
