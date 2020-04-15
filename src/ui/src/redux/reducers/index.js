import { combineReducers } from 'redux'
import global from './global'
import login from './login'
import socket from './socket'
import serviceStatus from './worktop/serviceStatus'
import istio from './istio'
import serviceList from './service/serviceList'
import taskTemplate from './service/taskTemplate'
import createTask from './service/createTask'

export default combineReducers({
  global,
  socket,
  login,
  serviceStatus,
  istio,
  serviceList,
  taskTemplate,
  createTask
})
