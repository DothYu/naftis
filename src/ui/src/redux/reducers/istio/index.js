import { TYPE } from '../../actions/istio'

const initState = {
  components: [],
  pods: []
}

export default (state = initState, action) => {
  switch (action.type) {
    case TYPE.DIAGNOSIS_DATA:
      let components = JSON.parse(JSON.stringify(action.payload.components))
      let pods = JSON.parse(JSON.stringify(action.payload.pods))
      return Object.assign({}, state, {components, pods})
    default:
      return state
  }
}
