import { PROFILE_MODIFIED, PROFILE_GET } from "../types";

export default function profile(state = {}, action = {}) {
  switch (action.type) {
    case PROFILE_GET:
      return action.profile;
    default:
      return state;
  }
}
