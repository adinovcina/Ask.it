import * as types from "../actionTypes";

const gradeReducer = (state = {}, action) => {
  switch (action.type) {
    case types.GET_GRADES:
      return action.payload;
    default:
      return state;
  }
};

export default gradeReducer;
