import * as types from "../actionTypes";

const myQuestionsReducer = (state = {}, action) => {
  switch (action.type) {
    case types.MY_QUESTIONS:
      return action.payload;
    default:
      return state;
  }
};

export default myQuestionsReducer;
