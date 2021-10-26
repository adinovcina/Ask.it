import React, { Component } from "react";
import { connect } from "react-redux";
import loginImg from "../../login.svg";
// import { login } from "../actions/loginAction";
// import { passwordChange } from "../actions/passwordAction";

export class Profile extends Component {
  constructor(props) {
    super(props);
    this.state = {
      username: "",
      oldpassword: "",
      newpassword: "",
    };
    this.onChange = this.onChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  componentWillMount() {
    this.setState({ username: this.props.user.username });
  }

  onSubmit(e) {
    e.preventDefault();
    const passChange = {
      username: this.state.username,
      oldpassword: this.state.oldpassword,
      newpassword: this.state.newpassword,
      isOldPasswordValid: false,
    };
    this.props.passwordChange(passChange);
  }

  onChange(e) {
    this.setState({ [e.target.name]: e.target.value });
  }

  componentWillReceiveProps(nextProp) {
    if (nextProp.mypassword.message !== undefined) {
      this.setState({ isOldPasswordValid: false });
      document.getElementById("errorMsg").style.display = "block";
      setTimeout(() => {
        document.getElementById("errorMsg").style.display = "none";
      }, 1000);
    } else if (nextProp.mypassword === this.state.username) {
      this.setState({ isOldPasswordValid: true });
      document.getElementById("errorMsg").style.display = "block";
      document.getElementById("errorMsg").style.color = "green";

      document.getElementById("errorMsg").innerHTML =
        "Password successfully changed";
      setTimeout(() => {
        document.getElementById("errorMsg").style.display = "none";
      }, 1000);
      window.location.href = "/";
    }
  }

  render() {
    return (
      <form
        onSubmit={this.onSubmit}
        className="base-container"
        ref={this.props.containerRef}
        id="form"
        style={{
          margin: "auto",
          backgroundColor: "rgb(198,234,213)",
          width: "30%",
          paddingBottom: "20px",
          paddingTop: "20px",
          marginTop: "50px",
          borderRadius: "50px",
        }}
      >
        <div className="header">
          Hello, <b>{this.state.username}</b>
        </div>
        <div className="content">
          <div className="image">
            <img src={loginImg} alt="img" />
          </div>
          <div className="form">
            <div className="form-group">
              <label style={{ fontSize: "17px" }}>Username</label>
              <input
                onChange={this.onChange}
                type="text"
                name="username"
                value={this.state.username}
                readOnly
                placeholder="Username"
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
                placeholder="Password"
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
                placeholder="Password"
                autoComplete="off"
              />
            </div>
            <p
              id="errorMsg"
              style={{
                color: "red",
                fontSize: "14px",
                marginTop: "10px",
                display: "none",
              }}
            >
              Incorrect old password
            </p>
          </div>
        </div>
        <div className="footer">
          <button
            type="submit"
            className="btn"
            style={{ backgroundColor: "rgb(125,212,162)", fontWeight: "bold" }}
          >
            Change password
          </button>
        </div>
      </form>
    );
  }
}

const mapStateToProps = (state) => {
  return {
    user: state.login,
    mypassword: state.password,
  };
};

export default connect(mapStateToProps, {})(Profile);
