import { setToken } from './Auth';

export function AuthContextProvider(props) {
    const [user, setUser] = useState();

    async function signIn({login, password}) {
        // TODO: Validations here

        // const res = await Api.adminSignIn({login, password});

        // if (!res.auth) {
        //     throw new Error(res.message);
        // }

        // if (res.token) {
            setToken(login);
        // }

        // const {name} = res.user;

        setUser({
            login
        });
    }

    return (
        <AuthContext.Provider value={{user, signIn}}>
            {props.children}
        </AuthContext.Provider>
    );
}