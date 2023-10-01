import React, {useState} from 'react';
import RegisterForm from '../components/Auth/Register'; 
import LoginForm from '../components/Auth/Login';
const HomePage: React.FC = () => {
    const [isRegister, setRegister] = React.useState(false);

    const toggleForm = () => {
        setRegister(!isRegister);
    }
    return (
        <section className="hero is-info is-fullheight has-background-dark">
          <div className="hero-body">
            <div className="container">
              <div className="columns is-centered">
                <div className="column is-three-fifths-desktop is-three-quarters-tablet has-text-centered">
                  <p className="bd-notification is-primary">
                    <p className="title">{isRegister ? 'Register' : 'Login'}</p>
                    <div className="card">
                      <div className="card-content">
                        {isRegister ? <RegisterForm  onSwitchForm={toggleForm} /> : <LoginForm  onSwitchForm={toggleForm} />}
                        <hr />
                        <p>
                          <a href="#" className="has-text-link has-text-weight-bold" onClick={toggleForm}>
                            {isRegister ? 'Already a member?' : 'Not yet registered?'}
                          </a>
                        </p>
                      </div>
                    </div>
                  </p>
                </div>
              </div>
            </div>
          </div>
        </section>
      );
    };

export default HomePage;