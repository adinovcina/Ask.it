import * as types from "../actionTypes";

const userReducer = (state = {}, action) => {
  switch (action.type) {
    case types.GET_USER:
      return action.payload;
    default:
      return state;
  }
};

export default userReducer;
