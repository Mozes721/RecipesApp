import { createUserWithEmailAndPassword  , Auth } from 'firebase/auth';
import {storeUser} from "../../hooks/storeUser";

export function signIn(authInstance: Auth, loginEmail: string, loginPassword: string): Promise<void> {
    return new Promise((resolve, reject) => {
        createUserWithEmailAndPassword(authInstance, loginEmail, loginPassword)
            .then((userCredential) => {
                const userID = userCredential.user.uid;
                const userEmail = userCredential.user.email;
                userCredential.user.getIdToken(true)
                    .then(idToken => {
                        storeUser(userID, userEmail, idToken);
                        resolve();
                    })
                    .catch(error => {
                        console.error('Error getting ID token:', error);
                        reject(error)
                    });
            });
    });
}
