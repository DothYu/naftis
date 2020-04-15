import { createStore, applyMiddleware } from 'redux'
import thunk from 'redux-thunk'
import { composeWithDevTools } from 'redux-devtools-extension'
import reducers from '../reducers'

const initialState = {}

const configureStore = () => {
  let enhancer = applyMiddleware(thunk)
  const store = createStore(reducers, initialState, composeWithDevTools(enhancer))

  if (process.env.NODE_ENV === 'development') {
    if (module.hot) {
      module.hot.accept('../reducers/index.js', () => {
        store.replaceReducer(require('../reducers/index.js').default)
      })
    }
  }
  return store
}

export default configureStore
