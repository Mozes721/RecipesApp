import { GoogleAuthProvider, signInWithPopup, Auth } from 'firebase/auth';
import { storeUser } from "../../hooks/storeUser";


const provider = new GoogleAuthProvider();
provider.setCustomParameters({ prompt: 'select_account' });

// Function to sign in with Google
export function signInWithGoogle(authInstance: Auth): Promise<void> {
    return new Promise((resolve, reject) => {
        signInWithPopup(authInstance, provider)
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

export { provider };
