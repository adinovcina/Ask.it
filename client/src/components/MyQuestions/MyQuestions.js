import React, { Component } from "react";
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";
import { getMyQuestions } from "../../actions/myQuestionsAction";
import { connect } from "react-redux";
import _ from "lodash";

class MyQuestions extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loadMore: 5,
    };
    this.handleLoadMore = this.handleLoadMore.bind(this);
  }

  handleLoadMore() {
    this.setState({ loadMore: this.state.loadMore + 5 });
  }

  componentWillMount() {
    this.props.getMyQuestions();
  }

  renderQuestions() {
    return _.map(this.props.myQuestions, (post, key) => {
      return (
        <Card key={key} style={{ width: "85%", margin: "10px auto" }}>
          <Card.Header>Question</Card.Header>
          <Card.Body>
            <blockquote className="blockquote mb-0">
              <p>{post.title}</p>
              <footer className="blockquote-footer">
                <cite title="Source Title">
                  {post.User.firstname + " " + post.User.lastname}
                </cite>
              </footer>
            </blockquote>
          </Card.Body>
        </Card>
      );
    }).slice(0, this.state.loadMore);
  }

  render() {
    return (
      <>
        {this.renderQuestions()}
        {this.props.myQuestions.length > 3 ? (
          <Button
            variant="secondary"
            id="loadMorePosts"
            onClick={this.handleLoadMore}
          >
            Load more
          </Button>
        ) : null}
      </>
    );
  }
}

function mapStateToProps(state) {
  return {
    myQuestions: state.myQuestions,
  };
}

export default connect(mapStateToProps, { getMyQuestions })(MyQuestions);
