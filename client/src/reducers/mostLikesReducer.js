import * as types from "../actionTypes";

const mostLikesReducer = (state = {}, action) => {
  switch (action.type) {
    case types.GET_MOST_LIKES:
      return action.payload;
    default:
      return state;
  }
};

export default mostLikesReducer;
