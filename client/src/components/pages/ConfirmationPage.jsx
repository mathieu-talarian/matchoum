import React, { Component } from "react";
import { Message, Icon } from "semantic-ui-react";
import { Link } from "react-router-dom";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { confirm } from "../../actions/auth";

class ConfirmationPage extends Component {
  state = {
    loading: true,
    success: false
  };

  componentDidMount() {
    this.props
      .confirm(this.props.match.params.token)
      .then(() => this.setState({ loading: false, success: true }))
      .catch(() => this.setState({ loading: false, success: false }));
  }

  render() {
    const { loading, success } = this.state;
    return (
      <div>
        {loading && (
          <Message icon>
            <Icon name="circle notched" loading />
            <Message.Header>
              Votre email est en cours de validation
            </Message.Header>
          </Message>
        )}

        {!loading &&
          success && (
            <Message icon>
              <Icon name="checkmark" />
              <Message.Content>
                <Message.Header>Votre Email est validé</Message.Header>
                <Link to="/dashboard">Dashboard</Link>
              </Message.Content>
            </Message>
          )}

        {!loading &&
          !success && (
            <Message icon>
              <Icon name="checkmark" />
              <Message.Header>Un probleme est survenu</Message.Header>
            </Message>
          )}
      </div>
    );
  }
}

ConfirmationPage.propTypes = {
  confirm: PropTypes.func.isRequired,
  match: PropTypes.shape({
    params: PropTypes.shape({
      token: PropTypes.string.isRequired
    }).isRequired
  }).isRequired
};

export default connect(
  null,
  { confirm }
)(ConfirmationPage);
