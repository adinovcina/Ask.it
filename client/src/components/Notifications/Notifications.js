import React, { Component } from "react";
import { getUser } from "../../actions/userAction";
import { getNotifications } from "../../actions/notificationAction";
import { connect } from "react-redux";
import _ from "lodash";
import Card from "react-bootstrap/Card";
import NoDataFound from "../NoDataFound";
import moment from "moment";
import Button from "react-bootstrap/Button";

class Notifications extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loadMore: 7,
    };
    this.handleLoadMore = this.handleLoadMore.bind(this);
  }

  componentWillMount() {
    this.props.getUser();
    this.props.getNotifications(this.props.user.id);
  }

  handleLoadMore() {
    this.setState({ loadMore: this.state.loadMore + 5 });
  }

  renderNotifications() {
    return _.map(this.props.notifications, (not, key) => {
      console.log(not);
      return (
        <Card
          key={key}
          style={{ height: "80px", width: "80%", margin: "10px auto" }}
        >
          <Card.Body>
            <b>{not.User_sending}</b> replied to your post
            <p style={{ fontSize: "13px", color: "rgb(68, 196, 200)" }}>
              {moment(not.commentdate).fromNow()}
            </p>
          </Card.Body>
        </Card>
      );
    })
      .reverse()
      .slice(0, this.state.loadMore);
  }

  render() {
    return (
      <>
        {this.renderNotifications()}
        {this.props.notifications.length > this.state.loadMore ? (
          <Button
            variant="secondary"
            id="loadMorePosts"
            onClick={this.handleLoadMore}
          >
            Load more
          </Button>
        ) : null}
        {this.props.notifications.length == 0 ? <NoDataFound /> : null}
      </>
    );
  }
}

function mapStateToProps(state) {
  return {
    user: state.login,
    notifications: state.notifications,
  };
}

export default connect(mapStateToProps, { getUser, getNotifications })(
  Notifications
);
