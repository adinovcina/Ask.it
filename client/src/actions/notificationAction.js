import * as types from "../actionTypes";
import axios from "axios";

export const getNotifications = (userId) => (dispatch) => {
  console.log("userid", userId);
  axios
    .get("/notification/", {
      params: {
        userId: userId,
      },
    })
    .then((res) => res.data.data)
    .then((answers) =>
      dispatch({
        type: types.GET_NOTIFICATIONS,
        payload: answers,
      })
    );
};
