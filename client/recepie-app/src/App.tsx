import React from 'react';
import  Nav  from "./components/Navbar/Nav";
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { Home } from './pages/HomePage';
import { Keto } from './pages/KetoPage';
import { Meat } from './pages/MeatPage';
import { Popular } from './pages/PopularPage';
import { Yumms } from './pages/YummsPage';




function App() {
  return (
    <>

    <Router>
        <Nav />
        <section className="hero is-fullheight">
          <div>
            <Routes>
              <Route path="/" element={<Home />} />
              <Route path="keto" element={<Keto />} />
              <Route path="meat" element={<Meat />} />
              <Route path="popular" element={<Popular />} />
              <Route path="yumms" element={<Yumms />} />
            </Routes>
          </div>
        </section>
    </Router>
    <script async type="text/javascript" src="../js/bulma.js"></script>
    </>
  );
}

export default App;
