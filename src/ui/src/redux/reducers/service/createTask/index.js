import { TYPE } from '../../../actions/service/createTask'

const initState = {
  currentStep: 0,
  createTaskList: [],
  createStatus: ''
}

export default (state = initState, action) => {
  switch (action.type) {
    case TYPE.SET_CURRENT_STEP:
      return Object.assign({}, state, {currentStep: action.payload})
    case TYPE.CREATE_TASK_LIST:
      let createTaskList = JSON.parse(JSON.stringify(action.payload))
      return Object.assign({}, state, {createTaskList})
    case TYPE.CREATE_TASK_STATUS_DATA:
      let createStatus = action.payload
      return Object.assign({}, state, {createStatus})
    default:
      return state
  }
}
