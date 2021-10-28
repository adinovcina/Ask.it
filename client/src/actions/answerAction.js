import * as types from "../actionTypes";
import axios from "axios";

export const getAnswers = () => (dispatch) => {
  axios
    .get("/answer/")
    .then((res) => res.data.data)
    .then((answers) =>
      dispatch({
        type: types.GET_ANSWERS,
        payload: answers,
      })
    );
};

export const createAnswer = (comment) => (dispatch) => {
  axios
    .post("/answer/", comment)
    .then((res) => res.data.data)
    .then((comm) =>
      dispatch({
        type: types.ADD_ANSWERS,
        payload: comm,
      })
    );
};

export const update = (postData) => (dispatch) => {
  axios
    .put("/answer/answerGrade", postData)
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.UPDATE_ANSWER_MARK,
        payload: post,
      })
    );
};

export const editAnswer = (postData) => (dispatch) => {
  axios
    .put("/answer/", postData)
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.UPDATE_ANSWER,
        payload: post,
      })
    );
};

export const deleteAnswer = (postData) => (dispatch) => {
  axios
    .delete(`/answer/${postData}`)
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.DELETE_ANSWER,
        payload: post,
      })
    );
};
