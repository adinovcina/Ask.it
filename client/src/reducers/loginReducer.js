import * as types from "../actionTypes";

const loginReducer = (state = {}, action) => {
  switch (action.type) {
    case types.LOGIN:
      return action.payload;
    default:
      return state;
  }
};

export default loginReducer;
