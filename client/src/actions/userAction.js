import * as types from "../actionTypes";
import { store } from "../store";
import _ from "lodash";

export const getUser = () => (dispatch) => {
  store.dispatch(() => {
    if (!_.isEmpty(store.getState().login)) {
      dispatch({
        type: types.GET_USER,
        payload: store.getState().login,
      });
    }
  });
};
