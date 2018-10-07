import React, { Component } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { Header, Button, Modal } from "semantic-ui-react";

import TagForm from "../forms/TagForm";

export class AddTagModal extends Component {
  static propTypes = {
    submit: PropTypes.func.isRequired
  };

  state = {
    modalOpen: false
  };

  handleOpen = () => this.setState({ modalOpen: true });

  handleClose = () => this.setState({ modalOpen: false });

  render() {
    const { modalOpen } = this.state;
    const { submit } = this.props;
    return (
      <Modal
        trigger={
          <Button circular onClick={this.handleOpen} icon="plus" size="small" />
        }
        open={modalOpen}
        basic
        size="small"
        closeIcon
        onClose={this.handleClose}
      >
        <Header icon="plus" content="Ajouter un tag" />
        <Modal.Actions>
          <TagForm submit={submit} handleClose={this.handleClose} />
        </Modal.Actions>
      </Modal>
    );
  }
}

const mapStateToProps = state => ({});

const mapDispatchToProps = {};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AddTagModal);
