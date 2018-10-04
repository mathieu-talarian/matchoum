import { combineReducers } from "redux";

import user from "./reducers/user";
import profile from "./reducers/profile";

export default combineReducers({
  user,
  profile
});
