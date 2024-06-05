import { signInWithEmailAndPassword, Auth } from 'firebase/auth';
import {storeUser} from "../../hooks/storeUser";
import {cacheUserToken} from "./cacheToken";

export function login(authInstance: Auth, loginEmail: string, loginPassword: string): Promise<void> {
    return new Promise((resolve, reject) => {
        signInWithEmailAndPassword(authInstance, loginEmail, loginPassword)
            .then((userCredential) => {
                const userID = userCredential.user.uid;
                const userEmail = userCredential.user.email;
                userCredential.user.getIdToken(true)
                    .then(idToken => {
                        cacheUserToken(userID, idToken).catch(error => {
                            reject(error)
                        });
                        storeUser(userID, userEmail);
                        storeUser(userID, userEmail);
                    })
                    .catch(error => {
                        reject(error)
                    });
            });
    });
}

