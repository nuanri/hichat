import { combineReducers } from 'redux'

import accountReducer from './account'
import messageReducer from './message'

export default combineReducers({
  account: accountReducer,
  message: messageReducer,
})
