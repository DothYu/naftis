import axios from '../../../../commons/axios'

const TYPE = {
  SERVICE_TEMPLATE_LIST_DATA: 'SERVICE_TEMPLATE_LIST_DATA',
  SERVICE_KUBE_LIST_DATA: 'SERVICE_KUBE_LIST_DATA',
  SERVICE_MODULE_LIST_DATA: 'SERVICE_MODULE_LIST_DATA',
  SERVICE_ADD_PARAM_DATA: 'SERVICE_ADD_PARAM_DATA',
  SET_ADD_DATA: 'SET_ADD_DATA'
}

const setModuleListData = moduleList => ({
  type: TYPE.SERVICE_MODULE_LIST_DATA,
  payload: moduleList
})

const setAddParamData = (key, value) => ({
  type: TYPE.SERVICE_ADD_PARAM_DATA,
  payload: { key, value }
})

const setAddData = submitParam => ({
  type: TYPE.SET_ADD_DATA,
  payload: submitParam
})

const getServiceTemplateDataAjax = () => {
  return dispatch => {
    axios
      .getAjax({
        url: 'api/tasktmpls',
        type: 'GET',
        data: ''
      })
      .then(response => {
        if (response.code === 0) {
          response.data.push({ type: 'add' })
          dispatch({
            type: TYPE.SERVICE_TEMPLATE_LIST_DATA,
            payload: response.data
          })
        }
      })
  }
}

const getKubeInfoAjax = () => {
  return dispatch => {
    axios
      .getAjax({
        url: 'api/kube/info',
        type: 'GET',
        data: ''
      })
      .then(response => {
        if (response.code === 0) {
          let ns = []
          if (response.data.namespaces && response.data.namespaces.length) {
            response.data.namespaces.map((item, index) => {
              ns.push({
                id: index.toString(),
                name: item
              })
            })
          }
          dispatch({
            type: TYPE.SERVICE_KUBE_LIST_DATA,
            payload: {namespaces: ns}
          })
        }
      })
  }
}

const commitServiceTemplateDataAjax = (data, fn) => {
  return dispatch => {
    axios
      .getAjax({
        url: 'api/tasktmpls',
        type: 'POST',
        data: {
          name: data.name,
          brief: data.brief,
          content: data.content,
          vars: data.vars
        }
      })
      .then(response => {
        if (response.code === 0) {
          fn && fn()
        }
      })
  }
}

const deleteServiceTemplateDataAjax = (data, fn) => {
  return dispatch => {
    axios
      .getAjax({
        url: `api/tasktmpls/${data.tplID}`,
        type: 'DELETE'
      })
      .then(response => {
        if (response.code === 0) {
          fn && fn()
        }
      })
  }
}

const getTemplateDetailDataAjax = (data, fn) => {
  return dispatch => {
    axios
      .getAjax({
        url: 'api/tasktmpls',
        type: 'POST',
        data: {
          name: data.name,
          brief: data.brief,
          content: data.content,
          vars: data.vars
        }
      })
      .then(response => {
        if (response.code === 0) {
          fn && fn()
        }
      })
  }
}

export {
  getServiceTemplateDataAjax,
  commitServiceTemplateDataAjax,
  deleteServiceTemplateDataAjax,
  setAddData,
  setModuleListData,
  setAddParamData,
  getTemplateDetailDataAjax,
  getKubeInfoAjax,
  TYPE
}
