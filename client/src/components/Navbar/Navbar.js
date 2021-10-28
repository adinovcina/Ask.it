import React, { Component } from "react";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import Container from "react-bootstrap/Container";
import { connect } from "react-redux";
import { getUser } from "../../actions/userAction";
import { logout } from "../../actions/loginAction";
import { store } from "../../store";

class MyNavbar extends Component {
  componentWillMount() {
    this.props.getUser();
  }
  render() {
    return (
      <Navbar
        collapseOnSelect
        expand="lg"
        variant="dark"
        style={{ background: "#02223c" }}
      >
        <Container>
          <Navbar.Brand href="/">Home</Navbar.Brand>
          <Navbar.Toggle aria-controls="responsive-navbar-nav" />
          <Navbar.Collapse id="responsive-navbar-nav">
            <Nav className="me-auto">
              <Nav.Link href="/mostLikes">Hot questions</Nav.Link>
              <Nav.Link href="/mostAnswers">Most answers</Nav.Link>
              {this.props.user.id !== undefined ? (
                <Nav.Link href="/myQuestions">My questions</Nav.Link>
              ) : null}
            </Nav>
            <Nav>
              {this.props.user.id !== undefined ? (
                <Nav.Link href="/profile">Profile</Nav.Link>
              ) : null}
              {this.props.user.id !== undefined ? (
                <Nav.Link href="/login" onClick={() => this.props.logout()}>
                  Logout
                </Nav.Link>
              ) : (
                <Nav.Link href="/login">Login</Nav.Link>
              )}
            </Nav>
            {/* {store.subscribe(() => {
              console.log(store.getState());
            })} */}
          </Navbar.Collapse>
        </Container>
      </Navbar>
    );
  }
}

function mapStateToProps(state) {
  return {
    user: state.login,
  };
}

export default connect(mapStateToProps, { logout, getUser })(MyNavbar);
