import React from 'react'
import { Icon, Tooltip } from '@hi-ui/hiui/es'
import './index.scss'

class NavMenu extends React.Component {
  render () {
    return (
      <div className='nav-menu'>
        <a href='http://bbs.xiaomi.cn' key='1'><Icon name='search' /></a>
        <a href='https://github.com/xiaomi-info' key='2'><Tooltip title='Tutorial' placement='bottom'><Icon name='edit' /></Tooltip></a>
        <a href='http://www.mi.com' key='3'><Icon name='info-circle-o' /></a>
      </div>
    )
  }
}

export default NavMenu
