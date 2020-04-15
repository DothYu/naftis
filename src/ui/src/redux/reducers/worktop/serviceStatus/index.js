import { TYPE } from '../../../actions/worktop/serviceStatus'

const initState = {
  podCount: 0,
  serviceCount: 0,
  globalSuccessRateData: [],
  fixxsData: [],
  foxxsData: [],
  fixxsByServiceData: [],
  foxxsByServiceData: []

}

export default (state = initState, action) => {
  state = JSON.parse(JSON.stringify(state))
  let payload = action.payload
  switch (action.type) {
    case TYPE.SERVICE_AND_PODS_DATA:
      return Object.assign({}, {
        ...state,
        ...payload
      })
    case TYPE.GLOBAL_SUCCESS_RATE_DATA:
      return Object.assign({}, {
        ...state,
        globalSuccessRateData: payload.result[0] ? payload.result[0].values : []
      })
    case TYPE.FIXXXS_BY_SERVICE_DATA:
      return Object.assign({}, {
        ...state,
        fixxsByServiceData: payload.result[0] ? payload.result : []
      })
    case TYPE.FOXXXS_BY_SERVICE_DATA:
      return Object.assign({}, {
        ...state,
        foxxsByServiceData: payload.result[0] ? payload.result : []
      })
    case TYPE.FIXXS_DATA:
      return Object.assign({}, {
        ...state,
        fixxsData: payload.result[0] ? payload.result[0].values : []
      })
    case TYPE.FOXXS_DATA:
      return Object.assign({}, {
        ...state,
        foxxsData: payload.result[0] ? payload.result[0].values : []
      })
    default:
      return state
  }
}
