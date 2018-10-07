import axios from "axios";

export const url = "/api/v1";

export default {
  user: {
    login: credentials =>
      axios.post(`${url}/auth`, { credentials }).then(res => res.data.user),
    signup: user =>
      axios.post(`${url}/users`, { user }).then(res => res.data.user),
    confirm: token =>
      axios
        .post(`${url}/auth/confirmation`, { token })
        .then(res => res.data.user),
    resetPasswordRequest: email =>
      axios.post(`${url}/auth/reset_password_request`, email),
    validateToken: token => axios.post(`${url}/auth/validate_token`, { token }),
    resetPassword: data => axios.post(`${url}/auth/reset_password`, { data })
  },
  tags: {
    searchTags: query =>
      axios.get(`${url}/tags/search?q=${query}`).then(res => res.data),
    fetchAll: () => axios.get(`${url}/tags/all`).then(res => res.data),
    newTag: tag => axios.post(`${url}/tags`, { tag })
  },
  profile: {
    add: profile => axios.post(`${url}/profile`, { profile }),
    get: () => axios.get(`${url}/profile`).then(res => res.data.profile),
    upadte: profile => axios.put(`${url}/profile`, { profile })
  }
};
