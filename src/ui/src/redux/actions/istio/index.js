import axios from '../../../commons/axios'

const TYPE = {
  DIAGNOSIS_DATA: 'DIAGNOSIS_DATA'
}

const getDiagnosisDataAjax = (username) => {
  return dispatch => {
    axios.getAjax({
      url: 'api/diagnose',
      type: 'GET',
      data: {}
    }).then(response => {
      if (response.code === 0) {
        dispatch({
          type: TYPE.DIAGNOSIS_DATA,
          payload: response.data
        })
      }
    })
  }
}

export {
  getDiagnosisDataAjax,
  TYPE
}
