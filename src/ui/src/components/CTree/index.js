import React, { Component } from 'react'
import './index.scss'

class CTree extends Component {
  constructor (props) {
    super(props)
    this.state = {
      treeList: [],
      name: '',
      nameSearch: '',
      lastChoose: ''
    }
  }

  componentDidMount () {
    const {treeList, name} = this.props
    this.setState({
      treeList,
      name
    })
  }

  componentWillReceiveProps (props) {
    this.setState({
      treeList: props.treeList,
      name: props.name,
      open: props.open,
      nameSearch: props.nameSearch
    })
    if (props.lastChoose && !this.props.lastChoose) {
      this.setState({
        lastChoose: props.lastChoose
      })
    }
  }

  renderItem = (treeList, fn) => {
    const {onClick} = this.props
    const {name, lastChoose} = this.state
    return (
      (treeList && treeList.length) ? treeList.map((item, index) => {
        return (
          <div className='CTree-item' key={index}>
            <div className='item-content'>
              <div className={`content-word ${lastChoose === item[name] ? 'word-choose' : ''}`} dangerouslySetInnerHTML={{__html: item[name]}}
                onClick={(e) => {
                  e.stopPropagation()
                  this.setState({
                    lastChoose: item[name]
                  })
                  onClick(item, index)
                }} />
              <div className='item-dot' onClick={(e) => {
                e.stopPropagation()
                if (item.children && item.children.length) {
                  treeList[index].isOpen = !item.isOpen
                  fn && fn()
                }
              }}>
                <div className={(item.children && item.children.length) ? (item.isOpen ? 'item-open item-strangle' : 'item-strangle') : null} />
              </div>
            </div>
            <div className='item-child'>
              {
                (item.isOpen && item.children && item.children.length) ? this.renderItem(item.children, fn) : ''
              }
            </div>
          </div>
        )
      }) : null
    )
  }

  render () {
    const {treeList} = this.state
    return (
      <div className='CTree-wrap'>
        {
          treeList && treeList.length && this.renderItem(treeList, () => {
            this.setState({treeList})
          })
        }
      </div>
    )
  }
}

export default CTree
