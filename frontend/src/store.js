import { createStore } from 'vuex';

export const store = createStore({
    state: {
        credentials: {
            email: localStorage.getItem('email') || "",
            host: localStorage.getItem('host') || "",
            database: localStorage.getItem('database') || "",
            user: localStorage.getItem('user') || "",
            password: localStorage.getItem('password') || ""
        }
    },
    mutations: {
        setCredentials(state, creds) {
            state.credentials = creds;

            localStorage.setItem('email', creds.email);
            localStorage.setItem('host', creds.host);
            localStorage.setItem('database', creds.database);
            localStorage.setItem('user', creds.user);
            localStorage.setItem('password', creds.password);
        }
    }
});