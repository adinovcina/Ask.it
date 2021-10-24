import * as types from "../actionTypes";

const answerGradeReducer = (state = {}, action) => {
  switch (action.type) {
    case types.GET_ANSWER_MARK:
      return action.payload;
    default:
      return state;
  }
};

export default answerGradeReducer;
