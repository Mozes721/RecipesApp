import { GoogleAuthProvider, signInWithPopup, Auth } from 'firebase/auth';
import { storeUser } from "../../hooks/storeUser";
import {cacheUserToken} from "./cacheToken";


const provider = new GoogleAuthProvider();
provider.setCustomParameters({ prompt: 'select_account' });

// Function to sign in with Google
export function signInWithGoogle(authInstance: Auth): Promise<void> {
    return new Promise((resolve, reject) => {
        signInWithPopup(authInstance, provider)
            .then((userCredential) => {
                const userID = userCredential.user.uid;
                const userEmail = userCredential.user.email;
                userCredential.user.getIdToken(true)
                    .then(idToken => {
                        cacheUserToken(userID, idToken).catch(error => {
                            reject(error)
                        });
                        storeUser(userID, userEmail);
                    })
                    .catch(error => {
                        reject(error)
                    });
            })
    });
}


