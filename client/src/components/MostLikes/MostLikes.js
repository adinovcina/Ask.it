import React, { Component } from "react";
import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";
import { connect } from "react-redux";
import _ from "lodash";
import { getMostLikes } from "../../actions/mostLikesAction";
import dateFormat from "dateformat";

class MostLikes extends Component {
  componentWillMount() {
    this.props.getMostLikes();
  }

  renderPosts() {
    return _.map(this.props.mostLikes, (post, key) => {
      return (
        <Card
          border="secondary"
          key={key}
          style={{ margin: "10px 0px 0px 170px", width: "80%" }}
        >
          <Card.Header>{dateFormat(post.postdate)}</Card.Header>
          <Card.Body>
            <Card.Title>{post.title}</Card.Title>
            <Card.Text>
              <i
                className="fa fa-thumbs-up fa-like"
                style={{
                  color: "green",
                }}
              >
                {post.likes}
              </i>
            </Card.Text>
          </Card.Body>
        </Card>
      );
    });
  }

  render() {
    return this.renderPosts();
  }
}

function mapStateToProps(state) {
  return {
    mostLikes: state.mostLikes,
  };
}

export default connect(mapStateToProps, { getMostLikes })(MostLikes);
