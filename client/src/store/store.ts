import { createStore, Reducer } from 'redux';
import { AuthenticationUserStates, User } from '../types/global';
import { persistStore, persistReducer } from 'redux-persist'
import storage from 'redux-persist/lib/storage'

const initialState: AuthenticationUserStates = {
    authenticated: false,
    authToken: '',
    userID: '',
    email: '',
};

// Define your reducer function
const authReducer: Reducer<AuthenticationUserStates, any> = (
    state = initialState,
    action: any
) => {
    switch (action.type) {
        case 'SET_USER':
            console.log("SET_USER action dispatched");
            return {
                authenticated: true,
                authToken: action.payload.accessToken,
                userID: action.payload.userID,
                email: action.payload.email,
            };
        case 'GET_USER':
            return {
                ...state,
            };
        case 'LOGOUT':
            return initialState;
        default:
            return state;
    }
};

const persistConfig = {
    key: 'root',
    storage,
    whitelist: ['userID'],
}

const persistedReducer = persistReducer(persistConfig, authReducer)

const store = createStore(persistedReducer);
const persistor = persistStore(store);

export { store, persistor };
