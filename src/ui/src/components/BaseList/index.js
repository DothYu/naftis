import React, { Component } from 'react'
import './index.scss'

class BaseList extends Component {
  render () {
    const {list, onClick} = this.props
    return (
      <div className='menu-wrap'>
        {
          list.length ? list.map((item, index) => {
            return (
              <div className={item.disabled ? 'menu-item menu-item-disabled' : 'menu-item'} key={item.id} onClick={() => {
                if (item.disabled) return
                onClick(item, index)
              }}>
                {item.prefix}
                <span className='menu-item-title'>{item.title}</span>
              </div>
            )
          }) : null
        }
      </div>
    )
  }
}

export default BaseList
