import React from 'react';
import  Nav  from "./components/Navbar/Nav";
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import  HomePage  from './pages/HomePage';
import  KetoPaeloPage  from './pages/KetoPaleoPage';
import  WholePopularPage from './pages/WholePopularPage';
import  VegPage from './pages/VegPage';
import  YummsPage from './pages/YummsPage';

 const App: React.FC = () => {
  return (
    <>
    <Router>
        <Nav />
        <section className="hero is-fullheight">
            <Routes>
                <Route path="/" element={<HomePage />} />
                <Route path="keto-paelo" element={<KetoPaeloPage />} />
                <Route path="whole-popular" element={<WholePopularPage />} />
                <Route path="vegan-vegetarian" element={<VegPage />} />
                <Route path="yumms" element={<YummsPage />} />
            </Routes>
        </section>
    </Router>
    <script async type="text/javascript" src="../js/bulma.js"></script>
    </>
  );
}

export default App;
