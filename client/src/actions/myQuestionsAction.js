import * as types from "../actionTypes";
import axios from "axios";

export const getMyQuestions = () => (dispatch) => {
  axios
    .get("/post/myPosts")
    .then((res) => res.data.data)
    .then((posts) =>
      dispatch({
        type: types.MY_QUESTIONS,
        payload: posts,
      })
    );
};
