import * as types from "../actionTypes";
import axios from "axios";

export const getMostLikes = () => (dispatch) => {
  axios
    .get("/post/mostLikes")
    .then((res) => res.data.data)
    .then((answers) =>
      dispatch({
        type: types.GET_MOST_LIKES,
        payload: answers,
      })
    );
};
