import * as types from "../actionTypes";
import axios from "axios";

export const passwordChange = (data) => (dispatch) => {
  axios
    .post("/changepassword/", data)
    .then((res) => res.data.data)
    .then((user) =>
      dispatch({
        type: types.CHANGE_PASSWORD,
        payload: user,
      })
    )
    .catch((err) =>
      dispatch({
        type: types.CHANGE_PASSWORD,
        payload: err,
      })
    );
};
