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

export const postAnswerGrade = (postData) => (dispatch) => {
  axios
    .post("/answer/grade/", postData)
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.NEW_POST_ANSWER_GRADE,
        payload: post,
      })
    );
};

export const updateAnswerGrade = (postData) => (dispatch) => {
  axios
    .put("/answer/grade/", postData)
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.UPDATE_POST_ANSWER_GRADE,
        payload: post,
      })
    );
};
