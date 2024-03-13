import { signInWithEmailAndPassword, Auth } from 'firebase/auth';
import {storeUser} from "../../hooks/storeUser";
import {AuthenticationUserStates} from "../../types/global";

export function login(authInstance: Auth, loginEmail: string, loginPassword: string): Promise<void> {
    return new Promise((resolve, reject) => {
        signInWithEmailAndPassword(authInstance, loginEmail, loginPassword)
            .then((userCredential) => {
                userCredential.user.getIdToken(true)
                    .then(idToken => {
                        let cachedUser: AuthenticationUserStates = {
                            authenticated: true,
                            authToken: idToken,
                            userID: userCredential.user.uid,
                            email: userCredential.user.email || ''
                        };
                        storeUser(cachedUser);
                        resolve();
                    })
                    .catch(error => {
                        console.error('Error getting ID token:', error);
                        reject(error)
                    });
            });
    });
}

