import { GoogleAuthProvider, signInWithPopup, Auth } from 'firebase/auth';

// Initialize the GoogleAuthProvider
const provider = new GoogleAuthProvider();
provider.setCustomParameters({ prompt: 'select_account' });

// Function to sign in with Google
export function signInWithGoogle(authInstance: Auth) {
    return new Promise((resolve, reject) => {
        signInWithPopup(authInstance, provider)
            .then((userCredential) => {
                // Handle successful Google sign-in here
                const user = userCredential.user;
                resolve(user);
            })
            .catch((error) => {
                // Handle error here
                reject(error);
            });
    });
}

// Export the initialized provider and the signInWithGoogle function
export { provider };
