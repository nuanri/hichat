import React from 'react'
import ReactDOM from 'react-dom'
import { createStore, combineReducers } from 'redux'
import { Provider } from 'react-redux'
import { Router, Route } from 'react-router'
import { createHistory } from 'history'
import { syncReduxAndRouter, routeReducer } from 'redux-simple-router'
//import reducers from '<project-path>/reducers'

import { reducer as formReducer } from 'redux-form'

import AuthSignUpForm from './components/signupform'
import App from './containers/app'

const reducer = combineReducers(Object.assign({},  {
  routing: routeReducer,
  form: formReducer,
}))
const store = createStore(reducer)
const history = createHistory()

syncReduxAndRouter(history, store)

console.log('App = ', App)
ReactDOM.render(
  <Provider store={store}>
    <Router history={history}>
      <Route path="/" component={App}>
        <Route path="auth/signup" component={AuthSignUpForm}/>
      </Route>
    </Router>
  </Provider>,
  document.getElementById('root')
)
