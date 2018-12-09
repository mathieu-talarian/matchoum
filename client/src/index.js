import React from "react";
import ReactDOM from "react-dom";
import { BrowserRouter, Route } from "react-router-dom";
import "semantic-ui-css/semantic.min.css";
import { applyMiddleware, createStore } from "redux";
import { Provider } from "react-redux";
import { composeWithDevTools } from "redux-devtools-extension";
import thunk from "redux-thunk";
import decode from "jwt-decode";

import App from "./App";
import registerServiceWorker from "./registerServiceWorker";

import rootReducer from "./rootReducer";
import { userLoggedIn } from "./actions/auth";

import setAuthorizationHeader from "./utils/setAuthorizationHeader";

const store = createStore(
  rootReducer,
  composeWithDevTools(applyMiddleware(thunk))
);

if (localStorage.JWT) {
  const payload = decode(localStorage.JWT);
  const user = {
    token: localStorage.JWT,
    email: payload.email,
    confirmed: payload.confirmed,
    profile: payload.profile
  };
  setAuthorizationHeader(localStorage.JWT);
  store.dispatch(userLoggedIn(user));
}
console.log("test")
ReactDOM.render(
  <BrowserRouter>
    <Provider store={store}>
      <Route component={App} />
    </Provider>
  </BrowserRouter>,
  document.getElementById("root")
);
registerServiceWorker();
