import React, { Component } from "react";
import PropTypes from "prop-types";
import { Form, Button, Message } from "semantic-ui-react";
import { isEmail } from "validator";
import _ from "lodash";

import InlineError from "../messages/InlineError";

export default class ForgotPasswordForm extends Component {
  static propTypes = {
    submit: PropTypes.func.isRequired
  };

  state = {
    data: {
      email: ""
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
    const { state } = this;
    const errors = this.validate(state.data);
    this.setState({ errors });
    if (_.isEmpty(errors)) {
      this.setState({ loading: true });
      this.props
        .submit(state.data)
        .catch(err =>
          this.setState({ errors: err.response.data.errors, loading: false })
        );
    }
  };

  validate = data => {
    const errors = {};
    if (!isEmail(data.email)) errors.email = "email invalide";
    return errors;
  };

  render() {
    const { errors, data, loading } = this.state;

    return (
      <Form onSubmit={this.onSubmit} loading={loading}>
        {!!errors.global && <Message negative>{errors.global}</Message>}
        <Form.Field error={!!errors.email}>
          <label htmlFor="email">
            Email
            <input
              type="email"
              id="email"
              name="email"
              placeholder="Votre email"
              value={data.email}
              onChange={this.onChange}
            />
          </label>
          {!!errors.email && <InlineError text={errors.email} />}
        </Form.Field>
        <Button primary>Reinitialiser le mot de passe</Button>
      </Form>
    );
  }
}
