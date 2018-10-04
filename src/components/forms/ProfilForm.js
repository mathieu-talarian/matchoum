import React, { Component } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { Form, Button, Radio, Dropdown, TextArea } from "semantic-ui-react";

import SearchTagsForm from "./SearchTagsForm";
import InlineError from "../messages/InlineError";
import { OrientationForm } from "./OrientationForm";

const opGender = ["Homme", "Femme"];

class ProfilForm extends Component {
  static propTypes = {
    profile: PropTypes.shape({
      bio: PropTypes.string,
      gender: PropTypes.number,
      orientation: PropTypes.number,
      tags: PropTypes.arrayOf(PropTypes.number)
    }).isRequired,
    submit: PropTypes.func.isRequired
  };

  state = {
    data: this.props.profile,
    loading: false,
    errors: {}
  };

  componentDidMount = () => {};

  onChange = (e, datas) =>
    datas
      ? this.setState({
          data: { ...this.state.data, [datas.name]: datas.value }
        })
      : this.setState({
          data: { ...this.state.data, [e.target.name]: e.target.value }
        });

  onSubmit = e => {
    e.preventDefault();
    this.props.submit(this.state.data);
  };

  onSearchChange = (e, data) => {
    clearTimeout(this.timer);
    this.setState({
      tags: {
        query: [data]
      }
    });
    this.timer = setTimeout(this.fetchOptions, 1000);
  };

  onRadioChange = (e, data) => {
    this.setState({
      data: { ...this.state.data, [e.target.name]: data.checked ? 1 : 0 }
    });
  };

  onTagsChange = (e, datas) => {
    console.log(datas);
    const { value } = datas;
    this.setState({
      data: { ...this.state.data, tags: value }
    });
  };

  printState = e => {
    e.preventDefault();
    console.log(this.state);
  };

  fetchOptions = () => {
    const { tags } = this.state;
    if (!tags.query) return;
    this.setState({ tags: { loading: true } });
  };

  render() {
    const { data, loading, errors } = this.state;
    const gender = `Sexe ${opGender[data.gender]}`;
    return (
      <Form
        className="attached fluid segment"
        loading={loading}
        onSubmit={this.onSubmit}
      >
        <Form.Field>
          <label>Genre</label>
          <Radio
            id="gender"
            name="gender"
            toggle
            label={gender}
            onChange={this.onRadioChange}
          />
        </Form.Field>
        <Form.Field error={!!errors.orientation}>
          <OrientationForm
            orientation={data.orientation}
            onChange={this.onChange}
          />
          {errors.orientation && <InlineError text={errors.orientation} />}
        </Form.Field>
        <Form.Field error={!!errors.bio}>
          <label htmlFor="bio"> Bio </label>
          <TextArea
            type="textarea"
            id="bio"
            name="bio"
            placeholder="Dites en plus sur vous"
            value={data.bio}
            onChange={this.onChange}
          />

          {errors.bio && <InlineError text={errors.bio} />}
        </Form.Field>
        <SearchTagsForm tags={data.tags} onTagsChange={this.onTagsChange} />

        <Button primary>Valider</Button>
        <button onClick={this.printState}>x</button>
      </Form>
    );
  }
}

const mapStateToProps = state => ({});

const mapDispatchToProps = {};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ProfilForm);
