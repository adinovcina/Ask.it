import React, { Component } from "react";
import Navbar from "react-bootstrap/Navbar";
import Nav from "react-bootstrap/Nav";
import Container from "react-bootstrap/Container";
import NavDropdown from "react-bootstrap/NavDropdown";
import { connect } from "react-redux";
import { getUser } from "../../actions/userAction";
import { logout } from "../../actions/loginAction";
import Pusher from "pusher-js";
import "./navbar.css";

class MyNavbar extends Component {
  constructor(props) {
    super(props);
    this.state = {
      notifications: [],
      isClicked: false,
    };
    this.readNotifications = this.readNotifications.bind(this);
  }

  componentWillMount() {
    this.props.getUser();
    Pusher.logToConsole = true;
    const this2 = this;

    const pusher = new Pusher("d88f98f756818528eb43", {
      cluster: "eu",
    });

    const channel = pusher.subscribe("notification");
    channel.bind("sendNotif", function (data) {
      const notif_array = this2.state.notifications;
      notif_array.push(data);
      this2.setState({ notifications: notif_array, isClicked: false });
    });
  }

  readNotifications() {
    this.setState({ isClicked: true });
  }

  showAllNotifications() {
    window.location.href = "/notifications";
  }

  render() {
    let count = 0;
    return (
      <>
        {this.state.notifications.map((not) => {
          if (not.User_receiving === this.props.user.id) {
            count++;
          }
        })}
        <Navbar collapseOnSelect expand="lg" variant="dark" id="navbar">
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
              {this.props.user.id !== undefined ? (
                <Nav>
                  <NavDropdown
                    onClick={this.readNotifications}
                    title={
                      <span
                        id="spanNotif"
                        style={{
                          color:
                            count !== 0 && !this.state.isClicked
                              ? "rgb(68, 196, 200)"
                              : "rgba(255,255,255,.55)",
                        }}
                      >
                        <i className="fa fa-bell" aria-hidden="true"></i>
                        <span className="num">{count}</span>
                      </span>
                    }
                  >
                    <NavDropdown.Item
                      id="showAll"
                      onClick={this.showAllNotifications}
                    >
                      Show all
                    </NavDropdown.Item>
                    <NavDropdown.Divider />
                    {this.state.notifications.map((not, key) => {
                      if (not.User_receiving === this.props.user.id) {
                        return (
                          <div key={key}>
                            <NavDropdown.Item id="item">
                              <b>{not.User_sending} </b>replied to your post
                            </NavDropdown.Item>
                            <NavDropdown.Divider />
                          </div>
                        );
                      }
                    })}
                  </NavDropdown>
                </Nav>
              ) : null}
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
            </Navbar.Collapse>
          </Container>
        </Navbar>
      </>
    );
  }
}

function mapStateToProps(state) {
  return {
    user: state.login,
  };
}

export default connect(mapStateToProps, { logout, getUser })(MyNavbar);
