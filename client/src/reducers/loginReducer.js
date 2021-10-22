import * as types from "../actionTypes";

const loginReducer = (state = {}, action) => {
  switch (action.type) {
    case types.LOGIN:
      return action.payload;
    // case types.LOGOUT:
    //   return (state = {});
    // case types.GET_USER:
    //   return state;
    case types.REGISTER:
      return action.payload;
    default:
      return state;
  }
};

export default loginReducer;
