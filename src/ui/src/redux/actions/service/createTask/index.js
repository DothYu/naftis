import axios from '../../../../commons/axios'

const TYPE = {
  SET_CURRENT_STEP: 'SET_CURRENT_STEP',
  CREATE_TASK_LIST: 'CREATE_TASK_LIST',
  CREATE_TASK_STATUS_DATA: 'CREATE_TASK_STATUS_DATA',
  SET_CREATE_TASK_ITEM: 'SET_CREATE_TASK_ITEM'
}

const setCurrentStepData = (currentStep) => ({
  type: TYPE.SET_CURRENT_STEP,
  payload: currentStep
})

const setCreateTaskListData = (createTaskList) => ({
  type: TYPE.CREATE_TASK_LIST,
  payload: createTaskList
})

const setCreateItemData = (createItem) => ({
  type: TYPE.SET_CREATE_TASK_ITEM,
  payload: createItem
})

const setCreateStatusData = (createStatus) => ({
  type: TYPE.CREATE_TASK_STATUS_DATA,
  payload: createStatus
})

const submitCreateTempAjax = (data, fn) => {
  return dispatch => {
    axios.getAjax({
      url: 'api/tasks',
      type: 'POST',
      data: data
    }).then(response => {
      fn && fn(response)
    })
  }
}

export {
  submitCreateTempAjax,
  setCurrentStepData,
  setCreateTaskListData,
  setCreateItemData,
  setCreateStatusData,
  TYPE
}
