import axios from "axios";

export default (token = null) => {
  if (token) {
    axios.defaults.headers.common.authorization = `Bearer ${token}`;
  } else {
    delete axios.default.headers.common.authorization;
  }
};
