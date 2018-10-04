import React, { Component } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";
import { Dropdown, Container } from "semantic-ui-react";

import AddTagModal from "../modals/AddTagModal";
import InlineError from "../messages/InlineError";

import { fetchTags, addTag } from "../../actions/tags";

export class SearchTagsForm extends Component {
  static propTypes = {
    tags: PropTypes.arrayOf(PropTypes.number).isRequired,
    onTagsChange: PropTypes.func.isRequired,
    fetchTags: PropTypes.func.isRequired,
    addTag: PropTypes.func.isRequired
  };

  constructor(props) {
    super(props);
    this.state = {
      searchQuery: "",
      options: [],
      loading: false,
      disabled: true,
      error: undefined
    };
  }

  componentDidMount = () => {
    this.setState({ loading: true, disabled: true });
    this.fillTags();
  };

  onSearchChange = (e, data) => {
    this.setState({
      searchQuery: data.searchQuery
    });
  };

  submit = tag => this.props.addTag(tag).then(() => this.fillTags());

  fillTags = () =>
    this.props
      .fetchTags()
      .then(data => {
        console.log(data);
        if (data && data.result.tags.length > 0) {
          const options = [];
          data.result.tags.forEach(id => {
            options.push({
              key: id,
              value: id,
              text: data.entities.tag[id].tag
            });
          });
          this.setState({ disabled: false, loading: false, options });
        } else {
          this.setState({
            disabled: true,
            loading: false,
            error: "Il n'y a pas encore de tags, veillez en rajouter"
          });
        }
      })
      .catch(err => {
        this.setState({
          disabled: true,
          loading: false,
          error: err.response.data.errors.global + " can't get tags"
        });
      });

  render() {
    const { searchQuery, options, loading, disabled, error } = this.state;
    const { tags, onTagsChange } = this.props;
    return (
      <Container>
        <Dropdown
          disabled={disabled}
          placeholder="Interets"
          fluid
          search
          selection
          name="op"
          multiple
          searchQuery={searchQuery}
          value={tags}
          loading={loading}
          options={options}
          onChange={onTagsChange}
          onSearchChange={this.onSearchChange}
        />
        {error && <InlineError text={error} />}
        <AddTagModal submit={this.submit} />
      </Container>
    );
  }
}

const mapStateToProps = state => ({});

const mapDispatchToProps = {
  fetchTags,
  addTag
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SearchTagsForm);
