import React, {useEffect, useState} from 'react';
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

 const App: React.FC = () => {
     const isAuthenticated = useSelector((state: AuthenticationUserStates) => state.authenticated);
     const userID = useSelector((state: AuthenticationUserStates) => state.userID);
     const [isSessionExpired, setIsSessionExpired] = useState(false);

     useEffect(() => {
         const fetchSessionStatus = async () => {
             if (isAuthenticated && userID) {
                 try {
                     const expired = await checkSessionExpiration(userID);
                     setIsSessionExpired(expired);
                 } catch (error) {
                     console.error('Failed to fetch session status:', error);
                 }
             }
         };
         fetchSessionStatus();
     }, [isAuthenticated, userID]);

     return (
         <Provider store={store}>
             <PersistGate loading={null} persistor={ persistor }>
                 <>
                     <Router>
                         <Nav expired={isSessionExpired}/>
                         <section className="hero is-fullheight">
                             <Routes>
                                 <Route path="/" element={<Home />} />
                                 {isAuthenticated  && !isSessionExpired ?
                                     null :
                                     (<Route path="/login-register" element={<LoginRegister />} />)
                                 }
                                 <Route path="keto-paelo" element={<KetoPaelo />} />
                                 <Route path="whole-popular" element={<WholePopular />} />
                                 <Route path="vegan-vegetarian" element={<Veg />} />
                                 {isAuthenticated  && !isSessionExpired ?
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
