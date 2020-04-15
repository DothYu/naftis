import React, { Component } from 'react'
import { connect } from 'react-redux'
import { Table, Panel } from '@hi-ui/hiui/es'
import '@hi-ui/hiui/es/table/style/index.css'
import { setBreadCrumbs } from '../../redux/actions/global'
import * as Actions from '../../redux/actions/istio'
import './index.scss'

class Istio extends Component {
  constructor (props) {
    super(props)
    this.componentColumns = [
      {title: T('app.common.tb.svcName'), dataIndex: 'name', key: 'name'},
      {title: T('app.common.tb.svcType'), dataIndex: 'type', key: 'type'},
      {title: T('app.common.tb.svcClusterIP'), dataIndex: 'clusterIP', key: 'clusterIP'},
      {title: T('app.common.tb.svcExternalIP'), dataIndex: 'externalIP', key: 'externalIP', render: (text) => text || 'None'},
      {title: T('app.common.tb.svcPorts'), dataIndex: 'ports', key: 'ports'},
      {title: T('app.common.tb.svcAge'), dataIndex: 'age', key: 'age'}]
    this.podColumns = [
      {title: T('app.common.tb.podName'), dataIndex: 'name', key: 'name'},
      {title: T('app.common.tb.podReady'), dataIndex: 'ready', key: 'ready'},
      {title: T('app.common.tb.podStatus'),
        dataIndex: 'status',
        key: 'status',
        render: (text) => {
          let colorValue = ''
          switch (text) {
            case 'Pending':
              colorValue = '#1890FF'
              break
            case 'Running' || 'Succeeded':
              colorValue = '#52C41A'
              break
            case 'Failed':
              colorValue = '#F5222D'
              break
            default:
              colorValue = '#1890FF'
          }
          return (
            <div>
              <span style={{ color: colorValue, marginRight: '6px' }}>●</span>
              {text}
            </div>
          )
        }},
      {title: T('app.common.tb.podRestarts'), dataIndex: 'restarts', key: 'restarts'},
      {title: T('app.common.tb.podAge'), dataIndex: 'age', key: 'age'}]
  }
  componentDidMount () {
    this.props.getDiagnosisDataAjax()
    const crumbsItems = [
      {title: T('app.menu.istio'), to: '/istio'}
    ]
    setBreadCrumbs(crumbsItems)
  }

  // istio-file-link
  renderTop = () => {
    return (
      <div className='Diagnosis-top-wrap'>
        <h2>Istio</h2>
        <p>{T('app.common.istioCmt')}</p>
        <div className='istio-file'><a href='https://istio.io/docs/'>{T('app.common.istioDoc')}</a></div>
      </div>
    )
  }

  renderCenter = () => {
    const {components = [], pods = []} = this.props
    return (
      <div className='navMenu-wrap'>
        <div className='component-status'>
          <Panel title={<div className='col-panel-title'>{T('app.common.services')}</div>}>
            {components && components.length > 0
              ? <Table columns={this.componentColumns} data={components} /> : null
            }
          </Panel>
        </div>
        <div className='pods-status'>
          <Panel title={<div className='col-panel-title'>{T('app.common.pods')}</div>}>
            {pods && pods.length > 0
              ? <Table columns={this.podColumns} data={pods} /> : null
            }
          </Panel>
        </div>
      </div>
    )
  }

  render () {
    return (
      <div className='Diagnosis-wrap'>
        { this.renderTop() }
        { this.renderCenter() }
      </div>
    )
  }
}

const mapStateToProps = state => ({
  components: state.istio.components,
  pods: state.istio.pods
})

export default connect(mapStateToProps, Actions)(Istio)
