import React, { Component } from 'react'
import { connect } from 'react-redux'
import * as Actions from '../../../redux/actions/service/serviceList'
import '../../../components/Forcegraph/index.css'
const forcegraph = require('../../../components/Forcegraph/index')

class Graph extends Component {
  componentDidMount () {
    this.getGraphData()
  }

  getGraphData = () => {
    const {graphData} = this.props
    if (!graphData || !graphData.nodes) {
      return
    }
    graphData && forcegraph.init(graphData)
  }

  componentWillUnmount () {
    // clear data
    forcegraph.clearData()
  }

  render () {
    this.getGraphData()

    return (
      <div id='total'>
        <div id='graph' />
        <div id='info'>
          <a>Close</a>
          <div id='incoming' className='conn-table'>
            <table>
              <tbody>
                <tr>
                  <th>1</th>
                  <th>2</th>
                </tr>
                <tr>
                  <td>3</td>
                  <td>4</td>
                </tr>
              </tbody>
            </table>
          </div>
          <div id='outgoing' className='conn-table'>
            <table>
              <tbody>
                <tr>
                  <th>1</th>
                  <th>2</th>
                </tr>
                <tr>
                  <td>3</td>
                  <td>4</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    )
  }
}

const mapStateToProps = state => ({
  graphData: state.serviceList.graphData
})

export default connect(mapStateToProps, Actions)(Graph)
