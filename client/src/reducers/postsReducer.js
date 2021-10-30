import * as types from "../actionTypes";

const postReducer = (state = [], action) => {
  switch (action.type) {
    case types.GET_POSTS:
      return action.payload;
    case types.NEW_POST:
      return [...state, action.payload];
    // case types.UPDATE:
    //   return state.map((post) => {
    //     if (post.id === action.payload.id) {
    //       return {
    //         ...post,
    //         ...action.payload,
    //       };
    //     } else {
    //       return post;
    //     }
    //   });

    default:
      return state;
  }
};

export default postReducer;
