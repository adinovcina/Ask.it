import React, { Component } from "react";
import { connect } from "react-redux";
import loginImg from "../../login.svg";
import { getUser } from "../../actions/userAction";
import Button from "react-bootstrap/Button";
import "./profile.css";
import { passwordChange } from "../../actions/passwordAction";
import { logout } from "../../actions/loginAction";

export class Profile extends Component {
  constructor(props) {
    super(props);
    this.state = {
      firstname: "",
      lastname: "",
      oldpassword: "",
      email: "",
      newpassword: "",
      showFirstName: "",
      showLastName: "",
    };
    this.onChange = this.onChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  componentWillMount() {
    this.props.getUser();
    this.setState({
      firstname: this.props.user.firstname,
      lastname: this.props.user.lastname,
      email: this.props.user.email,
      showFirstName: this.props.user.firstname,
      showLastName: this.props.user.lastname,
    });
  }

  onSubmit(e) {
    e.preventDefault();
    const passChange = {
      firstname: this.state.firstname,
      lastname: this.state.lastname,
      email: this.state.email,
      oldpassword: this.state.oldpassword,
      newpassword: this.state.newpassword,
    };
    this.props.passwordChange(passChange);
  }

  onChange(e) {
    this.setState({ [e.target.name]: e.target.value });
  }

  componentWillReceiveProps(nextProp) {
    console.log(nextProp);
    if (nextProp.mypassword.message !== undefined) {
      document.getElementById("msg").style.display = "block";
    } else if (nextProp.mypassword === this.state.email) {
      document.getElementById("msg").style.display = "block";
      document.getElementById("msg").style.color = "green";
      document.getElementById("msg").innerHTML =
        "Password successfully changed";
      this.props.logout();
      window.location.href = "/login";
    }
  }

  render() {
    const { firstname, lastname } = this.state;
    return (
      <form
        onSubmit={this.onSubmit}
        className="base-container"
        ref={this.props.containerRef}
        id="formProfile"
      >
        <div className="header" id="head">
          Hello,{" "}
          <b>{this.state.showFirstName + " " + this.state.showLastName}</b>
        </div>
        <div className="content">
          <div className="image">
            <img src={loginImg} alt="img" />
          </div>
          <div className="form">
            <div className="form-group">
              <label style={{ fontSize: "17px" }}>First name</label>
              <input
                onChange={this.onChange}
                type="text"
                name="firstname"
                required
                minLength={3}
                value={firstname}
                placeholder="First name"
                autoComplete="off"
              />
            </div>
            <div className="form-group">
              <label style={{ fontSize: "17px" }}>Last name</label>
              <input
                onChange={this.onChange}
                type="text"
                name="lastname"
                required
                minLength={3}
                value={lastname}
                placeholder="Last name"
                autoComplete="off"
              />
            </div>
            <div className="form-group">
              <label style={{ fontSize: "17px" }}>Old Password</label>
              <input
                onChange={this.onChange}
                type="password"
                name="oldpassword"
                required
                minLength={5}
                placeholder="Old password"
                autoComplete="off"
              />
            </div>
            <div className="form-group">
              <label style={{ fontSize: "17px" }}>New Password</label>
              <input
                onChange={this.onChange}
                type="password"
                name="newpassword"
                required
                minLength={5}
                placeholder="New password"
                autoComplete="off"
              />
            </div>
            <p id="msg">Incorrect old password</p>
          </div>
        </div>
        <div className="footer">
          <Button type="submit" id="btnChange">
            Change profile
          </Button>
        </div>
      </form>
    );
  }
}

const mapStateToProps = (state) => {
  return {
    user: state.user,
    mypassword: state.password,
  };
};

export default connect(mapStateToProps, { getUser, passwordChange, logout })(
  Profile
);
