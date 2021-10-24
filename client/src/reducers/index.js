import { combineReducers } from "redux";
import loginReducer from "./loginReducer";
import postReducer from "./postsReducer";
import gradeReducer from "./gradeReducer";
import userReducer from "./userReducer";
import answerReducer from "./answerReducer";
import answerGradeReducer from "./answerGradeReducer";
import { persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage";

const persistConfig = {
  key: "root",
  storage,
};

const rootReducer = combineReducers({
  login: loginReducer,
  posts: postReducer,
  answers: answerReducer,
  grades: gradeReducer,
  user: userReducer,
  answerGrades: answerGradeReducer,
});

export default persistReducer(persistConfig, rootReducer);
