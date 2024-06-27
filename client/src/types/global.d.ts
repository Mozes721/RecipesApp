interface AuthenticationUserStates {
    authenticated?: boolean;
    authToken?: string;
    userID?: string;
    email?: string;
}

interface Recepie {
    id?: string;
    userID?: string;
    title?: string;
    url?: string;
    imageUrl?: string;
    made?: boolean;
    rating?: number;
}

interface ReviewedRecepie {
    id?: string;
    rating?: number;
}

interface Recepies {
    recepies: Recepie[];
}

export {  AuthenticationUserStates, Recepie, Recepies, ReviewedRecepie };
