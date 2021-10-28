import * as types from "../actionTypes";

const passwordReducer = (state = {}, action) => {
  switch (action.type) {
    case types.CHANGE_PASSWORD:
      return action.payload;
    default:
      return state;
  }
};

export default passwordReducer;
