import React, { Component } from 'react'
import { hot } from 'react-hot-loader'
import Index from './views/index'

class App extends Component {
  render () {
    return (
      <Index />
    )
  }
}

export default hot(module)(App)
