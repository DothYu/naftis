import { store } from '../../index'

export const SET_BREAD_CRUMBS = 'SET_BREAD_CRUMBS'

export const setBreadCrumbs = crumbsItems => {
  const dispatch = store.dispatch
  dispatch({type: SET_BREAD_CRUMBS, crumbsItems})
}
