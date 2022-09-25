import React from 'react';
import  Nav  from "./components/Navbar/Nav";
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import  HomePage  from './pages/HomePage';
import  KetoPaeloPage  from './pages/KetoPaleoPage';
import  WholePopularPage from './pages/WholePopularPage';
import  VegPage from './pages/VegPage';
import CousinePage from './pages/CousinesPage';
import  YummsPage from './pages/YummsPage';



 
const App: React.FC = () => {
  return (
    <>
    <Router>
        <Nav />
        <section className="hero is-fullheight">
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="ketoPaelo" element={<KetoPaeloPage />} />
              <Route path="meat" element={<Meat />} />
              <Route path="popular" element={<Popular />} />
              <Route path="yumms" element={<Yumms />} />
            </Routes>
        </section>
    </Router>
    <script async type="text/javascript" src="../js/bulma.js"></script>
    </>
  );
}

export default App;
