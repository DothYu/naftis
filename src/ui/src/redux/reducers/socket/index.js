const initState = {
  socketData: '',
  connectStatus: false
}

export default (state = initState, action) => {
  switch (action.type) {
    case 'SET_SOCKET_DATA':
      return Object.assign({}, state, {socketData: action.payload})
    case 'SET_SOCKET_STATUS':
      return Object.assign({}, state, {connectStatus: action.payload})
    default:
      return state
  }
}
