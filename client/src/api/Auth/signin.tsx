import { createUserWithEmailAndPassword  , Auth } from 'firebase/auth';

export function signIn(authInstance: Auth, loginEmail: string, loginPassword: string) {
    return new Promise((resolve, reject) => {
        createUserWithEmailAndPassword(authInstance, loginEmail, loginPassword)
            .then((userCredential) => {
                const user = userCredential.user;
                resolve(user); // Resolve with user data on success
            })
            .catch((error) => {
                reject(error); // Reject with the error on failure
            });
    });
}