import { c_account } from '../actions/constants'

export default function accountReducer(state={
  current: {},
  error: undefined,
}, action) {

  switch (action.type) {

  case c_account.READ_SUCCESS:
    return Object.assign({}, state, {
      current: action.body,
    })

  case c_account.READ_FAILURE:
    return Object.assign({}, state, {
      error: action.error,
    })

  default:
    return state

  }

}
