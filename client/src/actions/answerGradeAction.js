import * as types from "../actionTypes";
import axios from "axios";

export const getAnswerGrades = () => (dispatch) => {
  axios
    .get("/answer/grade/")
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.GET_ANSWER_MARK,
        payload: post,
      })
    );
};
