import React, {useEffect } from 'react';
import  Nav  from "./components/Navbar/Nav";
import { Provider, useSelector } from 'react-redux';
import { PersistGate } from "redux-persist/integration/react";
import { store, persistor } from './store/store';
import { AuthenticationUserStates } from './types/global'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import  Home from './pages/Index';
import  LoginRegister  from './pages/LoginRegister';
import  KetoPaelo  from './pages/KetoPaleo';
import  WholePopular from './pages/WholePopular';
import  Veg from './pages/Veg';
import  Yumms from './pages/Yumms';
import { checkSessionExpiration } from "./api/Auth/checkExpiration";
import { logoutSession } from "./hooks/removeUser";

 const App: React.FC = () => {
     const isAuthenticated = useSelector((state: AuthenticationUserStates) => state.authenticated);
     const userID = useSelector((state: AuthenticationUserStates) => state.userID);

     useEffect(() => {
         const fetchSessionStatus = async () => {
             const expired  = await checkSessionExpiration(userID);
             if (expired) {
                 logoutSession();
             }
         };
         fetchSessionStatus();
     }, []);


     return (
         <Provider store={store}>
             <PersistGate loading={null} persistor={ persistor }>
                 <>
                     <Router>
                         <Nav />
                         <section className="hero is-fullheight">
                             <Routes>
                                 <Route path="/" element={<Home />} />
                                 {isAuthenticated ?
                                     null :
                                     (<Route path="/login-register" element={<LoginRegister />} />)
                                 }
                                 <Route path="keto-paelo" element={<KetoPaelo />} />
                                 <Route path="whole-popular" element={<WholePopular />} />
                                 <Route path="vegan-vegetarian" element={<Veg />} />
                                 {isAuthenticated ?
                                     (<Route path="yumms" element={<Yumms /> } />)
                                     : null }
                             </Routes>
                         </section>
                     </Router>
                     <script async type="text/javascript" src="../js/bulma.js"></script>
                 </>
             </PersistGate>
         </Provider>
     );
 }

export default App;
