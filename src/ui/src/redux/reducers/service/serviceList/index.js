import { TYPE } from '../../../actions/service/serviceList'

const initState = {
  podsInfo: {
    page: {
      pageIndex: 0,
      pageSize: 10
    },
    podsList: []
  },
  taskInfo: {
    page: {
      pageIndex: 0,
      pageSize: 10
    },
    taskList: []
  },
  keyPodsInfo: {
    page: {
      pageIndex: 0,
      pageSize: 10
    },
    podsList: []
  },
  topology: {},
  lastServiceItem: {},
  graphData: '',
  treeList: [],
  podsKey: []
}

export default (state = initState, action) => {
  switch (action.type) {
    case TYPE.SERVICE_PODS_DATA:
      let podsInfo = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {podsInfo})
    case TYPE.SERVICE_TASK_PAGE_LIST_DATA:
      let taskInfo = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {taskInfo})
    case TYPE.SERVICE_GRAPH_DATA:
      let topology = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {topology})
    case TYPE.SET_TREE_LIST:
      let treeList = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {treeList})
    case TYPE.GET_LAST_SERVICE_ITEM:
      let lastServiceItem = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {lastServiceItem})
    case TYPE.SET_GRAPH_DATA:
      let graphData = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {graphData})
    case TYPE.SET_SERVICE_KEY_DATA:
      let podsKey = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {podsKey})
    case TYPE.SET_SERVICE_KEY_PODS_DATA:
      let keyPodsInfo = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {keyPodsInfo})
    default:
      return state
  }
}
