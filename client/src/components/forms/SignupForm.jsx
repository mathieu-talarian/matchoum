import React, { Component } from "react";
import PropTypes from "prop-types";
import { Form, Button } from "semantic-ui-react";
import _ from "lodash";
import Validator from "validator";

import InlineError from "../messages/InlineError";

class SignupForm extends Component {
  state = {
    data: {
      firstname: "",
      lastname: "",
      pseudo: "",
      email: "",
      password: "",
      confirmPassword: ""
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
          this.setState({ errors: err.response.data.errors.global, loading: false })
        );
    }
  };

  validate = data => {
    const errors = {};
    if (!Validator.isEmail(data.email)) errors.email = "Invalid email";
    if (!data.password) errors.password = "Can't be blank";
    return errors;
  };

  render() {
    const { onSubmit } = this;
    const { data, loading, errors } = this.state;
    return (
      <Form onSubmit={onSubmit} loading={loading}>
        <Form.Field error={!!errors.firstname}>
          <label htmlFor="firstname">
            Prenom
            <input
              type="text"
              id="firstname"
              name="firstname"
              placeholder="Votre prenom"
              value={data.firstname}
              onChange={this.onChange}
            />
          </label>
          {/* {errors.firstname && <InlineError text={errors.firstname} />} */}
        </Form.Field>

        <Form.Field error={!!errors.lastname}>
          <label htmlFor="lastname">
            {" "}
            Nom
            <input
              type="text"
              id="lastname"
              name="lastname"
              placeholder="Votre nom"
              value={data.lastname}
              onChange={this.onChange}
            />
          </label>
          {/* {errors.lastname && <InlineError text={errors.lastname} />} */}
        </Form.Field>

        <Form.Field error={!!errors.pseudo}>
          <label htmlFor="pseudo">
            Pseudo
            <input
              type="text"
              id="pseudo"
              name="pseudo"
              placeholder="placeholder"
              value={data.pseudo}
              onChange={this.onChange}
            />
          </label>
          {/* {errors.pseudo && <InlineError text={errors.pseudo} />} */}
        </Form.Field>

        <Form.Field error={!!errors.email}>
          <label htmlFor="email">
            Email
            <input
              type="email"
              id="email"
              name="email"
              placeholder="your email here"
              value={data.email}
              onChange={this.onChange}
            />
          </label>
          {errors.email && <InlineError text={errors.email} />}
        </Form.Field>
        <Form.Field error={!!errors.password}>
          <label htmlFor="password">
            Password
            <input
              type="password"
              id="password"
              name="password"
              placeholder="your pwd here"
              value={data.password}
              onChange={this.onChange}
            />
          </label>
          {/* {errors.password && <InlineError text={errors.password} />} */}
        </Form.Field>

        <Form.Field error={!!errors.confirmPassword}>
          <label htmlFor="confirmPassword">
            {" "}
            Confirmez votre mot de passe
            <input
              type="password"
              id="confirmPassword"
              name="confirmPassword"
              placeholder="placeholder"
              value={data.confirmPassword}
              onChange={this.onChange}
            />
          </label>
          {/* {errors.confirmPassword && (
            <InlineError text={errors.confirmPassword} />
          )} */}
        </Form.Field>
        <Button primary>Signup</Button>
      </Form>
    );
  }
}

SignupForm.propTypes = {
  submit: PropTypes.func.isRequired
};

export default SignupForm;
