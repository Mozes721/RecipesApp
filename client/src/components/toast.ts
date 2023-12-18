// Import the toast function
import * as bulmaToast from 'bulma-toast'
import {toast} from "bulma-toast";
import {type} from "os";


export const toastNotification = (message: string, type: any) => {
    bulmaToast.toast({
        message: message,
        type: type,
        dismissible: true,
        duration: 1000,
        position: 'top-left',
        closeOnClick: true,
    });
};


