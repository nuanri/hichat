import { c_message } from '../actions/constants'

export default function messageReducer(state={
  lists: [],
  onlineUsers: [],
  error: undefined,
}, action) {

  switch (action.type) {

  case c_message.LIST_SUCCESS:
    return Object.assign({}, state, {
      lists: action.data.body,
      onlineUsers: action.data.onlineusers,
    })

  case c_message.LIST_FAILURE:
    return Object.assign({}, state, {
      error: action.error,
    })

  case c_message.SENDING:
    return Object.assign({}, state, {
      lists: [...state.lists, {msg: action.data.body}],
    })

  default:
    return state

  }

}
