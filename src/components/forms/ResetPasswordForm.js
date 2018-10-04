import React, { Component } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { Form, Button } from "semantic-ui-react";
import _ from "lodash";

import InlineError from "../messages/InlineError";

class ResetPasswordForm extends Component {
  static propTypes = {
    token: PropTypes.string.isRequired,
    submit: PropTypes.func.isRequired
  };

  state = {
    data: {
      token: this.props.token,
      password: "",
      passwordConfirmation: ""
    },
    loading: false,
    errors: {}
  };

  onChange = e =>
    this.setState({
      data: { ...this.state.data, [e.target.name]: e.target.value }
    });

  onSubmit = e => {
    e.preventDefault();
    const { state, props } = this;
    const errors = this.validate(state.data);
    this.setState({ errors });
    if (_.isEmpty(errors)) {
      this.setState({ loading: true });
      props
        .submit(state.data)
        .catch(err =>
          this.setState({ errors: err.response.data.errors, loading: false })
        );
    }
  };

  validate = data => {
    const errors = {};
    if (!data.password) errors.password = "Ne doit pas etre vide";
    if (data.password !== data.passwordConfirmation)
      errors.password = "Les mots de passe ne correspondent pas";
    return errors;
  };

  render() {
    const { data, errors, loading } = this.state;
    return (
      <Form loading={loading} onSubmit={this.onSubmit}>
        <Form.Field error={!!errors.password}>
          <label htmlFor="password">
            Mot de passe
            <input
              type="password"
              id="password"
              name="password"
              placeholder="Votre mot de passe"
              value={data.password}
              onChange={this.onChange}
            />
          </label>
          {errors.password && <InlineError text={errors.password} />}
        </Form.Field>

        <Form.Field error={!!errors.passwordConfirmation}>
          <label htmlFor="passwordConfirmation">
            Confirmer le mot de passe
            <input
              type="password"
              id="passwordConfirmation"
              name="passwordConfirmation"
              placeholder="Confirmez votre mot de passe"
              value={data.passwordConfirmation}
              onChange={this.onChange}
            />
          </label>
        </Form.Field>
        <Button primary>Reset password</Button>
      </Form>
    );
  }
}

const mapStateToProps = state => ({});

const mapDispatchToProps = {};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ResetPasswordForm);
