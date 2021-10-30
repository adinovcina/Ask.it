import React, { Component } from "react";
import "./login.css";
import loginImg from "../../login.svg";
import { register } from "../../actions/registerAction";
import { connect } from "react-redux";
import SweetAlert from "react-bootstrap-sweetalert";

class Register extends Component {
  constructor(props) {
    super(props);
    this.state = {
      firstname: "",
      lastname: "",
      email: "",
      password: "",
      show: false,
    };
    this.onChange = this.onChange.bind(this);
    this.onSubmit = this.onSubmit.bind(this);
    this.onConfirm = this.onConfirm.bind(this);
  }
  onChange(e) {
    this.setState({ error: false });
    this.setState({ [e.target.name]: e.target.value });
  }

  onSubmit(e) {
    e.preventDefault();
    const registerData = {
      firstname: this.state.firstname,
      lastname: this.state.lastname,
      password: this.state.password,
      email: this.state.email,
    };
    this.props.register(registerData);
  }

  componentWillReceiveProps(nextProps) {
    if (nextProps.register.message !== undefined) {
      document.getElementById("emailErrorMsg").style.display = "block";
    } else this.setState({ show: true });
  }

  onConfirm() {
    this.setState({ show: false });
  }

  successMessage() {
    return (
      <SweetAlert
        success
        title="Congrats, new account successfully created!"
        onConfirm={this.onConfirm}
      >
        I did it!
      </SweetAlert>
    );
  }

  render() {
    return (
      <form onSubmit={this.onSubmit}>
        <div className="base-container" ref={this.props.containerRef}>
          <div className="header">Register</div>
          <div className="content">
            <div className="image">
              <img src={loginImg} alt="img" />
            </div>
            <div className="form">
              <div className="form-group">
                <label>First name</label>
                <input
                  onChange={this.onChange}
                  type="text"
                  name="firstname"
                  placeholder="First name"
                  required
                  minLength={2}
                  autoComplete="off"
                  value={this.state.firstname}
                />
              </div>
              <div className="form-group">
                <label>Last name</label>
                <input
                  onChange={this.onChange}
                  type="text"
                  name="lastname"
                  placeholder="Last name"
                  required
                  minLength={2}
                  autoComplete="off"
                  value={this.state.lastname}
                />
              </div>
              <div className="form-group">
                <label>Email</label>
                <input
                  onChange={this.onChange}
                  type="email"
                  name="email"
                  placeholder="Email"
                  required
                  minLength={5}
                  autoComplete="off"
                  value={this.state.email}
                />
                <p id="emailErrorMsg">
                  This email address is already being used
                </p>
              </div>
              <div className="form-group">
                <label>Password</label>
                <input
                  onChange={this.onChange}
                  type="password"
                  autoComplete="off"
                  required
                  minLength={5}
                  name="password"
                  placeholder="Password"
                  value={this.state.password}
                />
              </div>
            </div>
          </div>
          <div className="footer">
            <button
              type="submit"
              className="btn"
              id="registerBtn"
              style={{ color: "white" }}
            >
              Register
            </button>
          </div>
          {this.state.show ? this.successMessage() : null}
        </div>
      </form>
    );
  }
}

function mapStateToProps(state) {
  return {
    register: state.register,
  };
}

export default connect(mapStateToProps, { register })(Register);
