import { SET_BREAD_CRUMBS } from '../actions/global'

const initState = {
  crumbsItems: [
    {title: 'Index', to: '/'}
  ]
}

export default (state = initState, action) => {
  const { type, crumbsItems } = action
  state = JSON.parse(JSON.stringify(state))

  switch (type) {
    case SET_BREAD_CRUMBS:
      return Object.assign({}, state, {crumbsItems})
    default:
      return state
  }
}
