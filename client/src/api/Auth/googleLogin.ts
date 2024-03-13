import { GoogleAuthProvider, signInWithPopup, Auth } from 'firebase/auth';
import { storeUser } from "../../hooks/storeUser";
import {AuthenticationUserStates} from "../../types/global";


const provider = new GoogleAuthProvider();
provider.setCustomParameters({ prompt: 'select_account' });

// Function to sign in with Google
export function signInWithGoogle(authInstance: Auth): Promise<void> {
    return new Promise((resolve, reject) => {
        signInWithPopup(authInstance, provider)
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
            })
    });
}


