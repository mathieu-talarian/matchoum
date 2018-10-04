import React from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";

import { Menu, Dropdown, Divider } from "semantic-ui-react";
import { Link } from "react-router-dom";
import Avatar from "react-avatar";

import * as actions from "../../actions/auth";

const TopNavigation = ({ user, logout }) => (
  <Menu secondary pointing>
    <Menu.Item active={false}>{user.email}</Menu.Item>
    <Menu.Item as={Link} to="/dashboard">
      Dashboard
    </Menu.Item>
    <Menu.Menu position="right">
      <Dropdown trigger={<Avatar facebookId="100008343750912" size="150" />}>
        <Dropdown.Menu>
          <Dropdown.Item as={Link} to="/profil">
            Profil
          </Dropdown.Item>
          <Dropdown.Item as={Link} to="/account">
            Parametres de votre compte
          </Dropdown.Item>
          <Divider />
          <Dropdown.Item onClick={() => logout()}>Logout</Dropdown.Item>
        </Dropdown.Menu>
      </Dropdown>
    </Menu.Menu>
  </Menu>
);

TopNavigation.propTypes = {
  user: PropTypes.shape({
    email: PropTypes.string.isRequired
  }).isRequired,
  logout: PropTypes.func.isRequired
};

function mapStateToProps(state) {
  return {
    user: state.user
  };
}

export default connect(
  mapStateToProps,
  { logout: actions.logout }
)(TopNavigation);
