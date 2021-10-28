import React, { Component } from "react";
import Card from "react-bootstrap/Card";
import { connect } from "react-redux";
import _ from "lodash";
import { getMostAnswers } from "../../actions/mostAnswersAction";

class MostAnswers extends Component {
  componentWillMount() {
    this.props.getMostAnswers();
  }

  renderPosts() {
    return _.map(this.props.mostAnswers, (ans, key) => {
      return (
        <Card key={key} style={{ width: "80%", margin: "10px auto" }}>
          <Card.Body>
            <b>{ans.user.firstname + " " + ans.user.lastname}</b> with number of
            answers : {ans.numberofanswers}
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
    mostAnswers: state.mostAnswers,
  };
}

export default connect(mapStateToProps, { getMostAnswers })(MostAnswers);
