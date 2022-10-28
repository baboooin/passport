export const TOKEN_KEY = 'XYZ';
export const getToken = () => localStorage.getItem(TOKEN_KEY);
export const isAuthenticated = () => getToken() !== null; // I thought about validating the token right here in the app
export const setToken = (token) => {
    localStorage.setItem(TOKEN_KEY, token);
};
export const logout = () => {
    localStorage.removeItem(TOKEN_KEY);
};
