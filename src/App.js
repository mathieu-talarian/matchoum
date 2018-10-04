import React from "react";
import { Route, Switch } from "react-router-dom";
import { connect } from "react-redux";
import PropTypes from "prop-types";

import HomePage from "./components/pages/HomePage";
import LoginPage from "./components/pages/LoginPage";
import SignupPage from "./components/pages/SignupPage";
import ConfirmationPage from "./components/pages/ConfirmationPage";
import ForgotPasswordPage from "./components/pages/ForgotPasswordPage";
import ResetPasswordPage from "./components/pages/ResetPasswordPage";
import ProfilPage from "./components/pages/ProfilPage";
import ErrorPage from "./components/pages/ErrorPage";

import TopNavigation from "./components/navigation/TopNavigation";

import DashboardPage from "./components/pages/DashboardPage";
import UserRoute from "./components/routes/UserRoute";
import GuestRoute from "./components/routes/GuestRoute";

const App = ({ isAuthenticated, location }) => (
  <div className="ui container">
    {isAuthenticated && <TopNavigation />}
    <Switch>
      <GuestRoute location={location} path="/" exact component={HomePage} />
      <Route
        location={location}
        path="/confirmation/:token"
        exact
        component={ConfirmationPage}
      />
      <GuestRoute
        location={location}
        path="/login"
        exact
        component={LoginPage}
      />
      <GuestRoute
        location={location}
        path="/signup"
        exact
        component={SignupPage}
      />

      <GuestRoute
        location={location}
        path="/forgot_password"
        exact
        component={ForgotPasswordPage}
      />
      <GuestRoute
        location={location}
        path="/reset_password/:token"
        exact
        component={ResetPasswordPage}
      />
      <UserRoute
        location={location}
        path="/dashboard"
        exact
        component={DashboardPage}
      />
      <UserRoute
        location={location}
        path="/profil"
        exaxt
        component={ProfilPage}
      />
      <Route component={ErrorPage} />
    </Switch>
  </div>
);

App.propTypes = {
  location: PropTypes.shape({
    pathname: PropTypes.string.isRequired
  }).isRequired,
  isAuthenticated: PropTypes.bool.isRequired
};

const mapStateToProps = state => ({
  isAuthenticated: !!state.user.email
});

export default connect(
  mapStateToProps,
  null
)(App);
