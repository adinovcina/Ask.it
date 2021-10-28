import * as types from "../actionTypes";
import axios from "axios";

export const getGrades = () => (dispatch) => {
  axios
    .get("/grade/")
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.GET_GRADES,
        payload: post,
      })
    );
};

export const postGrade = (postData) => (dispatch) => {
  axios
    .post("/grade/", postData)
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.NEW_POST_GRADE,
        payload: post,
      })
    );
};

export const updateGrade = (postData) => (dispatch) => {
  axios
    .put("/grade/", postData)
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.UPDATE_POST_GRADE,
        payload: post,
      })
    );
};
