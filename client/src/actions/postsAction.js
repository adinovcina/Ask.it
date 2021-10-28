import * as types from "../actionTypes";
import axios from "axios";

export const getPosts = () => (dispatch) => {
  axios
    .get("/post/")
    .then((res) => res.data.data)
    .then((posts) =>
      dispatch({
        type: types.GET_POSTS,
        payload: posts,
      })
    );
};

export const createPost = (postData) => (dispatch) => {
  axios
    .post("/post/", postData)
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.NEW_POST,
        payload: post,
      })
    );
};

export const update = (postData) => (dispatch) => {
  axios
    .put("/post/", postData)
    .then((res) => res.data.data)
    .then((post) =>
      dispatch({
        type: types.UPDATE,
        payload: post,
      })
    );
};

// export const update = (postData) => (dispatch) => {
//   axios
//     .post("/grade/", postData)
//     .then((res) => res.data.data)
//     .then((post) =>
//       dispatch({
//         type: types.UPDATE,
//         payload: post,
//       })
//     );
// };
