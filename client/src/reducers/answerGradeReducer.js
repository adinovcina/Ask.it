import * as types from "../actionTypes";

const answerGradeReducer = (state = [], action) => {
  switch (action.type) {
    case types.GET_ANSWER_MARK:
      return action.payload;
    case types.NEW_POST_ANSWER_GRADE:
      if (action.payload.id !== undefined) return [...state, action.payload];
    case types.UPDATE_POST_ANSWER_GRADE:
      return state.map((post) => {
        if (post.id === action.payload.id) {
          return {
            ...post,
            ...action.payload,
          };
        } else {
          return post;
        }
      });
    default:
      return state;
  }
};

export default answerGradeReducer;
