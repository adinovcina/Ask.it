import React, { Component } from "react";
import { connect } from "react-redux";
import Card from "react-bootstrap/Card";
import { getAnswerGrades } from "../../actions/answerGradeAction";
import { update } from "../../actions/answerAction";
import { getUser } from "../../actions/userAction";
import {
  getAnswers,
  createAnswer,
  editAnswer,
  deleteAnswer,
} from "../../actions/answerAction";
import _ from "lodash";
import "./answer.css";
import Button from "react-bootstrap/Button";
import $ from "jquery";
import SweetAlert from "react-bootstrap-sweetalert";

class Answer extends Component {
  constructor(props) {
    super(props);
    this.state = {
      loadMore: 3,
      comment: "",
      show: false,
      editComment: "",
    };
    this.handleLoadMore = this.handleLoadMore.bind(this);
    this.handleHide = this.handleHide.bind(this);
    this.openCommentInput = this.openCommentInput.bind(this);
    this.handleChange = this.handleChange.bind(this);
    this.handleClick = this.handleClick.bind(this);
    this.onConfirm = this.onConfirm.bind(this);
    this.handleEditComment = this.handleEditComment.bind(this);
    this.onChange = this.onChange.bind(this);
    this.handleCancel = this.handleCancel.bind(this);
    this.handleDelete = this.handleDelete.bind(this);
  }

  handleEditComment(answer, e) {
    const inputElement = e.target.parentElement.children[2];
    const cancelBtn = e.target.parentElement.children[8];
    inputElement.value = answer;
    cancelBtn.style.display = "block";
    inputElement.style.display = "block";
  }

  handleDelete(e, ansId) {
    this.props.deleteAnswer(ansId);
  }

  handleCancel(answer, e) {
    const inputElement = e.target.parentElement.children[2];
    inputElement.style.display = "none";
    inputElement.value = answer;
    e.target.style.display = "none";
  }

  handleChange(e) {
    this.setState({ comment: e.target.value });
  }

  onChange(ansId, e) {
    const editAnswer = {
      id: ansId,
      answer: e.target.value,
    };
    const inputBtn = e.target;
    const editBtn = e.target.parentElement.children[8];
    inputBtn.style.display = "none";
    editBtn.style.display = "none";
    this.props.editAnswer(editAnswer);
  }

  handleClick(e) {
    var postId = $(e.target.parentElement.parentElement.parentElement).find(
      "#postId"
    );
    var p = parseInt(postId[0].innerHTML);
    const answer = {
      answer: this.state.comment,
      postid: p,
    };
    this.props.createAnswer(answer);
    this.setState({ comment: "" });
  }

  handleLike(id, postid) {
    const grade = {
      answerid: id,
      postid: postid,
      grade: 1,
    };
    this.props.update(grade);
    setTimeout(() => {
      this.props.getAnswerGrades();
      this.props.getAnswers();
    }, 100);
  }

  onConfirm() {
    this.setState({ show: false });
  }

  successMessage() {
    return (
      <SweetAlert
        danger
        title="Sorry, you must be logged in to give comment!"
        onConfirm={this.onConfirm}
      ></SweetAlert>
    );
  }

  openCommentInput(e) {
    if (this.props.user.id !== undefined) {
      var commentDiv = $(e.target.parentElement).find("#commentDiv");
      if (commentDiv[0].style.display === "none")
        commentDiv[0].style.display = "block";
      else commentDiv[0].style.display = "none";
    } else {
      this.setState({ show: true });
    }
  }

  handleDislike(id, postid) {
    const grade = {
      answerid: id,
      postid: postid,
      grade: -1,
    };
    this.props.update(grade);
    setTimeout(() => {
      this.props.getAnswerGrades();
      this.props.getAnswers();
    }, 100);
  }

  handleHide() {
    this.setState({ loadMore: 0 });
  }

  handleLoadMore() {
    var increase = (this.state.loadMore += 3);
    this.setState({ loadMore: increase });
  }

  componentWillMount() {
    this.props.getAnswers();
    this.props.getUser();
    this.props.getAnswerGrades();
  }

  getDifferenceInDays(date1, date2) {
    return Math.round((date1 - date2) / (1000 * 60 * 60 * 24), 1);
  }

  renderOneAnswer() {
    if (_.isEmpty(this.props.user)) {
      var post = this.props.postId;
      var filter = _.filter(this.props.answers, function (a) {
        return post === a.postid;
      });
      return filter
        .map((ans, key) => {
          let days = this.getDifferenceInDays(
            new Date(),
            new Date(ans.postdate)
          );
          return (
            <Card.Text key={key}>
              <span id="answerUser">
                {ans.User.firstname + " " + ans.User.lastname + " : "}
              </span>
              {ans.answer}
              <i
                className="fa fa-thumbs-up fa-like"
                id="thumbUp"
                style={{ paddingLeft: "15px" }}
              >
                {ans.likes}
              </i>
              <i className="fa fa-thumbs-down fa-dislike" id="thumbDown">
                {ans.dislikes}
              </i>
              <i id="date" style={{ marginLeft: "25px" }}>
                {days === 1
                  ? days + " day ago"
                  : days === 0
                  ? "today"
                  : days + " days ago"}
              </i>
            </Card.Text>
          );
        })
        .slice(0, this.state.loadMore);
    } else {
      var postLogged = this.props.postId;
      var filterLogged = _.filter(this.props.answers, function (a) {
        return postLogged === a.postid;
      });
      return filterLogged
        .map((ans, key) => {
          let days = this.getDifferenceInDays(
            new Date(),
            new Date(ans.postdate)
          );

          var userId = this.props.user.id;
          var gradeFilter = _.filter(this.props.answerGrades, function (a) {
            return (
              a.postid === ans.postid &&
              a.userid === userId &&
              a.answerid === ans.id
            );
          });
          return (
            <Card.Text key={key}>
              <span id="answerUser">
                {userId === ans.userid
                  ? "You : "
                  : ans.User.firstname + " " + ans.User.lastname + " : "}
              </span>
              {ans.answer}
              <br />
              <input
                style={{ width: "100%", display: "none" }}
                defaultValue={ans.answer}
                onKeyPress={(e) => {
                  if (e.key === "Enter") {
                    e.preventDefault();
                    this.onChange(ans.id, e);
                  }
                }}
              />
              <i
                className="fa fa-thumbs-up fa-like"
                style={{
                  paddingLeft: "15px",
                  color:
                    gradeFilter[0] !== undefined && gradeFilter[0].grade === 1
                      ? "green"
                      : null,
                }}
                onClick={() => this.handleLike(ans.id, ans.postid)}
              >
                {ans.likes}
              </i>
              <i
                className="fa fa-thumbs-down fa-dislike"
                style={{
                  paddingLeft: "20px",
                  color:
                    gradeFilter[0] !== undefined && gradeFilter[0].grade === -1
                      ? "red"
                      : null,
                }}
                onClick={() => this.handleDislike(ans.id, ans.postid)}
              >
                {ans.dislikes}
              </i>
              <i id="date" style={{ marginLeft: "25px" }}>
                {days === 1
                  ? days + " day ago"
                  : days === 0
                  ? "today"
                  : days + " days ago"}
              </i>
              {userId === ans.userid ? (
                <i
                  className="fa fa-trash"
                  aria-hidden="true"
                  id="trashBtn"
                  onClick={(e) => this.handleDelete(e, ans.id)}
                ></i>
              ) : null}
              {userId === ans.userid ? (
                <i
                  className="fa fa-edit"
                  id="editBtn"
                  onClick={(e) => this.handleEditComment(ans.answer, e)}
                ></i>
              ) : null}
              {userId === ans.userid ? (
                <i
                  className="fa fa-times"
                  aria-hidden="true"
                  id="cancelBtn"
                  onClick={(e) => this.handleCancel(ans.answer, e)}
                ></i>
              ) : null}
            </Card.Text>
          );
        })
        .reverse()
        .slice(0, this.state.loadMore);
    }
  }

  render() {
    return (
      <>
        {this.state.show ? this.successMessage() : null}
        {this.renderOneAnswer()}
        {this.props.numberOfComments > this.state.loadMore ? (
          <Button
            variant="outline-secondary"
            id="loadMore"
            onClick={this.handleLoadMore}
          >
            Load more
          </Button>
        ) : null}
        {this.props.numberOfComments > 0 ? (
          <Button
            variant="outline-secondary"
            id="hide"
            onClick={this.handleHide}
          >
            Hide comments
          </Button>
        ) : null}
        {this.props.numberOfComments > 0 ? (
          <Button
            variant="outline-secondary"
            id="commentBtn"
            onClick={this.openCommentInput}
          >
            Comment
          </Button>
        ) : (
          <Button
            variant="outline-secondary"
            id="noCommentBtn"
            onClick={this.openCommentInput}
          >
            Comment
          </Button>
        )}
        <div id="commentDiv" style={{ display: "none" }}>
          <p id="commentTitle">Give your thoughts :</p>
          <textarea
            id="commentInput"
            value={this.state.comment}
            onChange={this.handleChange}
            onKeyPress={(e) => {
              if (e.key === "Enter") {
                e.preventDefault();
                this.handleClick(e);
              }
            }}
          />
        </div>
      </>
    );
  }
}

function mapStateToProps(state) {
  return {
    answers: state.answers,
    answerGrades: state.answerGrades,
    user: state.user,
  };
}

export default connect(mapStateToProps, {
  getAnswers,
  getUser,
  getAnswerGrades,
  update,
  createAnswer,
  editAnswer,
  deleteAnswer,
})(Answer);
