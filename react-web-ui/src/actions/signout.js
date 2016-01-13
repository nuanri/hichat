import { api_get } from '../api'
import { pushPath } from 'redux-simple-router'
import { readAccountInfo } from './account_read'


export function signoutFromServer() {
  return dispatch => {
    return api_get('/auth/signout')
    .then(function(json) {
      if (!json.error) {
        localStorage.removeItem('sid')
        dispatch(readAccountInfo())
        dispatch(pushPath("/auth/signup"))
      }
    })
    .catch(error => console.log('signout err = ', error))
  }
}
