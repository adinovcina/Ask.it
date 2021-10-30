import { combineReducers } from "redux";
import loginReducer from "./loginReducer";
import postReducer from "./postsReducer";
import gradeReducer from "./gradeReducer";
import userReducer from "./userReducer";
import mosLikesReducer from "./mostLikesReducer";
import mostAnswersReducer from "./mostAnswersReducer";
import notificationReducer from "./notificationReducer";
import passwordReducer from "./passwordReducer";
import answerReducer from "./answerReducer";
import registerReducer from "./registerReducer";
import myQuestionsReducer from "./myQuestionsReducer";
import answerGradeReducer from "./answerGradeReducer";
import { persistReducer } from "redux-persist";
import storage from "redux-persist/lib/storage";

const persistConfig = {
  key: "root",
  storage,
};

const rootReducer = combineReducers({
  login: loginReducer,
  register: registerReducer,
  posts: postReducer,
  answers: answerReducer,
  grades: gradeReducer,
  user: userReducer,
  answerGrades: answerGradeReducer,
  mostLikes: mosLikesReducer,
  mostAnswers: mostAnswersReducer,
  myQuestions: myQuestionsReducer,
  password: passwordReducer,
  notifications: notificationReducer,
});

export default persistReducer(persistConfig, rootReducer);
