import api from "../api";
import { PROFILE_MODIFIED, PROFILE_GET } from "../types";

const profileGet = profile => ({
  type: PROFILE_GET,
  profile
});

const profileMod = profile => ({
  type: PROFILE_MODIFIED,
  profile
});

export const add = profile => () => api.profile.add(profile);

export const get = () => dispatch =>
  api.profile.get().then(profile => {
    console.log(profile);
    dispatch(profileGet(profile));
  });

export const update = profile => api.profile.update(profile);

export const updateStatus = user => dispatch => {
  console.log(user);
  dispatch(profileMod(user));
};
