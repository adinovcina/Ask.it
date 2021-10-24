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
