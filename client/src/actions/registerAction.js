import * as types from "../actionTypes";
import axios from "axios";

export const register = (postData) => (dispatch) => {
  axios
    .post("/auth/register", postData, {
      method: "POST",
      headers: {
        "content-type": "application/json",
      },
    })
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.REGISTER,
        payload: post,
      })
    )
    .catch((err) =>
      dispatch({
        type: types.REGISTER,
        payload: err,
      })
    );
};
