import React from 'react';
import logoHome from '../assets/images/bg-image-home.jpg';

const Home: React.FC = () => {
    return (
        <section className="hero is-info is-fullheight">
            <div className="hero-body">
                <div className="container">
                    <div className="columns is-vcentered">
                        <div className="column is-half">
                            <img src={logoHome} alt="RecepiesApp" className="img-responsive" />
                        </div>
                        <div className="column is-half has-text-centered">
                            <h1 className="title">
                                Start storing your recipes
                            </h1>
                            <h2 className="subtitle">
                                Here you will be able to find recipes for your liking and give ratings!
                            </h2>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    );
};

export default Home;


