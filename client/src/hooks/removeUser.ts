import { store } from '../store/store'

export function removeUser(event: React.MouseEvent<HTMLSpanElement, MouseEvent>): Promise<void> {
    event.stopPropagation();
    return new Promise((resolve) => {
        store.dispatch({
            type: 'LOGOUT',
        });
        resolve();
    });
}

export function logoutSession() {
    store.dispatch({
        type: 'LOGOUT'
    });
}
