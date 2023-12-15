interface User {
    userID?: string;
}

interface AuthenticationUserStates {
    authenticated?: boolean;
    authToken?: string;
    userID?: string;
    email?: string;
}

interface Recepie {
    userID?: string;
    title?: string;
    url?: string;
    made?: boolean;
    rating?: number;
}


export { User, AuthenticationUserStates, Recepie };
