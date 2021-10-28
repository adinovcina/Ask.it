import * as types from "../actionTypes";

const answerReducer = (state = [], action) => {
  switch (action.type) {
    case types.GET_ANSWERS:
      return action.payload;
    case types.DELETE_ANSWER:
      return state.filter(({ id }) => id !== action.payload.id);
    case types.ADD_ANSWERS:
      return [...state, action.payload];
    case types.UPDATE_ANSWER:
      return state.map((ans) => {
        if (ans.id === action.payload.id) {
          return {
            ...ans,
            ...action.payload,
          };
        } else {
          return ans;
        }
      });
    case types.UPDATE_ANSWER_MARK:
      return state.map((ans) => {
        if (ans.id === action.payload.id) {
          return {
            ...ans,
            ...action.payload,
          };
        } else {
          return ans;
        }
      });
    default:
      return state;
  }
};

export default answerReducer;
