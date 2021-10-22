import allReducers from "./reducers";
import thunk from "redux-thunk";
import { createStore } from "redux";
import { applyMiddleware } from "redux";
import { compose } from "redux";
import { persistStore } from "redux-persist";

const middleware = [thunk];

export const store = createStore(
  allReducers,
  compose(
    applyMiddleware(...middleware),
    window.__REDUX_DEVTOOLS_EXTENSION__
      ? window.__REDUX_DEVTOOLS_EXTENSION__()
      : (f) => f
  )
);

export const persistor = persistStore(store);

const exportedObject = { store, persistor };

export default exportedObject;
