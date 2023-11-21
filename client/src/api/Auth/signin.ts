import { createUserWithEmailAndPassword  , Auth } from 'firebase/auth';
import {storeUser} from "../../hooks/storeUser";

export function signIn(authInstance: Auth, loginEmail: string, loginPassword: string): Promise<void> {
    return new Promise((resolve, reject) => {
        createUserWithEmailAndPassword(authInstance, loginEmail, loginPassword)
            .then((userCredential) => {
                const userID = userCredential.user.uid;
                const userEmail = userCredential.user.email;
                const userToken = userCredential.user.refreshToken;
                storeUser(userID, userEmail, userToken);
                resolve();
            })
            .catch((error) => {
                reject(error);
            });
    });
}
