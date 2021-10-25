import * as types from "../actionTypes";
import axios from "axios";

export const getMostAnswers = () => (dispatch) => {
  axios
    .get("/answer/mostAnswers")
    .then((res) => res.data.data)
    .then((answers) =>
      dispatch({
        type: types.GET_MOST_ANSWERS,
        payload: answers,
      })
    );
};
