import * as types from "../actionTypes";

const registerReducer = (state = {}, action) => {
  switch (action.type) {
    case types.REGISTER:
      return action.payload;
    default:
      return state;
  }
};

export default registerReducer;
