import React from 'react';
import  Nav  from "./components/Navbar/Nav";
import { Provider, useSelector } from 'react-redux';
import store from './store/store';
import { AuthenticationUserStates } from './types/global'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import  LoginRegister  from './pages/LoginRegister';
import  KetoPaelo  from './pages/KetoPaleo';
import  WholePopular from './pages/WholePopular';
import  Veg from './pages/Veg';
import  Yumms from './pages/Yumms';

 const App: React.FC = () => {
     const isAuthenticated = useSelector((state: AuthenticationUserStates) => state.authenticated);
     return (
      <Provider store={store}>
    <>
    <Router>
        <Nav />
        <section className="hero is-fullheight">
            <Routes>
                <Route path="/" element={<LoginRegister />} />
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
      </Provider>
  );
}   

export default App;
