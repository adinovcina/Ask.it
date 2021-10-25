import * as types from "../actionTypes";

const mostAnswerReducer = (state = {}, action) => {
  switch (action.type) {
    case types.GET_MOST_ANSWERS:
      return action.payload;
    default:
      return state;
  }
};

export default mostAnswerReducer;
