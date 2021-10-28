import * as types from "../actionTypes";

const gradeReducer = (state = [], action) => {
  switch (action.type) {
    case types.GET_GRADES:
      return action.payload;
    case types.NEW_POST_GRADE:
      if (action.payload.id !== undefined) return [...state, action.payload];
    case types.UPDATE_POST_GRADE:
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

export default gradeReducer;
