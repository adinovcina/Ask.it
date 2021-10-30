import * as types from "../actionTypes";
import axios from "axios";
import { persistor } from "../store";

export const login = (postData) => (dispatch) => {
  axios
    .post("/auth/login", postData, {
      method: "POST",
      headers: {
        "content-type": "application/json",
      },
    })
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.LOGIN,
        payload: post,
      })
    )
    .catch((err) =>
      dispatch({
        type: types.LOGIN,
        payload: err,
      })
    );
};

export const logout = () => (dispatch) => {
  persistor.pause();
  persistor.flush().then(() => {
    return persistor.purge();
  });
  dispatch({
    type: types.LOGOUT,
    payload: {},
  });
};
