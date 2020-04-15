import { TYPE } from '../../../actions/service/taskTemplate'

const initState = {
  templateList: [],
  moduleList: [],
  submitParam: {
    name: '',
    brief: '',
    content: '',
    vars: []
  },
  kubeinfo: {}
}

export default (state = initState, action) => {
  switch (action.type) {
    case TYPE.SERVICE_TEMPLATE_LIST_DATA:
      let templateList = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {templateList})
    case TYPE.SERVICE_MODULE_LIST_DATA:
      let moduleList = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {moduleList})
    case TYPE.SERVICE_ADD_PARAM_DATA:
      const {key, value} = action.payload
      let submitParam = JSON.parse(JSON.stringify(state.submitParam))
      submitParam[key] = value
      return Object.assign({}, state, {submitParam})
    case TYPE.SET_ADD_DATA:
      let addData = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {submitParam: addData})
    case TYPE.SERVICE_KUBE_LIST_DATA:
      let kubeinfo = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {kubeinfo: kubeinfo})
    default:
      return state
  }
}
