import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";
import { update, getPosts, createPost } from "./actions/postsAction";
import { getAnswers } from "./actions/answerAction";
import { getGrades } from "./actions/gradeAction";
import { getUser } from "./actions/userAction";
import { connect } from "react-redux";
import "./app.css";
import React, { Component } from "react";
import _ from "lodash";
import Answer from "./components/Answer/Answer";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loadMore: 10,
      post: {
        title: "",
      },
    };
    this.handleLoadMore = this.handleLoadMore.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(e) {
    const post = { ...this.state.post, title: e.target.value };
    this.setState({ post });
  }

  handleSubmit(e) {
    e.preventDefault();
    this.props.createPost(this.state.post);
    const post = { ...this.state.post, title: "" };
    this.setState({ post });
  }

  handleLike(postId) {
    const grade = {
      postid: postId,
      grade: 1,
    };
    this.props.update(grade);
    setTimeout(() => {
      this.props.getGrades();
      this.props.getPosts();
    }, 100);
  }

  handleDislike(postId) {
    const grade = {
      postid: postId,
      grade: -1,
    };
    this.props.update(grade);
    setTimeout(() => {
      this.props.getGrades();
      this.props.getPosts();
    }, 100);
  }

  handleLoadMore() {
    var increase = (this.state.loadMore += 5);
    this.setState({ loadMore: increase });
  }

  getDifferenceInDays(date1, date2) {
    return Math.round((date1 - date2) / (1000 * 60 * 60 * 24), 1);
  }

  componentWillMount() {
    this.props.getGrades();
    this.props.getPosts();
    this.props.getUser();
  }

  renderCountAnswers(postid) {
    var filter = _.filter(this.props.answers, function (a) {
      return a.postid === postid;
    });
    return filter.length;
  }

  renderPosts() {
    if (_.isEmpty(this.props.user)) {
      return _.map(this.props.posts, (post, key) => {
        let days = this.getDifferenceInDays(
          new Date(),
          new Date(post.postdate)
        );
        return (
          <Card id="card" key={key}>
            <Card.Text id="cardText">
              <b style={{ fontSize: "20px" }}>{post.title}</b>
            </Card.Text>
            <Card.Body>
              <span id="date">
                {days === 1
                  ? days + " day ago"
                  : days === 0
                  ? "Today"
                  : days + " days ago"}
                <span id="comments">
                  {this.renderCountAnswers(post.id) === 1
                    ? this.renderCountAnswers(post.id) + " comment"
                    : this.renderCountAnswers(post.id) + " comments"}
                </span>
              </span>
              <br />
              <br />
              <i className="fa fa-thumbs-up fa-like" id="thumbUp">
                {post.likes}
              </i>
              <i className="fa fa-thumbs-down fa-dislike" id="thumbDown">
                {post.dislikes}
              </i>
              <Answer
                postId={post.id}
                numberOfComments={this.renderCountAnswers(post.id)}
              />
            </Card.Body>
            <span>
              <i>{"post by" + " : "}</i>
              <i id="postBy">
                {post.User.firstname + " " + post.User.lastname}
              </i>
            </span>
          </Card>
        );
      }).slice(0, this.state.loadMore);
    } else {
      return _.map(this.props.posts, (post, key) => {
        var userId = this.props.user.id;
        var gradeFilter = _.filter(this.props.grades, function (a) {
          return a.postid === post.id && a.userid === userId;
        });

        let days = this.getDifferenceInDays(
          new Date(),
          new Date(post.postdate)
        );
        return (
          <Card id="card" key={key}>
            <i id="postId" style={{ display: "none" }}>
              {post.id}
            </i>
            <Card.Text id="cardText">
              <b style={{ fontSize: "20px" }}>{post.title}</b>
            </Card.Text>
            <Card.Body>
              <span id="date">
                {days === 1
                  ? days + " day ago"
                  : days === 0
                  ? "Today"
                  : days + " days ago"}
                <span id="comments">
                  {this.renderCountAnswers(post.id) === 1
                    ? this.renderCountAnswers(post.id) + " comment"
                    : this.renderCountAnswers(post.id) + " comments"}
                </span>
              </span>
              <br />
              <br />
              <i
                className="fa fa-thumbs-up fa-like"
                style={{
                  color:
                    gradeFilter[0] !== undefined && gradeFilter[0].grade === 1
                      ? "green"
                      : null,
                }}
                onClick={() => this.handleLike(post.id)}
              >
                {post.likes}{" "}
              </i>
              <i
                className="fa fa-thumbs-down fa-dislike"
                style={{
                  paddingLeft: "20px",
                  paddingBottom: "20px",
                  color:
                    gradeFilter[0] !== undefined && gradeFilter[0].grade === -1
                      ? "red"
                      : null,
                }}
                onClick={() => this.handleDislike(post.id)}
              >
                {post.dislikes}
              </i>
              <Answer
                postId={post.id}
                numberOfComments={this.renderCountAnswers(post.id)}
              />
            </Card.Body>
            <span>
              <i>{"post by" + " : "}</i>
              <i id="postBy">
                {userId === post.Userid
                  ? "You"
                  : post.User.firstname + " " + post.User.lastname}
              </i>
            </span>
          </Card>
        );
      })
        .reverse()
        .slice(0, this.state.loadMore);
    }
  }

  render() {
    return (
      <>
        {this.props.user.id !== undefined ? (
          <div id="postForm">
            <header className="header">
              <i
                className="fa fa-question"
                aria-hidden="true"
                id="questionnaire"
              >
                ask...
              </i>
            </header>
            <div className="question">
              <input
                id="question"
                onChange={this.handleChange}
                value={this.state.post.title}
              />
            </div>
            <button
              name="button"
              type="submit"
              id="btnSend"
              className="btn btn-danger"
              onClick={this.handleSubmit}
            >
              <i className="fa fa-paper-plane" aria-hidden="true"></i>
            </button>
          </div>
        ) : null}
        {this.renderPosts()}
        <Button
          variant="secondary"
          id="loadMorePosts"
          onClick={this.handleLoadMore}
        >
          Load more
        </Button>
      </>
    );
  }
}

function mapStateToProps(state) {
  return {
    posts: state.posts,
    answers: state.answers,
    grades: state.grades,
    user: state.user,
  };
}

export default connect(mapStateToProps, {
  getPosts,
  getAnswers,
  getGrades,
  getUser,
  update,
  createPost,
})(App);
