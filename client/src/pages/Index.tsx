import React from 'react';
import '../assets/style/main.css'
import logoHome from '../assets/images/bg-image-home.jpg'

const Home: React.FC = () => {
    return (
        <section className="hero is-info is-fullheight" style={{ position: 'relative', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
            <div style={{ position: 'absolute', zIndex: 1, textAlign: 'center' }}>
                <img src={logoHome} alt="RecepiesApp" className="bg-home-img" />

                <div className="container has-text-centered">
                    <h1 className="title">
                        Start storing your recipes
                    </h1>
                    <h2 className="subtitle">
                        Here you will be able to find recipes for your liking and give ratings!
                    </h2>
                </div>
            </div>
        </section>
    );
};

export default Home;

