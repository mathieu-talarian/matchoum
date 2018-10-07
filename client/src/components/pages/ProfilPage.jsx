import React, { Component } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";
import { Container, Header, Icon, Message, Button } from "semantic-ui-react";
import { isEmpty } from "lodash";

import ProfilForm from "../forms/ProfilForm";

import * as profileActions from "../../actions/profile";

class ProfilPage extends Component {
  static propTypes = {
    isConfirmed: PropTypes.bool.isRequired,
    profile: PropTypes.shape({}).isRequired,
    add: PropTypes.func.isRequired,
    get: PropTypes.func.isRequired,
    update: PropTypes.func.isRequired,
    updateState: PropTypes.func.isRequired
  };

  constructor(props) {
    const { get, profile } = props;
    super(props);
    if (isEmpty(profile)) {
      get().then(prof =>
        this.setState({
          profile: prof
        })
      );
    }
  }

  componentDidMount = () => {};

  submit = data => this.props.add(data);

  mod = e => {
    const { user, updateState } = this.props;
    e.preventDefault();
    updateState({
      ...user,
      profile: !user.profile
    });
  };

  render() {
    const { isConfirmed, profile } = this.props;
    return (
      <Container>
        {!isConfirmed && <Redirect to="/dashboard" />}
        <Container textAlign="center">
          <Header as="h2" icon>
            <Icon name="settings" />
            Parametres de votre profil
            <Header.Subheader>
              Gerez ici vos informations personnelles
            </Header.Subheader>
          </Header>
        </Container>
        {!profile && (
          <Message
            warning
            floating
            attached
            header="Attention"
            content="Vous devez remplir vos infos avant de commencer a Matcher"
          />
        )}
        <ProfilForm submit={this.submit} profile={this.props.profile} />
        <Button onClick={this.mod}>Click</Button>
      </Container>
    );
  }
}

const mapStateToProps = state => ({
  isConfirmed: !!state.user.confirmed,
  profile: state.profile,
  user: state.user
});

const mapDispatchToProps = {
  add: profileActions.add,
  get: profileActions.get,
  update: profileActions.update,
  updateState: profileActions.updateStatus
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ProfilPage);
