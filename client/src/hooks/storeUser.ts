import { signInWithEmailAndPassword, Auth } from 'firebase/auth';

export function storeUser(userID: string, email: string | null, accessToken: string) {
    console.log(userID, email, accessToken);
}