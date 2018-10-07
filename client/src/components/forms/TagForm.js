import React, { Component } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { Form, Input } from "semantic-ui-react";
import { isEmpty } from "lodash";
import InlineError from "../messages/InlineError";

export class TagForm extends Component {
  static propTypes = {
    submit: PropTypes.func.isRequired,
    handleClose: PropTypes.func.isRequired
  };

  state = {
    data: {
      tag: ""
    },
    loading: false,
    errors: {}
  };

  componentDidMount = () => {
    this.focus();
  };

  onSubmit = e => {
    this.setState({
      loading: true
    });
    const { submit, handleClose } = this.props;
    const { data } = this.state;
    console.log(data);
    e.preventDefault();
    const errors = this.validate(data.tag);
    if (isEmpty(errors)) {
      submit(data.tag)
        .then(() => handleClose())
        .catch(err =>
          this.setState({
            loading: false,
            errors: { tag: err.response.data.errors.global }
          })
        );
    } else {
      this.setState({
        loading: false,
        errors
      });
    }
  };

  onChange = e =>
    this.setState({
      data: { ...this.state.data, [e.target.name]: e.target.value }
    });

  validate = tag => {
    const errors = {};
    if (!tag) errors.tag = "Remplissez avec un nouveau tag";
    return errors;
  };

  focus = () => {
    this.inputRef.focus();
  };

  render() {
    const { data, loading, errors } = this.state;
    return (
      <Form onSubmit={this.onSubmit}>
        <Form.Field error={!!errors.tag}>
          <Input
            ref={c => {
              this.inputRef = c;
            }}
            loading={loading}
            action="ajouter un nouveau tag"
            label="Nouveau tag"
            type="text"
            id="tag"
            name="tag"
            inverted
            placeholder="tag"
            focus
            value={data.tag}
            onChange={this.onChange}
          />
          {errors.tag && <InlineError text={errors.tag} />}
        </Form.Field>
      </Form>
    );
  }
}

const mapStateToProps = state => ({});

const mapDispatchToProps = {};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(TagForm);
