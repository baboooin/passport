import React,{useState} from "react";


const AuthFn = {
    isAuthenticated: false,
    signin(login, password, cb) {
        //connect to server 
        


        AuthFn.isAuthenticated = true;
      setTimeout(cb, 100); // fake async
    },
    signout(cb) {
        AuthFn.isAuthenticated = false;
      setTimeout(cb, 100);
    }
  };

  export default function ProvideAuth({ children }) {
    const auth = useProvideAuth();
    return (
      <authContext.Provider value={auth}>
        {children}
      </authContext.Provider>
    );
  }
  
  function useAuth() {
    return useContext(authContext);
  }
  
  

export default function useProvideAuth() {
    const [user, setUser] = useState(null);
  
    const signin = cb => {
      return fakeAuth.signin(() => {
        setUser("user");
        cb();
      });
    };
  
    const signout = cb => {
      return fakeAuth.signout(() => {
        setUser(null);
        cb();
      });
    };
  
    return {
      user,
      signin,
      signout
    };
  }


  