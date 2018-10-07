import React, { Component } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { Form, Dropdown } from "semantic-ui-react";

export class OrientationForm extends Component {
  opOrientation = [
    {
      key: 0,
      text: "Heterosexuel",
      value: 0
    },
    {
      key: 1,
      text: "Homosexuel",
      value: 1
    },
    {
      key: 2,
      text: "Bisexuel",
      value: 2
    }
  ];

  static propTypes = {
    orientation: PropTypes.number.isRequired,
    onChange: PropTypes.func.isRequired
  };

  render() {
    const { orientation, onChange } = this.props;
    const { opOrientation } = this;
    return (
      <Form.Field>
        <label> Orientation Sexuelle </label>
        <Dropdown
          fluid
          selection
          name="orientation"
          options={opOrientation}
          value={orientation}
          onChange={onChange}
        />
      </Form.Field>
    );
  }
}

const mapStateToProps = state => ({});

const mapDispatchToProps = {};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(OrientationForm);
