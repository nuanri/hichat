import { combineReducers, createStore, applyMiddleware, compose } from 'redux'
import { routeReducer } from 'redux-simple-router'

import thunkMiddleware from 'redux-thunk'
import createLogger from 'redux-logger';

import mainReducer from '../reducers/index'
import { reducer as formReducer } from 'redux-form'

const reducer = combineReducers(Object.assign({},  {
  routing: routeReducer,
  form: formReducer,
  main: mainReducer,
}))

const logger = createLogger();
let middlewares = [thunkMiddleware, logger]

const finalCreateStore = compose(
  applyMiddleware(...middlewares)
)(createStore)

export default function configureStore(initialState) {
  return finalCreateStore(reducer, initialState)
}
