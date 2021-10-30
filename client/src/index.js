import React from "react";
import ReactDOM from "react-dom";
import App from "./App";
import { Provider } from "react-redux";
import LoginRegister from "./components/Login&Register/LoginRegister";
import reportWebVitals from "./reportWebVitals";
import { PersistGate } from "redux-persist/integration/react";
import { Switch, Route, BrowserRouter } from "react-router-dom";
import { store, persistor } from "./store";
import axios from "axios";
import "bootstrap/dist/css/bootstrap.min.css";
import Navbar from "./components/Navbar/Navbar";
import MostLikes from "./components/MostLikes/MostLikes";
import MostAnswers from "./components/MostAnswers/MostAnswers";
import MyQuestions from "./components/MyQuestions/MyQuestions";
import Profile from "./components/Profile/Profile";
import Notifications from "./components/Notifications/Notifications";

axios.defaults.baseURL = "http://localhost:8080/api/";
// axios.defaults.baseURL = "https://askit-golang.herokuapp.com/api/";

store.subscribe(() => {
  axios.defaults.headers = {
    Authorization: store.getState().login.token,
  };
});
ReactDOM.render(
  <React.StrictMode>
    <Provider store={store}>
      <BrowserRouter>
        <PersistGate persistor={persistor}>
          <Switch>
            <>
              <Navbar />
              <Route path="/login" component={LoginRegister} exact />
              <Route path="/mostLikes" component={MostLikes} exact />
              <Route path="/mostAnswers" component={MostAnswers} exact />
              <Route path="/myQuestions" component={MyQuestions} exact />
              <Route path="/profile" component={Profile} exact />
              <Route path="/notifications" component={Notifications} exact />
              <Route path="/" component={App} exact />
            </>
          </Switch>
        </PersistGate>
      </BrowserRouter>
    </Provider>
  </React.StrictMode>,
  document.getElementById("root")
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
