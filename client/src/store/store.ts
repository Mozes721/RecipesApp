import { createStore } from 'redux';

export interface AuthenticationUserStates {
    readonly state: 'UNAUTHENTICATED' | 'AUTHENTICATED';
    authToken?: string;
    userID? : string;
    email? : string;
}

