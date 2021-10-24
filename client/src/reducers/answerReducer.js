import * as types from "../actionTypes";

const answerReducer = (state = {}, action) => {
  switch (action.type) {
    case types.GET_ANSWERS:
      return action.payload;
    // case types.ADD_ANSWERS:
    //   return {
    //     ...state,
    //     item: action.payload,
    //   };
    default:
      return state;
  }
};

export default answerReducer;
