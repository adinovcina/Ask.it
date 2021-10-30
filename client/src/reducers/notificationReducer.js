import * as types from "../actionTypes";

const notificationReducer = (state = [], action) => {
  switch (action.type) {
    case types.GET_NOTIFICATIONS:
      return action.payload;
    default:
      return state;
  }
};

export default notificationReducer;
