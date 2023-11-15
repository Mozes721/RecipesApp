import { signInWithEmailAndPassword, Auth } from 'firebase/auth';

export function login(authInstance: Auth, loginEmail: string, loginPassword: string) {
    return new Promise((resolve, reject) => {
        signInWithEmailAndPassword(authInstance, loginEmail, loginPassword)
            .then((userCredential) => {
                const user = userCredential.user;
                resolve(user.uid); 
            })
            .catch((error) => {
                reject(error); 
            });
    });
}
