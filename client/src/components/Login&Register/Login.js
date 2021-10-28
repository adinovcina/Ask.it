import React, { Component } from "react";
import "./login.css";
import loginImg from "../../login.svg";
import { login } from "../../actions/loginAction";
import { connect } from "react-redux";
import PropTypes from "prop-types";
import { Redirect } from "react-router";

class Login extends Component {
  constructor(props) {
    super(props);
    this.state = {
      email: "",
      password: "",
      isLogged: false,
    };
    this.onChange = this.onChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
  }

  onChange(e) {
    this.setState({ [e.target.name]: e.target.value });
  }

  onSubmit(e) {
    e.preventDefault();
    const loginData = {
      email: this.state.email,
      password: this.state.password,
    };
    this.props.login(loginData);
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.user.message !== undefined) {
      document.getElementById("errorMsg").style.display = "block";
    } else {
      this.setState({ isLogged: true });
    }
  }

  render() {
    return (
      <>
        {this.state.isLogged ? (
          <Redirect
            push
            to={{
              pathname: "/",
            }}
          />
        ) : null}
        <form
          className="base-container"
          ref={this.props.containerRef}
          onSubmit={this.onSubmit}
        >
          <div className="header">Login</div>
          <div className="content">
            <div className="image">
              <img src={loginImg} alt="img" />
            </div>
            <div className="form">
              <div className="form-group">
                <label>Email</label>
                <input
                  onChange={this.onChange}
                  type="text"
                  name="email"
                  value={this.state.email}
                  placeholder="Email"
                  autoComplete="off"
                  required
                />
              </div>
              <div className="form-group">
                <label>Password</label>
                <input
                  onChange={this.onChange}
                  type="password"
                  name="password"
                  value={this.state.password}
                  placeholder="Password"
                  autoComplete="off"
                  required
                />
              </div>
              <p id="errorMsg">Invalid credentials. Please try again.</p>
            </div>
          </div>
          <div className="footer">
            <button
              type="submit"
              className="btn"
              id="loginBtn"
              style={{ color: "white" }}
            >
              Login
            </button>
          </div>
        </form>
      </>
    );
  }
}

Login.propTypes = {
  login: PropTypes.func.isRequired,
};

function mapStateToProps(state) {
  return {
    user: state.login,
  };
}

export default connect(mapStateToProps, { login })(Login);
