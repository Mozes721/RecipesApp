import { GoogleAuthProvider, signInWithPopup, Auth } from 'firebase/auth';


const provider = new GoogleAuthProvider();
provider.setCustomParameters({ prompt: 'select_account' });

// Function to sign in with Google
export function signInWithGoogle(authInstance: Auth) {
    return new Promise((resolve, reject) => {
        signInWithPopup(authInstance, provider)
            .then((userCredential) => {
                const user = userCredential.user;
                resolve(user);
            })
            .catch((error) => {
                reject(error);
            });
    });
}

export { provider };
