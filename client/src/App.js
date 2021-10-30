import Card from "react-bootstrap/Card";
import Button from "react-bootstrap/Button";
import { update, getPosts, createPost } from "./actions/postsAction";
import { getAnswers } from "./actions/answerAction";
import { getGrades, updateGrade, postGrade } from "./actions/gradeAction";
import { getUser } from "./actions/userAction";
import { connect } from "react-redux";
import "./app.css";
import React, { Component } from "react";
import _ from "lodash";
import Answer from "./components/Answer/Answer";
import moment from "moment";

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
    if (this.state.post.title.length > 0) {
      this.props.createPost(this.state.post);
      const post = { ...this.state.post, title: "" };
      this.setState({ post });
    }
  }

  handleLike(postId) {
    const newGrade = {
      postid: postId,
      grade: 1,
    };
    // this.props.update(newGrade);
    this.props.postGrade(newGrade);
    this.props.updateGrade(newGrade);
  }

  handleDislike(postId) {
    const newGrade = {
      postid: postId,
      grade: -1,
    };

    // this.props.update(newGrade);
    this.props.postGrade(newGrade);
    this.props.updateGrade(newGrade);
  }

  handleLoadMore() {
    this.setState({ loadMore: this.state.loadMore + 5 });
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
        var gradeFilterLikes = _.filter(this.props.grades, function (a) {
          return a.postid === post.id && a.grade === 1;
        });
        var gradeFilterDislikes = _.filter(this.props.grades, function (a) {
          return a.postid === post.id && a.grade === -1;
        });
        return (
          <Card id="card" key={key}>
            <Card.Text id="cardText">
              <b style={{ fontSize: "20px" }}>{post.title}</b>
            </Card.Text>
            <Card.Body>
              <span id="date">
                {moment(post.postdate).fromNow()}
                <span id="comments">
                  {this.renderCountAnswers(post.id) === 1
                    ? this.renderCountAnswers(post.id) + " comment"
                    : this.renderCountAnswers(post.id) + " comments"}
                </span>
              </span>
              <br />
              <br />
              <i className="fa fa-thumbs-up fa-like" id="thumbUp">
                {gradeFilterLikes.length}
              </i>
              <i className="fa fa-thumbs-down fa-dislike" id="thumbDown">
                {gradeFilterDislikes.length}
              </i>
              <hr />
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

        var gradeFilterLikes = _.filter(this.props.grades, function (a) {
          return a.postid === post.id && a.grade === 1;
        });
        var gradeFilterDislikes = _.filter(this.props.grades, function (a) {
          return a.postid === post.id && a.grade === -1;
        });

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
                {moment(post.postdate).fromNow()}
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
                  cursor: "pointer",
                  fontSize: "18px",
                  color:
                    gradeFilter[0] !== undefined && gradeFilter[0].grade === 1
                      ? "green"
                      : null,
                }}
                onClick={() => this.handleLike(post.id)}
              >
                {gradeFilterLikes.length}
              </i>
              <i
                className="fa fa-thumbs-down fa-dislike"
                style={{
                  cursor: "pointer",
                  fontSize: "18px",
                  marginLeft: "20px",
                  marginBottom: "20px",
                  color:
                    gradeFilter[0] !== undefined && gradeFilter[0].grade === -1
                      ? "red"
                      : null,
                }}
                onClick={() => this.handleDislike(post.id)}
              >
                {gradeFilterDislikes.length}
              </i>
              <hr />
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
            <div className="question" style={{ background: "ghostwhite" }}>
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
        {this.props.posts.length > this.state.loadMore ? (
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
  updateGrade,
  postGrade,
})(App);
